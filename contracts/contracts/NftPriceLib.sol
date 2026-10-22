// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "@openzeppelin/contracts/utils/math/Math.sol";

/// @notice 盲盒价格计算库，减小 Nft 合约体积
library NftPriceLib {
    uint256 internal constant DECIMALS_18 = 1e18;
    uint256 internal constant MONTH_SECONDS = 30 days;

    /// @notice 线性增长价格增量（external 使库单独部署，减小主合约）
    function linearPriceIncrement(
        uint256 base,
        uint256 rateE18,
        uint256 elapsed
    ) external pure returns (uint256 inc) {
        uint256 numerator = rateE18 * elapsed;
        inc = Math.mulDiv(base, numerator, DECIMALS_18 * MONTH_SECONDS);
    }

}
