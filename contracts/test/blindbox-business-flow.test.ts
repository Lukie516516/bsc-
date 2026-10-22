import assert from "node:assert/strict";
import { describe, it } from "node:test";

import { network } from "hardhat";
import { decodeEventLog, parseEther, zeroAddress } from "viem";

const E18 = 10n ** 18n;
const ONE_MONTH_PLUS = 31 * 24 * 60 * 60;
const MONTH_SECONDS = 30n * 24n * 60n * 60n;
const DEFAULT_ADMIN_ROLE =
  "0x0000000000000000000000000000000000000000000000000000000000000000";
const secondarySaleSplitEventAbi = {
  type: "event",
  name: "SecondarySaleSplit",
  inputs: [
    { name: "buyer", type: "address", indexed: true },
    { name: "seller", type: "address", indexed: true },
    { name: "tokenId", type: "uint256", indexed: true },
    { name: "b", type: "address", indexed: false },
    { name: "c", type: "address", indexed: false },
    { name: "priceUSDT", type: "uint256", indexed: false },
    { name: "toSeller", type: "uint256", indexed: false },
    { name: "toB", type: "uint256", indexed: false },
    { name: "toC", type: "uint256", indexed: false },
    { name: "toProject", type: "uint256", indexed: false },
    { name: "feePaidInB", type: "bool", indexed: false },
    { name: "feeUSDT", type: "uint256", indexed: false },
    { name: "feeB", type: "uint256", indexed: false },
  ],
} as const;

describe("Blind box business flow", async function () {
  const { viem } = await network.connect();
  const publicClient = await viem.getPublicClient();
  const wallets = await viem.getWalletClients();

  async function waitFor(hash: `0x${string}`) {
    return publicClient.waitForTransactionReceipt({ hash });
  }

  async function increaseTime(seconds: number) {
    await (publicClient.transport as any).request({
      method: "evm_increaseTime",
      params: [seconds],
    });
    await (publicClient.transport as any).request({
      method: "evm_mine",
      params: [],
    });
  }

  async function expectRevert(run: Promise<unknown>, label: string) {
    await assert.rejects(run, undefined, label);
  }

  async function deployFixture() {
    const [
      deployer,
      superAdmin,
      admin,
      manager,
      project,
      sellerA,
      buyerB,
      uplineB,
      topD,
      vaultOwner,
      quantReceiver,
    ] = wallets as any[];

    const nftPriceLib = await viem.deployContract("NftPriceLib");
    const nftFeeLib = await viem.deployContract("NftFeeLib");

    const usdt = await viem.deployContract("MockUSDT", [
      deployer.account.address,
      parseEther("100000000"),
    ]);
    const piece = await viem.deployContract("Piece", [
      "Piece",
      "PIECE",
      deployer.account.address,
      parseEther("1000000000"),
    ]);
    const hrx = await viem.deployContract("Hrx", ["Hrx", "HRX"]);
    const market = await viem.deployContract("BondingCurvePrimaryMarket", [
      usdt.address,
      hrx.address,
      deployer.account.address,
      project.account.address,
    ]);
    const stake = await viem.deployContract("UmbrellaStakeDividend", [
      hrx.address,
      usdt.address,
      project.account.address,
      parseEther("300"),
    ]);
    const nft = await viem.deployContract(
      "BlindBoxNFT",
      [
        "BlindBox",
        "BOX",
        usdt.address,
        hrx.address,
        project.account.address,
        uplineB.account.address,
        admin.account.address,
        topD.account.address,
        0n,
        stake.address,
      ],
      {
        libraries: {
          NftPriceLib: nftPriceLib.address,
          NftFeeLib: nftFeeLib.address,
        },
      } as any,
    );
    const swap = await viem.deployContract("SwapHrxUsdt", [
      hrx.address,
      usdt.address,
      deployer.account.address,
      topD.account.address,
      3,
      100,
    ]);
    const listTemplate = await viem.deployContract("BlindBoxUserVault");
    const factory = await viem.deployContract("BlindBoxVaultFactory", [
      listTemplate.address,
      nft.address,
      usdt.address,
    ]);

    await waitFor(await swap.write.setPrimaryMarket([market.address]));
    await waitFor(await nft.write.setFeeQuotePool([swap.address]));
    await waitFor(await nft.write.setPieceToken([piece.address]));
    await waitFor(await nft.write.setPieceValueUsdtE18([E18]));
    await waitFor(await nft.write.setStakeReferralReader([stake.address]));
    await waitFor(await piece.write.setNFT([nft.address]));
    await waitFor(await stake.write.setNftContract([nft.address]));
    await waitFor(await stake.write.setPieceToken([piece.address]));

    const recipients = [
      sellerA.account.address,
      buyerB.account.address,
      uplineB.account.address,
      topD.account.address,
      vaultOwner.account.address,
      project.account.address,
    ];
    const amounts = recipients.map(() => parseEther("100000"));
    await waitFor(await usdt.write.mintBatch([recipients, amounts]));

    for (const wallet of [sellerA, buyerB, uplineB, topD, vaultOwner]) {
      await waitFor(await hrx.write.transfer([wallet.account.address, parseEther("100000")]));
      await waitFor(await usdt.write.approve([nft.address, parseEther("100000000")], { account: wallet.account }));
      await waitFor(await usdt.write.approve([market.address, parseEther("100000000")], { account: wallet.account }));
      await waitFor(await hrx.write.approve([nft.address, parseEther("100000000")], { account: wallet.account }));
      await waitFor(await hrx.write.approve([stake.address, parseEther("100000000")], { account: wallet.account }));
      await waitFor(await hrx.write.approve([swap.address, parseEther("100000000")], { account: wallet.account }));
    }

    await waitFor(await hrx.write.transfer([swap.address, parseEther("100000")]));
    await waitFor(await usdt.write.transfer([swap.address, parseEther("100000")]));
    await waitFor(await swap.write.mintInitial([0n]));

    await waitFor(await nft.write.setUserExtraOpenCount([sellerA.account.address, 20n]));

    const superRole = await nft.read.SUPER_ADMIN_ROLE();
    const adminRole = await nft.read.ADMIN_ROLE();
    const managerRole = await stake.read.MANAGER_ROLE();
    const migratorRole = await stake.read.MIGRATOR_ROLE();

    await waitFor(await nft.write.grantRole([superRole, superAdmin.account.address]));
    await waitFor(await nft.write.grantRole([adminRole, admin.account.address]));
    await waitFor(await stake.write.grantRole([managerRole, manager.account.address]));
    await waitFor(await stake.write.grantRole([migratorRole, manager.account.address]));

    return {
      deployer,
      superAdmin,
      admin,
      manager,
      project,
      sellerA,
      buyerB,
      uplineB,
      topD,
      vaultOwner,
      quantReceiver,
      usdt,
      piece,
      hrx,
      market,
      stake,
      nft,
      swap,
      factory,
    };
  }

  it("covers the core contract business flow end to end", async function () {
    const fx = await deployFixture();
    const {
      admin,
      buyerB,
      manager,
      nft,
      stake,
      superAdmin,
      project,
      quantReceiver,
      sellerA,
      topD,
      uplineB,
      usdt,
      hrx,
      vaultOwner,
      factory,
    } = fx;

    await expectRevert(
      stake.write.setProjectAutoSettlement([true], { account: manager.account }),
      "manager should not control project settlement mode",
    );
    await expectRevert(
      nft.write.transferHeldUsdtToQuant([parseEther("1"), "unauthorized"], { account: admin.account }),
      "ordinary admin must not use quant transfer hook",
    );

    await waitFor(await stake.write.register([zeroAddress], { account: topD.account }));
    assert.equal(await stake.read.isTopEligible([topD.account.address]), false);
    const topStatusBefore = await stake.read.getTopEligibilityStatus([topD.account.address]);
    assert.equal(topStatusBefore[0], false);
    assert.equal(topStatusBefore[1], false);
    assert.equal(topStatusBefore[2], 0n);
    assert.equal(topStatusBefore[3], parseEther("300"));
    assert.equal(topStatusBefore[4], false);

    await waitFor(await stake.write.setTopEnabled([topD.account.address, true], { account: manager.account }));
    assert.equal(await stake.read.topEnabled([topD.account.address]), true);
    assert.equal(await stake.read.isTopEligible([topD.account.address]), false);

    await waitFor(await stake.write.setTopManualEligible([topD.account.address, true], { account: manager.account }));
    assert.equal(await stake.read.isTopEligible([topD.account.address]), true);
    const topStatusAfterManual = await stake.read.getTopEligibilityStatus([topD.account.address]);
    assert.equal(topStatusAfterManual[1], true);
    assert.equal(topStatusAfterManual[4], true);
    await expectRevert(
      stake.write.setMinStakeA([0n]),
      "min stake should never be reset to zero",
    );

    await waitFor(await stake.write.register([topD.account.address], { account: uplineB.account }));
    await waitFor(await stake.write.register([uplineB.account.address], { account: sellerA.account }));

    await expectRevert(
      nft.write.mint([0], { account: sellerA.account }),
      "tier zero mint should fail",
    );
    await expectRevert(
      nft.write.mintBatch([1, 0n], { account: sellerA.account }),
      "zero quantity batch mint should fail",
    );

    const slotCostUsdt = await nft.read.quoteStorageSlotCost([buyerB.account.address, 2n]);
    assert.equal(slotCostUsdt, 0n);
    const buyerUsdtBeforeSlotBuy = await usdt.read.balanceOf([buyerB.account.address]);
    await waitFor(await nft.write.buyStorageSlots([2n], { account: buyerB.account }));
    assert.equal(await nft.read.extraStorageSlots([buyerB.account.address]), 0n);
    assert.equal(await usdt.read.balanceOf([buyerB.account.address]), buyerUsdtBeforeSlotBuy);

    const slotCostB = await nft.read.quoteSlotCostInB([buyerB.account.address, 1n]);
    assert.equal(slotCostB, 0n);
    const buyerHrxBeforeSlotBuy = await hrx.read.balanceOf([buyerB.account.address]);
    await waitFor(await nft.write.buyStorageSlotsPayHRX([1n, slotCostB], { account: buyerB.account }));
    assert.equal(await nft.read.extraStorageSlots([buyerB.account.address]), 0n);
    assert.equal(await hrx.read.balanceOf([buyerB.account.address]), buyerHrxBeforeSlotBuy);

    await waitFor(await nft.write.setOpenCountCostConfigB([parseEther("1"), parseEther("1"), true]));
    const openCountFixedCost = await nft.read.quoteOpenCountCostInB([buyerB.account.address, 2n]);
    assert.equal(openCountFixedCost, parseEther("3"));
    await expectRevert(
      nft.write.buyOpenCountPayHRX([2n, parseEther("2.999999999999999999")], { account: buyerB.account }),
      "maxCostB guard should reject insufficient cap",
    );
    await waitFor(await nft.write.buyOpenCountPayHRX([2n, openCountFixedCost], { account: buyerB.account }));
    assert.equal(await nft.read.extraOpenCount([buyerB.account.address]), 2n);

    await waitFor(await nft.write.setUserExtraOpenCount([buyerB.account.address, 20n]));

    const accountABeforeMint = await usdt.read.balanceOf([project.account.address]);
    const stakeBeforeMint = await usdt.read.balanceOf([stake.address]);
    await waitFor(await nft.write.mintBatch([1, 15n], { account: uplineB.account }));
    const accountAAfterMint = await usdt.read.balanceOf([project.account.address]);
    const stakeAfterMint = await usdt.read.balanceOf([stake.address]);
    assert.equal(accountAAfterMint - accountABeforeMint, parseEther("285"));
    assert.equal(stakeAfterMint - stakeBeforeMint, parseEther("15"));
    assert.equal(await nft.read.holdingValueOf([uplineB.account.address]), parseEther("300"));

    const sellerMintCountBefore = await nft.read.mintedTokenIdsLength();
    await waitFor(await nft.write.mintBatch([1, 7n], { account: sellerA.account }));
    const sellerMintCountAfter = await nft.read.mintedTokenIdsLength();
    assert.equal(sellerMintCountAfter - sellerMintCountBefore, 7n);

    const mintedInfos = await nft.read.getMintedByPage([15n, 7n]);
    const directSaleTokenId = mintedInfos[0].tokenId;
    const bFeeSaleTokenId = mintedInfos[1].tokenId;
    const selfOpenTokenIds = mintedInfos.slice(2).map((item: any) => item.tokenId);

    await increaseTime(ONE_MONTH_PLUS);

    await waitFor(await nft.write.list([directSaleTokenId], { account: sellerA.account }));
    await expectRevert(
      nft.write.unlist([directSaleTokenId], { account: buyerB.account }),
      "non seller should not unlist",
    );
    await expectRevert(
      nft.write.buyPayFeeInUSDT([directSaleTokenId], { account: sellerA.account }),
      "seller should not buy own listing",
    );

    const directSalePrice = await nft.read.secondaryBuyPrice([directSaleTokenId]);
    const directFeeUsdt = directSalePrice / 100n;
    const sellerUsdtBefore = await usdt.read.balanceOf([sellerA.account.address]);
    const uplineUsdtBefore = await usdt.read.balanceOf([uplineB.account.address]);
    const topUsdtBefore = await usdt.read.balanceOf([topD.account.address]);
    const projectUsdtBefore = await usdt.read.balanceOf([admin.account.address]);

    const directSaleReceipt = await waitFor(
      await nft.write.buyPayFeeInUSDT([directSaleTokenId], { account: buyerB.account }),
    );
    let directSaleSplitArgs: any;
    for (const log of directSaleReceipt.logs) {
      try {
        const decoded = decodeEventLog({
          abi: [secondarySaleSplitEventAbi],
          data: log.data,
          topics: log.topics,
        });
        if (decoded.eventName === "SecondarySaleSplit") {
          directSaleSplitArgs = decoded.args;
          break;
        }
      } catch {}
    }
    assert.ok(directSaleSplitArgs, "secondary sale split event should be emitted");

    assert.equal(
      (await nft.read.ownerOf([directSaleTokenId])).toLowerCase(),
      buyerB.account.address.toLowerCase(),
    );
    const minElapsed = BigInt(ONE_MONTH_PLUS);
    const splitDenominator = 10_000n * MONTH_SECONDS;
    const minSellerDelta = parseEther("20") + (parseEther("20") * 350n * minElapsed) / splitDenominator;
    const minUplineDelta = (parseEther("20") * 30n * minElapsed) / splitDenominator;
    const minTopDelta = (parseEther("20") * 20n * minElapsed) / splitDenominator;
    const sellerSaleDelta = (await usdt.read.balanceOf([sellerA.account.address])) - sellerUsdtBefore;
    const uplineSaleDelta = (await usdt.read.balanceOf([uplineB.account.address])) - uplineUsdtBefore;
    const topSaleDelta = (await usdt.read.balanceOf([topD.account.address])) - topUsdtBefore;
    assert.ok(sellerSaleDelta >= minSellerDelta);
    assert.ok(sellerSaleDelta - minSellerDelta < parseEther("0.001"));
    assert.ok(uplineSaleDelta >= minUplineDelta);
    assert.ok(uplineSaleDelta - minUplineDelta < parseEther("0.001"));
    assert.ok(topSaleDelta >= minTopDelta);
    assert.ok(topSaleDelta - minTopDelta < parseEther("0.001"));
    const projectSaleDelta = (await usdt.read.balanceOf([admin.account.address])) - projectUsdtBefore;
    const projectSaleExpectedFloor =
      directSalePrice - sellerSaleDelta - uplineSaleDelta - topSaleDelta + directFeeUsdt;
    assert.ok(projectSaleDelta >= projectSaleExpectedFloor);
    assert.ok(projectSaleDelta - projectSaleExpectedFloor < parseEther("0.001"));
    assert.equal(directSaleSplitArgs.b.toLowerCase(), uplineB.account.address.toLowerCase());
    assert.equal(directSaleSplitArgs.c.toLowerCase(), topD.account.address.toLowerCase());
    assert.equal(directSaleSplitArgs.toSeller, sellerSaleDelta);
    assert.equal(directSaleSplitArgs.toB, uplineSaleDelta);
    assert.equal(directSaleSplitArgs.toC, topSaleDelta);
    assert.equal(directSaleSplitArgs.toProject + directSaleSplitArgs.feeUSDT, projectSaleDelta);
    assert.equal(directSaleSplitArgs.toProject, 0n);
    assert.equal(directSaleSplitArgs.feeUSDT, directSaleSplitArgs.priceUSDT / 100n);
    assert.equal(directSaleSplitArgs.feeB, 0n);

    await waitFor(await nft.write.list([bFeeSaleTokenId], { account: sellerA.account }));
    const bModePrice = await nft.read.secondaryBuyPrice([bFeeSaleTokenId]);
    const bModeFeeUsdtEquivalent = (bModePrice * 50n) / 10000n;
    const quotedBModeFee = await nft.read.quoteFeeInB([bModeFeeUsdtEquivalent]);
    const accountDBefore = await hrx.read.balanceOf([topD.account.address]);
    await waitFor(
      await nft.write.buyPayFeeInB([bFeeSaleTokenId, quotedBModeFee + parseEther("1")], { account: buyerB.account }),
    );
    assert.equal(
      (await nft.read.ownerOf([bFeeSaleTokenId])).toLowerCase(),
      buyerB.account.address.toLowerCase(),
    );
    const topFeeDelta = (await hrx.read.balanceOf([topD.account.address])) - accountDBefore;
    assert.ok(topFeeDelta >= quotedBModeFee);
    assert.ok(topFeeDelta - quotedBModeFee < parseEther("0.001"));

    await expectRevert(
      nft.write.openBoxPayUsdt([selfOpenTokenIds[0]], { account: sellerA.account }),
      "self minted box must not open via USDT",
    );
    await waitFor(await nft.write.openBoxPayHRX([selfOpenTokenIds[0], parseEther("10")], { account: sellerA.account }));
    await waitFor(await nft.write.openBoxPayHRX([selfOpenTokenIds[1], parseEther("10")], { account: sellerA.account }));
    await waitFor(await nft.write.openBoxPayHRX([selfOpenTokenIds[2], parseEther("10")], { account: sellerA.account }));
    await waitFor(await nft.write.openBoxPayHRX([selfOpenTokenIds[3], parseEther("10")], { account: sellerA.account }));
    await waitFor(await nft.write.openBoxPayHRX([selfOpenTokenIds[4], parseEther("10")], { account: sellerA.account }));

    const buyerUsdtBeforePurchasedOpen = await usdt.read.balanceOf([buyerB.account.address]);
    const buyerHrxBeforePurchasedOpen = await hrx.read.balanceOf([buyerB.account.address]);
    const accountDHrxBeforePurchasedOpen = await hrx.read.balanceOf([topD.account.address]);
    await waitFor(await nft.write.openBoxPayUsdt([directSaleTokenId], { account: buyerB.account }));
    await waitFor(await nft.write.openBoxPayHRX([bFeeSaleTokenId, 0n], { account: buyerB.account }));
    assert.equal(await usdt.read.balanceOf([buyerB.account.address]), buyerUsdtBeforePurchasedOpen);
    assert.equal(await hrx.read.balanceOf([buyerB.account.address]), buyerHrxBeforePurchasedOpen);
    assert.equal(await hrx.read.balanceOf([topD.account.address]), accountDHrxBeforePurchasedOpen);

    assert.equal(await nft.read.open([selfOpenTokenIds[0]]), true);
    const openUsage = await nft.read.openUsageNow([buyerB.account.address]);
    assert.ok(openUsage[1] >= 2n);
    await expectRevert(
      nft.write.openBoxPayUsdt([directSaleTokenId], { account: buyerB.account }),
      "opened box should not open twice",
    );
    await expectRevert(
      nft.write.openBoxPayUsdt([bFeeSaleTokenId], { account: sellerA.account }),
      "non owner must not open",
    );

    await waitFor(await usdt.write.transfer([nft.address, parseEther("500")]));
    await waitFor(await usdt.write.transfer([stake.address, parseEther("500")]));

    const sellerWalletBeforeNegative = await usdt.read.balanceOf([sellerA.account.address]);
    await waitFor(await nft.write.setRewardOnce([selfOpenTokenIds[0], -parseEther("5")], { account: admin.account }));
    assert.equal((await usdt.read.balanceOf([sellerA.account.address])) - sellerWalletBeforeNegative, 0n);

    const sellerWalletBeforePartial = await usdt.read.balanceOf([sellerA.account.address]);
    await waitFor(await nft.write.setRewardOnce([selfOpenTokenIds[1], parseEther("10")], { account: admin.account }));
    assert.equal((await usdt.read.balanceOf([sellerA.account.address])) - sellerWalletBeforePartial, parseEther("10"));

    const sellerWalletBeforeFull = await usdt.read.balanceOf([sellerA.account.address]);
    await waitFor(await nft.write.setRewardOnce([selfOpenTokenIds[2], parseEther("20")], { account: admin.account }));
    assert.equal((await usdt.read.balanceOf([sellerA.account.address])) - sellerWalletBeforeFull, parseEther("20"));

    await waitFor(await nft.write.setRewardOnce([selfOpenTokenIds[3], parseEther("30")], { account: admin.account }));
    assert.equal(await stake.read.pendingUsdt([sellerA.account.address]), parseEther("9"));
    assert.equal(await stake.read.pendingUsdt([uplineB.account.address]), parseEther("0.2"));
    assert.equal(await stake.read.pendingUsdt([topD.account.address]), parseEther("0.4"));
    assert.equal(await stake.read.pendingUsdt([project.account.address]), parseEther("0.4"));

    const sellerClaimBalanceBefore = await usdt.read.balanceOf([sellerA.account.address]);
    await waitFor(await stake.write.claimAll([sellerA.account.address], { account: sellerA.account }));
    assert.equal((await usdt.read.balanceOf([sellerA.account.address])) - sellerClaimBalanceBefore, parseEther("9"));

    const projectPendingPayoutBefore = await usdt.read.balanceOf([project.account.address]);
    await waitFor(await stake.write.settleProjectPending());
    assert.equal(await stake.read.pendingUsdt([project.account.address]), 0n);
    assert.equal(
      (await usdt.read.balanceOf([project.account.address])) - projectPendingPayoutBefore,
      parseEther("0.4"),
    );

    await waitFor(await stake.write.setProjectAutoSettlement([true]));
    const projectAutoBefore = await usdt.read.balanceOf([project.account.address]);
    await waitFor(await nft.write.setRewardOnce([selfOpenTokenIds[4], parseEther("30")], { account: admin.account }));
    assert.equal(await stake.read.pendingUsdt([project.account.address]), 0n);
    assert.equal((await usdt.read.balanceOf([project.account.address])) - projectAutoBefore, parseEther("0.4"));

    await waitFor(await nft.write.setQuantReceiver([quantReceiver.account.address]));
    await waitFor(await usdt.write.transfer([nft.address, parseEther("5")]));
    await waitFor(await hrx.write.transfer([nft.address, parseEther("5")]));

    const quantUsdtBefore = await usdt.read.balanceOf([quantReceiver.account.address]);
    const quantHrxBefore = await hrx.read.balanceOf([quantReceiver.account.address]);
    await waitFor(
      await nft.write.transferHeldUsdtToQuant([parseEther("1"), "usdt-reserve"], { account: superAdmin.account }),
    );
    await waitFor(
      await nft.write.transferHeldERC20ToQuant([hrx.address, parseEther("2"), "hrx-reserve"], { account: superAdmin.account }),
    );
    assert.equal((await usdt.read.balanceOf([quantReceiver.account.address])) - quantUsdtBefore, parseEther("1"));
    assert.equal((await hrx.read.balanceOf([quantReceiver.account.address])) - quantHrxBefore, parseEther("2"));

    await waitFor(await factory.write.createVault({ account: vaultOwner.account } as any));
    const vaultAddress = await factory.read.vaultOf([vaultOwner.account.address]);
    assert.notEqual(vaultAddress, zeroAddress);

    await waitFor(await nft.write.mint([1], { account: vaultOwner.account }));
    const latestMintedCount = await nft.read.mintedTokenIdsLength();
    const latestMinted = await nft.read.getMintedByPage([latestMintedCount - 1n, 1n]);
    const vaultTokenId = latestMinted[0].tokenId;
    await waitFor(await nft.write.setApprovalForAll([vaultAddress, true], { account: vaultOwner.account }));

    const vault = await viem.getContractAt("BlindBoxUserVault", vaultAddress as `0x${string}`);
    await waitFor(await vault.write.batchDepositAndList([[vaultTokenId]], { account: vaultOwner.account }));
    assert.equal((await nft.read.ownerOf([vaultTokenId])).toLowerCase(), nft.address.toLowerCase());

    const vaultOwnerUsdtBefore = await usdt.read.balanceOf([vaultOwner.account.address]);
    const vaultUsdtBefore = await usdt.read.balanceOf([vaultAddress as `0x${string}`]);
    await waitFor(await nft.write.buyPayFeeInUSDT([vaultTokenId], { account: buyerB.account }));
    assert.equal((await nft.read.ownerOf([vaultTokenId])).toLowerCase(), buyerB.account.address.toLowerCase());
    assert.equal((await usdt.read.balanceOf([vaultAddress as `0x${string}`])) - vaultUsdtBefore, 0n);
    const vaultOwnerDelta = (await usdt.read.balanceOf([vaultOwner.account.address])) - vaultOwnerUsdtBefore;
    assert.ok(vaultOwnerDelta >= parseEther("20"));
    assert.ok(vaultOwnerDelta < parseEther("20.01"));

    await expectRevert(
      nft.write.transferHeldUsdtToQuant([0n, "zero"], { account: superAdmin.account }),
      "zero amount quant transfer should fail",
    );

    const defaultAdminOnNft = await nft.read.hasRole([DEFAULT_ADMIN_ROLE, fx.deployer.account.address]);
    assert.equal(defaultAdminOnNft, true);
  });

  it("imports legacy referral relationships until migration is closed", async function () {
    const fx = await deployFixture();
    const { buyerB, manager, quantReceiver, sellerA, stake, topD, uplineB } = fx;

    await waitFor(await stake.write.importReferral([topD.account.address, zeroAddress], { account: manager.account }));
    await waitFor(
      await stake.write.importReferralBatch(
        [
          [uplineB.account.address, sellerA.account.address],
          [topD.account.address, uplineB.account.address],
        ],
        { account: manager.account },
      ),
    );

    assert.equal(await stake.read.registered([topD.account.address]), true);
    assert.equal(await stake.read.registered([sellerA.account.address]), true);
    assert.equal((await stake.read.parentOf([sellerA.account.address])).toLowerCase(), uplineB.account.address.toLowerCase());
    assert.equal((await stake.read.topOf([sellerA.account.address])).toLowerCase(), topD.account.address.toLowerCase());
    assert.equal(await stake.read.childrenLength([uplineB.account.address]), 1n);

    await expectRevert(
      stake.write.importReferral([sellerA.account.address, topD.account.address], { account: manager.account }),
      "legacy import must not overwrite existing members",
    );

    await expectRevert(
      stake.write.importReferralBatch(
        [[buyerB.account.address], [topD.account.address, uplineB.account.address]],
        { account: manager.account },
      ),
      "legacy batch import must require matching input lengths",
    );

    await waitFor(await stake.write.closeMigration());
    assert.equal(await stake.read.migrationClosed(), true);

    await expectRevert(
      stake.write.importReferral([buyerB.account.address, topD.account.address], { account: manager.account }),
      "legacy import must stop after migration is closed",
    );

    await waitFor(await stake.write.register([topD.account.address], { account: quantReceiver.account }));
    assert.equal(await stake.read.registered([quantReceiver.account.address]), true);
    assert.equal((await stake.read.parentOf([quantReceiver.account.address])).toLowerCase(), topD.account.address.toLowerCase());
  });

  it("keeps storage unlimited and caps daily secondary buys at open allowance", async function () {
    const fx = await deployFixture();
    const { nft, sellerA, buyerB } = fx;

    await waitFor(await nft.write.setBaseDailyOpens([1n]));
    assert.equal(await nft.read.baseDailyOpensOf([buyerB.account.address]), 1n);
    await waitFor(await nft.write.setUserBaseDailyOpens([buyerB.account.address, 3n, true]));
    assert.equal(await nft.read.baseDailyOpensOf([buyerB.account.address]), 3n);

    await waitFor(await nft.write.mintBatch([1, 3n], { account: sellerA.account }));
    const mintedCount = await nft.read.mintedTokenIdsLength();
    const mintedInfos = await nft.read.getMintedByPage([mintedCount - 3n, 3n]);
    const tokenIds = mintedInfos.map((item: any) => item.tokenId);

    assert.equal(await nft.read.holdingValueOf([sellerA.account.address]), parseEther("60"));

    await waitFor(await nft.write.list([tokenIds[0]], { account: sellerA.account }));
    await waitFor(await nft.write.list([tokenIds[1]], { account: sellerA.account }));
    await waitFor(await nft.write.list([tokenIds[2]], { account: sellerA.account }));

    await waitFor(await nft.write.buyPayFeeInUSDT([tokenIds[0]], { account: buyerB.account }));
    await waitFor(await nft.write.buyPayFeeInUSDT([tokenIds[1]], { account: buyerB.account }));
    await waitFor(await nft.write.buyPayFeeInUSDT([tokenIds[2]], { account: buyerB.account }));

    const buyUsage = await nft.read.secondaryBuyUsageNow([buyerB.account.address]);
    assert.equal(buyUsage[1], 3n);
    assert.equal(buyUsage[2], 3n);
    assert.equal(buyUsage[3], 0n);

    await waitFor(await nft.write.setUserBaseDailyOpens([buyerB.account.address, 0n, false]));
    assert.equal(await nft.read.baseDailyOpensOf([buyerB.account.address]), 1n);

    await expectRevert(
      nft.write.buyPayFeeInUSDT([tokenIds[0]], { account: buyerB.account }),
      "daily secondary purchases must not exceed daily open allowance",
    );
  });
});
