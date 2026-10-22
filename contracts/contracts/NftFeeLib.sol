// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "@openzeppelin/contracts/utils/math/Math.sol";

interface IToken0Token1AmmQuote {
    function token0() external view returns (address);
    function token1() external view returns (address);
    function liquidityInited() external view returns (bool);
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 ts);
    function feeNum() external view returns (uint32);
    function feeDen() external view returns (uint32);
}

/// @notice 报价换算库：USDT 金额 -> TokenB 数量，减小 Nft 合约体积
library NftFeeLib {
    /// @notice 按 AMM 池子将 amount1InUSDT 换算为 token0(B) 数量（扣手续费后）
    function quoteFeeInB(
        IToken0Token1AmmQuote pool,
        uint256 amount1InUSDT,
        address expectedUsdt,
        address expectedTokenB
    ) external view returns (uint256 out0NetB) {
        require(address(pool) != address(0), "quote pool not set");
        require(amount1InUSDT > 0, "ZERO_IN");
        require(pool.liquidityInited(), "POOL_NOT_INIT");
        require(pool.token1() == expectedUsdt, "pool.token1!=USDT");
        require(pool.token0() == expectedTokenB, "pool.token0!=B");

        (uint112 r0, uint112 r1,) = pool.getReserves();
        require(r0 > 0 && r1 > 0, "NO_LIQ");
        uint256 grossMax = Math.mulDiv(uint256(r0), amount1InUSDT, uint256(r1) + amount1InUSDT);
        require(grossMax > 0, "INSUFFICIENT_IN");
        uint256 feeOut = Math.mulDiv(grossMax, uint256(pool.feeNum()), uint256(pool.feeDen()));
        return grossMax - feeOut;
    }
}
