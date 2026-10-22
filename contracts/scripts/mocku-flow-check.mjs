import {
  createPublicClient,
  createWalletClient,
  decodeEventLog,
  formatEther,
  getAddress,
  http,
  parseAbi,
  parseEther,
  parseUnits,
} from "viem";
import { generatePrivateKey, privateKeyToAccount } from "viem/accounts";
import { bsc } from "viem/chains";

const rpc = process.env.BSC_RPC_URL || "https://bnb56743.allnodes.me:8545/or6NBNSPsU9uDxP3";
const pk = (process.env.TEST_PRIVATE_KEY || "").trim().replace(/^0x/i, "");
if (!pk) throw new Error("TEST_PRIVATE_KEY is required");

const owner = privateKeyToAccount(`0x${pk}`);
const secondPk = generatePrivateKey();
const second = privateKeyToAccount(secondPk);

const publicClient = createPublicClient({ chain: bsc, transport: http(rpc) });
const ownerClient = createWalletClient({ account: owner, chain: bsc, transport: http(rpc) });
const secondClient = createWalletClient({ account: second, chain: bsc, transport: http(rpc) });

const ADDR = {
  mockUsdt: getAddress("0x2674DE1Fe0c884D45cd128b3F67Eb03BFD649ff4"),
  nft: getAddress("0x73F76b11ebF03718B71c163BE2B5445ECFC25Ce6"),
};

const erc20Abi = parseAbi([
  "function balanceOf(address) view returns (uint256)",
  "function approve(address spender, uint256 amount) returns (bool)",
  "function transfer(address to, uint256 amount) returns (bool)",
  "function mint(address to, uint256 amount)",
]);

const nftAbi = parseAbi([
  "function mint(uint8 tier) returns (uint256 tokenId)",
  "function ownerOf(uint256 tokenId) view returns (address)",
  "function safeTransferFrom(address from, address to, uint256 tokenId)",
  "function openBoxPayUsdt(uint256 tokenId)",
  "function setRewardOnce(uint256 tokenId, int256 reward_)",
  "function boxInfo(uint256 tokenId) view returns (uint8 tier, uint256 usdtPaid, uint64 mintedAt, uint64 openedAt, int256 reward, uint64 rewardSetAt, uint256 tokenIdInStruct)",
  "event Minted(address indexed minter, uint256 indexed tokenId, uint8 indexed tier, uint256 usdtPaid)",
  "event Opened(address indexed owner, uint256 indexed tokenId, uint256 timestamp, uint256 usdtFeePaid, uint256 feeBPaid)",
  "event RewardSetOnce(address indexed admin, uint256 indexed tokenId, int256 reward, uint256 timestamp, uint256 profitNotified)",
]);

function print(name, payload) {
  console.log(JSON.stringify({ step: name, ...payload }, null, 2));
}

async function wait(hash) {
  return publicClient.waitForTransactionReceipt({ hash });
}

async function main() {
  print("accounts", {
    owner: owner.address,
    second: second.address,
    secondPrivateKey: secondPk,
    rpc,
  });

  const ownerMintAmount = parseUnits("200", 18);
  const secondMintAmount = parseUnits("20", 18);
  const rewardFundAmount = parseUnits("5", 18);
  const mintCost = parseUnits("20", 18);
  const openFee = parseUnits("1", 18);
  const rewardAmount = parseUnits("1", 18);

  let hash = await ownerClient.writeContract({
    address: ADDR.mockUsdt,
    abi: erc20Abi,
    functionName: "mint",
    args: [owner.address, ownerMintAmount],
  });
  await wait(hash);
  print("mint_mockusdt_owner", { hash, amount: ownerMintAmount.toString() });

  hash = await ownerClient.writeContract({
    address: ADDR.mockUsdt,
    abi: erc20Abi,
    functionName: "mint",
    args: [second.address, secondMintAmount],
  });
  await wait(hash);
  print("mint_mockusdt_second", { hash, amount: secondMintAmount.toString() });

  hash = await ownerClient.sendTransaction({
    account: owner,
    to: second.address,
    value: parseEther("0.003"),
  });
  await wait(hash);
  print("fund_second_bnb", { hash, amount: parseEther("0.003").toString() });

  hash = await ownerClient.writeContract({
    address: ADDR.mockUsdt,
    abi: erc20Abi,
    functionName: "approve",
    args: [ADDR.nft, mintCost],
  });
  await wait(hash);
  print("approve_owner_mockusdt", { hash, amount: mintCost.toString() });

  hash = await ownerClient.writeContract({
    address: ADDR.nft,
    abi: nftAbi,
    functionName: "mint",
    args: [1],
  });
  const mintReceipt = await wait(hash);
  let tokenId = null;
  for (const log of mintReceipt.logs) {
    try {
      const parsed = decodeEventLog({ abi: nftAbi, data: log.data, topics: log.topics });
      if (parsed.eventName === "Minted") tokenId = parsed.args.tokenId;
    } catch {}
  }
  if (tokenId == null) throw new Error("Minted tokenId not found");
  print("mint_box", { hash, tokenId: tokenId.toString(), status: mintReceipt.status });

  hash = await ownerClient.writeContract({
    address: ADDR.nft,
    abi: nftAbi,
    functionName: "safeTransferFrom",
    args: [owner.address, second.address, tokenId],
  });
  await wait(hash);
  print("transfer_box_to_second", { hash, tokenId: tokenId.toString() });

  hash = await secondClient.writeContract({
    account: second,
    address: ADDR.mockUsdt,
    abi: erc20Abi,
    functionName: "approve",
    args: [ADDR.nft, openFee],
  });
  await wait(hash);
  print("approve_second_open_fee", { hash, amount: openFee.toString() });

  hash = await secondClient.writeContract({
    account: second,
    address: ADDR.nft,
    abi: nftAbi,
    functionName: "openBoxPayUsdt",
    args: [tokenId],
  });
  const openReceipt = await wait(hash);
  print("open_box", { hash, tokenId: tokenId.toString(), status: openReceipt.status });

  hash = await ownerClient.writeContract({
    address: ADDR.mockUsdt,
    abi: erc20Abi,
    functionName: "transfer",
    args: [ADDR.nft, rewardFundAmount],
  });
  await wait(hash);
  print("fund_nft_with_mockusdt", { hash, amount: rewardFundAmount.toString() });

  hash = await ownerClient.writeContract({
    address: ADDR.nft,
    abi: nftAbi,
    functionName: "setRewardOnce",
    args: [tokenId, rewardAmount],
  });
  const rewardReceipt = await wait(hash);
  print("set_reward_once", { hash, tokenId: tokenId.toString(), status: rewardReceipt.status, rewardAmount: rewardAmount.toString() });

  const ownerUsdt = await publicClient.readContract({ address: ADDR.mockUsdt, abi: erc20Abi, functionName: "balanceOf", args: [owner.address] });
  const secondUsdt = await publicClient.readContract({ address: ADDR.mockUsdt, abi: erc20Abi, functionName: "balanceOf", args: [second.address] });
  const nftUsdt = await publicClient.readContract({ address: ADDR.mockUsdt, abi: erc20Abi, functionName: "balanceOf", args: [ADDR.nft] });
  const boxInfo = await publicClient.readContract({ address: ADDR.nft, abi: nftAbi, functionName: "boxInfo", args: [tokenId] });
  const ownerOf = await publicClient.readContract({ address: ADDR.nft, abi: nftAbi, functionName: "ownerOf", args: [tokenId] });

  print("final_state", {
    tokenId: tokenId.toString(),
    ownerOf,
    ownerMockUsdt: ownerUsdt.toString(),
    secondMockUsdt: secondUsdt.toString(),
    nftMockUsdt: nftUsdt.toString(),
    openedAt: boxInfo[3].toString(),
    reward: boxInfo[4].toString(),
    rewardSetAt: boxInfo[5].toString(),
    secondBnb: formatEther(await publicClient.getBalance({ address: second.address })),
  });
}

main().catch((err) => {
  console.error(err);
  process.exitCode = 1;
});
