/**
 * 计算 BSC 部署花费的 BNB
 * 运行：npx hardhat run scripts/deployment-cost.ts --network bsc
 */
import { createPublicClient, http } from "viem";
import { bsc } from "viem/chains";

const TX_HASHES = [
  "0xb64380ad4eeb490c88e0a11236b62765f262bb159cfabfd31e54bcda8502f173",
  "0x77a483846f47b0b9c56ced32e4ac607b717137c58e6f037623c01ba215804263",
  "0x2998c23832797c082502fefc6cf3224cb6ac42c419f1b73d2ded04ee0083c450",
  "0x4e9c5abbcf55f5a3dca58310b7f19cd467e49058f2b76545ded6aa6e123384ed",
  "0x27bc1379a18bca0c2b8c88dd67505db0383f2ecf154c04299e8c590548184dff",
  "0x2246aae0c92761ceb57e4849d69c2be87120818e01928c513a00d5413cf25893",
  "0xe47de3dd4fd3defb6687be901ac50025ede2f8a910e45506cdf34e953ed57b95",
  "0x3d1380e74d473efcce942f47c679e003775fd85add3c981d0b031b22669a9a7c",
  "0xab4c550f5aa04e2c8e073d40331381680f83e7ee10a831d401f75911cdf2a8cb",
  "0x265eed63744338c92d2ce68a378ec515eb1b45beaa816453b207e338a28c968a",
  "0x31e3014e375ca79eee9522b705acf99ebe2badf81b66c10ccd490d2d919ddc8d",
  "0xe1d7a3ab2943b8024ab3753fe683097cee1ce1e5e7885577c8409231332994f6",
  "0x211ccd8bb891962be506806ca724ddb6045db30d4dbe830904fa90ad7168015f",
  "0x5f162a1328360d78b7519aff5b170bcaa9d318ef44ad36a571e72baa013b934a",
  "0x23bbf8d064eca163367ba034f18fac8d59c1b1d5f33cda654b5a025d242ae483",
  "0x6a23f03421d03d7c0da1c5adc2e18d322363ac1dce0ba30395b26713a1e2ddc5",
  "0xffea31e983f92a767e567face453c8d4d328ba173fd9a8d47fd9192448b5aa9a",
  "0x63dc6b26903732ad28df1133dd420162db38954204dfdb618442354593a1ed72",
  "0x81aa5e60c2aa2ba2b5c0b6fafab4b69687738678e05122606a52b80a20133eeb",
];

async function main() {
  const rpc = process.env.BSC_RPC_URL || "https://bsc.publicnode.com";
  const client = createPublicClient({ chain: bsc, transport: http(rpc) });

  let totalGasUsed = 0n;
  let totalEffectiveGasPrice = 0n;
  const receipts: { hash: string; gasUsed: bigint; effectiveGasPrice: bigint }[] = [];

  for (const hash of TX_HASHES) {
    const receipt = await client.getTransactionReceipt({ hash: hash as `0x${string}` });
    if (receipt) {
      totalGasUsed += receipt.gasUsed;
      totalEffectiveGasPrice += receipt.effectiveGasPrice;
      receipts.push({
        hash: hash.slice(0, 18) + "...",
        gasUsed: receipt.gasUsed,
        effectiveGasPrice: receipt.effectiveGasPrice,
      });
    }
  }

  // 用第一笔的 effectiveGasPrice 估算（通常相同）
  const avgGasPrice = receipts.length > 0 ? receipts[0].effectiveGasPrice : 0n;
  const totalWei = totalGasUsed * avgGasPrice;
  const bnb = Number(totalWei) / 1e18;

  console.log("\n=== BSC 部署 Gas 消耗 ===\n");
  console.log(`  交易笔数: ${receipts.length}`);
  console.log(`  总 Gas:   ${totalGasUsed.toString()}`);
  console.log(`  Gas 单价: ${avgGasPrice.toString()} wei (${Number(avgGasPrice) / 1e9} Gwei)`);
  console.log(`  总花费:   ${totalWei.toString()} wei`);
  console.log(`\n  约 ${bnb.toFixed(6)} BNB\n`);
}

main().catch(console.error);
