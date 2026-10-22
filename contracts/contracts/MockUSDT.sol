// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { ERC20Permit } from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";
import { AccessControl } from "@openzeppelin/contracts/access/AccessControl.sol";

/// @title MockUSDT
/// @notice 测试用 USDT 模拟合约，可 mint 供测试
/// @dev 与 BSC 主网 USDT 接口兼容（18 位小数），仅用于测试网
contract MockUSDT is ERC20, ERC20Permit, AccessControl {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    constructor(address initialRecipient_, uint256 initialSupply_)
        ERC20("Mock USDT", "USDT")
        ERC20Permit("Mock USDT")
    {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(MINTER_ROLE, msg.sender);
        if (initialRecipient_ != address(0) && initialSupply_ > 0) {
            _mint(initialRecipient_, initialSupply_);
        }
    }

    function decimals() public pure override returns (uint8) {
        return 18;
    }

    /// @notice 测试用：mint 代币给指定地址
    function mint(address to, uint256 amount) external onlyRole(MINTER_ROLE) {
        _mint(to, amount);
    }

    /// @notice 批量 mint 给多个地址
    function mintBatch(address[] calldata recipients, uint256[] calldata amounts) external onlyRole(MINTER_ROLE) {
        require(recipients.length == amounts.length, "LEN_MISMATCH");
        for (uint256 i = 0; i < recipients.length; i++) {
            _mint(recipients[i], amounts[i]);
        }
    }
}
