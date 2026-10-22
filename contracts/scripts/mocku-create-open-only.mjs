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
  "function approve(address spender, uint256 amount) returns (bool)",
  "function mint(address to, uint256 amount)",
]);

const nftAbi = parseAbi([
  "function mint(uint8 tier) returns (uint256 tokenId)",
  "function safeTransferFrom(address from, address to, uint256 tokenId)",
  "function openBoxPayUsdt(uint256 tokenId)",
  "event Minted(address indexed minter, uint256 indexed tokenId, uint8 indexed tier, uint256 usdtPaid)",
]);

async function wait(hash) {
  return publicClient.waitForTransactionReceipt({ hash });
}

function out(name, payload) {
  console.log(JSON.stringify({ step: name, ...payload }, null, 2));
}

async function main() {
  const secondMockUsdt = parseUnits("20", 18);
  const mintCost = parseUnits("20", 18);
  const openFee = parseUnits("1", 18);

  let hash = await ownerClient.writeContract({
    address: ADDR.mockUsdt,
    abi: erc20Abi,
    functionName: "mint",
    args: [second.address, secondMockUsdt],
  });
  await wait(hash);

  hash = await ownerClient.sendTransaction({
    account: owner,
    to: second.address,
    value: parseEther("0.003"),
  });
  await wait(hash);

  hash = await ownerClient.writeContract({
    address: ADDR.mockUsdt,
    abi: erc20Abi,
    functionName: "approve",
    args: [ADDR.nft, mintCost],
  });
  await wait(hash);

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

  hash = await ownerClient.writeContract({
    address: ADDR.nft,
    abi: nftAbi,
    functionName: "safeTransferFrom",
    args: [owner.address, second.address, tokenId],
  });
  await wait(hash);

  hash = await secondClient.writeContract({
    account: second,
    address: ADDR.mockUsdt,
    abi: erc20Abi,
    functionName: "approve",
    args: [ADDR.nft, openFee],
  });
  await wait(hash);

  hash = await secondClient.writeContract({
    account: second,
    address: ADDR.nft,
    abi: nftAbi,
    functionName: "openBoxPayUsdt",
    args: [tokenId],
  });
  const openReceipt = await wait(hash);

  out("result", {
    tokenId: tokenId.toString(),
    secondAddress: second.address,
    secondPrivateKey: secondPk,
    openHash: hash,
    openStatus: openReceipt.status,
    secondBnb: formatEther(await publicClient.getBalance({ address: second.address })),
  });
}

main().catch((err) => {
  console.error(err);
  process.exitCode = 1;
});
