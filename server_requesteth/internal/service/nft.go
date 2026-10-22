// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package service

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BlindBoxNFTBoxInfo is an auto generated low-level Go binding around an user-defined struct.
type BlindBoxNFTBoxInfo struct {
	Tier        uint8
	UsdtPaid    *big.Int
	MintedAt    uint64
	OpenedAt    uint64
	Reward      *big.Int
	RewardSetAt uint64
	TokenId     *big.Int
}

// NftMetaData contains all meta data concerning the Nft contract.
var NftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burnAfterReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"burnAfterRewardBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"buyOpenSlots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeeB\",\"type\":\"uint256\"}],\"name\":\"buyPayFeeInB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buyPayFeeInUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"holdingAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"list\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"mintBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"openBox\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAccountA\",\"type\":\"address\"}],\"name\":\"setAccountA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAToken\",\"type\":\"address\"}],\"name\":\"setAToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rateTwo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseTwo\",\"type\":\"uint256\"}],\"name\":\"setFeeConfigB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"}],\"name\":\"setFeeConfigUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"setFeeQuotePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCost\",\"type\":\"uint256\"}],\"name\":\"setGlobalSlotStepCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeWei\",\"type\":\"uint256\"}],\"name\":\"setOpenFeeBNB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeWei\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"setOpenFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"setOpenFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"usdt_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accountA_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"globalSlotStepCost_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"staking_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newToken\",\"type\":\"address\"}],\"name\":\"ATokenChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAccountA\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAccountA\",\"type\":\"address\"}],\"name\":\"AccountAChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rateTwo\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseTwo\",\"type\":\"uint256\"}],\"name\":\"FeeConfigBChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"}],\"name\":\"FeeConfigUSDTChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0B\",\"type\":\"address\"}],\"name\":\"FeeQuotePoolChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCost\",\"type\":\"uint256\"}],\"name\":\"GlobalSlotStepCostChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Listed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdtPaid\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"openFeeBNB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"openFeeReceiver\",\"type\":\"address\"}],\"name\":\"OpenFeeConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bnbPaid\",\"type\":\"uint256\"}],\"name\":\"Opened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bpsDenominator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountAFeeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"monthlyRateDefaultE18\",\"type\":\"uint256\"}],\"name\":\"ParamsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueAddedUsdt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPieceHoldingValue\",\"type\":\"uint256\"}],\"name\":\"PieceHoldingAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPiece\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPiece\",\"type\":\"address\"}],\"name\":\"PieceTokenChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValueUsdtE18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValueUsdtE18\",\"type\":\"uint256\"}],\"name\":\"PieceValueChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceUSDT\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"feePaidInB\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeUSDT\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeB\",\"type\":\"uint256\"}],\"name\":\"Purchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"reward\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profitNotified\",\"type\":\"uint256\"}],\"name\":\"RewardSetOnce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"denom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aFeeBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"monthlyRateDefault\",\"type\":\"uint256\"}],\"name\":\"setParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPiece\",\"type\":\"address\"}],\"name\":\"setPieceToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValueUsdtE18\",\"type\":\"uint256\"}],\"name\":\"setPieceValueUsdtE18\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"reward_\",\"type\":\"int256\"}],\"name\":\"setRewardOnce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newStaking\",\"type\":\"address\"}],\"name\":\"setStakingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"setTierPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stepCostLocked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExtraSlots\",\"type\":\"uint256\"}],\"name\":\"SlotsPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldStaking\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newStaking\",\"type\":\"address\"}],\"name\":\"StakingContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TierPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"unlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Unlisted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountAFeeBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_DAILY_OPENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"boxInfo\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"usdtPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"mintedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"openedAt\",\"type\":\"uint64\"},{\"internalType\":\"int256\",\"name\":\"reward\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"rewardSetAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bpsDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAY_SHIFT_SECONDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DECIMALS_18\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"extraSlots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBaseTwo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeQuotePool\",\"outputs\":[{\"internalType\":\"contractIToken0Token1AmmQuote\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRateTwo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTokenB\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getListedByPage\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"usdtPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"mintedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"openedAt\",\"type\":\"uint64\"},{\"internalType\":\"int256\",\"name\":\"reward\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"rewardSetAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structBlindBoxNFT.BoxInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getMintedByPage\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"usdtPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"mintedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"openedAt\",\"type\":\"uint64\"},{\"internalType\":\"int256\",\"name\":\"reward\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"rewardSetAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structBlindBoxNFT.BoxInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getMyOpenedBoxesByPage\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"usdtPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"mintedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"openedAt\",\"type\":\"uint64\"},{\"internalType\":\"int256\",\"name\":\"reward\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"rewardSetAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structBlindBoxNFT.BoxInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getMyUnopenedBoxesByPage\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"usdtPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"mintedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"openedAt\",\"type\":\"uint64\"},{\"internalType\":\"int256\",\"name\":\"reward\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"rewardSetAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structBlindBoxNFT.BoxInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"GetOpenedByPage\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"usdtPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"mintedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"openedAt\",\"type\":\"uint64\"},{\"internalType\":\"int256\",\"name\":\"reward\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"rewardSetAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structBlindBoxNFT.BoxInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getOwnerTokensByPage\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"usdtPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"mintedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"openedAt\",\"type\":\"uint64\"},{\"internalType\":\"int256\",\"name\":\"reward\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"rewardSetAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structBlindBoxNFT.BoxInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getSellerSoldByPage\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"usdtPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"mintedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"openedAt\",\"type\":\"uint64\"},{\"internalType\":\"int256\",\"name\":\"reward\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"rewardSetAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structBlindBoxNFT.BoxInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalSlotStepCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"holdingBreakdownOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unopenedSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pieceSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"holdingValueOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listedTokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listedTokenIdsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"listedAt\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mintedTokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintedTokenIdsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MONTH_SECONDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"monthlyRateDefaultE18\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"myOpenedBoxesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"myUnopenedBoxesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"open\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openActionTokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OpenActionTokenIdsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openFeeBNB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"openUsageNow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dayIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"pieceHoldingOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pieceToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pieceValueUsdtE18\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"previewSlotCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"step\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount1InUSDT\",\"type\":\"uint256\"}],\"name\":\"quoteFeeInB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"out0NetB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"secondaryBuyPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"sellerSoldTokenIdsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"slotStepForUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContract\",\"outputs\":[{\"internalType\":\"contractIUmbrellaStakeDividend\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"tierPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenMonthlyRateE18\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userSlotStepCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NftABI is the input ABI used to generate the binding from.
// Deprecated: Use NftMetaData.ABI instead.
var NftABI = NftMetaData.ABI

// Nft is an auto generated Go binding around an Ethereum contract.
type Nft struct {
	NftCaller     // Read-only binding to the contract
	NftTransactor // Write-only binding to the contract
	NftFilterer   // Log filterer for contract events
}

// NftCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftSession struct {
	Contract     *Nft              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftCallerSession struct {
	Contract *NftCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftTransactorSession struct {
	Contract     *NftTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftRaw struct {
	Contract *Nft // Generic contract binding to access the raw methods on
}

// NftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftCallerRaw struct {
	Contract *NftCaller // Generic read-only contract binding to access the raw methods on
}

// NftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftTransactorRaw struct {
	Contract *NftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNft creates a new instance of Nft, bound to a specific deployed contract.
func NewNft(address common.Address, backend bind.ContractBackend) (*Nft, error) {
	contract, err := bindNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nft{NftCaller: NftCaller{contract: contract}, NftTransactor: NftTransactor{contract: contract}, NftFilterer: NftFilterer{contract: contract}}, nil
}

// NewNftCaller creates a new read-only instance of Nft, bound to a specific deployed contract.
func NewNftCaller(address common.Address, caller bind.ContractCaller) (*NftCaller, error) {
	contract, err := bindNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftCaller{contract: contract}, nil
}

// NewNftTransactor creates a new write-only instance of Nft, bound to a specific deployed contract.
func NewNftTransactor(address common.Address, transactor bind.ContractTransactor) (*NftTransactor, error) {
	contract, err := bindNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftTransactor{contract: contract}, nil
}

// NewNftFilterer creates a new log filterer instance of Nft, bound to a specific deployed contract.
func NewNftFilterer(address common.Address, filterer bind.ContractFilterer) (*NftFilterer, error) {
	contract, err := bindNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftFilterer{contract: contract}, nil
}

// bindNft binds a generic wrapper to an already deployed contract.
func bindNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nft *NftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nft.Contract.NftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nft *NftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nft.Contract.NftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nft *NftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nft.Contract.NftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nft *NftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nft *NftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nft *NftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nft.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Nft *NftCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Nft *NftSession) ADMINROLE() ([32]byte, error) {
	return _Nft.Contract.ADMINROLE(&_Nft.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Nft *NftCallerSession) ADMINROLE() ([32]byte, error) {
	return _Nft.Contract.ADMINROLE(&_Nft.CallOpts)
}

// BASEDAILYOPENS is a free data retrieval call binding the contract method 0x1c84c602.
//
// Solidity: function BASE_DAILY_OPENS() view returns(uint256)
func (_Nft *NftCaller) BASEDAILYOPENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "BASE_DAILY_OPENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEDAILYOPENS is a free data retrieval call binding the contract method 0x1c84c602.
//
// Solidity: function BASE_DAILY_OPENS() view returns(uint256)
func (_Nft *NftSession) BASEDAILYOPENS() (*big.Int, error) {
	return _Nft.Contract.BASEDAILYOPENS(&_Nft.CallOpts)
}

// BASEDAILYOPENS is a free data retrieval call binding the contract method 0x1c84c602.
//
// Solidity: function BASE_DAILY_OPENS() view returns(uint256)
func (_Nft *NftCallerSession) BASEDAILYOPENS() (*big.Int, error) {
	return _Nft.Contract.BASEDAILYOPENS(&_Nft.CallOpts)
}

// DAYSHIFTSECONDS is a free data retrieval call binding the contract method 0x2d184b2b.
//
// Solidity: function DAY_SHIFT_SECONDS() view returns(uint256)
func (_Nft *NftCaller) DAYSHIFTSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "DAY_SHIFT_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DAYSHIFTSECONDS is a free data retrieval call binding the contract method 0x2d184b2b.
//
// Solidity: function DAY_SHIFT_SECONDS() view returns(uint256)
func (_Nft *NftSession) DAYSHIFTSECONDS() (*big.Int, error) {
	return _Nft.Contract.DAYSHIFTSECONDS(&_Nft.CallOpts)
}

// DAYSHIFTSECONDS is a free data retrieval call binding the contract method 0x2d184b2b.
//
// Solidity: function DAY_SHIFT_SECONDS() view returns(uint256)
func (_Nft *NftCallerSession) DAYSHIFTSECONDS() (*big.Int, error) {
	return _Nft.Contract.DAYSHIFTSECONDS(&_Nft.CallOpts)
}

// DECIMALS18 is a free data retrieval call binding the contract method 0x16ec8257.
//
// Solidity: function DECIMALS_18() view returns(uint256)
func (_Nft *NftCaller) DECIMALS18(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "DECIMALS_18")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DECIMALS18 is a free data retrieval call binding the contract method 0x16ec8257.
//
// Solidity: function DECIMALS_18() view returns(uint256)
func (_Nft *NftSession) DECIMALS18() (*big.Int, error) {
	return _Nft.Contract.DECIMALS18(&_Nft.CallOpts)
}

// DECIMALS18 is a free data retrieval call binding the contract method 0x16ec8257.
//
// Solidity: function DECIMALS_18() view returns(uint256)
func (_Nft *NftCallerSession) DECIMALS18() (*big.Int, error) {
	return _Nft.Contract.DECIMALS18(&_Nft.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Nft *NftCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Nft *NftSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Nft.Contract.DEFAULTADMINROLE(&_Nft.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Nft *NftCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Nft.Contract.DEFAULTADMINROLE(&_Nft.CallOpts)
}

// GetOpenedByPage is a free data retrieval call binding the contract method 0x4cdcb01c.
//
// Solidity: function GetOpenedByPage(uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftCaller) GetOpenedByPage(opts *bind.CallOpts, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "GetOpenedByPage", start, count)

	if err != nil {
		return *new([]BlindBoxNFTBoxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]BlindBoxNFTBoxInfo)).(*[]BlindBoxNFTBoxInfo)

	return out0, err

}

// GetOpenedByPage is a free data retrieval call binding the contract method 0x4cdcb01c.
//
// Solidity: function GetOpenedByPage(uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftSession) GetOpenedByPage(start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	return _Nft.Contract.GetOpenedByPage(&_Nft.CallOpts, start, count)
}

// GetOpenedByPage is a free data retrieval call binding the contract method 0x4cdcb01c.
//
// Solidity: function GetOpenedByPage(uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftCallerSession) GetOpenedByPage(start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	return _Nft.Contract.GetOpenedByPage(&_Nft.CallOpts, start, count)
}

// MONTHSECONDS is a free data retrieval call binding the contract method 0xf4a814fd.
//
// Solidity: function MONTH_SECONDS() view returns(uint256)
func (_Nft *NftCaller) MONTHSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "MONTH_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MONTHSECONDS is a free data retrieval call binding the contract method 0xf4a814fd.
//
// Solidity: function MONTH_SECONDS() view returns(uint256)
func (_Nft *NftSession) MONTHSECONDS() (*big.Int, error) {
	return _Nft.Contract.MONTHSECONDS(&_Nft.CallOpts)
}

// MONTHSECONDS is a free data retrieval call binding the contract method 0xf4a814fd.
//
// Solidity: function MONTH_SECONDS() view returns(uint256)
func (_Nft *NftCallerSession) MONTHSECONDS() (*big.Int, error) {
	return _Nft.Contract.MONTHSECONDS(&_Nft.CallOpts)
}

// OpenActionTokenIdsLength is a free data retrieval call binding the contract method 0x099f6192.
//
// Solidity: function OpenActionTokenIdsLength() view returns(uint256)
func (_Nft *NftCaller) OpenActionTokenIdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "OpenActionTokenIdsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenActionTokenIdsLength is a free data retrieval call binding the contract method 0x099f6192.
//
// Solidity: function OpenActionTokenIdsLength() view returns(uint256)
func (_Nft *NftSession) OpenActionTokenIdsLength() (*big.Int, error) {
	return _Nft.Contract.OpenActionTokenIdsLength(&_Nft.CallOpts)
}

// OpenActionTokenIdsLength is a free data retrieval call binding the contract method 0x099f6192.
//
// Solidity: function OpenActionTokenIdsLength() view returns(uint256)
func (_Nft *NftCallerSession) OpenActionTokenIdsLength() (*big.Int, error) {
	return _Nft.Contract.OpenActionTokenIdsLength(&_Nft.CallOpts)
}

// AToken is a free data retrieval call binding the contract method 0xa0c1f15e.
//
// Solidity: function aToken() view returns(address)
func (_Nft *NftCaller) AToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "aToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AToken is a free data retrieval call binding the contract method 0xa0c1f15e.
//
// Solidity: function aToken() view returns(address)
func (_Nft *NftSession) AToken() (common.Address, error) {
	return _Nft.Contract.AToken(&_Nft.CallOpts)
}

// AToken is a free data retrieval call binding the contract method 0xa0c1f15e.
//
// Solidity: function aToken() view returns(address)
func (_Nft *NftCallerSession) AToken() (common.Address, error) {
	return _Nft.Contract.AToken(&_Nft.CallOpts)
}

// AccountA is a free data retrieval call binding the contract method 0x06867f7d.
//
// Solidity: function accountA() view returns(address)
func (_Nft *NftCaller) AccountA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "accountA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountA is a free data retrieval call binding the contract method 0x06867f7d.
//
// Solidity: function accountA() view returns(address)
func (_Nft *NftSession) AccountA() (common.Address, error) {
	return _Nft.Contract.AccountA(&_Nft.CallOpts)
}

// AccountA is a free data retrieval call binding the contract method 0x06867f7d.
//
// Solidity: function accountA() view returns(address)
func (_Nft *NftCallerSession) AccountA() (common.Address, error) {
	return _Nft.Contract.AccountA(&_Nft.CallOpts)
}

// AccountAFeeBps is a free data retrieval call binding the contract method 0x06bb1d2d.
//
// Solidity: function accountAFeeBps() view returns(uint256)
func (_Nft *NftCaller) AccountAFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "accountAFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountAFeeBps is a free data retrieval call binding the contract method 0x06bb1d2d.
//
// Solidity: function accountAFeeBps() view returns(uint256)
func (_Nft *NftSession) AccountAFeeBps() (*big.Int, error) {
	return _Nft.Contract.AccountAFeeBps(&_Nft.CallOpts)
}

// AccountAFeeBps is a free data retrieval call binding the contract method 0x06bb1d2d.
//
// Solidity: function accountAFeeBps() view returns(uint256)
func (_Nft *NftCallerSession) AccountAFeeBps() (*big.Int, error) {
	return _Nft.Contract.AccountAFeeBps(&_Nft.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nft *NftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nft *NftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Nft.Contract.BalanceOf(&_Nft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nft *NftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Nft.Contract.BalanceOf(&_Nft.CallOpts, owner)
}

// BoxInfo is a free data retrieval call binding the contract method 0x260d68fa.
//
// Solidity: function boxInfo(uint256 ) view returns(uint8 tier, uint256 usdtPaid, uint64 mintedAt, uint64 openedAt, int256 reward, uint64 rewardSetAt, uint256 tokenId)
func (_Nft *NftCaller) BoxInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Tier        uint8
	UsdtPaid    *big.Int
	MintedAt    uint64
	OpenedAt    uint64
	Reward      *big.Int
	RewardSetAt uint64
	TokenId     *big.Int
}, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "boxInfo", arg0)

	outstruct := new(struct {
		Tier        uint8
		UsdtPaid    *big.Int
		MintedAt    uint64
		OpenedAt    uint64
		Reward      *big.Int
		RewardSetAt uint64
		TokenId     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tier = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.UsdtPaid = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MintedAt = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.OpenedAt = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.Reward = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RewardSetAt = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.TokenId = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BoxInfo is a free data retrieval call binding the contract method 0x260d68fa.
//
// Solidity: function boxInfo(uint256 ) view returns(uint8 tier, uint256 usdtPaid, uint64 mintedAt, uint64 openedAt, int256 reward, uint64 rewardSetAt, uint256 tokenId)
func (_Nft *NftSession) BoxInfo(arg0 *big.Int) (struct {
	Tier        uint8
	UsdtPaid    *big.Int
	MintedAt    uint64
	OpenedAt    uint64
	Reward      *big.Int
	RewardSetAt uint64
	TokenId     *big.Int
}, error) {
	return _Nft.Contract.BoxInfo(&_Nft.CallOpts, arg0)
}

// BoxInfo is a free data retrieval call binding the contract method 0x260d68fa.
//
// Solidity: function boxInfo(uint256 ) view returns(uint8 tier, uint256 usdtPaid, uint64 mintedAt, uint64 openedAt, int256 reward, uint64 rewardSetAt, uint256 tokenId)
func (_Nft *NftCallerSession) BoxInfo(arg0 *big.Int) (struct {
	Tier        uint8
	UsdtPaid    *big.Int
	MintedAt    uint64
	OpenedAt    uint64
	Reward      *big.Int
	RewardSetAt uint64
	TokenId     *big.Int
}, error) {
	return _Nft.Contract.BoxInfo(&_Nft.CallOpts, arg0)
}

// BpsDenominator is a free data retrieval call binding the contract method 0xf60632a0.
//
// Solidity: function bpsDenominator() view returns(uint256)
func (_Nft *NftCaller) BpsDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "bpsDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BpsDenominator is a free data retrieval call binding the contract method 0xf60632a0.
//
// Solidity: function bpsDenominator() view returns(uint256)
func (_Nft *NftSession) BpsDenominator() (*big.Int, error) {
	return _Nft.Contract.BpsDenominator(&_Nft.CallOpts)
}

// BpsDenominator is a free data retrieval call binding the contract method 0xf60632a0.
//
// Solidity: function bpsDenominator() view returns(uint256)
func (_Nft *NftCallerSession) BpsDenominator() (*big.Int, error) {
	return _Nft.Contract.BpsDenominator(&_Nft.CallOpts)
}

// ExtraSlots is a free data retrieval call binding the contract method 0x1bee6dd5.
//
// Solidity: function extraSlots(address ) view returns(uint256)
func (_Nft *NftCaller) ExtraSlots(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "extraSlots", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraSlots is a free data retrieval call binding the contract method 0x1bee6dd5.
//
// Solidity: function extraSlots(address ) view returns(uint256)
func (_Nft *NftSession) ExtraSlots(arg0 common.Address) (*big.Int, error) {
	return _Nft.Contract.ExtraSlots(&_Nft.CallOpts, arg0)
}

// ExtraSlots is a free data retrieval call binding the contract method 0x1bee6dd5.
//
// Solidity: function extraSlots(address ) view returns(uint256)
func (_Nft *NftCallerSession) ExtraSlots(arg0 common.Address) (*big.Int, error) {
	return _Nft.Contract.ExtraSlots(&_Nft.CallOpts, arg0)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_Nft *NftCaller) FeeBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "feeBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_Nft *NftSession) FeeBase() (*big.Int, error) {
	return _Nft.Contract.FeeBase(&_Nft.CallOpts)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_Nft *NftCallerSession) FeeBase() (*big.Int, error) {
	return _Nft.Contract.FeeBase(&_Nft.CallOpts)
}

// FeeBaseTwo is a free data retrieval call binding the contract method 0xf2464a58.
//
// Solidity: function feeBaseTwo() view returns(uint256)
func (_Nft *NftCaller) FeeBaseTwo(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "feeBaseTwo")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBaseTwo is a free data retrieval call binding the contract method 0xf2464a58.
//
// Solidity: function feeBaseTwo() view returns(uint256)
func (_Nft *NftSession) FeeBaseTwo() (*big.Int, error) {
	return _Nft.Contract.FeeBaseTwo(&_Nft.CallOpts)
}

// FeeBaseTwo is a free data retrieval call binding the contract method 0xf2464a58.
//
// Solidity: function feeBaseTwo() view returns(uint256)
func (_Nft *NftCallerSession) FeeBaseTwo() (*big.Int, error) {
	return _Nft.Contract.FeeBaseTwo(&_Nft.CallOpts)
}

// FeeQuotePool is a free data retrieval call binding the contract method 0x434cee9c.
//
// Solidity: function feeQuotePool() view returns(address)
func (_Nft *NftCaller) FeeQuotePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "feeQuotePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeQuotePool is a free data retrieval call binding the contract method 0x434cee9c.
//
// Solidity: function feeQuotePool() view returns(address)
func (_Nft *NftSession) FeeQuotePool() (common.Address, error) {
	return _Nft.Contract.FeeQuotePool(&_Nft.CallOpts)
}

// FeeQuotePool is a free data retrieval call binding the contract method 0x434cee9c.
//
// Solidity: function feeQuotePool() view returns(address)
func (_Nft *NftCallerSession) FeeQuotePool() (common.Address, error) {
	return _Nft.Contract.FeeQuotePool(&_Nft.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Nft *NftCaller) FeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Nft *NftSession) FeeRate() (*big.Int, error) {
	return _Nft.Contract.FeeRate(&_Nft.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Nft *NftCallerSession) FeeRate() (*big.Int, error) {
	return _Nft.Contract.FeeRate(&_Nft.CallOpts)
}

// FeeRateTwo is a free data retrieval call binding the contract method 0x64a10261.
//
// Solidity: function feeRateTwo() view returns(uint256)
func (_Nft *NftCaller) FeeRateTwo(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "feeRateTwo")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRateTwo is a free data retrieval call binding the contract method 0x64a10261.
//
// Solidity: function feeRateTwo() view returns(uint256)
func (_Nft *NftSession) FeeRateTwo() (*big.Int, error) {
	return _Nft.Contract.FeeRateTwo(&_Nft.CallOpts)
}

// FeeRateTwo is a free data retrieval call binding the contract method 0x64a10261.
//
// Solidity: function feeRateTwo() view returns(uint256)
func (_Nft *NftCallerSession) FeeRateTwo() (*big.Int, error) {
	return _Nft.Contract.FeeRateTwo(&_Nft.CallOpts)
}

// FeeTokenB is a free data retrieval call binding the contract method 0xe18bbce2.
//
// Solidity: function feeTokenB() view returns(address)
func (_Nft *NftCaller) FeeTokenB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "feeTokenB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenB is a free data retrieval call binding the contract method 0xe18bbce2.
//
// Solidity: function feeTokenB() view returns(address)
func (_Nft *NftSession) FeeTokenB() (common.Address, error) {
	return _Nft.Contract.FeeTokenB(&_Nft.CallOpts)
}

// FeeTokenB is a free data retrieval call binding the contract method 0xe18bbce2.
//
// Solidity: function feeTokenB() view returns(address)
func (_Nft *NftCallerSession) FeeTokenB() (common.Address, error) {
	return _Nft.Contract.FeeTokenB(&_Nft.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Nft *NftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Nft *NftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Nft.Contract.GetApproved(&_Nft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Nft *NftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Nft.Contract.GetApproved(&_Nft.CallOpts, tokenId)
}

// GetListedByPage is a free data retrieval call binding the contract method 0x2d331eda.
//
// Solidity: function getListedByPage(uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftCaller) GetListedByPage(opts *bind.CallOpts, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "getListedByPage", start, count)

	if err != nil {
		return *new([]BlindBoxNFTBoxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]BlindBoxNFTBoxInfo)).(*[]BlindBoxNFTBoxInfo)

	return out0, err

}

// GetListedByPage is a free data retrieval call binding the contract method 0x2d331eda.
//
// Solidity: function getListedByPage(uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftSession) GetListedByPage(start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	return _Nft.Contract.GetListedByPage(&_Nft.CallOpts, start, count)
}

// GetListedByPage is a free data retrieval call binding the contract method 0x2d331eda.
//
// Solidity: function getListedByPage(uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftCallerSession) GetListedByPage(start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	return _Nft.Contract.GetListedByPage(&_Nft.CallOpts, start, count)
}

// GetMintedByPage is a free data retrieval call binding the contract method 0x4b93d9e6.
//
// Solidity: function getMintedByPage(uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftCaller) GetMintedByPage(opts *bind.CallOpts, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "getMintedByPage", start, count)

	if err != nil {
		return *new([]BlindBoxNFTBoxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]BlindBoxNFTBoxInfo)).(*[]BlindBoxNFTBoxInfo)

	return out0, err

}

// GetMintedByPage is a free data retrieval call binding the contract method 0x4b93d9e6.
//
// Solidity: function getMintedByPage(uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftSession) GetMintedByPage(start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	return _Nft.Contract.GetMintedByPage(&_Nft.CallOpts, start, count)
}

// GetMintedByPage is a free data retrieval call binding the contract method 0x4b93d9e6.
//
// Solidity: function getMintedByPage(uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftCallerSession) GetMintedByPage(start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	return _Nft.Contract.GetMintedByPage(&_Nft.CallOpts, start, count)
}

// GetMyOpenedBoxesByPage is a free data retrieval call binding the contract method 0xa8bffdc2.
//
// Solidity: function getMyOpenedBoxesByPage(address user, uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftCaller) GetMyOpenedBoxesByPage(opts *bind.CallOpts, user common.Address, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "getMyOpenedBoxesByPage", user, start, count)

	if err != nil {
		return *new([]BlindBoxNFTBoxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]BlindBoxNFTBoxInfo)).(*[]BlindBoxNFTBoxInfo)

	return out0, err

}

// GetMyOpenedBoxesByPage is a free data retrieval call binding the contract method 0xa8bffdc2.
//
// Solidity: function getMyOpenedBoxesByPage(address user, uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftSession) GetMyOpenedBoxesByPage(user common.Address, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	return _Nft.Contract.GetMyOpenedBoxesByPage(&_Nft.CallOpts, user, start, count)
}

// GetMyOpenedBoxesByPage is a free data retrieval call binding the contract method 0xa8bffdc2.
//
// Solidity: function getMyOpenedBoxesByPage(address user, uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftCallerSession) GetMyOpenedBoxesByPage(user common.Address, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	return _Nft.Contract.GetMyOpenedBoxesByPage(&_Nft.CallOpts, user, start, count)
}

// GetMyUnopenedBoxesByPage is a free data retrieval call binding the contract method 0xd3ac2471.
//
// Solidity: function getMyUnopenedBoxesByPage(address user, uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftCaller) GetMyUnopenedBoxesByPage(opts *bind.CallOpts, user common.Address, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "getMyUnopenedBoxesByPage", user, start, count)

	if err != nil {
		return *new([]BlindBoxNFTBoxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]BlindBoxNFTBoxInfo)).(*[]BlindBoxNFTBoxInfo)

	return out0, err

}

// GetMyUnopenedBoxesByPage is a free data retrieval call binding the contract method 0xd3ac2471.
//
// Solidity: function getMyUnopenedBoxesByPage(address user, uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftSession) GetMyUnopenedBoxesByPage(user common.Address, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	return _Nft.Contract.GetMyUnopenedBoxesByPage(&_Nft.CallOpts, user, start, count)
}

// GetMyUnopenedBoxesByPage is a free data retrieval call binding the contract method 0xd3ac2471.
//
// Solidity: function getMyUnopenedBoxesByPage(address user, uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftCallerSession) GetMyUnopenedBoxesByPage(user common.Address, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	return _Nft.Contract.GetMyUnopenedBoxesByPage(&_Nft.CallOpts, user, start, count)
}

// GetOwnerTokensByPage is a free data retrieval call binding the contract method 0x185b5de0.
//
// Solidity: function getOwnerTokensByPage(address owner, uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftCaller) GetOwnerTokensByPage(opts *bind.CallOpts, owner common.Address, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "getOwnerTokensByPage", owner, start, count)

	if err != nil {
		return *new([]BlindBoxNFTBoxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]BlindBoxNFTBoxInfo)).(*[]BlindBoxNFTBoxInfo)

	return out0, err

}

// GetOwnerTokensByPage is a free data retrieval call binding the contract method 0x185b5de0.
//
// Solidity: function getOwnerTokensByPage(address owner, uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftSession) GetOwnerTokensByPage(owner common.Address, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	return _Nft.Contract.GetOwnerTokensByPage(&_Nft.CallOpts, owner, start, count)
}

// GetOwnerTokensByPage is a free data retrieval call binding the contract method 0x185b5de0.
//
// Solidity: function getOwnerTokensByPage(address owner, uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftCallerSession) GetOwnerTokensByPage(owner common.Address, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	return _Nft.Contract.GetOwnerTokensByPage(&_Nft.CallOpts, owner, start, count)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Nft *NftCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Nft *NftSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Nft.Contract.GetRoleAdmin(&_Nft.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Nft *NftCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Nft.Contract.GetRoleAdmin(&_Nft.CallOpts, role)
}

// GetSellerSoldByPage is a free data retrieval call binding the contract method 0x029ae463.
//
// Solidity: function getSellerSoldByPage(address seller, uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftCaller) GetSellerSoldByPage(opts *bind.CallOpts, seller common.Address, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "getSellerSoldByPage", seller, start, count)

	if err != nil {
		return *new([]BlindBoxNFTBoxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]BlindBoxNFTBoxInfo)).(*[]BlindBoxNFTBoxInfo)

	return out0, err

}

// GetSellerSoldByPage is a free data retrieval call binding the contract method 0x029ae463.
//
// Solidity: function getSellerSoldByPage(address seller, uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftSession) GetSellerSoldByPage(seller common.Address, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	return _Nft.Contract.GetSellerSoldByPage(&_Nft.CallOpts, seller, start, count)
}

// GetSellerSoldByPage is a free data retrieval call binding the contract method 0x029ae463.
//
// Solidity: function getSellerSoldByPage(address seller, uint256 start, uint256 count) view returns((uint8,uint256,uint64,uint64,int256,uint64,uint256)[] infos)
func (_Nft *NftCallerSession) GetSellerSoldByPage(seller common.Address, start *big.Int, count *big.Int) ([]BlindBoxNFTBoxInfo, error) {
	return _Nft.Contract.GetSellerSoldByPage(&_Nft.CallOpts, seller, start, count)
}

// GlobalSlotStepCost is a free data retrieval call binding the contract method 0xb49ff99c.
//
// Solidity: function globalSlotStepCost() view returns(uint256)
func (_Nft *NftCaller) GlobalSlotStepCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "globalSlotStepCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalSlotStepCost is a free data retrieval call binding the contract method 0xb49ff99c.
//
// Solidity: function globalSlotStepCost() view returns(uint256)
func (_Nft *NftSession) GlobalSlotStepCost() (*big.Int, error) {
	return _Nft.Contract.GlobalSlotStepCost(&_Nft.CallOpts)
}

// GlobalSlotStepCost is a free data retrieval call binding the contract method 0xb49ff99c.
//
// Solidity: function globalSlotStepCost() view returns(uint256)
func (_Nft *NftCallerSession) GlobalSlotStepCost() (*big.Int, error) {
	return _Nft.Contract.GlobalSlotStepCost(&_Nft.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Nft *NftCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Nft *NftSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Nft.Contract.HasRole(&_Nft.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Nft *NftCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Nft.Contract.HasRole(&_Nft.CallOpts, role, account)
}

// HoldingBreakdownOf is a free data retrieval call binding the contract method 0x27d74167.
//
// Solidity: function holdingBreakdownOf(address user) view returns(uint256 unopenedSum, uint256 pieceSum, uint256 total)
func (_Nft *NftCaller) HoldingBreakdownOf(opts *bind.CallOpts, user common.Address) (struct {
	UnopenedSum *big.Int
	PieceSum    *big.Int
	Total       *big.Int
}, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "holdingBreakdownOf", user)

	outstruct := new(struct {
		UnopenedSum *big.Int
		PieceSum    *big.Int
		Total       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UnopenedSum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PieceSum = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Total = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// HoldingBreakdownOf is a free data retrieval call binding the contract method 0x27d74167.
//
// Solidity: function holdingBreakdownOf(address user) view returns(uint256 unopenedSum, uint256 pieceSum, uint256 total)
func (_Nft *NftSession) HoldingBreakdownOf(user common.Address) (struct {
	UnopenedSum *big.Int
	PieceSum    *big.Int
	Total       *big.Int
}, error) {
	return _Nft.Contract.HoldingBreakdownOf(&_Nft.CallOpts, user)
}

// HoldingBreakdownOf is a free data retrieval call binding the contract method 0x27d74167.
//
// Solidity: function holdingBreakdownOf(address user) view returns(uint256 unopenedSum, uint256 pieceSum, uint256 total)
func (_Nft *NftCallerSession) HoldingBreakdownOf(user common.Address) (struct {
	UnopenedSum *big.Int
	PieceSum    *big.Int
	Total       *big.Int
}, error) {
	return _Nft.Contract.HoldingBreakdownOf(&_Nft.CallOpts, user)
}

// HoldingValueOf is a free data retrieval call binding the contract method 0xe579aca3.
//
// Solidity: function holdingValueOf(address user) view returns(uint256)
func (_Nft *NftCaller) HoldingValueOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "holdingValueOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HoldingValueOf is a free data retrieval call binding the contract method 0xe579aca3.
//
// Solidity: function holdingValueOf(address user) view returns(uint256)
func (_Nft *NftSession) HoldingValueOf(user common.Address) (*big.Int, error) {
	return _Nft.Contract.HoldingValueOf(&_Nft.CallOpts, user)
}

// HoldingValueOf is a free data retrieval call binding the contract method 0xe579aca3.
//
// Solidity: function holdingValueOf(address user) view returns(uint256)
func (_Nft *NftCallerSession) HoldingValueOf(user common.Address) (*big.Int, error) {
	return _Nft.Contract.HoldingValueOf(&_Nft.CallOpts, user)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Nft *NftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Nft *NftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Nft.Contract.IsApprovedForAll(&_Nft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Nft *NftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Nft.Contract.IsApprovedForAll(&_Nft.CallOpts, owner, operator)
}

// ListedTokenIds is a free data retrieval call binding the contract method 0xb6feadc8.
//
// Solidity: function listedTokenIds(uint256 ) view returns(uint256)
func (_Nft *NftCaller) ListedTokenIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "listedTokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ListedTokenIds is a free data retrieval call binding the contract method 0xb6feadc8.
//
// Solidity: function listedTokenIds(uint256 ) view returns(uint256)
func (_Nft *NftSession) ListedTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Nft.Contract.ListedTokenIds(&_Nft.CallOpts, arg0)
}

// ListedTokenIds is a free data retrieval call binding the contract method 0xb6feadc8.
//
// Solidity: function listedTokenIds(uint256 ) view returns(uint256)
func (_Nft *NftCallerSession) ListedTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Nft.Contract.ListedTokenIds(&_Nft.CallOpts, arg0)
}

// ListedTokenIdsLength is a free data retrieval call binding the contract method 0xd949ed57.
//
// Solidity: function listedTokenIdsLength() view returns(uint256)
func (_Nft *NftCaller) ListedTokenIdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "listedTokenIdsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ListedTokenIdsLength is a free data retrieval call binding the contract method 0xd949ed57.
//
// Solidity: function listedTokenIdsLength() view returns(uint256)
func (_Nft *NftSession) ListedTokenIdsLength() (*big.Int, error) {
	return _Nft.Contract.ListedTokenIdsLength(&_Nft.CallOpts)
}

// ListedTokenIdsLength is a free data retrieval call binding the contract method 0xd949ed57.
//
// Solidity: function listedTokenIdsLength() view returns(uint256)
func (_Nft *NftCallerSession) ListedTokenIdsLength() (*big.Int, error) {
	return _Nft.Contract.ListedTokenIdsLength(&_Nft.CallOpts)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, uint64 listedAt, bool active)
func (_Nft *NftCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller   common.Address
	ListedAt uint64
	Active   bool
}, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "listings", arg0)

	outstruct := new(struct {
		Seller   common.Address
		ListedAt uint64
		Active   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ListedAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Active = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, uint64 listedAt, bool active)
func (_Nft *NftSession) Listings(arg0 *big.Int) (struct {
	Seller   common.Address
	ListedAt uint64
	Active   bool
}, error) {
	return _Nft.Contract.Listings(&_Nft.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, uint64 listedAt, bool active)
func (_Nft *NftCallerSession) Listings(arg0 *big.Int) (struct {
	Seller   common.Address
	ListedAt uint64
	Active   bool
}, error) {
	return _Nft.Contract.Listings(&_Nft.CallOpts, arg0)
}

// MintedTokenIds is a free data retrieval call binding the contract method 0xb95b63f8.
//
// Solidity: function mintedTokenIds(uint256 ) view returns(uint256)
func (_Nft *NftCaller) MintedTokenIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "mintedTokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintedTokenIds is a free data retrieval call binding the contract method 0xb95b63f8.
//
// Solidity: function mintedTokenIds(uint256 ) view returns(uint256)
func (_Nft *NftSession) MintedTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Nft.Contract.MintedTokenIds(&_Nft.CallOpts, arg0)
}

// MintedTokenIds is a free data retrieval call binding the contract method 0xb95b63f8.
//
// Solidity: function mintedTokenIds(uint256 ) view returns(uint256)
func (_Nft *NftCallerSession) MintedTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Nft.Contract.MintedTokenIds(&_Nft.CallOpts, arg0)
}

// MintedTokenIdsLength is a free data retrieval call binding the contract method 0x0ff896de.
//
// Solidity: function mintedTokenIdsLength() view returns(uint256)
func (_Nft *NftCaller) MintedTokenIdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "mintedTokenIdsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintedTokenIdsLength is a free data retrieval call binding the contract method 0x0ff896de.
//
// Solidity: function mintedTokenIdsLength() view returns(uint256)
func (_Nft *NftSession) MintedTokenIdsLength() (*big.Int, error) {
	return _Nft.Contract.MintedTokenIdsLength(&_Nft.CallOpts)
}

// MintedTokenIdsLength is a free data retrieval call binding the contract method 0x0ff896de.
//
// Solidity: function mintedTokenIdsLength() view returns(uint256)
func (_Nft *NftCallerSession) MintedTokenIdsLength() (*big.Int, error) {
	return _Nft.Contract.MintedTokenIdsLength(&_Nft.CallOpts)
}

// MonthlyRateDefaultE18 is a free data retrieval call binding the contract method 0x14cd11bd.
//
// Solidity: function monthlyRateDefaultE18() view returns(uint256)
func (_Nft *NftCaller) MonthlyRateDefaultE18(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "monthlyRateDefaultE18")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MonthlyRateDefaultE18 is a free data retrieval call binding the contract method 0x14cd11bd.
//
// Solidity: function monthlyRateDefaultE18() view returns(uint256)
func (_Nft *NftSession) MonthlyRateDefaultE18() (*big.Int, error) {
	return _Nft.Contract.MonthlyRateDefaultE18(&_Nft.CallOpts)
}

// MonthlyRateDefaultE18 is a free data retrieval call binding the contract method 0x14cd11bd.
//
// Solidity: function monthlyRateDefaultE18() view returns(uint256)
func (_Nft *NftCallerSession) MonthlyRateDefaultE18() (*big.Int, error) {
	return _Nft.Contract.MonthlyRateDefaultE18(&_Nft.CallOpts)
}

// MyOpenedBoxesLength is a free data retrieval call binding the contract method 0xc2f5b597.
//
// Solidity: function myOpenedBoxesLength(address user) view returns(uint256)
func (_Nft *NftCaller) MyOpenedBoxesLength(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "myOpenedBoxesLength", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyOpenedBoxesLength is a free data retrieval call binding the contract method 0xc2f5b597.
//
// Solidity: function myOpenedBoxesLength(address user) view returns(uint256)
func (_Nft *NftSession) MyOpenedBoxesLength(user common.Address) (*big.Int, error) {
	return _Nft.Contract.MyOpenedBoxesLength(&_Nft.CallOpts, user)
}

// MyOpenedBoxesLength is a free data retrieval call binding the contract method 0xc2f5b597.
//
// Solidity: function myOpenedBoxesLength(address user) view returns(uint256)
func (_Nft *NftCallerSession) MyOpenedBoxesLength(user common.Address) (*big.Int, error) {
	return _Nft.Contract.MyOpenedBoxesLength(&_Nft.CallOpts, user)
}

// MyUnopenedBoxesLength is a free data retrieval call binding the contract method 0x72fb5ee6.
//
// Solidity: function myUnopenedBoxesLength(address user) view returns(uint256)
func (_Nft *NftCaller) MyUnopenedBoxesLength(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "myUnopenedBoxesLength", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyUnopenedBoxesLength is a free data retrieval call binding the contract method 0x72fb5ee6.
//
// Solidity: function myUnopenedBoxesLength(address user) view returns(uint256)
func (_Nft *NftSession) MyUnopenedBoxesLength(user common.Address) (*big.Int, error) {
	return _Nft.Contract.MyUnopenedBoxesLength(&_Nft.CallOpts, user)
}

// MyUnopenedBoxesLength is a free data retrieval call binding the contract method 0x72fb5ee6.
//
// Solidity: function myUnopenedBoxesLength(address user) view returns(uint256)
func (_Nft *NftCallerSession) MyUnopenedBoxesLength(user common.Address) (*big.Int, error) {
	return _Nft.Contract.MyUnopenedBoxesLength(&_Nft.CallOpts, user)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nft *NftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nft *NftSession) Name() (string, error) {
	return _Nft.Contract.Name(&_Nft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nft *NftCallerSession) Name() (string, error) {
	return _Nft.Contract.Name(&_Nft.CallOpts)
}

// Open is a free data retrieval call binding the contract method 0x690e7c09.
//
// Solidity: function open(uint256 ) view returns(bool)
func (_Nft *NftCaller) Open(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "open", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Open is a free data retrieval call binding the contract method 0x690e7c09.
//
// Solidity: function open(uint256 ) view returns(bool)
func (_Nft *NftSession) Open(arg0 *big.Int) (bool, error) {
	return _Nft.Contract.Open(&_Nft.CallOpts, arg0)
}

// Open is a free data retrieval call binding the contract method 0x690e7c09.
//
// Solidity: function open(uint256 ) view returns(bool)
func (_Nft *NftCallerSession) Open(arg0 *big.Int) (bool, error) {
	return _Nft.Contract.Open(&_Nft.CallOpts, arg0)
}

// OpenActionTokenIds is a free data retrieval call binding the contract method 0xd9f51f65.
//
// Solidity: function openActionTokenIds(uint256 ) view returns(uint256)
func (_Nft *NftCaller) OpenActionTokenIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "openActionTokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenActionTokenIds is a free data retrieval call binding the contract method 0xd9f51f65.
//
// Solidity: function openActionTokenIds(uint256 ) view returns(uint256)
func (_Nft *NftSession) OpenActionTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Nft.Contract.OpenActionTokenIds(&_Nft.CallOpts, arg0)
}

// OpenActionTokenIds is a free data retrieval call binding the contract method 0xd9f51f65.
//
// Solidity: function openActionTokenIds(uint256 ) view returns(uint256)
func (_Nft *NftCallerSession) OpenActionTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Nft.Contract.OpenActionTokenIds(&_Nft.CallOpts, arg0)
}

// OpenFeeBNB is a free data retrieval call binding the contract method 0xcaae60b5.
//
// Solidity: function openFeeBNB() view returns(uint256)
func (_Nft *NftCaller) OpenFeeBNB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "openFeeBNB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenFeeBNB is a free data retrieval call binding the contract method 0xcaae60b5.
//
// Solidity: function openFeeBNB() view returns(uint256)
func (_Nft *NftSession) OpenFeeBNB() (*big.Int, error) {
	return _Nft.Contract.OpenFeeBNB(&_Nft.CallOpts)
}

// OpenFeeBNB is a free data retrieval call binding the contract method 0xcaae60b5.
//
// Solidity: function openFeeBNB() view returns(uint256)
func (_Nft *NftCallerSession) OpenFeeBNB() (*big.Int, error) {
	return _Nft.Contract.OpenFeeBNB(&_Nft.CallOpts)
}

// OpenFeeReceiver is a free data retrieval call binding the contract method 0xa3494545.
//
// Solidity: function openFeeReceiver() view returns(address)
func (_Nft *NftCaller) OpenFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "openFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OpenFeeReceiver is a free data retrieval call binding the contract method 0xa3494545.
//
// Solidity: function openFeeReceiver() view returns(address)
func (_Nft *NftSession) OpenFeeReceiver() (common.Address, error) {
	return _Nft.Contract.OpenFeeReceiver(&_Nft.CallOpts)
}

// OpenFeeReceiver is a free data retrieval call binding the contract method 0xa3494545.
//
// Solidity: function openFeeReceiver() view returns(address)
func (_Nft *NftCallerSession) OpenFeeReceiver() (common.Address, error) {
	return _Nft.Contract.OpenFeeReceiver(&_Nft.CallOpts)
}

// OpenUsageNow is a free data retrieval call binding the contract method 0x9401f811.
//
// Solidity: function openUsageNow(address user) view returns(uint256 dayIndex, uint256 used, uint256 limit, uint256 remaining)
func (_Nft *NftCaller) OpenUsageNow(opts *bind.CallOpts, user common.Address) (struct {
	DayIndex  *big.Int
	Used      *big.Int
	Limit     *big.Int
	Remaining *big.Int
}, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "openUsageNow", user)

	outstruct := new(struct {
		DayIndex  *big.Int
		Used      *big.Int
		Limit     *big.Int
		Remaining *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DayIndex = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Used = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Limit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Remaining = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OpenUsageNow is a free data retrieval call binding the contract method 0x9401f811.
//
// Solidity: function openUsageNow(address user) view returns(uint256 dayIndex, uint256 used, uint256 limit, uint256 remaining)
func (_Nft *NftSession) OpenUsageNow(user common.Address) (struct {
	DayIndex  *big.Int
	Used      *big.Int
	Limit     *big.Int
	Remaining *big.Int
}, error) {
	return _Nft.Contract.OpenUsageNow(&_Nft.CallOpts, user)
}

// OpenUsageNow is a free data retrieval call binding the contract method 0x9401f811.
//
// Solidity: function openUsageNow(address user) view returns(uint256 dayIndex, uint256 used, uint256 limit, uint256 remaining)
func (_Nft *NftCallerSession) OpenUsageNow(user common.Address) (struct {
	DayIndex  *big.Int
	Used      *big.Int
	Limit     *big.Int
	Remaining *big.Int
}, error) {
	return _Nft.Contract.OpenUsageNow(&_Nft.CallOpts, user)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Nft *NftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Nft *NftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Nft.Contract.OwnerOf(&_Nft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Nft *NftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Nft.Contract.OwnerOf(&_Nft.CallOpts, tokenId)
}

// PieceHoldingOf is a free data retrieval call binding the contract method 0xf3613b2c.
//
// Solidity: function pieceHoldingOf(address user) view returns(uint256)
func (_Nft *NftCaller) PieceHoldingOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "pieceHoldingOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PieceHoldingOf is a free data retrieval call binding the contract method 0xf3613b2c.
//
// Solidity: function pieceHoldingOf(address user) view returns(uint256)
func (_Nft *NftSession) PieceHoldingOf(user common.Address) (*big.Int, error) {
	return _Nft.Contract.PieceHoldingOf(&_Nft.CallOpts, user)
}

// PieceHoldingOf is a free data retrieval call binding the contract method 0xf3613b2c.
//
// Solidity: function pieceHoldingOf(address user) view returns(uint256)
func (_Nft *NftCallerSession) PieceHoldingOf(user common.Address) (*big.Int, error) {
	return _Nft.Contract.PieceHoldingOf(&_Nft.CallOpts, user)
}

// PieceToken is a free data retrieval call binding the contract method 0x55d41692.
//
// Solidity: function pieceToken() view returns(address)
func (_Nft *NftCaller) PieceToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "pieceToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PieceToken is a free data retrieval call binding the contract method 0x55d41692.
//
// Solidity: function pieceToken() view returns(address)
func (_Nft *NftSession) PieceToken() (common.Address, error) {
	return _Nft.Contract.PieceToken(&_Nft.CallOpts)
}

// PieceToken is a free data retrieval call binding the contract method 0x55d41692.
//
// Solidity: function pieceToken() view returns(address)
func (_Nft *NftCallerSession) PieceToken() (common.Address, error) {
	return _Nft.Contract.PieceToken(&_Nft.CallOpts)
}

// PieceValueUsdtE18 is a free data retrieval call binding the contract method 0x84c08e4b.
//
// Solidity: function pieceValueUsdtE18() view returns(uint256)
func (_Nft *NftCaller) PieceValueUsdtE18(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "pieceValueUsdtE18")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PieceValueUsdtE18 is a free data retrieval call binding the contract method 0x84c08e4b.
//
// Solidity: function pieceValueUsdtE18() view returns(uint256)
func (_Nft *NftSession) PieceValueUsdtE18() (*big.Int, error) {
	return _Nft.Contract.PieceValueUsdtE18(&_Nft.CallOpts)
}

// PieceValueUsdtE18 is a free data retrieval call binding the contract method 0x84c08e4b.
//
// Solidity: function pieceValueUsdtE18() view returns(uint256)
func (_Nft *NftCallerSession) PieceValueUsdtE18() (*big.Int, error) {
	return _Nft.Contract.PieceValueUsdtE18(&_Nft.CallOpts)
}

// PreviewSlotCost is a free data retrieval call binding the contract method 0xc7abcf9d.
//
// Solidity: function previewSlotCost(address user, uint256 count) view returns(uint256 cost, uint256 step)
func (_Nft *NftCaller) PreviewSlotCost(opts *bind.CallOpts, user common.Address, count *big.Int) (struct {
	Cost *big.Int
	Step *big.Int
}, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "previewSlotCost", user, count)

	outstruct := new(struct {
		Cost *big.Int
		Step *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cost = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Step = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PreviewSlotCost is a free data retrieval call binding the contract method 0xc7abcf9d.
//
// Solidity: function previewSlotCost(address user, uint256 count) view returns(uint256 cost, uint256 step)
func (_Nft *NftSession) PreviewSlotCost(user common.Address, count *big.Int) (struct {
	Cost *big.Int
	Step *big.Int
}, error) {
	return _Nft.Contract.PreviewSlotCost(&_Nft.CallOpts, user, count)
}

// PreviewSlotCost is a free data retrieval call binding the contract method 0xc7abcf9d.
//
// Solidity: function previewSlotCost(address user, uint256 count) view returns(uint256 cost, uint256 step)
func (_Nft *NftCallerSession) PreviewSlotCost(user common.Address, count *big.Int) (struct {
	Cost *big.Int
	Step *big.Int
}, error) {
	return _Nft.Contract.PreviewSlotCost(&_Nft.CallOpts, user, count)
}

// QuoteFeeInB is a free data retrieval call binding the contract method 0xcedac22c.
//
// Solidity: function quoteFeeInB(uint256 amount1InUSDT) view returns(uint256 out0NetB)
func (_Nft *NftCaller) QuoteFeeInB(opts *bind.CallOpts, amount1InUSDT *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "quoteFeeInB", amount1InUSDT)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteFeeInB is a free data retrieval call binding the contract method 0xcedac22c.
//
// Solidity: function quoteFeeInB(uint256 amount1InUSDT) view returns(uint256 out0NetB)
func (_Nft *NftSession) QuoteFeeInB(amount1InUSDT *big.Int) (*big.Int, error) {
	return _Nft.Contract.QuoteFeeInB(&_Nft.CallOpts, amount1InUSDT)
}

// QuoteFeeInB is a free data retrieval call binding the contract method 0xcedac22c.
//
// Solidity: function quoteFeeInB(uint256 amount1InUSDT) view returns(uint256 out0NetB)
func (_Nft *NftCallerSession) QuoteFeeInB(amount1InUSDT *big.Int) (*big.Int, error) {
	return _Nft.Contract.QuoteFeeInB(&_Nft.CallOpts, amount1InUSDT)
}

// SecondaryBuyPrice is a free data retrieval call binding the contract method 0xafbb5263.
//
// Solidity: function secondaryBuyPrice(uint256 tokenId) view returns(uint256 price)
func (_Nft *NftCaller) SecondaryBuyPrice(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "secondaryBuyPrice", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondaryBuyPrice is a free data retrieval call binding the contract method 0xafbb5263.
//
// Solidity: function secondaryBuyPrice(uint256 tokenId) view returns(uint256 price)
func (_Nft *NftSession) SecondaryBuyPrice(tokenId *big.Int) (*big.Int, error) {
	return _Nft.Contract.SecondaryBuyPrice(&_Nft.CallOpts, tokenId)
}

// SecondaryBuyPrice is a free data retrieval call binding the contract method 0xafbb5263.
//
// Solidity: function secondaryBuyPrice(uint256 tokenId) view returns(uint256 price)
func (_Nft *NftCallerSession) SecondaryBuyPrice(tokenId *big.Int) (*big.Int, error) {
	return _Nft.Contract.SecondaryBuyPrice(&_Nft.CallOpts, tokenId)
}

// SellerSoldTokenIdsLength is a free data retrieval call binding the contract method 0x4dbfd62b.
//
// Solidity: function sellerSoldTokenIdsLength(address seller) view returns(uint256)
func (_Nft *NftCaller) SellerSoldTokenIdsLength(opts *bind.CallOpts, seller common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "sellerSoldTokenIdsLength", seller)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellerSoldTokenIdsLength is a free data retrieval call binding the contract method 0x4dbfd62b.
//
// Solidity: function sellerSoldTokenIdsLength(address seller) view returns(uint256)
func (_Nft *NftSession) SellerSoldTokenIdsLength(seller common.Address) (*big.Int, error) {
	return _Nft.Contract.SellerSoldTokenIdsLength(&_Nft.CallOpts, seller)
}

// SellerSoldTokenIdsLength is a free data retrieval call binding the contract method 0x4dbfd62b.
//
// Solidity: function sellerSoldTokenIdsLength(address seller) view returns(uint256)
func (_Nft *NftCallerSession) SellerSoldTokenIdsLength(seller common.Address) (*big.Int, error) {
	return _Nft.Contract.SellerSoldTokenIdsLength(&_Nft.CallOpts, seller)
}

// SlotStepForUser is a free data retrieval call binding the contract method 0x29dd4d4f.
//
// Solidity: function slotStepForUser(address user) view returns(uint256)
func (_Nft *NftCaller) SlotStepForUser(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "slotStepForUser", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotStepForUser is a free data retrieval call binding the contract method 0x29dd4d4f.
//
// Solidity: function slotStepForUser(address user) view returns(uint256)
func (_Nft *NftSession) SlotStepForUser(user common.Address) (*big.Int, error) {
	return _Nft.Contract.SlotStepForUser(&_Nft.CallOpts, user)
}

// SlotStepForUser is a free data retrieval call binding the contract method 0x29dd4d4f.
//
// Solidity: function slotStepForUser(address user) view returns(uint256)
func (_Nft *NftCallerSession) SlotStepForUser(user common.Address) (*big.Int, error) {
	return _Nft.Contract.SlotStepForUser(&_Nft.CallOpts, user)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Nft *NftCaller) StakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "stakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Nft *NftSession) StakingContract() (common.Address, error) {
	return _Nft.Contract.StakingContract(&_Nft.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Nft *NftCallerSession) StakingContract() (common.Address, error) {
	return _Nft.Contract.StakingContract(&_Nft.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nft *NftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nft *NftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Nft.Contract.SupportsInterface(&_Nft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nft *NftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Nft.Contract.SupportsInterface(&_Nft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nft *NftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nft *NftSession) Symbol() (string, error) {
	return _Nft.Contract.Symbol(&_Nft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nft *NftCallerSession) Symbol() (string, error) {
	return _Nft.Contract.Symbol(&_Nft.CallOpts)
}

// TierPrice is a free data retrieval call binding the contract method 0x8b542526.
//
// Solidity: function tierPrice(uint8 ) view returns(uint256)
func (_Nft *NftCaller) TierPrice(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "tierPrice", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierPrice is a free data retrieval call binding the contract method 0x8b542526.
//
// Solidity: function tierPrice(uint8 ) view returns(uint256)
func (_Nft *NftSession) TierPrice(arg0 uint8) (*big.Int, error) {
	return _Nft.Contract.TierPrice(&_Nft.CallOpts, arg0)
}

// TierPrice is a free data retrieval call binding the contract method 0x8b542526.
//
// Solidity: function tierPrice(uint8 ) view returns(uint256)
func (_Nft *NftCallerSession) TierPrice(arg0 uint8) (*big.Int, error) {
	return _Nft.Contract.TierPrice(&_Nft.CallOpts, arg0)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Nft *NftCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Nft *NftSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Nft.Contract.TokenByIndex(&_Nft.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Nft *NftCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Nft.Contract.TokenByIndex(&_Nft.CallOpts, index)
}

// TokenMonthlyRateE18 is a free data retrieval call binding the contract method 0xba32dff1.
//
// Solidity: function tokenMonthlyRateE18(uint256 ) view returns(uint256)
func (_Nft *NftCaller) TokenMonthlyRateE18(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "tokenMonthlyRateE18", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMonthlyRateE18 is a free data retrieval call binding the contract method 0xba32dff1.
//
// Solidity: function tokenMonthlyRateE18(uint256 ) view returns(uint256)
func (_Nft *NftSession) TokenMonthlyRateE18(arg0 *big.Int) (*big.Int, error) {
	return _Nft.Contract.TokenMonthlyRateE18(&_Nft.CallOpts, arg0)
}

// TokenMonthlyRateE18 is a free data retrieval call binding the contract method 0xba32dff1.
//
// Solidity: function tokenMonthlyRateE18(uint256 ) view returns(uint256)
func (_Nft *NftCallerSession) TokenMonthlyRateE18(arg0 *big.Int) (*big.Int, error) {
	return _Nft.Contract.TokenMonthlyRateE18(&_Nft.CallOpts, arg0)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Nft *NftCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Nft *NftSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Nft.Contract.TokenOfOwnerByIndex(&_Nft.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Nft *NftCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Nft.Contract.TokenOfOwnerByIndex(&_Nft.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Nft *NftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Nft *NftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Nft.Contract.TokenURI(&_Nft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Nft *NftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Nft.Contract.TokenURI(&_Nft.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nft *NftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nft *NftSession) TotalSupply() (*big.Int, error) {
	return _Nft.Contract.TotalSupply(&_Nft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nft *NftCallerSession) TotalSupply() (*big.Int, error) {
	return _Nft.Contract.TotalSupply(&_Nft.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Nft *NftCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Nft *NftSession) Usdt() (common.Address, error) {
	return _Nft.Contract.Usdt(&_Nft.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Nft *NftCallerSession) Usdt() (common.Address, error) {
	return _Nft.Contract.Usdt(&_Nft.CallOpts)
}

// UserSlotStepCost is a free data retrieval call binding the contract method 0x2051cff1.
//
// Solidity: function userSlotStepCost(address ) view returns(uint256)
func (_Nft *NftCaller) UserSlotStepCost(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "userSlotStepCost", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserSlotStepCost is a free data retrieval call binding the contract method 0x2051cff1.
//
// Solidity: function userSlotStepCost(address ) view returns(uint256)
func (_Nft *NftSession) UserSlotStepCost(arg0 common.Address) (*big.Int, error) {
	return _Nft.Contract.UserSlotStepCost(&_Nft.CallOpts, arg0)
}

// UserSlotStepCost is a free data retrieval call binding the contract method 0x2051cff1.
//
// Solidity: function userSlotStepCost(address ) view returns(uint256)
func (_Nft *NftCallerSession) UserSlotStepCost(arg0 common.Address) (*big.Int, error) {
	return _Nft.Contract.UserSlotStepCost(&_Nft.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Nft *NftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Nft *NftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.Approve(&_Nft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Nft *NftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.Approve(&_Nft.TransactOpts, to, tokenId)
}

// BurnAfterReward is a paid mutator transaction binding the contract method 0xf8b85ec3.
//
// Solidity: function burnAfterReward(uint256 tokenId) returns()
func (_Nft *NftTransactor) BurnAfterReward(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "burnAfterReward", tokenId)
}

// BurnAfterReward is a paid mutator transaction binding the contract method 0xf8b85ec3.
//
// Solidity: function burnAfterReward(uint256 tokenId) returns()
func (_Nft *NftSession) BurnAfterReward(tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.BurnAfterReward(&_Nft.TransactOpts, tokenId)
}

// BurnAfterReward is a paid mutator transaction binding the contract method 0xf8b85ec3.
//
// Solidity: function burnAfterReward(uint256 tokenId) returns()
func (_Nft *NftTransactorSession) BurnAfterReward(tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.BurnAfterReward(&_Nft.TransactOpts, tokenId)
}

// BurnAfterRewardBatch is a paid mutator transaction binding the contract method 0x4cbecef0.
//
// Solidity: function burnAfterRewardBatch(uint256[] tokenIds) returns()
func (_Nft *NftTransactor) BurnAfterRewardBatch(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "burnAfterRewardBatch", tokenIds)
}

// BurnAfterRewardBatch is a paid mutator transaction binding the contract method 0x4cbecef0.
//
// Solidity: function burnAfterRewardBatch(uint256[] tokenIds) returns()
func (_Nft *NftSession) BurnAfterRewardBatch(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Nft.Contract.BurnAfterRewardBatch(&_Nft.TransactOpts, tokenIds)
}

// BurnAfterRewardBatch is a paid mutator transaction binding the contract method 0x4cbecef0.
//
// Solidity: function burnAfterRewardBatch(uint256[] tokenIds) returns()
func (_Nft *NftTransactorSession) BurnAfterRewardBatch(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Nft.Contract.BurnAfterRewardBatch(&_Nft.TransactOpts, tokenIds)
}

// BuyOpenSlots is a paid mutator transaction binding the contract method 0xfd73eb41.
//
// Solidity: function buyOpenSlots(uint256 count) returns()
func (_Nft *NftTransactor) BuyOpenSlots(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "buyOpenSlots", count)
}

// BuyOpenSlots is a paid mutator transaction binding the contract method 0xfd73eb41.
//
// Solidity: function buyOpenSlots(uint256 count) returns()
func (_Nft *NftSession) BuyOpenSlots(count *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.BuyOpenSlots(&_Nft.TransactOpts, count)
}

// BuyOpenSlots is a paid mutator transaction binding the contract method 0xfd73eb41.
//
// Solidity: function buyOpenSlots(uint256 count) returns()
func (_Nft *NftTransactorSession) BuyOpenSlots(count *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.BuyOpenSlots(&_Nft.TransactOpts, count)
}

// BuyPayFeeInB is a paid mutator transaction binding the contract method 0x346d08e0.
//
// Solidity: function buyPayFeeInB(uint256 tokenId, uint256 maxFeeB) returns()
func (_Nft *NftTransactor) BuyPayFeeInB(opts *bind.TransactOpts, tokenId *big.Int, maxFeeB *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "buyPayFeeInB", tokenId, maxFeeB)
}

// BuyPayFeeInB is a paid mutator transaction binding the contract method 0x346d08e0.
//
// Solidity: function buyPayFeeInB(uint256 tokenId, uint256 maxFeeB) returns()
func (_Nft *NftSession) BuyPayFeeInB(tokenId *big.Int, maxFeeB *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.BuyPayFeeInB(&_Nft.TransactOpts, tokenId, maxFeeB)
}

// BuyPayFeeInB is a paid mutator transaction binding the contract method 0x346d08e0.
//
// Solidity: function buyPayFeeInB(uint256 tokenId, uint256 maxFeeB) returns()
func (_Nft *NftTransactorSession) BuyPayFeeInB(tokenId *big.Int, maxFeeB *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.BuyPayFeeInB(&_Nft.TransactOpts, tokenId, maxFeeB)
}

// BuyPayFeeInUSDT is a paid mutator transaction binding the contract method 0xf7b4e561.
//
// Solidity: function buyPayFeeInUSDT(uint256 tokenId) returns()
func (_Nft *NftTransactor) BuyPayFeeInUSDT(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "buyPayFeeInUSDT", tokenId)
}

// BuyPayFeeInUSDT is a paid mutator transaction binding the contract method 0xf7b4e561.
//
// Solidity: function buyPayFeeInUSDT(uint256 tokenId) returns()
func (_Nft *NftSession) BuyPayFeeInUSDT(tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.BuyPayFeeInUSDT(&_Nft.TransactOpts, tokenId)
}

// BuyPayFeeInUSDT is a paid mutator transaction binding the contract method 0xf7b4e561.
//
// Solidity: function buyPayFeeInUSDT(uint256 tokenId) returns()
func (_Nft *NftTransactorSession) BuyPayFeeInUSDT(tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.BuyPayFeeInUSDT(&_Nft.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Nft *NftTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Nft *NftSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nft.Contract.GrantRole(&_Nft.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Nft *NftTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nft.Contract.GrantRole(&_Nft.TransactOpts, role, account)
}

// HoldingAdd is a paid mutator transaction binding the contract method 0xbea1d461.
//
// Solidity: function holdingAdd(address user, uint256 burnAmount) returns()
func (_Nft *NftTransactor) HoldingAdd(opts *bind.TransactOpts, user common.Address, burnAmount *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "holdingAdd", user, burnAmount)
}

// HoldingAdd is a paid mutator transaction binding the contract method 0xbea1d461.
//
// Solidity: function holdingAdd(address user, uint256 burnAmount) returns()
func (_Nft *NftSession) HoldingAdd(user common.Address, burnAmount *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.HoldingAdd(&_Nft.TransactOpts, user, burnAmount)
}

// HoldingAdd is a paid mutator transaction binding the contract method 0xbea1d461.
//
// Solidity: function holdingAdd(address user, uint256 burnAmount) returns()
func (_Nft *NftTransactorSession) HoldingAdd(user common.Address, burnAmount *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.HoldingAdd(&_Nft.TransactOpts, user, burnAmount)
}

// List is a paid mutator transaction binding the contract method 0x80c9419e.
//
// Solidity: function list(uint256 tokenId) returns()
func (_Nft *NftTransactor) List(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "list", tokenId)
}

// List is a paid mutator transaction binding the contract method 0x80c9419e.
//
// Solidity: function list(uint256 tokenId) returns()
func (_Nft *NftSession) List(tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.List(&_Nft.TransactOpts, tokenId)
}

// List is a paid mutator transaction binding the contract method 0x80c9419e.
//
// Solidity: function list(uint256 tokenId) returns()
func (_Nft *NftTransactorSession) List(tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.List(&_Nft.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x6ecd2306.
//
// Solidity: function mint(uint8 tier) returns(uint256 tokenId)
func (_Nft *NftTransactor) Mint(opts *bind.TransactOpts, tier uint8) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "mint", tier)
}

// Mint is a paid mutator transaction binding the contract method 0x6ecd2306.
//
// Solidity: function mint(uint8 tier) returns(uint256 tokenId)
func (_Nft *NftSession) Mint(tier uint8) (*types.Transaction, error) {
	return _Nft.Contract.Mint(&_Nft.TransactOpts, tier)
}

// Mint is a paid mutator transaction binding the contract method 0x6ecd2306.
//
// Solidity: function mint(uint8 tier) returns(uint256 tokenId)
func (_Nft *NftTransactorSession) Mint(tier uint8) (*types.Transaction, error) {
	return _Nft.Contract.Mint(&_Nft.TransactOpts, tier)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd6d88d4e.
//
// Solidity: function mintBatch(uint8 tier, uint256 quantity) returns(uint256[] tokenIds)
func (_Nft *NftTransactor) MintBatch(opts *bind.TransactOpts, tier uint8, quantity *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "mintBatch", tier, quantity)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd6d88d4e.
//
// Solidity: function mintBatch(uint8 tier, uint256 quantity) returns(uint256[] tokenIds)
func (_Nft *NftSession) MintBatch(tier uint8, quantity *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.MintBatch(&_Nft.TransactOpts, tier, quantity)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd6d88d4e.
//
// Solidity: function mintBatch(uint8 tier, uint256 quantity) returns(uint256[] tokenIds)
func (_Nft *NftTransactorSession) MintBatch(tier uint8, quantity *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.MintBatch(&_Nft.TransactOpts, tier, quantity)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 tokenId) payable returns()
func (_Nft *NftTransactor) OpenBox(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "openBox", tokenId)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 tokenId) payable returns()
func (_Nft *NftSession) OpenBox(tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.OpenBox(&_Nft.TransactOpts, tokenId)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 tokenId) payable returns()
func (_Nft *NftTransactorSession) OpenBox(tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.OpenBox(&_Nft.TransactOpts, tokenId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Nft *NftTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Nft *NftSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Nft.Contract.RenounceRole(&_Nft.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Nft *NftTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Nft.Contract.RenounceRole(&_Nft.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Nft *NftTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Nft *NftSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nft.Contract.RevokeRole(&_Nft.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Nft *NftTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nft.Contract.RevokeRole(&_Nft.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Nft *NftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Nft *NftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SafeTransferFrom(&_Nft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Nft *NftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SafeTransferFrom(&_Nft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Nft *NftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Nft *NftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Nft.Contract.SafeTransferFrom0(&_Nft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Nft *NftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Nft.Contract.SafeTransferFrom0(&_Nft.TransactOpts, from, to, tokenId, data)
}

// SetAToken is a paid mutator transaction binding the contract method 0x5f09d84c.
//
// Solidity: function setAToken(address newAToken) returns()
func (_Nft *NftTransactor) SetAToken(opts *bind.TransactOpts, newAToken common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setAToken", newAToken)
}

// SetAToken is a paid mutator transaction binding the contract method 0x5f09d84c.
//
// Solidity: function setAToken(address newAToken) returns()
func (_Nft *NftSession) SetAToken(newAToken common.Address) (*types.Transaction, error) {
	return _Nft.Contract.SetAToken(&_Nft.TransactOpts, newAToken)
}

// SetAToken is a paid mutator transaction binding the contract method 0x5f09d84c.
//
// Solidity: function setAToken(address newAToken) returns()
func (_Nft *NftTransactorSession) SetAToken(newAToken common.Address) (*types.Transaction, error) {
	return _Nft.Contract.SetAToken(&_Nft.TransactOpts, newAToken)
}

// SetAccountA is a paid mutator transaction binding the contract method 0x8eaeb531.
//
// Solidity: function setAccountA(address newAccountA) returns()
func (_Nft *NftTransactor) SetAccountA(opts *bind.TransactOpts, newAccountA common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setAccountA", newAccountA)
}

// SetAccountA is a paid mutator transaction binding the contract method 0x8eaeb531.
//
// Solidity: function setAccountA(address newAccountA) returns()
func (_Nft *NftSession) SetAccountA(newAccountA common.Address) (*types.Transaction, error) {
	return _Nft.Contract.SetAccountA(&_Nft.TransactOpts, newAccountA)
}

// SetAccountA is a paid mutator transaction binding the contract method 0x8eaeb531.
//
// Solidity: function setAccountA(address newAccountA) returns()
func (_Nft *NftTransactorSession) SetAccountA(newAccountA common.Address) (*types.Transaction, error) {
	return _Nft.Contract.SetAccountA(&_Nft.TransactOpts, newAccountA)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nft *NftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nft *NftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nft.Contract.SetApprovalForAll(&_Nft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nft *NftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nft.Contract.SetApprovalForAll(&_Nft.TransactOpts, operator, approved)
}

// SetFeeConfigB is a paid mutator transaction binding the contract method 0x160bfb34.
//
// Solidity: function setFeeConfigB(uint256 rateTwo, uint256 baseTwo) returns()
func (_Nft *NftTransactor) SetFeeConfigB(opts *bind.TransactOpts, rateTwo *big.Int, baseTwo *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setFeeConfigB", rateTwo, baseTwo)
}

// SetFeeConfigB is a paid mutator transaction binding the contract method 0x160bfb34.
//
// Solidity: function setFeeConfigB(uint256 rateTwo, uint256 baseTwo) returns()
func (_Nft *NftSession) SetFeeConfigB(rateTwo *big.Int, baseTwo *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetFeeConfigB(&_Nft.TransactOpts, rateTwo, baseTwo)
}

// SetFeeConfigB is a paid mutator transaction binding the contract method 0x160bfb34.
//
// Solidity: function setFeeConfigB(uint256 rateTwo, uint256 baseTwo) returns()
func (_Nft *NftTransactorSession) SetFeeConfigB(rateTwo *big.Int, baseTwo *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetFeeConfigB(&_Nft.TransactOpts, rateTwo, baseTwo)
}

// SetFeeConfigUSDT is a paid mutator transaction binding the contract method 0x92baf3c4.
//
// Solidity: function setFeeConfigUSDT(uint256 rate, uint256 base) returns()
func (_Nft *NftTransactor) SetFeeConfigUSDT(opts *bind.TransactOpts, rate *big.Int, base *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setFeeConfigUSDT", rate, base)
}

// SetFeeConfigUSDT is a paid mutator transaction binding the contract method 0x92baf3c4.
//
// Solidity: function setFeeConfigUSDT(uint256 rate, uint256 base) returns()
func (_Nft *NftSession) SetFeeConfigUSDT(rate *big.Int, base *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetFeeConfigUSDT(&_Nft.TransactOpts, rate, base)
}

// SetFeeConfigUSDT is a paid mutator transaction binding the contract method 0x92baf3c4.
//
// Solidity: function setFeeConfigUSDT(uint256 rate, uint256 base) returns()
func (_Nft *NftTransactorSession) SetFeeConfigUSDT(rate *big.Int, base *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetFeeConfigUSDT(&_Nft.TransactOpts, rate, base)
}

// SetFeeQuotePool is a paid mutator transaction binding the contract method 0xf6c89e15.
//
// Solidity: function setFeeQuotePool(address pool) returns()
func (_Nft *NftTransactor) SetFeeQuotePool(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setFeeQuotePool", pool)
}

// SetFeeQuotePool is a paid mutator transaction binding the contract method 0xf6c89e15.
//
// Solidity: function setFeeQuotePool(address pool) returns()
func (_Nft *NftSession) SetFeeQuotePool(pool common.Address) (*types.Transaction, error) {
	return _Nft.Contract.SetFeeQuotePool(&_Nft.TransactOpts, pool)
}

// SetFeeQuotePool is a paid mutator transaction binding the contract method 0xf6c89e15.
//
// Solidity: function setFeeQuotePool(address pool) returns()
func (_Nft *NftTransactorSession) SetFeeQuotePool(pool common.Address) (*types.Transaction, error) {
	return _Nft.Contract.SetFeeQuotePool(&_Nft.TransactOpts, pool)
}

// SetGlobalSlotStepCost is a paid mutator transaction binding the contract method 0x6521b058.
//
// Solidity: function setGlobalSlotStepCost(uint256 newCost) returns()
func (_Nft *NftTransactor) SetGlobalSlotStepCost(opts *bind.TransactOpts, newCost *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setGlobalSlotStepCost", newCost)
}

// SetGlobalSlotStepCost is a paid mutator transaction binding the contract method 0x6521b058.
//
// Solidity: function setGlobalSlotStepCost(uint256 newCost) returns()
func (_Nft *NftSession) SetGlobalSlotStepCost(newCost *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetGlobalSlotStepCost(&_Nft.TransactOpts, newCost)
}

// SetGlobalSlotStepCost is a paid mutator transaction binding the contract method 0x6521b058.
//
// Solidity: function setGlobalSlotStepCost(uint256 newCost) returns()
func (_Nft *NftTransactorSession) SetGlobalSlotStepCost(newCost *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetGlobalSlotStepCost(&_Nft.TransactOpts, newCost)
}

// SetOpenFeeBNB is a paid mutator transaction binding the contract method 0x775859d0.
//
// Solidity: function setOpenFeeBNB(uint256 feeWei) returns()
func (_Nft *NftTransactor) SetOpenFeeBNB(opts *bind.TransactOpts, feeWei *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setOpenFeeBNB", feeWei)
}

// SetOpenFeeBNB is a paid mutator transaction binding the contract method 0x775859d0.
//
// Solidity: function setOpenFeeBNB(uint256 feeWei) returns()
func (_Nft *NftSession) SetOpenFeeBNB(feeWei *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetOpenFeeBNB(&_Nft.TransactOpts, feeWei)
}

// SetOpenFeeBNB is a paid mutator transaction binding the contract method 0x775859d0.
//
// Solidity: function setOpenFeeBNB(uint256 feeWei) returns()
func (_Nft *NftTransactorSession) SetOpenFeeBNB(feeWei *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetOpenFeeBNB(&_Nft.TransactOpts, feeWei)
}

// SetOpenFeeConfig is a paid mutator transaction binding the contract method 0x12b811fa.
//
// Solidity: function setOpenFeeConfig(uint256 feeWei, address receiver) returns()
func (_Nft *NftTransactor) SetOpenFeeConfig(opts *bind.TransactOpts, feeWei *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setOpenFeeConfig", feeWei, receiver)
}

// SetOpenFeeConfig is a paid mutator transaction binding the contract method 0x12b811fa.
//
// Solidity: function setOpenFeeConfig(uint256 feeWei, address receiver) returns()
func (_Nft *NftSession) SetOpenFeeConfig(feeWei *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Nft.Contract.SetOpenFeeConfig(&_Nft.TransactOpts, feeWei, receiver)
}

// SetOpenFeeConfig is a paid mutator transaction binding the contract method 0x12b811fa.
//
// Solidity: function setOpenFeeConfig(uint256 feeWei, address receiver) returns()
func (_Nft *NftTransactorSession) SetOpenFeeConfig(feeWei *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Nft.Contract.SetOpenFeeConfig(&_Nft.TransactOpts, feeWei, receiver)
}

// SetOpenFeeReceiver is a paid mutator transaction binding the contract method 0xd59d8dd2.
//
// Solidity: function setOpenFeeReceiver(address receiver) returns()
func (_Nft *NftTransactor) SetOpenFeeReceiver(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setOpenFeeReceiver", receiver)
}

// SetOpenFeeReceiver is a paid mutator transaction binding the contract method 0xd59d8dd2.
//
// Solidity: function setOpenFeeReceiver(address receiver) returns()
func (_Nft *NftSession) SetOpenFeeReceiver(receiver common.Address) (*types.Transaction, error) {
	return _Nft.Contract.SetOpenFeeReceiver(&_Nft.TransactOpts, receiver)
}

// SetOpenFeeReceiver is a paid mutator transaction binding the contract method 0xd59d8dd2.
//
// Solidity: function setOpenFeeReceiver(address receiver) returns()
func (_Nft *NftTransactorSession) SetOpenFeeReceiver(receiver common.Address) (*types.Transaction, error) {
	return _Nft.Contract.SetOpenFeeReceiver(&_Nft.TransactOpts, receiver)
}

// SetParams is a paid mutator transaction binding the contract method 0x5a0ce676.
//
// Solidity: function setParams(uint256 denom, uint256 aFeeBps, uint256 monthlyRateDefault) returns()
func (_Nft *NftTransactor) SetParams(opts *bind.TransactOpts, denom *big.Int, aFeeBps *big.Int, monthlyRateDefault *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setParams", denom, aFeeBps, monthlyRateDefault)
}

// SetParams is a paid mutator transaction binding the contract method 0x5a0ce676.
//
// Solidity: function setParams(uint256 denom, uint256 aFeeBps, uint256 monthlyRateDefault) returns()
func (_Nft *NftSession) SetParams(denom *big.Int, aFeeBps *big.Int, monthlyRateDefault *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetParams(&_Nft.TransactOpts, denom, aFeeBps, monthlyRateDefault)
}

// SetParams is a paid mutator transaction binding the contract method 0x5a0ce676.
//
// Solidity: function setParams(uint256 denom, uint256 aFeeBps, uint256 monthlyRateDefault) returns()
func (_Nft *NftTransactorSession) SetParams(denom *big.Int, aFeeBps *big.Int, monthlyRateDefault *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetParams(&_Nft.TransactOpts, denom, aFeeBps, monthlyRateDefault)
}

// SetPieceToken is a paid mutator transaction binding the contract method 0x6ef54d1e.
//
// Solidity: function setPieceToken(address newPiece) returns()
func (_Nft *NftTransactor) SetPieceToken(opts *bind.TransactOpts, newPiece common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setPieceToken", newPiece)
}

// SetPieceToken is a paid mutator transaction binding the contract method 0x6ef54d1e.
//
// Solidity: function setPieceToken(address newPiece) returns()
func (_Nft *NftSession) SetPieceToken(newPiece common.Address) (*types.Transaction, error) {
	return _Nft.Contract.SetPieceToken(&_Nft.TransactOpts, newPiece)
}

// SetPieceToken is a paid mutator transaction binding the contract method 0x6ef54d1e.
//
// Solidity: function setPieceToken(address newPiece) returns()
func (_Nft *NftTransactorSession) SetPieceToken(newPiece common.Address) (*types.Transaction, error) {
	return _Nft.Contract.SetPieceToken(&_Nft.TransactOpts, newPiece)
}

// SetPieceValueUsdtE18 is a paid mutator transaction binding the contract method 0x013e5afc.
//
// Solidity: function setPieceValueUsdtE18(uint256 newValueUsdtE18) returns()
func (_Nft *NftTransactor) SetPieceValueUsdtE18(opts *bind.TransactOpts, newValueUsdtE18 *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setPieceValueUsdtE18", newValueUsdtE18)
}

// SetPieceValueUsdtE18 is a paid mutator transaction binding the contract method 0x013e5afc.
//
// Solidity: function setPieceValueUsdtE18(uint256 newValueUsdtE18) returns()
func (_Nft *NftSession) SetPieceValueUsdtE18(newValueUsdtE18 *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetPieceValueUsdtE18(&_Nft.TransactOpts, newValueUsdtE18)
}

// SetPieceValueUsdtE18 is a paid mutator transaction binding the contract method 0x013e5afc.
//
// Solidity: function setPieceValueUsdtE18(uint256 newValueUsdtE18) returns()
func (_Nft *NftTransactorSession) SetPieceValueUsdtE18(newValueUsdtE18 *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetPieceValueUsdtE18(&_Nft.TransactOpts, newValueUsdtE18)
}

// SetRewardOnce is a paid mutator transaction binding the contract method 0xe78d18e4.
//
// Solidity: function setRewardOnce(uint256 tokenId, int256 reward_) returns()
func (_Nft *NftTransactor) SetRewardOnce(opts *bind.TransactOpts, tokenId *big.Int, reward_ *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setRewardOnce", tokenId, reward_)
}

// SetRewardOnce is a paid mutator transaction binding the contract method 0xe78d18e4.
//
// Solidity: function setRewardOnce(uint256 tokenId, int256 reward_) returns()
func (_Nft *NftSession) SetRewardOnce(tokenId *big.Int, reward_ *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetRewardOnce(&_Nft.TransactOpts, tokenId, reward_)
}

// SetRewardOnce is a paid mutator transaction binding the contract method 0xe78d18e4.
//
// Solidity: function setRewardOnce(uint256 tokenId, int256 reward_) returns()
func (_Nft *NftTransactorSession) SetRewardOnce(tokenId *big.Int, reward_ *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetRewardOnce(&_Nft.TransactOpts, tokenId, reward_)
}

// SetStakingContract is a paid mutator transaction binding the contract method 0x9dd373b9.
//
// Solidity: function setStakingContract(address newStaking) returns()
func (_Nft *NftTransactor) SetStakingContract(opts *bind.TransactOpts, newStaking common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setStakingContract", newStaking)
}

// SetStakingContract is a paid mutator transaction binding the contract method 0x9dd373b9.
//
// Solidity: function setStakingContract(address newStaking) returns()
func (_Nft *NftSession) SetStakingContract(newStaking common.Address) (*types.Transaction, error) {
	return _Nft.Contract.SetStakingContract(&_Nft.TransactOpts, newStaking)
}

// SetStakingContract is a paid mutator transaction binding the contract method 0x9dd373b9.
//
// Solidity: function setStakingContract(address newStaking) returns()
func (_Nft *NftTransactorSession) SetStakingContract(newStaking common.Address) (*types.Transaction, error) {
	return _Nft.Contract.SetStakingContract(&_Nft.TransactOpts, newStaking)
}

// SetTierPrice is a paid mutator transaction binding the contract method 0x4a5c6f8c.
//
// Solidity: function setTierPrice(uint8 tier, uint256 newPrice) returns()
func (_Nft *NftTransactor) SetTierPrice(opts *bind.TransactOpts, tier uint8, newPrice *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setTierPrice", tier, newPrice)
}

// SetTierPrice is a paid mutator transaction binding the contract method 0x4a5c6f8c.
//
// Solidity: function setTierPrice(uint8 tier, uint256 newPrice) returns()
func (_Nft *NftSession) SetTierPrice(tier uint8, newPrice *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetTierPrice(&_Nft.TransactOpts, tier, newPrice)
}

// SetTierPrice is a paid mutator transaction binding the contract method 0x4a5c6f8c.
//
// Solidity: function setTierPrice(uint8 tier, uint256 newPrice) returns()
func (_Nft *NftTransactorSession) SetTierPrice(tier uint8, newPrice *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetTierPrice(&_Nft.TransactOpts, tier, newPrice)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Nft *NftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Nft *NftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.TransferFrom(&_Nft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Nft *NftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.TransferFrom(&_Nft.TransactOpts, from, to, tokenId)
}

// Unlist is a paid mutator transaction binding the contract method 0xbf5a4dd3.
//
// Solidity: function unlist(uint256 tokenId) returns()
func (_Nft *NftTransactor) Unlist(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "unlist", tokenId)
}

// Unlist is a paid mutator transaction binding the contract method 0xbf5a4dd3.
//
// Solidity: function unlist(uint256 tokenId) returns()
func (_Nft *NftSession) Unlist(tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.Unlist(&_Nft.TransactOpts, tokenId)
}

// Unlist is a paid mutator transaction binding the contract method 0xbf5a4dd3.
//
// Solidity: function unlist(uint256 tokenId) returns()
func (_Nft *NftTransactorSession) Unlist(tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.Unlist(&_Nft.TransactOpts, tokenId)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address token, address to, uint256 amount) returns()
func (_Nft *NftTransactor) WithdrawERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "withdrawERC20", token, to, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address token, address to, uint256 amount) returns()
func (_Nft *NftSession) WithdrawERC20(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.WithdrawERC20(&_Nft.TransactOpts, token, to, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address token, address to, uint256 amount) returns()
func (_Nft *NftTransactorSession) WithdrawERC20(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.WithdrawERC20(&_Nft.TransactOpts, token, to, amount)
}

// NftATokenChangedIterator is returned from FilterATokenChanged and is used to iterate over the raw logs and unpacked data for ATokenChanged events raised by the Nft contract.
type NftATokenChangedIterator struct {
	Event *NftATokenChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftATokenChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftATokenChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftATokenChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftATokenChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftATokenChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftATokenChanged represents a ATokenChanged event raised by the Nft contract.
type NftATokenChanged struct {
	OldToken common.Address
	NewToken common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterATokenChanged is a free log retrieval operation binding the contract event 0x4f3853b044c37479b3c01138086947d855c93851c189818335d10df28ab9cdd6.
//
// Solidity: event ATokenChanged(address indexed oldToken, address indexed newToken)
func (_Nft *NftFilterer) FilterATokenChanged(opts *bind.FilterOpts, oldToken []common.Address, newToken []common.Address) (*NftATokenChangedIterator, error) {

	var oldTokenRule []interface{}
	for _, oldTokenItem := range oldToken {
		oldTokenRule = append(oldTokenRule, oldTokenItem)
	}
	var newTokenRule []interface{}
	for _, newTokenItem := range newToken {
		newTokenRule = append(newTokenRule, newTokenItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "ATokenChanged", oldTokenRule, newTokenRule)
	if err != nil {
		return nil, err
	}
	return &NftATokenChangedIterator{contract: _Nft.contract, event: "ATokenChanged", logs: logs, sub: sub}, nil
}

// WatchATokenChanged is a free log subscription operation binding the contract event 0x4f3853b044c37479b3c01138086947d855c93851c189818335d10df28ab9cdd6.
//
// Solidity: event ATokenChanged(address indexed oldToken, address indexed newToken)
func (_Nft *NftFilterer) WatchATokenChanged(opts *bind.WatchOpts, sink chan<- *NftATokenChanged, oldToken []common.Address, newToken []common.Address) (event.Subscription, error) {

	var oldTokenRule []interface{}
	for _, oldTokenItem := range oldToken {
		oldTokenRule = append(oldTokenRule, oldTokenItem)
	}
	var newTokenRule []interface{}
	for _, newTokenItem := range newToken {
		newTokenRule = append(newTokenRule, newTokenItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "ATokenChanged", oldTokenRule, newTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftATokenChanged)
				if err := _Nft.contract.UnpackLog(event, "ATokenChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseATokenChanged is a log parse operation binding the contract event 0x4f3853b044c37479b3c01138086947d855c93851c189818335d10df28ab9cdd6.
//
// Solidity: event ATokenChanged(address indexed oldToken, address indexed newToken)
func (_Nft *NftFilterer) ParseATokenChanged(log types.Log) (*NftATokenChanged, error) {
	event := new(NftATokenChanged)
	if err := _Nft.contract.UnpackLog(event, "ATokenChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftAccountAChangedIterator is returned from FilterAccountAChanged and is used to iterate over the raw logs and unpacked data for AccountAChanged events raised by the Nft contract.
type NftAccountAChangedIterator struct {
	Event *NftAccountAChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftAccountAChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftAccountAChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftAccountAChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftAccountAChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftAccountAChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftAccountAChanged represents a AccountAChanged event raised by the Nft contract.
type NftAccountAChanged struct {
	OldAccountA common.Address
	NewAccountA common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAccountAChanged is a free log retrieval operation binding the contract event 0xb44514b5290f12a9b78e7ffbf5849781ed3bd4a82c5e4fa5d098d1bda3f3f9d8.
//
// Solidity: event AccountAChanged(address indexed oldAccountA, address indexed newAccountA)
func (_Nft *NftFilterer) FilterAccountAChanged(opts *bind.FilterOpts, oldAccountA []common.Address, newAccountA []common.Address) (*NftAccountAChangedIterator, error) {

	var oldAccountARule []interface{}
	for _, oldAccountAItem := range oldAccountA {
		oldAccountARule = append(oldAccountARule, oldAccountAItem)
	}
	var newAccountARule []interface{}
	for _, newAccountAItem := range newAccountA {
		newAccountARule = append(newAccountARule, newAccountAItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "AccountAChanged", oldAccountARule, newAccountARule)
	if err != nil {
		return nil, err
	}
	return &NftAccountAChangedIterator{contract: _Nft.contract, event: "AccountAChanged", logs: logs, sub: sub}, nil
}

// WatchAccountAChanged is a free log subscription operation binding the contract event 0xb44514b5290f12a9b78e7ffbf5849781ed3bd4a82c5e4fa5d098d1bda3f3f9d8.
//
// Solidity: event AccountAChanged(address indexed oldAccountA, address indexed newAccountA)
func (_Nft *NftFilterer) WatchAccountAChanged(opts *bind.WatchOpts, sink chan<- *NftAccountAChanged, oldAccountA []common.Address, newAccountA []common.Address) (event.Subscription, error) {

	var oldAccountARule []interface{}
	for _, oldAccountAItem := range oldAccountA {
		oldAccountARule = append(oldAccountARule, oldAccountAItem)
	}
	var newAccountARule []interface{}
	for _, newAccountAItem := range newAccountA {
		newAccountARule = append(newAccountARule, newAccountAItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "AccountAChanged", oldAccountARule, newAccountARule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftAccountAChanged)
				if err := _Nft.contract.UnpackLog(event, "AccountAChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountAChanged is a log parse operation binding the contract event 0xb44514b5290f12a9b78e7ffbf5849781ed3bd4a82c5e4fa5d098d1bda3f3f9d8.
//
// Solidity: event AccountAChanged(address indexed oldAccountA, address indexed newAccountA)
func (_Nft *NftFilterer) ParseAccountAChanged(log types.Log) (*NftAccountAChanged, error) {
	event := new(NftAccountAChanged)
	if err := _Nft.contract.UnpackLog(event, "AccountAChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Nft contract.
type NftApprovalIterator struct {
	Event *NftApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftApproval represents a Approval event raised by the Nft contract.
type NftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Nft *NftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NftApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftApprovalIterator{contract: _Nft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Nft *NftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftApproval)
				if err := _Nft.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Nft *NftFilterer) ParseApproval(log types.Log) (*NftApproval, error) {
	event := new(NftApproval)
	if err := _Nft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Nft contract.
type NftApprovalForAllIterator struct {
	Event *NftApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftApprovalForAll represents a ApprovalForAll event raised by the Nft contract.
type NftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Nft *NftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NftApprovalForAllIterator{contract: _Nft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Nft *NftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftApprovalForAll)
				if err := _Nft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Nft *NftFilterer) ParseApprovalForAll(log types.Log) (*NftApprovalForAll, error) {
	event := new(NftApprovalForAll)
	if err := _Nft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the Nft contract.
type NftBurnedIterator struct {
	Event *NftBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftBurned represents a Burned event raised by the Nft contract.
type NftBurned struct {
	Owner     common.Address
	TokenId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x23ff0e75edf108e3d0392d92e13e8c8a868ef19001bd49f9e94876dc46dff87f.
//
// Solidity: event Burned(address indexed owner, uint256 indexed tokenId, uint256 timestamp)
func (_Nft *NftFilterer) FilterBurned(opts *bind.FilterOpts, owner []common.Address, tokenId []*big.Int) (*NftBurnedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Burned", ownerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftBurnedIterator{contract: _Nft.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x23ff0e75edf108e3d0392d92e13e8c8a868ef19001bd49f9e94876dc46dff87f.
//
// Solidity: event Burned(address indexed owner, uint256 indexed tokenId, uint256 timestamp)
func (_Nft *NftFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *NftBurned, owner []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Burned", ownerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftBurned)
				if err := _Nft.contract.UnpackLog(event, "Burned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurned is a log parse operation binding the contract event 0x23ff0e75edf108e3d0392d92e13e8c8a868ef19001bd49f9e94876dc46dff87f.
//
// Solidity: event Burned(address indexed owner, uint256 indexed tokenId, uint256 timestamp)
func (_Nft *NftFilterer) ParseBurned(log types.Log) (*NftBurned, error) {
	event := new(NftBurned)
	if err := _Nft.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftERC20WithdrawnIterator is returned from FilterERC20Withdrawn and is used to iterate over the raw logs and unpacked data for ERC20Withdrawn events raised by the Nft contract.
type NftERC20WithdrawnIterator struct {
	Event *NftERC20Withdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftERC20WithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftERC20Withdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftERC20Withdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftERC20WithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftERC20WithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftERC20Withdrawn represents a ERC20Withdrawn event raised by the Nft contract.
type NftERC20Withdrawn struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC20Withdrawn is a free log retrieval operation binding the contract event 0xbfed55bdcd242e3dd0f60ddd7d1e87c67f61c34cd9527b3e6455d841b1025362.
//
// Solidity: event ERC20Withdrawn(address indexed token, address indexed to, uint256 amount)
func (_Nft *NftFilterer) FilterERC20Withdrawn(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*NftERC20WithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "ERC20Withdrawn", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NftERC20WithdrawnIterator{contract: _Nft.contract, event: "ERC20Withdrawn", logs: logs, sub: sub}, nil
}

// WatchERC20Withdrawn is a free log subscription operation binding the contract event 0xbfed55bdcd242e3dd0f60ddd7d1e87c67f61c34cd9527b3e6455d841b1025362.
//
// Solidity: event ERC20Withdrawn(address indexed token, address indexed to, uint256 amount)
func (_Nft *NftFilterer) WatchERC20Withdrawn(opts *bind.WatchOpts, sink chan<- *NftERC20Withdrawn, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "ERC20Withdrawn", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftERC20Withdrawn)
				if err := _Nft.contract.UnpackLog(event, "ERC20Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseERC20Withdrawn is a log parse operation binding the contract event 0xbfed55bdcd242e3dd0f60ddd7d1e87c67f61c34cd9527b3e6455d841b1025362.
//
// Solidity: event ERC20Withdrawn(address indexed token, address indexed to, uint256 amount)
func (_Nft *NftFilterer) ParseERC20Withdrawn(log types.Log) (*NftERC20Withdrawn, error) {
	event := new(NftERC20Withdrawn)
	if err := _Nft.contract.UnpackLog(event, "ERC20Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftFeeConfigBChangedIterator is returned from FilterFeeConfigBChanged and is used to iterate over the raw logs and unpacked data for FeeConfigBChanged events raised by the Nft contract.
type NftFeeConfigBChangedIterator struct {
	Event *NftFeeConfigBChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftFeeConfigBChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftFeeConfigBChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftFeeConfigBChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftFeeConfigBChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftFeeConfigBChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftFeeConfigBChanged represents a FeeConfigBChanged event raised by the Nft contract.
type NftFeeConfigBChanged struct {
	RateTwo *big.Int
	BaseTwo *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeConfigBChanged is a free log retrieval operation binding the contract event 0x1dea14f468e41a49c408d03736cf0cac2a5b408cf24f37228bdad4a98e984220.
//
// Solidity: event FeeConfigBChanged(uint256 rateTwo, uint256 baseTwo)
func (_Nft *NftFilterer) FilterFeeConfigBChanged(opts *bind.FilterOpts) (*NftFeeConfigBChangedIterator, error) {

	logs, sub, err := _Nft.contract.FilterLogs(opts, "FeeConfigBChanged")
	if err != nil {
		return nil, err
	}
	return &NftFeeConfigBChangedIterator{contract: _Nft.contract, event: "FeeConfigBChanged", logs: logs, sub: sub}, nil
}

// WatchFeeConfigBChanged is a free log subscription operation binding the contract event 0x1dea14f468e41a49c408d03736cf0cac2a5b408cf24f37228bdad4a98e984220.
//
// Solidity: event FeeConfigBChanged(uint256 rateTwo, uint256 baseTwo)
func (_Nft *NftFilterer) WatchFeeConfigBChanged(opts *bind.WatchOpts, sink chan<- *NftFeeConfigBChanged) (event.Subscription, error) {

	logs, sub, err := _Nft.contract.WatchLogs(opts, "FeeConfigBChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftFeeConfigBChanged)
				if err := _Nft.contract.UnpackLog(event, "FeeConfigBChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeConfigBChanged is a log parse operation binding the contract event 0x1dea14f468e41a49c408d03736cf0cac2a5b408cf24f37228bdad4a98e984220.
//
// Solidity: event FeeConfigBChanged(uint256 rateTwo, uint256 baseTwo)
func (_Nft *NftFilterer) ParseFeeConfigBChanged(log types.Log) (*NftFeeConfigBChanged, error) {
	event := new(NftFeeConfigBChanged)
	if err := _Nft.contract.UnpackLog(event, "FeeConfigBChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftFeeConfigUSDTChangedIterator is returned from FilterFeeConfigUSDTChanged and is used to iterate over the raw logs and unpacked data for FeeConfigUSDTChanged events raised by the Nft contract.
type NftFeeConfigUSDTChangedIterator struct {
	Event *NftFeeConfigUSDTChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftFeeConfigUSDTChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftFeeConfigUSDTChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftFeeConfigUSDTChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftFeeConfigUSDTChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftFeeConfigUSDTChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftFeeConfigUSDTChanged represents a FeeConfigUSDTChanged event raised by the Nft contract.
type NftFeeConfigUSDTChanged struct {
	Rate *big.Int
	Base *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFeeConfigUSDTChanged is a free log retrieval operation binding the contract event 0x1a91e08ce5bde2175a1eb46d7203f4716d772b59bb4d3b32690bf72e43eada08.
//
// Solidity: event FeeConfigUSDTChanged(uint256 rate, uint256 base)
func (_Nft *NftFilterer) FilterFeeConfigUSDTChanged(opts *bind.FilterOpts) (*NftFeeConfigUSDTChangedIterator, error) {

	logs, sub, err := _Nft.contract.FilterLogs(opts, "FeeConfigUSDTChanged")
	if err != nil {
		return nil, err
	}
	return &NftFeeConfigUSDTChangedIterator{contract: _Nft.contract, event: "FeeConfigUSDTChanged", logs: logs, sub: sub}, nil
}

// WatchFeeConfigUSDTChanged is a free log subscription operation binding the contract event 0x1a91e08ce5bde2175a1eb46d7203f4716d772b59bb4d3b32690bf72e43eada08.
//
// Solidity: event FeeConfigUSDTChanged(uint256 rate, uint256 base)
func (_Nft *NftFilterer) WatchFeeConfigUSDTChanged(opts *bind.WatchOpts, sink chan<- *NftFeeConfigUSDTChanged) (event.Subscription, error) {

	logs, sub, err := _Nft.contract.WatchLogs(opts, "FeeConfigUSDTChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftFeeConfigUSDTChanged)
				if err := _Nft.contract.UnpackLog(event, "FeeConfigUSDTChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeConfigUSDTChanged is a log parse operation binding the contract event 0x1a91e08ce5bde2175a1eb46d7203f4716d772b59bb4d3b32690bf72e43eada08.
//
// Solidity: event FeeConfigUSDTChanged(uint256 rate, uint256 base)
func (_Nft *NftFilterer) ParseFeeConfigUSDTChanged(log types.Log) (*NftFeeConfigUSDTChanged, error) {
	event := new(NftFeeConfigUSDTChanged)
	if err := _Nft.contract.UnpackLog(event, "FeeConfigUSDTChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftFeeQuotePoolChangedIterator is returned from FilterFeeQuotePoolChanged and is used to iterate over the raw logs and unpacked data for FeeQuotePoolChanged events raised by the Nft contract.
type NftFeeQuotePoolChangedIterator struct {
	Event *NftFeeQuotePoolChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftFeeQuotePoolChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftFeeQuotePoolChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftFeeQuotePoolChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftFeeQuotePoolChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftFeeQuotePoolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftFeeQuotePoolChanged represents a FeeQuotePoolChanged event raised by the Nft contract.
type NftFeeQuotePoolChanged struct {
	Pool    common.Address
	Token0B common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeQuotePoolChanged is a free log retrieval operation binding the contract event 0x0e2c82b5a1dfc893709c58ba5b3657d0c588b7016dec3a20313d6f5b5a31f971.
//
// Solidity: event FeeQuotePoolChanged(address indexed pool, address indexed token0B)
func (_Nft *NftFilterer) FilterFeeQuotePoolChanged(opts *bind.FilterOpts, pool []common.Address, token0B []common.Address) (*NftFeeQuotePoolChangedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var token0BRule []interface{}
	for _, token0BItem := range token0B {
		token0BRule = append(token0BRule, token0BItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "FeeQuotePoolChanged", poolRule, token0BRule)
	if err != nil {
		return nil, err
	}
	return &NftFeeQuotePoolChangedIterator{contract: _Nft.contract, event: "FeeQuotePoolChanged", logs: logs, sub: sub}, nil
}

// WatchFeeQuotePoolChanged is a free log subscription operation binding the contract event 0x0e2c82b5a1dfc893709c58ba5b3657d0c588b7016dec3a20313d6f5b5a31f971.
//
// Solidity: event FeeQuotePoolChanged(address indexed pool, address indexed token0B)
func (_Nft *NftFilterer) WatchFeeQuotePoolChanged(opts *bind.WatchOpts, sink chan<- *NftFeeQuotePoolChanged, pool []common.Address, token0B []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var token0BRule []interface{}
	for _, token0BItem := range token0B {
		token0BRule = append(token0BRule, token0BItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "FeeQuotePoolChanged", poolRule, token0BRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftFeeQuotePoolChanged)
				if err := _Nft.contract.UnpackLog(event, "FeeQuotePoolChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeQuotePoolChanged is a log parse operation binding the contract event 0x0e2c82b5a1dfc893709c58ba5b3657d0c588b7016dec3a20313d6f5b5a31f971.
//
// Solidity: event FeeQuotePoolChanged(address indexed pool, address indexed token0B)
func (_Nft *NftFilterer) ParseFeeQuotePoolChanged(log types.Log) (*NftFeeQuotePoolChanged, error) {
	event := new(NftFeeQuotePoolChanged)
	if err := _Nft.contract.UnpackLog(event, "FeeQuotePoolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftGlobalSlotStepCostChangedIterator is returned from FilterGlobalSlotStepCostChanged and is used to iterate over the raw logs and unpacked data for GlobalSlotStepCostChanged events raised by the Nft contract.
type NftGlobalSlotStepCostChangedIterator struct {
	Event *NftGlobalSlotStepCostChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftGlobalSlotStepCostChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftGlobalSlotStepCostChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftGlobalSlotStepCostChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftGlobalSlotStepCostChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftGlobalSlotStepCostChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftGlobalSlotStepCostChanged represents a GlobalSlotStepCostChanged event raised by the Nft contract.
type NftGlobalSlotStepCostChanged struct {
	OldCost *big.Int
	NewCost *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGlobalSlotStepCostChanged is a free log retrieval operation binding the contract event 0x3ed52a2bd772aa485432d718581b74559c35af88dbbe26b28522da8cede4926e.
//
// Solidity: event GlobalSlotStepCostChanged(uint256 oldCost, uint256 newCost)
func (_Nft *NftFilterer) FilterGlobalSlotStepCostChanged(opts *bind.FilterOpts) (*NftGlobalSlotStepCostChangedIterator, error) {

	logs, sub, err := _Nft.contract.FilterLogs(opts, "GlobalSlotStepCostChanged")
	if err != nil {
		return nil, err
	}
	return &NftGlobalSlotStepCostChangedIterator{contract: _Nft.contract, event: "GlobalSlotStepCostChanged", logs: logs, sub: sub}, nil
}

// WatchGlobalSlotStepCostChanged is a free log subscription operation binding the contract event 0x3ed52a2bd772aa485432d718581b74559c35af88dbbe26b28522da8cede4926e.
//
// Solidity: event GlobalSlotStepCostChanged(uint256 oldCost, uint256 newCost)
func (_Nft *NftFilterer) WatchGlobalSlotStepCostChanged(opts *bind.WatchOpts, sink chan<- *NftGlobalSlotStepCostChanged) (event.Subscription, error) {

	logs, sub, err := _Nft.contract.WatchLogs(opts, "GlobalSlotStepCostChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftGlobalSlotStepCostChanged)
				if err := _Nft.contract.UnpackLog(event, "GlobalSlotStepCostChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGlobalSlotStepCostChanged is a log parse operation binding the contract event 0x3ed52a2bd772aa485432d718581b74559c35af88dbbe26b28522da8cede4926e.
//
// Solidity: event GlobalSlotStepCostChanged(uint256 oldCost, uint256 newCost)
func (_Nft *NftFilterer) ParseGlobalSlotStepCostChanged(log types.Log) (*NftGlobalSlotStepCostChanged, error) {
	event := new(NftGlobalSlotStepCostChanged)
	if err := _Nft.contract.UnpackLog(event, "GlobalSlotStepCostChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftListedIterator is returned from FilterListed and is used to iterate over the raw logs and unpacked data for Listed events raised by the Nft contract.
type NftListedIterator struct {
	Event *NftListed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftListed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftListed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftListed represents a Listed event raised by the Nft contract.
type NftListed struct {
	Seller    common.Address
	TokenId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterListed is a free log retrieval operation binding the contract event 0x6e16e8064e3c8254cef11e01e152416398602d7e8785f743f1c2f93a9a15ff35.
//
// Solidity: event Listed(address indexed seller, uint256 indexed tokenId, uint256 timestamp)
func (_Nft *NftFilterer) FilterListed(opts *bind.FilterOpts, seller []common.Address, tokenId []*big.Int) (*NftListedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Listed", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftListedIterator{contract: _Nft.contract, event: "Listed", logs: logs, sub: sub}, nil
}

// WatchListed is a free log subscription operation binding the contract event 0x6e16e8064e3c8254cef11e01e152416398602d7e8785f743f1c2f93a9a15ff35.
//
// Solidity: event Listed(address indexed seller, uint256 indexed tokenId, uint256 timestamp)
func (_Nft *NftFilterer) WatchListed(opts *bind.WatchOpts, sink chan<- *NftListed, seller []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Listed", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftListed)
				if err := _Nft.contract.UnpackLog(event, "Listed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseListed is a log parse operation binding the contract event 0x6e16e8064e3c8254cef11e01e152416398602d7e8785f743f1c2f93a9a15ff35.
//
// Solidity: event Listed(address indexed seller, uint256 indexed tokenId, uint256 timestamp)
func (_Nft *NftFilterer) ParseListed(log types.Log) (*NftListed, error) {
	event := new(NftListed)
	if err := _Nft.contract.UnpackLog(event, "Listed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the Nft contract.
type NftMintedIterator struct {
	Event *NftMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftMinted represents a Minted event raised by the Nft contract.
type NftMinted struct {
	Minter   common.Address
	TokenId  *big.Int
	Tier     uint8
	UsdtPaid *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x0fcee5463734b554d0accfca29a4811ecd45e6c12b1cee0cb22d8d8eb3be1c33.
//
// Solidity: event Minted(address indexed minter, uint256 indexed tokenId, uint8 indexed tier, uint256 usdtPaid)
func (_Nft *NftFilterer) FilterMinted(opts *bind.FilterOpts, minter []common.Address, tokenId []*big.Int, tier []uint8) (*NftMintedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Minted", minterRule, tokenIdRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &NftMintedIterator{contract: _Nft.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x0fcee5463734b554d0accfca29a4811ecd45e6c12b1cee0cb22d8d8eb3be1c33.
//
// Solidity: event Minted(address indexed minter, uint256 indexed tokenId, uint8 indexed tier, uint256 usdtPaid)
func (_Nft *NftFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *NftMinted, minter []common.Address, tokenId []*big.Int, tier []uint8) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Minted", minterRule, tokenIdRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftMinted)
				if err := _Nft.contract.UnpackLog(event, "Minted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinted is a log parse operation binding the contract event 0x0fcee5463734b554d0accfca29a4811ecd45e6c12b1cee0cb22d8d8eb3be1c33.
//
// Solidity: event Minted(address indexed minter, uint256 indexed tokenId, uint8 indexed tier, uint256 usdtPaid)
func (_Nft *NftFilterer) ParseMinted(log types.Log) (*NftMinted, error) {
	event := new(NftMinted)
	if err := _Nft.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftOpenFeeConfigChangedIterator is returned from FilterOpenFeeConfigChanged and is used to iterate over the raw logs and unpacked data for OpenFeeConfigChanged events raised by the Nft contract.
type NftOpenFeeConfigChangedIterator struct {
	Event *NftOpenFeeConfigChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftOpenFeeConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftOpenFeeConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftOpenFeeConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftOpenFeeConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftOpenFeeConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftOpenFeeConfigChanged represents a OpenFeeConfigChanged event raised by the Nft contract.
type NftOpenFeeConfigChanged struct {
	OpenFeeBNB      *big.Int
	OpenFeeReceiver common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOpenFeeConfigChanged is a free log retrieval operation binding the contract event 0xe0e59605d0066759eb9ef5a7d5c286244c870fb42449cbcefb448d2836899a95.
//
// Solidity: event OpenFeeConfigChanged(uint256 openFeeBNB, address openFeeReceiver)
func (_Nft *NftFilterer) FilterOpenFeeConfigChanged(opts *bind.FilterOpts) (*NftOpenFeeConfigChangedIterator, error) {

	logs, sub, err := _Nft.contract.FilterLogs(opts, "OpenFeeConfigChanged")
	if err != nil {
		return nil, err
	}
	return &NftOpenFeeConfigChangedIterator{contract: _Nft.contract, event: "OpenFeeConfigChanged", logs: logs, sub: sub}, nil
}

// WatchOpenFeeConfigChanged is a free log subscription operation binding the contract event 0xe0e59605d0066759eb9ef5a7d5c286244c870fb42449cbcefb448d2836899a95.
//
// Solidity: event OpenFeeConfigChanged(uint256 openFeeBNB, address openFeeReceiver)
func (_Nft *NftFilterer) WatchOpenFeeConfigChanged(opts *bind.WatchOpts, sink chan<- *NftOpenFeeConfigChanged) (event.Subscription, error) {

	logs, sub, err := _Nft.contract.WatchLogs(opts, "OpenFeeConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftOpenFeeConfigChanged)
				if err := _Nft.contract.UnpackLog(event, "OpenFeeConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOpenFeeConfigChanged is a log parse operation binding the contract event 0xe0e59605d0066759eb9ef5a7d5c286244c870fb42449cbcefb448d2836899a95.
//
// Solidity: event OpenFeeConfigChanged(uint256 openFeeBNB, address openFeeReceiver)
func (_Nft *NftFilterer) ParseOpenFeeConfigChanged(log types.Log) (*NftOpenFeeConfigChanged, error) {
	event := new(NftOpenFeeConfigChanged)
	if err := _Nft.contract.UnpackLog(event, "OpenFeeConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftOpenedIterator is returned from FilterOpened and is used to iterate over the raw logs and unpacked data for Opened events raised by the Nft contract.
type NftOpenedIterator struct {
	Event *NftOpened // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftOpened)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftOpened)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftOpened represents a Opened event raised by the Nft contract.
type NftOpened struct {
	Owner     common.Address
	TokenId   *big.Int
	Timestamp *big.Int
	BnbPaid   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOpened is a free log retrieval operation binding the contract event 0x1ae84d4a14bea8146749d9831e35557d4318d019c88ef77238907396f34fb38e.
//
// Solidity: event Opened(address indexed owner, uint256 indexed tokenId, uint256 timestamp, uint256 bnbPaid)
func (_Nft *NftFilterer) FilterOpened(opts *bind.FilterOpts, owner []common.Address, tokenId []*big.Int) (*NftOpenedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Opened", ownerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftOpenedIterator{contract: _Nft.contract, event: "Opened", logs: logs, sub: sub}, nil
}

// WatchOpened is a free log subscription operation binding the contract event 0x1ae84d4a14bea8146749d9831e35557d4318d019c88ef77238907396f34fb38e.
//
// Solidity: event Opened(address indexed owner, uint256 indexed tokenId, uint256 timestamp, uint256 bnbPaid)
func (_Nft *NftFilterer) WatchOpened(opts *bind.WatchOpts, sink chan<- *NftOpened, owner []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Opened", ownerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftOpened)
				if err := _Nft.contract.UnpackLog(event, "Opened", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOpened is a log parse operation binding the contract event 0x1ae84d4a14bea8146749d9831e35557d4318d019c88ef77238907396f34fb38e.
//
// Solidity: event Opened(address indexed owner, uint256 indexed tokenId, uint256 timestamp, uint256 bnbPaid)
func (_Nft *NftFilterer) ParseOpened(log types.Log) (*NftOpened, error) {
	event := new(NftOpened)
	if err := _Nft.contract.UnpackLog(event, "Opened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftParamsChangedIterator is returned from FilterParamsChanged and is used to iterate over the raw logs and unpacked data for ParamsChanged events raised by the Nft contract.
type NftParamsChangedIterator struct {
	Event *NftParamsChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftParamsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftParamsChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftParamsChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftParamsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftParamsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftParamsChanged represents a ParamsChanged event raised by the Nft contract.
type NftParamsChanged struct {
	BpsDenominator        *big.Int
	AccountAFeeBps        *big.Int
	MonthlyRateDefaultE18 *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterParamsChanged is a free log retrieval operation binding the contract event 0xc67740ec6b8e01e85235967c5031201015f6bc252124415197f443c6d48a2a8f.
//
// Solidity: event ParamsChanged(uint256 bpsDenominator, uint256 accountAFeeBps, uint256 monthlyRateDefaultE18)
func (_Nft *NftFilterer) FilterParamsChanged(opts *bind.FilterOpts) (*NftParamsChangedIterator, error) {

	logs, sub, err := _Nft.contract.FilterLogs(opts, "ParamsChanged")
	if err != nil {
		return nil, err
	}
	return &NftParamsChangedIterator{contract: _Nft.contract, event: "ParamsChanged", logs: logs, sub: sub}, nil
}

// WatchParamsChanged is a free log subscription operation binding the contract event 0xc67740ec6b8e01e85235967c5031201015f6bc252124415197f443c6d48a2a8f.
//
// Solidity: event ParamsChanged(uint256 bpsDenominator, uint256 accountAFeeBps, uint256 monthlyRateDefaultE18)
func (_Nft *NftFilterer) WatchParamsChanged(opts *bind.WatchOpts, sink chan<- *NftParamsChanged) (event.Subscription, error) {

	logs, sub, err := _Nft.contract.WatchLogs(opts, "ParamsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftParamsChanged)
				if err := _Nft.contract.UnpackLog(event, "ParamsChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParamsChanged is a log parse operation binding the contract event 0xc67740ec6b8e01e85235967c5031201015f6bc252124415197f443c6d48a2a8f.
//
// Solidity: event ParamsChanged(uint256 bpsDenominator, uint256 accountAFeeBps, uint256 monthlyRateDefaultE18)
func (_Nft *NftFilterer) ParseParamsChanged(log types.Log) (*NftParamsChanged, error) {
	event := new(NftParamsChanged)
	if err := _Nft.contract.UnpackLog(event, "ParamsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftPieceHoldingAddedIterator is returned from FilterPieceHoldingAdded and is used to iterate over the raw logs and unpacked data for PieceHoldingAdded events raised by the Nft contract.
type NftPieceHoldingAddedIterator struct {
	Event *NftPieceHoldingAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftPieceHoldingAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftPieceHoldingAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftPieceHoldingAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftPieceHoldingAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftPieceHoldingAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftPieceHoldingAdded represents a PieceHoldingAdded event raised by the Nft contract.
type NftPieceHoldingAdded struct {
	User                 common.Address
	BurnAmount           *big.Int
	ValueAddedUsdt       *big.Int
	NewPieceHoldingValue *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPieceHoldingAdded is a free log retrieval operation binding the contract event 0x170bb387490bfead43cc0280c4cc0b6cc8c32e3e1e91b40645b000933d547b1c.
//
// Solidity: event PieceHoldingAdded(address indexed user, uint256 burnAmount, uint256 valueAddedUsdt, uint256 newPieceHoldingValue)
func (_Nft *NftFilterer) FilterPieceHoldingAdded(opts *bind.FilterOpts, user []common.Address) (*NftPieceHoldingAddedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "PieceHoldingAdded", userRule)
	if err != nil {
		return nil, err
	}
	return &NftPieceHoldingAddedIterator{contract: _Nft.contract, event: "PieceHoldingAdded", logs: logs, sub: sub}, nil
}

// WatchPieceHoldingAdded is a free log subscription operation binding the contract event 0x170bb387490bfead43cc0280c4cc0b6cc8c32e3e1e91b40645b000933d547b1c.
//
// Solidity: event PieceHoldingAdded(address indexed user, uint256 burnAmount, uint256 valueAddedUsdt, uint256 newPieceHoldingValue)
func (_Nft *NftFilterer) WatchPieceHoldingAdded(opts *bind.WatchOpts, sink chan<- *NftPieceHoldingAdded, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "PieceHoldingAdded", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftPieceHoldingAdded)
				if err := _Nft.contract.UnpackLog(event, "PieceHoldingAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePieceHoldingAdded is a log parse operation binding the contract event 0x170bb387490bfead43cc0280c4cc0b6cc8c32e3e1e91b40645b000933d547b1c.
//
// Solidity: event PieceHoldingAdded(address indexed user, uint256 burnAmount, uint256 valueAddedUsdt, uint256 newPieceHoldingValue)
func (_Nft *NftFilterer) ParsePieceHoldingAdded(log types.Log) (*NftPieceHoldingAdded, error) {
	event := new(NftPieceHoldingAdded)
	if err := _Nft.contract.UnpackLog(event, "PieceHoldingAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftPieceTokenChangedIterator is returned from FilterPieceTokenChanged and is used to iterate over the raw logs and unpacked data for PieceTokenChanged events raised by the Nft contract.
type NftPieceTokenChangedIterator struct {
	Event *NftPieceTokenChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftPieceTokenChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftPieceTokenChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftPieceTokenChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftPieceTokenChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftPieceTokenChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftPieceTokenChanged represents a PieceTokenChanged event raised by the Nft contract.
type NftPieceTokenChanged struct {
	OldPiece common.Address
	NewPiece common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPieceTokenChanged is a free log retrieval operation binding the contract event 0x03bc23adb4580907fc84238351b8c9a1079e0e17f10e83910b2ab2591b0b7831.
//
// Solidity: event PieceTokenChanged(address indexed oldPiece, address indexed newPiece)
func (_Nft *NftFilterer) FilterPieceTokenChanged(opts *bind.FilterOpts, oldPiece []common.Address, newPiece []common.Address) (*NftPieceTokenChangedIterator, error) {

	var oldPieceRule []interface{}
	for _, oldPieceItem := range oldPiece {
		oldPieceRule = append(oldPieceRule, oldPieceItem)
	}
	var newPieceRule []interface{}
	for _, newPieceItem := range newPiece {
		newPieceRule = append(newPieceRule, newPieceItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "PieceTokenChanged", oldPieceRule, newPieceRule)
	if err != nil {
		return nil, err
	}
	return &NftPieceTokenChangedIterator{contract: _Nft.contract, event: "PieceTokenChanged", logs: logs, sub: sub}, nil
}

// WatchPieceTokenChanged is a free log subscription operation binding the contract event 0x03bc23adb4580907fc84238351b8c9a1079e0e17f10e83910b2ab2591b0b7831.
//
// Solidity: event PieceTokenChanged(address indexed oldPiece, address indexed newPiece)
func (_Nft *NftFilterer) WatchPieceTokenChanged(opts *bind.WatchOpts, sink chan<- *NftPieceTokenChanged, oldPiece []common.Address, newPiece []common.Address) (event.Subscription, error) {

	var oldPieceRule []interface{}
	for _, oldPieceItem := range oldPiece {
		oldPieceRule = append(oldPieceRule, oldPieceItem)
	}
	var newPieceRule []interface{}
	for _, newPieceItem := range newPiece {
		newPieceRule = append(newPieceRule, newPieceItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "PieceTokenChanged", oldPieceRule, newPieceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftPieceTokenChanged)
				if err := _Nft.contract.UnpackLog(event, "PieceTokenChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePieceTokenChanged is a log parse operation binding the contract event 0x03bc23adb4580907fc84238351b8c9a1079e0e17f10e83910b2ab2591b0b7831.
//
// Solidity: event PieceTokenChanged(address indexed oldPiece, address indexed newPiece)
func (_Nft *NftFilterer) ParsePieceTokenChanged(log types.Log) (*NftPieceTokenChanged, error) {
	event := new(NftPieceTokenChanged)
	if err := _Nft.contract.UnpackLog(event, "PieceTokenChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftPieceValueChangedIterator is returned from FilterPieceValueChanged and is used to iterate over the raw logs and unpacked data for PieceValueChanged events raised by the Nft contract.
type NftPieceValueChangedIterator struct {
	Event *NftPieceValueChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftPieceValueChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftPieceValueChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftPieceValueChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftPieceValueChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftPieceValueChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftPieceValueChanged represents a PieceValueChanged event raised by the Nft contract.
type NftPieceValueChanged struct {
	OldValueUsdtE18 *big.Int
	NewValueUsdtE18 *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPieceValueChanged is a free log retrieval operation binding the contract event 0x8cac28b5de93b812c51392809a662b730d9b80efa707e864074fd8ceb706b054.
//
// Solidity: event PieceValueChanged(uint256 oldValueUsdtE18, uint256 newValueUsdtE18)
func (_Nft *NftFilterer) FilterPieceValueChanged(opts *bind.FilterOpts) (*NftPieceValueChangedIterator, error) {

	logs, sub, err := _Nft.contract.FilterLogs(opts, "PieceValueChanged")
	if err != nil {
		return nil, err
	}
	return &NftPieceValueChangedIterator{contract: _Nft.contract, event: "PieceValueChanged", logs: logs, sub: sub}, nil
}

// WatchPieceValueChanged is a free log subscription operation binding the contract event 0x8cac28b5de93b812c51392809a662b730d9b80efa707e864074fd8ceb706b054.
//
// Solidity: event PieceValueChanged(uint256 oldValueUsdtE18, uint256 newValueUsdtE18)
func (_Nft *NftFilterer) WatchPieceValueChanged(opts *bind.WatchOpts, sink chan<- *NftPieceValueChanged) (event.Subscription, error) {

	logs, sub, err := _Nft.contract.WatchLogs(opts, "PieceValueChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftPieceValueChanged)
				if err := _Nft.contract.UnpackLog(event, "PieceValueChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePieceValueChanged is a log parse operation binding the contract event 0x8cac28b5de93b812c51392809a662b730d9b80efa707e864074fd8ceb706b054.
//
// Solidity: event PieceValueChanged(uint256 oldValueUsdtE18, uint256 newValueUsdtE18)
func (_Nft *NftFilterer) ParsePieceValueChanged(log types.Log) (*NftPieceValueChanged, error) {
	event := new(NftPieceValueChanged)
	if err := _Nft.contract.UnpackLog(event, "PieceValueChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftPurchasedIterator is returned from FilterPurchased and is used to iterate over the raw logs and unpacked data for Purchased events raised by the Nft contract.
type NftPurchasedIterator struct {
	Event *NftPurchased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftPurchased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftPurchased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftPurchased represents a Purchased event raised by the Nft contract.
type NftPurchased struct {
	Buyer      common.Address
	Seller     common.Address
	TokenId    *big.Int
	PriceUSDT  *big.Int
	FeePaidInB bool
	FeeUSDT    *big.Int
	FeeB       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPurchased is a free log retrieval operation binding the contract event 0x975f20bb07cab97bb279c9ff5a165f88871f0a2f1d7a35e8ae4f423a253abcda.
//
// Solidity: event Purchased(address indexed buyer, address indexed seller, uint256 indexed tokenId, uint256 priceUSDT, bool feePaidInB, uint256 feeUSDT, uint256 feeB)
func (_Nft *NftFilterer) FilterPurchased(opts *bind.FilterOpts, buyer []common.Address, seller []common.Address, tokenId []*big.Int) (*NftPurchasedIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Purchased", buyerRule, sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftPurchasedIterator{contract: _Nft.contract, event: "Purchased", logs: logs, sub: sub}, nil
}

// WatchPurchased is a free log subscription operation binding the contract event 0x975f20bb07cab97bb279c9ff5a165f88871f0a2f1d7a35e8ae4f423a253abcda.
//
// Solidity: event Purchased(address indexed buyer, address indexed seller, uint256 indexed tokenId, uint256 priceUSDT, bool feePaidInB, uint256 feeUSDT, uint256 feeB)
func (_Nft *NftFilterer) WatchPurchased(opts *bind.WatchOpts, sink chan<- *NftPurchased, buyer []common.Address, seller []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Purchased", buyerRule, sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftPurchased)
				if err := _Nft.contract.UnpackLog(event, "Purchased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePurchased is a log parse operation binding the contract event 0x975f20bb07cab97bb279c9ff5a165f88871f0a2f1d7a35e8ae4f423a253abcda.
//
// Solidity: event Purchased(address indexed buyer, address indexed seller, uint256 indexed tokenId, uint256 priceUSDT, bool feePaidInB, uint256 feeUSDT, uint256 feeB)
func (_Nft *NftFilterer) ParsePurchased(log types.Log) (*NftPurchased, error) {
	event := new(NftPurchased)
	if err := _Nft.contract.UnpackLog(event, "Purchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftRewardSetOnceIterator is returned from FilterRewardSetOnce and is used to iterate over the raw logs and unpacked data for RewardSetOnce events raised by the Nft contract.
type NftRewardSetOnceIterator struct {
	Event *NftRewardSetOnce // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftRewardSetOnceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftRewardSetOnce)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftRewardSetOnce)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftRewardSetOnceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftRewardSetOnceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftRewardSetOnce represents a RewardSetOnce event raised by the Nft contract.
type NftRewardSetOnce struct {
	Admin          common.Address
	TokenId        *big.Int
	Reward         *big.Int
	Timestamp      *big.Int
	ProfitNotified *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardSetOnce is a free log retrieval operation binding the contract event 0x34c1d818a4247ed2f506dfa14100369260c0fc65727b0a48c96ff28119daad7b.
//
// Solidity: event RewardSetOnce(address indexed admin, uint256 indexed tokenId, int256 reward, uint256 timestamp, uint256 profitNotified)
func (_Nft *NftFilterer) FilterRewardSetOnce(opts *bind.FilterOpts, admin []common.Address, tokenId []*big.Int) (*NftRewardSetOnceIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "RewardSetOnce", adminRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftRewardSetOnceIterator{contract: _Nft.contract, event: "RewardSetOnce", logs: logs, sub: sub}, nil
}

// WatchRewardSetOnce is a free log subscription operation binding the contract event 0x34c1d818a4247ed2f506dfa14100369260c0fc65727b0a48c96ff28119daad7b.
//
// Solidity: event RewardSetOnce(address indexed admin, uint256 indexed tokenId, int256 reward, uint256 timestamp, uint256 profitNotified)
func (_Nft *NftFilterer) WatchRewardSetOnce(opts *bind.WatchOpts, sink chan<- *NftRewardSetOnce, admin []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "RewardSetOnce", adminRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftRewardSetOnce)
				if err := _Nft.contract.UnpackLog(event, "RewardSetOnce", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardSetOnce is a log parse operation binding the contract event 0x34c1d818a4247ed2f506dfa14100369260c0fc65727b0a48c96ff28119daad7b.
//
// Solidity: event RewardSetOnce(address indexed admin, uint256 indexed tokenId, int256 reward, uint256 timestamp, uint256 profitNotified)
func (_Nft *NftFilterer) ParseRewardSetOnce(log types.Log) (*NftRewardSetOnce, error) {
	event := new(NftRewardSetOnce)
	if err := _Nft.contract.UnpackLog(event, "RewardSetOnce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Nft contract.
type NftRoleAdminChangedIterator struct {
	Event *NftRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftRoleAdminChanged represents a RoleAdminChanged event raised by the Nft contract.
type NftRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Nft *NftFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NftRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NftRoleAdminChangedIterator{contract: _Nft.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Nft *NftFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NftRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftRoleAdminChanged)
				if err := _Nft.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Nft *NftFilterer) ParseRoleAdminChanged(log types.Log) (*NftRoleAdminChanged, error) {
	event := new(NftRoleAdminChanged)
	if err := _Nft.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Nft contract.
type NftRoleGrantedIterator struct {
	Event *NftRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftRoleGranted represents a RoleGranted event raised by the Nft contract.
type NftRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Nft *NftFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NftRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NftRoleGrantedIterator{contract: _Nft.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Nft *NftFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NftRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftRoleGranted)
				if err := _Nft.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Nft *NftFilterer) ParseRoleGranted(log types.Log) (*NftRoleGranted, error) {
	event := new(NftRoleGranted)
	if err := _Nft.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Nft contract.
type NftRoleRevokedIterator struct {
	Event *NftRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftRoleRevoked represents a RoleRevoked event raised by the Nft contract.
type NftRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Nft *NftFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NftRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NftRoleRevokedIterator{contract: _Nft.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Nft *NftFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NftRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftRoleRevoked)
				if err := _Nft.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Nft *NftFilterer) ParseRoleRevoked(log types.Log) (*NftRoleRevoked, error) {
	event := new(NftRoleRevoked)
	if err := _Nft.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftSlotsPurchasedIterator is returned from FilterSlotsPurchased and is used to iterate over the raw logs and unpacked data for SlotsPurchased events raised by the Nft contract.
type NftSlotsPurchasedIterator struct {
	Event *NftSlotsPurchased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftSlotsPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftSlotsPurchased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftSlotsPurchased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftSlotsPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftSlotsPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftSlotsPurchased represents a SlotsPurchased event raised by the Nft contract.
type NftSlotsPurchased struct {
	Buyer          common.Address
	Count          *big.Int
	StepCostLocked *big.Int
	TotalCost      *big.Int
	NewExtraSlots  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSlotsPurchased is a free log retrieval operation binding the contract event 0x71edde32e1af18692bf2a15a2fa1474806b883f05d27aefc7b04fad08d9d4386.
//
// Solidity: event SlotsPurchased(address indexed buyer, uint256 count, uint256 stepCostLocked, uint256 totalCost, uint256 newExtraSlots)
func (_Nft *NftFilterer) FilterSlotsPurchased(opts *bind.FilterOpts, buyer []common.Address) (*NftSlotsPurchasedIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "SlotsPurchased", buyerRule)
	if err != nil {
		return nil, err
	}
	return &NftSlotsPurchasedIterator{contract: _Nft.contract, event: "SlotsPurchased", logs: logs, sub: sub}, nil
}

// WatchSlotsPurchased is a free log subscription operation binding the contract event 0x71edde32e1af18692bf2a15a2fa1474806b883f05d27aefc7b04fad08d9d4386.
//
// Solidity: event SlotsPurchased(address indexed buyer, uint256 count, uint256 stepCostLocked, uint256 totalCost, uint256 newExtraSlots)
func (_Nft *NftFilterer) WatchSlotsPurchased(opts *bind.WatchOpts, sink chan<- *NftSlotsPurchased, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "SlotsPurchased", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftSlotsPurchased)
				if err := _Nft.contract.UnpackLog(event, "SlotsPurchased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlotsPurchased is a log parse operation binding the contract event 0x71edde32e1af18692bf2a15a2fa1474806b883f05d27aefc7b04fad08d9d4386.
//
// Solidity: event SlotsPurchased(address indexed buyer, uint256 count, uint256 stepCostLocked, uint256 totalCost, uint256 newExtraSlots)
func (_Nft *NftFilterer) ParseSlotsPurchased(log types.Log) (*NftSlotsPurchased, error) {
	event := new(NftSlotsPurchased)
	if err := _Nft.contract.UnpackLog(event, "SlotsPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftStakingContractChangedIterator is returned from FilterStakingContractChanged and is used to iterate over the raw logs and unpacked data for StakingContractChanged events raised by the Nft contract.
type NftStakingContractChangedIterator struct {
	Event *NftStakingContractChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftStakingContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftStakingContractChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftStakingContractChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftStakingContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftStakingContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftStakingContractChanged represents a StakingContractChanged event raised by the Nft contract.
type NftStakingContractChanged struct {
	OldStaking common.Address
	NewStaking common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStakingContractChanged is a free log retrieval operation binding the contract event 0x4502a0b31e425abf32e0aeda2ff2d41e7daf2be34bce1ff1749ca01f338cd399.
//
// Solidity: event StakingContractChanged(address indexed oldStaking, address indexed newStaking)
func (_Nft *NftFilterer) FilterStakingContractChanged(opts *bind.FilterOpts, oldStaking []common.Address, newStaking []common.Address) (*NftStakingContractChangedIterator, error) {

	var oldStakingRule []interface{}
	for _, oldStakingItem := range oldStaking {
		oldStakingRule = append(oldStakingRule, oldStakingItem)
	}
	var newStakingRule []interface{}
	for _, newStakingItem := range newStaking {
		newStakingRule = append(newStakingRule, newStakingItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "StakingContractChanged", oldStakingRule, newStakingRule)
	if err != nil {
		return nil, err
	}
	return &NftStakingContractChangedIterator{contract: _Nft.contract, event: "StakingContractChanged", logs: logs, sub: sub}, nil
}

// WatchStakingContractChanged is a free log subscription operation binding the contract event 0x4502a0b31e425abf32e0aeda2ff2d41e7daf2be34bce1ff1749ca01f338cd399.
//
// Solidity: event StakingContractChanged(address indexed oldStaking, address indexed newStaking)
func (_Nft *NftFilterer) WatchStakingContractChanged(opts *bind.WatchOpts, sink chan<- *NftStakingContractChanged, oldStaking []common.Address, newStaking []common.Address) (event.Subscription, error) {

	var oldStakingRule []interface{}
	for _, oldStakingItem := range oldStaking {
		oldStakingRule = append(oldStakingRule, oldStakingItem)
	}
	var newStakingRule []interface{}
	for _, newStakingItem := range newStaking {
		newStakingRule = append(newStakingRule, newStakingItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "StakingContractChanged", oldStakingRule, newStakingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftStakingContractChanged)
				if err := _Nft.contract.UnpackLog(event, "StakingContractChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakingContractChanged is a log parse operation binding the contract event 0x4502a0b31e425abf32e0aeda2ff2d41e7daf2be34bce1ff1749ca01f338cd399.
//
// Solidity: event StakingContractChanged(address indexed oldStaking, address indexed newStaking)
func (_Nft *NftFilterer) ParseStakingContractChanged(log types.Log) (*NftStakingContractChanged, error) {
	event := new(NftStakingContractChanged)
	if err := _Nft.contract.UnpackLog(event, "StakingContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTierPriceChangedIterator is returned from FilterTierPriceChanged and is used to iterate over the raw logs and unpacked data for TierPriceChanged events raised by the Nft contract.
type NftTierPriceChangedIterator struct {
	Event *NftTierPriceChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftTierPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTierPriceChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftTierPriceChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftTierPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTierPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTierPriceChanged represents a TierPriceChanged event raised by the Nft contract.
type NftTierPriceChanged struct {
	Tier     uint8
	OldPrice *big.Int
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTierPriceChanged is a free log retrieval operation binding the contract event 0xfb2caf0020041df26569f7905b97507e5f94247d803603171f364d0929fe0ff6.
//
// Solidity: event TierPriceChanged(uint8 indexed tier, uint256 oldPrice, uint256 newPrice)
func (_Nft *NftFilterer) FilterTierPriceChanged(opts *bind.FilterOpts, tier []uint8) (*NftTierPriceChangedIterator, error) {

	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "TierPriceChanged", tierRule)
	if err != nil {
		return nil, err
	}
	return &NftTierPriceChangedIterator{contract: _Nft.contract, event: "TierPriceChanged", logs: logs, sub: sub}, nil
}

// WatchTierPriceChanged is a free log subscription operation binding the contract event 0xfb2caf0020041df26569f7905b97507e5f94247d803603171f364d0929fe0ff6.
//
// Solidity: event TierPriceChanged(uint8 indexed tier, uint256 oldPrice, uint256 newPrice)
func (_Nft *NftFilterer) WatchTierPriceChanged(opts *bind.WatchOpts, sink chan<- *NftTierPriceChanged, tier []uint8) (event.Subscription, error) {

	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "TierPriceChanged", tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTierPriceChanged)
				if err := _Nft.contract.UnpackLog(event, "TierPriceChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTierPriceChanged is a log parse operation binding the contract event 0xfb2caf0020041df26569f7905b97507e5f94247d803603171f364d0929fe0ff6.
//
// Solidity: event TierPriceChanged(uint8 indexed tier, uint256 oldPrice, uint256 newPrice)
func (_Nft *NftFilterer) ParseTierPriceChanged(log types.Log) (*NftTierPriceChanged, error) {
	event := new(NftTierPriceChanged)
	if err := _Nft.contract.UnpackLog(event, "TierPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Nft contract.
type NftTransferIterator struct {
	Event *NftTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTransfer represents a Transfer event raised by the Nft contract.
type NftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Nft *NftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftTransferIterator{contract: _Nft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Nft *NftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTransfer)
				if err := _Nft.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Nft *NftFilterer) ParseTransfer(log types.Log) (*NftTransfer, error) {
	event := new(NftTransfer)
	if err := _Nft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftUnlistedIterator is returned from FilterUnlisted and is used to iterate over the raw logs and unpacked data for Unlisted events raised by the Nft contract.
type NftUnlistedIterator struct {
	Event *NftUnlisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftUnlistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftUnlisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftUnlisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftUnlistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftUnlistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftUnlisted represents a Unlisted event raised by the Nft contract.
type NftUnlisted struct {
	Operator common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnlisted is a free log retrieval operation binding the contract event 0xd1b3de83f2e838a8b6eb09447f2316d7bfd1a64d0059abe8f1160f490616fd83.
//
// Solidity: event Unlisted(address indexed operator, uint256 indexed tokenId)
func (_Nft *NftFilterer) FilterUnlisted(opts *bind.FilterOpts, operator []common.Address, tokenId []*big.Int) (*NftUnlistedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Unlisted", operatorRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftUnlistedIterator{contract: _Nft.contract, event: "Unlisted", logs: logs, sub: sub}, nil
}

// WatchUnlisted is a free log subscription operation binding the contract event 0xd1b3de83f2e838a8b6eb09447f2316d7bfd1a64d0059abe8f1160f490616fd83.
//
// Solidity: event Unlisted(address indexed operator, uint256 indexed tokenId)
func (_Nft *NftFilterer) WatchUnlisted(opts *bind.WatchOpts, sink chan<- *NftUnlisted, operator []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Unlisted", operatorRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftUnlisted)
				if err := _Nft.contract.UnpackLog(event, "Unlisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnlisted is a log parse operation binding the contract event 0xd1b3de83f2e838a8b6eb09447f2316d7bfd1a64d0059abe8f1160f490616fd83.
//
// Solidity: event Unlisted(address indexed operator, uint256 indexed tokenId)
func (_Nft *NftFilterer) ParseUnlisted(log types.Log) (*NftUnlisted, error) {
	event := new(NftUnlisted)
	if err := _Nft.contract.UnpackLog(event, "Unlisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
