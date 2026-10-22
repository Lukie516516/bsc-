/**
 * 验证 BSC 部署：读取链上合约状态确认部署成功
 * 运行：npx hardhat run scripts/verify-deployment.ts --network bsc
 */
import hre from "hardhat";
import { createPublicClient, http } from "viem";
import { bsc } from "viem/chains";

const ADDRESSES: Record<string, string> = {
  BlindBoxNFT: "0x98B085B80E24c25670F45dbF0040F2102c114fB6",
  Piece: "0xFEe7d410ac6D3ad2e30A5c3EF13cb5f01A459D11",
  Hrx: "0x7C8e78c3e53eDb62249fFA2e7b9a2F53Bf2085E2",
  BondingCurvePrimaryMarket: "0x341dC323E1958494Bb93aA0a13A646B4DEA3e3E6",
  SwapHrxUsdt: "0x20590B5c8739Edc25219d939767A17b6E31dDC53",
  UmbrellaStakeDividend: "0x6fc73dDAc8F0fd0B814b0C68fBeCeD793A8f1479",
  DealRouter: "0x46e654047A964801C419A649fe0796E2f92ac271",
  BlindBoxUserVault: "0x2DD55DF8603315c47c915B5e6F631E2Ded23ED26",
  BlindBoxVaultFactory: "0xf9ED59FeDc2A2b095fE9D5C0438eC77C1Af321ba",
  BlindBoxPrincipalVault: "0x5AB00DaD2a941Ba317FfAB4B3038B9B46949962F",
};

async function main() {
  const rpc = process.env.BSC_RPC_URL || "https://bsc.publicnode.com";
  const publicClient = createPublicClient({
    chain: bsc,
    transport: http(rpc),
  });

  console.log("\n=== 合约文件 -> 部署名称 -> 地址 ===\n");
  console.log("Stake.sol         -> UmbrellaStakeDividend ->", ADDRESSES.UmbrellaStakeDividend);
  console.log("Market.sol        -> BondingCurvePrimaryMarket ->", ADDRESSES.BondingCurvePrimaryMarket);
  console.log("StakeTwo.sol      -> BlindBoxPrincipalVault ->", ADDRESSES.BlindBoxPrincipalVault);
  console.log("ListTemplate.sol  -> BlindBoxUserVault ->", ADDRESSES.BlindBoxUserVault);
  console.log("FactoryList.sol   -> BlindBoxVaultFactory ->", ADDRESSES.BlindBoxVaultFactory);
  console.log("");

  // 1. 检查合约是否有代码（确认已部署）
  console.log("=== 1. 检查合约代码是否存在 ===\n");
  for (const [name, addr] of Object.entries(ADDRESSES)) {
    const code = await publicClient.getBytecode({ address: addr as `0x${string}` });
    const ok = code && code !== "0x" && code.length > 10;
    console.log(`  ${name.padEnd(25)} ${ok ? "✓ 已部署" : "✗ 无代码"}`);
  }

  // 2. 读取关键状态（通过 eth_call 模拟）
  console.log("\n=== 2. 读取关键状态 ===\n");

  const nftAbi = [
    { inputs: [], name: "name", outputs: [{ type: "string" }], stateMutability: "view", type: "function" },
  ] as const;
  const erc20Abi = [
    { inputs: [], name: "totalSupply", outputs: [{ type: "uint256" }], stateMutability: "view", type: "function" },
  ] as const;
  const stakeAbi = [
    { inputs: [], name: "nftContract", outputs: [{ type: "address" }], stateMutability: "view", type: "function" },
  ] as const;

  const [nftName, pieceSupply, hrxTotal, stakeNft] = await Promise.all([
    publicClient.readContract({
      address: ADDRESSES.BlindBoxNFT as `0x${string}`,
      abi: nftAbi,
      functionName: "name",
    }),
    publicClient.readContract({
      address: ADDRESSES.Piece as `0x${string}`,
      abi: erc20Abi,
      functionName: "totalSupply",
    }),
    publicClient.readContract({
      address: ADDRESSES.Hrx as `0x${string}`,
      abi: erc20Abi,
      functionName: "totalSupply",
    }),
    publicClient.readContract({
      address: ADDRESSES.UmbrellaStakeDividend as `0x${string}`,
      abi: stakeAbi,
      functionName: "nftContract",
    }),
  ]);

  console.log(`  BlindBoxNFT.name():        ${nftName}`);
  console.log(`  Piece.totalSupply():       ${pieceSupply}`);
  console.log(`  Hrx.totalSupply():         ${hrxTotal}`);
  console.log(`  Stake.nftContract():       ${stakeNft}`);
  console.log(`  (应与 BlindBoxNFT 地址一致: ${ADDRESSES.BlindBoxNFT})`);

  console.log("\n=== 验证完成 ✓ ===\n");
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
