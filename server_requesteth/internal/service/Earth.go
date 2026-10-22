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

// EarthMetaData contains all meta data concerning the Earth contract.
var EarthMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"backBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"backPool\",\"outputs\":[{\"internalType\":\"contractBackPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"backRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"buyCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dailyFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroyBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroyRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAccountFive\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAccountOne\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAccountTwo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fiveBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fiveRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fourBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fourRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalDailyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBuyPer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBuyPerDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oneBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oneRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_base\",\"type\":\"uint256\"}],\"name\":\"setBackRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_base\",\"type\":\"uint256\"}],\"name\":\"setBuyRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_base\",\"type\":\"uint256\"}],\"name\":\"setDestroyRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setFeeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setFeeAccountFive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setFeeAccountOne\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setFeeAccountTwo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_base\",\"type\":\"uint256\"}],\"name\":\"setFiveRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_base\",\"type\":\"uint256\"}],\"name\":\"setFourRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxBuyPer\",\"type\":\"uint256\"}],\"name\":\"setMaxBuyPer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxBuyPerDay\",\"type\":\"uint256\"}],\"name\":\"setMaxBuyPerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_base\",\"type\":\"uint256\"}],\"name\":\"setOneRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_base\",\"type\":\"uint256\"}],\"name\":\"setTwoRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBuy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalBuyAmountPerDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twoBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twoRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userTotalBuy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EarthABI is the input ABI used to generate the binding from.
// Deprecated: Use EarthMetaData.ABI instead.
var EarthABI = EarthMetaData.ABI

// Earth is an auto generated Go binding around an Ethereum contract.
type Earth struct {
	EarthCaller     // Read-only binding to the contract
	EarthTransactor // Write-only binding to the contract
	EarthFilterer   // Log filterer for contract events
}

// EarthCaller is an auto generated read-only Go binding around an Ethereum contract.
type EarthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EarthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EarthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EarthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EarthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EarthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EarthSession struct {
	Contract     *Earth            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EarthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EarthCallerSession struct {
	Contract *EarthCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EarthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EarthTransactorSession struct {
	Contract     *EarthTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EarthRaw is an auto generated low-level Go binding around an Ethereum contract.
type EarthRaw struct {
	Contract *Earth // Generic contract binding to access the raw methods on
}

// EarthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EarthCallerRaw struct {
	Contract *EarthCaller // Generic read-only contract binding to access the raw methods on
}

// EarthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EarthTransactorRaw struct {
	Contract *EarthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEarth creates a new instance of Earth, bound to a specific deployed contract.
func NewEarth(address common.Address, backend bind.ContractBackend) (*Earth, error) {
	contract, err := bindEarth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Earth{EarthCaller: EarthCaller{contract: contract}, EarthTransactor: EarthTransactor{contract: contract}, EarthFilterer: EarthFilterer{contract: contract}}, nil
}

// NewEarthCaller creates a new read-only instance of Earth, bound to a specific deployed contract.
func NewEarthCaller(address common.Address, caller bind.ContractCaller) (*EarthCaller, error) {
	contract, err := bindEarth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EarthCaller{contract: contract}, nil
}

// NewEarthTransactor creates a new write-only instance of Earth, bound to a specific deployed contract.
func NewEarthTransactor(address common.Address, transactor bind.ContractTransactor) (*EarthTransactor, error) {
	contract, err := bindEarth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EarthTransactor{contract: contract}, nil
}

// NewEarthFilterer creates a new log filterer instance of Earth, bound to a specific deployed contract.
func NewEarthFilterer(address common.Address, filterer bind.ContractFilterer) (*EarthFilterer, error) {
	contract, err := bindEarth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EarthFilterer{contract: contract}, nil
}

// bindEarth binds a generic wrapper to an already deployed contract.
func bindEarth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EarthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Earth *EarthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Earth.Contract.EarthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Earth *EarthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Earth.Contract.EarthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Earth *EarthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Earth.Contract.EarthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Earth *EarthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Earth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Earth *EarthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Earth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Earth *EarthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Earth.Contract.contract.Transact(opts, method, params...)
}

// DAY is a free data retrieval call binding the contract method 0x27cfe856.
//
// Solidity: function DAY() view returns(uint256)
func (_Earth *EarthCaller) DAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "DAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DAY is a free data retrieval call binding the contract method 0x27cfe856.
//
// Solidity: function DAY() view returns(uint256)
func (_Earth *EarthSession) DAY() (*big.Int, error) {
	return _Earth.Contract.DAY(&_Earth.CallOpts)
}

// DAY is a free data retrieval call binding the contract method 0x27cfe856.
//
// Solidity: function DAY() view returns(uint256)
func (_Earth *EarthCallerSession) DAY() (*big.Int, error) {
	return _Earth.Contract.DAY(&_Earth.CallOpts)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_Earth *EarthCaller) USDT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "USDT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_Earth *EarthSession) USDT() (common.Address, error) {
	return _Earth.Contract.USDT(&_Earth.CallOpts)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_Earth *EarthCallerSession) USDT() (common.Address, error) {
	return _Earth.Contract.USDT(&_Earth.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Earth *EarthCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Earth *EarthSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Earth.Contract.Allowance(&_Earth.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Earth *EarthCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Earth.Contract.Allowance(&_Earth.CallOpts, owner, spender)
}

// BackBase is a free data retrieval call binding the contract method 0x759dab3a.
//
// Solidity: function backBase() view returns(uint256)
func (_Earth *EarthCaller) BackBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "backBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BackBase is a free data retrieval call binding the contract method 0x759dab3a.
//
// Solidity: function backBase() view returns(uint256)
func (_Earth *EarthSession) BackBase() (*big.Int, error) {
	return _Earth.Contract.BackBase(&_Earth.CallOpts)
}

// BackBase is a free data retrieval call binding the contract method 0x759dab3a.
//
// Solidity: function backBase() view returns(uint256)
func (_Earth *EarthCallerSession) BackBase() (*big.Int, error) {
	return _Earth.Contract.BackBase(&_Earth.CallOpts)
}

// BackPool is a free data retrieval call binding the contract method 0xcb300137.
//
// Solidity: function backPool() view returns(address)
func (_Earth *EarthCaller) BackPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "backPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BackPool is a free data retrieval call binding the contract method 0xcb300137.
//
// Solidity: function backPool() view returns(address)
func (_Earth *EarthSession) BackPool() (common.Address, error) {
	return _Earth.Contract.BackPool(&_Earth.CallOpts)
}

// BackPool is a free data retrieval call binding the contract method 0xcb300137.
//
// Solidity: function backPool() view returns(address)
func (_Earth *EarthCallerSession) BackPool() (common.Address, error) {
	return _Earth.Contract.BackPool(&_Earth.CallOpts)
}

// BackRate is a free data retrieval call binding the contract method 0xdfc3a418.
//
// Solidity: function backRate() view returns(uint256)
func (_Earth *EarthCaller) BackRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "backRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BackRate is a free data retrieval call binding the contract method 0xdfc3a418.
//
// Solidity: function backRate() view returns(uint256)
func (_Earth *EarthSession) BackRate() (*big.Int, error) {
	return _Earth.Contract.BackRate(&_Earth.CallOpts)
}

// BackRate is a free data retrieval call binding the contract method 0xdfc3a418.
//
// Solidity: function backRate() view returns(uint256)
func (_Earth *EarthCallerSession) BackRate() (*big.Int, error) {
	return _Earth.Contract.BackRate(&_Earth.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Earth *EarthCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Earth *EarthSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Earth.Contract.BalanceOf(&_Earth.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Earth *EarthCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Earth.Contract.BalanceOf(&_Earth.CallOpts, account)
}

// BuyBase is a free data retrieval call binding the contract method 0x034a1aff.
//
// Solidity: function buyBase() view returns(uint256)
func (_Earth *EarthCaller) BuyBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "buyBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyBase is a free data retrieval call binding the contract method 0x034a1aff.
//
// Solidity: function buyBase() view returns(uint256)
func (_Earth *EarthSession) BuyBase() (*big.Int, error) {
	return _Earth.Contract.BuyBase(&_Earth.CallOpts)
}

// BuyBase is a free data retrieval call binding the contract method 0x034a1aff.
//
// Solidity: function buyBase() view returns(uint256)
func (_Earth *EarthCallerSession) BuyBase() (*big.Int, error) {
	return _Earth.Contract.BuyBase(&_Earth.CallOpts)
}

// BuyCount is a free data retrieval call binding the contract method 0x07cf02b2.
//
// Solidity: function buyCount(address , uint256 ) view returns(uint256)
func (_Earth *EarthCaller) BuyCount(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "buyCount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyCount is a free data retrieval call binding the contract method 0x07cf02b2.
//
// Solidity: function buyCount(address , uint256 ) view returns(uint256)
func (_Earth *EarthSession) BuyCount(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Earth.Contract.BuyCount(&_Earth.CallOpts, arg0, arg1)
}

// BuyCount is a free data retrieval call binding the contract method 0x07cf02b2.
//
// Solidity: function buyCount(address , uint256 ) view returns(uint256)
func (_Earth *EarthCallerSession) BuyCount(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Earth.Contract.BuyCount(&_Earth.CallOpts, arg0, arg1)
}

// BuyRate is a free data retrieval call binding the contract method 0xfc37987b.
//
// Solidity: function buyRate() view returns(uint256)
func (_Earth *EarthCaller) BuyRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "buyRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyRate is a free data retrieval call binding the contract method 0xfc37987b.
//
// Solidity: function buyRate() view returns(uint256)
func (_Earth *EarthSession) BuyRate() (*big.Int, error) {
	return _Earth.Contract.BuyRate(&_Earth.CallOpts)
}

// BuyRate is a free data retrieval call binding the contract method 0xfc37987b.
//
// Solidity: function buyRate() view returns(uint256)
func (_Earth *EarthCallerSession) BuyRate() (*big.Int, error) {
	return _Earth.Contract.BuyRate(&_Earth.CallOpts)
}

// DailyFee is a free data retrieval call binding the contract method 0x3bc1b433.
//
// Solidity: function dailyFee(uint256 ) view returns(uint256)
func (_Earth *EarthCaller) DailyFee(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "dailyFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyFee is a free data retrieval call binding the contract method 0x3bc1b433.
//
// Solidity: function dailyFee(uint256 ) view returns(uint256)
func (_Earth *EarthSession) DailyFee(arg0 *big.Int) (*big.Int, error) {
	return _Earth.Contract.DailyFee(&_Earth.CallOpts, arg0)
}

// DailyFee is a free data retrieval call binding the contract method 0x3bc1b433.
//
// Solidity: function dailyFee(uint256 ) view returns(uint256)
func (_Earth *EarthCallerSession) DailyFee(arg0 *big.Int) (*big.Int, error) {
	return _Earth.Contract.DailyFee(&_Earth.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Earth *EarthCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Earth *EarthSession) Decimals() (uint8, error) {
	return _Earth.Contract.Decimals(&_Earth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Earth *EarthCallerSession) Decimals() (uint8, error) {
	return _Earth.Contract.Decimals(&_Earth.CallOpts)
}

// DestroyAddress is a free data retrieval call binding the contract method 0xd12d9f15.
//
// Solidity: function destroyAddress() view returns(address)
func (_Earth *EarthCaller) DestroyAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "destroyAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DestroyAddress is a free data retrieval call binding the contract method 0xd12d9f15.
//
// Solidity: function destroyAddress() view returns(address)
func (_Earth *EarthSession) DestroyAddress() (common.Address, error) {
	return _Earth.Contract.DestroyAddress(&_Earth.CallOpts)
}

// DestroyAddress is a free data retrieval call binding the contract method 0xd12d9f15.
//
// Solidity: function destroyAddress() view returns(address)
func (_Earth *EarthCallerSession) DestroyAddress() (common.Address, error) {
	return _Earth.Contract.DestroyAddress(&_Earth.CallOpts)
}

// DestroyBase is a free data retrieval call binding the contract method 0xe3b914eb.
//
// Solidity: function destroyBase() view returns(uint256)
func (_Earth *EarthCaller) DestroyBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "destroyBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DestroyBase is a free data retrieval call binding the contract method 0xe3b914eb.
//
// Solidity: function destroyBase() view returns(uint256)
func (_Earth *EarthSession) DestroyBase() (*big.Int, error) {
	return _Earth.Contract.DestroyBase(&_Earth.CallOpts)
}

// DestroyBase is a free data retrieval call binding the contract method 0xe3b914eb.
//
// Solidity: function destroyBase() view returns(uint256)
func (_Earth *EarthCallerSession) DestroyBase() (*big.Int, error) {
	return _Earth.Contract.DestroyBase(&_Earth.CallOpts)
}

// DestroyRate is a free data retrieval call binding the contract method 0x2998a92b.
//
// Solidity: function destroyRate() view returns(uint256)
func (_Earth *EarthCaller) DestroyRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "destroyRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DestroyRate is a free data retrieval call binding the contract method 0x2998a92b.
//
// Solidity: function destroyRate() view returns(uint256)
func (_Earth *EarthSession) DestroyRate() (*big.Int, error) {
	return _Earth.Contract.DestroyRate(&_Earth.CallOpts)
}

// DestroyRate is a free data retrieval call binding the contract method 0x2998a92b.
//
// Solidity: function destroyRate() view returns(uint256)
func (_Earth *EarthCallerSession) DestroyRate() (*big.Int, error) {
	return _Earth.Contract.DestroyRate(&_Earth.CallOpts)
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() view returns(address)
func (_Earth *EarthCaller) FeeAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "feeAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() view returns(address)
func (_Earth *EarthSession) FeeAccount() (common.Address, error) {
	return _Earth.Contract.FeeAccount(&_Earth.CallOpts)
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() view returns(address)
func (_Earth *EarthCallerSession) FeeAccount() (common.Address, error) {
	return _Earth.Contract.FeeAccount(&_Earth.CallOpts)
}

// FeeAccountFive is a free data retrieval call binding the contract method 0x91d12ec0.
//
// Solidity: function feeAccountFive() view returns(address)
func (_Earth *EarthCaller) FeeAccountFive(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "feeAccountFive")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAccountFive is a free data retrieval call binding the contract method 0x91d12ec0.
//
// Solidity: function feeAccountFive() view returns(address)
func (_Earth *EarthSession) FeeAccountFive() (common.Address, error) {
	return _Earth.Contract.FeeAccountFive(&_Earth.CallOpts)
}

// FeeAccountFive is a free data retrieval call binding the contract method 0x91d12ec0.
//
// Solidity: function feeAccountFive() view returns(address)
func (_Earth *EarthCallerSession) FeeAccountFive() (common.Address, error) {
	return _Earth.Contract.FeeAccountFive(&_Earth.CallOpts)
}

// FeeAccountOne is a free data retrieval call binding the contract method 0xbefeb267.
//
// Solidity: function feeAccountOne() view returns(address)
func (_Earth *EarthCaller) FeeAccountOne(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "feeAccountOne")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAccountOne is a free data retrieval call binding the contract method 0xbefeb267.
//
// Solidity: function feeAccountOne() view returns(address)
func (_Earth *EarthSession) FeeAccountOne() (common.Address, error) {
	return _Earth.Contract.FeeAccountOne(&_Earth.CallOpts)
}

// FeeAccountOne is a free data retrieval call binding the contract method 0xbefeb267.
//
// Solidity: function feeAccountOne() view returns(address)
func (_Earth *EarthCallerSession) FeeAccountOne() (common.Address, error) {
	return _Earth.Contract.FeeAccountOne(&_Earth.CallOpts)
}

// FeeAccountTwo is a free data retrieval call binding the contract method 0xa4716686.
//
// Solidity: function feeAccountTwo() view returns(address)
func (_Earth *EarthCaller) FeeAccountTwo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "feeAccountTwo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAccountTwo is a free data retrieval call binding the contract method 0xa4716686.
//
// Solidity: function feeAccountTwo() view returns(address)
func (_Earth *EarthSession) FeeAccountTwo() (common.Address, error) {
	return _Earth.Contract.FeeAccountTwo(&_Earth.CallOpts)
}

// FeeAccountTwo is a free data retrieval call binding the contract method 0xa4716686.
//
// Solidity: function feeAccountTwo() view returns(address)
func (_Earth *EarthCallerSession) FeeAccountTwo() (common.Address, error) {
	return _Earth.Contract.FeeAccountTwo(&_Earth.CallOpts)
}

// FiveBase is a free data retrieval call binding the contract method 0x01dabcdc.
//
// Solidity: function fiveBase() view returns(uint256)
func (_Earth *EarthCaller) FiveBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "fiveBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FiveBase is a free data retrieval call binding the contract method 0x01dabcdc.
//
// Solidity: function fiveBase() view returns(uint256)
func (_Earth *EarthSession) FiveBase() (*big.Int, error) {
	return _Earth.Contract.FiveBase(&_Earth.CallOpts)
}

// FiveBase is a free data retrieval call binding the contract method 0x01dabcdc.
//
// Solidity: function fiveBase() view returns(uint256)
func (_Earth *EarthCallerSession) FiveBase() (*big.Int, error) {
	return _Earth.Contract.FiveBase(&_Earth.CallOpts)
}

// FiveRate is a free data retrieval call binding the contract method 0x2d82bb74.
//
// Solidity: function fiveRate() view returns(uint256)
func (_Earth *EarthCaller) FiveRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "fiveRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FiveRate is a free data retrieval call binding the contract method 0x2d82bb74.
//
// Solidity: function fiveRate() view returns(uint256)
func (_Earth *EarthSession) FiveRate() (*big.Int, error) {
	return _Earth.Contract.FiveRate(&_Earth.CallOpts)
}

// FiveRate is a free data retrieval call binding the contract method 0x2d82bb74.
//
// Solidity: function fiveRate() view returns(uint256)
func (_Earth *EarthCallerSession) FiveRate() (*big.Int, error) {
	return _Earth.Contract.FiveRate(&_Earth.CallOpts)
}

// FourBase is a free data retrieval call binding the contract method 0x7d5edbed.
//
// Solidity: function fourBase() view returns(uint256)
func (_Earth *EarthCaller) FourBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "fourBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FourBase is a free data retrieval call binding the contract method 0x7d5edbed.
//
// Solidity: function fourBase() view returns(uint256)
func (_Earth *EarthSession) FourBase() (*big.Int, error) {
	return _Earth.Contract.FourBase(&_Earth.CallOpts)
}

// FourBase is a free data retrieval call binding the contract method 0x7d5edbed.
//
// Solidity: function fourBase() view returns(uint256)
func (_Earth *EarthCallerSession) FourBase() (*big.Int, error) {
	return _Earth.Contract.FourBase(&_Earth.CallOpts)
}

// FourRate is a free data retrieval call binding the contract method 0x9fc62acc.
//
// Solidity: function fourRate() view returns(uint256)
func (_Earth *EarthCaller) FourRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "fourRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FourRate is a free data retrieval call binding the contract method 0x9fc62acc.
//
// Solidity: function fourRate() view returns(uint256)
func (_Earth *EarthSession) FourRate() (*big.Int, error) {
	return _Earth.Contract.FourRate(&_Earth.CallOpts)
}

// FourRate is a free data retrieval call binding the contract method 0x9fc62acc.
//
// Solidity: function fourRate() view returns(uint256)
func (_Earth *EarthCallerSession) FourRate() (*big.Int, error) {
	return _Earth.Contract.FourRate(&_Earth.CallOpts)
}

// GetGlobalDailyLimit is a free data retrieval call binding the contract method 0x888b1133.
//
// Solidity: function getGlobalDailyLimit() view returns(uint256)
func (_Earth *EarthCaller) GetGlobalDailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "getGlobalDailyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGlobalDailyLimit is a free data retrieval call binding the contract method 0x888b1133.
//
// Solidity: function getGlobalDailyLimit() view returns(uint256)
func (_Earth *EarthSession) GetGlobalDailyLimit() (*big.Int, error) {
	return _Earth.Contract.GetGlobalDailyLimit(&_Earth.CallOpts)
}

// GetGlobalDailyLimit is a free data retrieval call binding the contract method 0x888b1133.
//
// Solidity: function getGlobalDailyLimit() view returns(uint256)
func (_Earth *EarthCallerSession) GetGlobalDailyLimit() (*big.Int, error) {
	return _Earth.Contract.GetGlobalDailyLimit(&_Earth.CallOpts)
}

// MaxBuyPer is a free data retrieval call binding the contract method 0x6f2bc3cf.
//
// Solidity: function maxBuyPer() view returns(uint256)
func (_Earth *EarthCaller) MaxBuyPer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "maxBuyPer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBuyPer is a free data retrieval call binding the contract method 0x6f2bc3cf.
//
// Solidity: function maxBuyPer() view returns(uint256)
func (_Earth *EarthSession) MaxBuyPer() (*big.Int, error) {
	return _Earth.Contract.MaxBuyPer(&_Earth.CallOpts)
}

// MaxBuyPer is a free data retrieval call binding the contract method 0x6f2bc3cf.
//
// Solidity: function maxBuyPer() view returns(uint256)
func (_Earth *EarthCallerSession) MaxBuyPer() (*big.Int, error) {
	return _Earth.Contract.MaxBuyPer(&_Earth.CallOpts)
}

// MaxBuyPerDay is a free data retrieval call binding the contract method 0xe10fbc08.
//
// Solidity: function maxBuyPerDay() view returns(uint256)
func (_Earth *EarthCaller) MaxBuyPerDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "maxBuyPerDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBuyPerDay is a free data retrieval call binding the contract method 0xe10fbc08.
//
// Solidity: function maxBuyPerDay() view returns(uint256)
func (_Earth *EarthSession) MaxBuyPerDay() (*big.Int, error) {
	return _Earth.Contract.MaxBuyPerDay(&_Earth.CallOpts)
}

// MaxBuyPerDay is a free data retrieval call binding the contract method 0xe10fbc08.
//
// Solidity: function maxBuyPerDay() view returns(uint256)
func (_Earth *EarthCallerSession) MaxBuyPerDay() (*big.Int, error) {
	return _Earth.Contract.MaxBuyPerDay(&_Earth.CallOpts)
}

// MintAccount is a free data retrieval call binding the contract method 0xbe1dd152.
//
// Solidity: function mintAccount() view returns(address)
func (_Earth *EarthCaller) MintAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "mintAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MintAccount is a free data retrieval call binding the contract method 0xbe1dd152.
//
// Solidity: function mintAccount() view returns(address)
func (_Earth *EarthSession) MintAccount() (common.Address, error) {
	return _Earth.Contract.MintAccount(&_Earth.CallOpts)
}

// MintAccount is a free data retrieval call binding the contract method 0xbe1dd152.
//
// Solidity: function mintAccount() view returns(address)
func (_Earth *EarthCallerSession) MintAccount() (common.Address, error) {
	return _Earth.Contract.MintAccount(&_Earth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Earth *EarthCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Earth *EarthSession) Name() (string, error) {
	return _Earth.Contract.Name(&_Earth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Earth *EarthCallerSession) Name() (string, error) {
	return _Earth.Contract.Name(&_Earth.CallOpts)
}

// OneBase is a free data retrieval call binding the contract method 0xe26d2210.
//
// Solidity: function oneBase() view returns(uint256)
func (_Earth *EarthCaller) OneBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "oneBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OneBase is a free data retrieval call binding the contract method 0xe26d2210.
//
// Solidity: function oneBase() view returns(uint256)
func (_Earth *EarthSession) OneBase() (*big.Int, error) {
	return _Earth.Contract.OneBase(&_Earth.CallOpts)
}

// OneBase is a free data retrieval call binding the contract method 0xe26d2210.
//
// Solidity: function oneBase() view returns(uint256)
func (_Earth *EarthCallerSession) OneBase() (*big.Int, error) {
	return _Earth.Contract.OneBase(&_Earth.CallOpts)
}

// OneRate is a free data retrieval call binding the contract method 0xc354ebb0.
//
// Solidity: function oneRate() view returns(uint256)
func (_Earth *EarthCaller) OneRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "oneRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OneRate is a free data retrieval call binding the contract method 0xc354ebb0.
//
// Solidity: function oneRate() view returns(uint256)
func (_Earth *EarthSession) OneRate() (*big.Int, error) {
	return _Earth.Contract.OneRate(&_Earth.CallOpts)
}

// OneRate is a free data retrieval call binding the contract method 0xc354ebb0.
//
// Solidity: function oneRate() view returns(uint256)
func (_Earth *EarthCallerSession) OneRate() (*big.Int, error) {
	return _Earth.Contract.OneRate(&_Earth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Earth *EarthCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Earth *EarthSession) Owner() (common.Address, error) {
	return _Earth.Contract.Owner(&_Earth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Earth *EarthCallerSession) Owner() (common.Address, error) {
	return _Earth.Contract.Owner(&_Earth.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Earth *EarthCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Earth *EarthSession) StartTime() (*big.Int, error) {
	return _Earth.Contract.StartTime(&_Earth.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Earth *EarthCallerSession) StartTime() (*big.Int, error) {
	return _Earth.Contract.StartTime(&_Earth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Earth *EarthCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Earth *EarthSession) Symbol() (string, error) {
	return _Earth.Contract.Symbol(&_Earth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Earth *EarthCallerSession) Symbol() (string, error) {
	return _Earth.Contract.Symbol(&_Earth.CallOpts)
}

// TotalBuy is a free data retrieval call binding the contract method 0xb1c09b2a.
//
// Solidity: function totalBuy() view returns(uint256)
func (_Earth *EarthCaller) TotalBuy(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "totalBuy")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBuy is a free data retrieval call binding the contract method 0xb1c09b2a.
//
// Solidity: function totalBuy() view returns(uint256)
func (_Earth *EarthSession) TotalBuy() (*big.Int, error) {
	return _Earth.Contract.TotalBuy(&_Earth.CallOpts)
}

// TotalBuy is a free data retrieval call binding the contract method 0xb1c09b2a.
//
// Solidity: function totalBuy() view returns(uint256)
func (_Earth *EarthCallerSession) TotalBuy() (*big.Int, error) {
	return _Earth.Contract.TotalBuy(&_Earth.CallOpts)
}

// TotalBuyAmountPerDay is a free data retrieval call binding the contract method 0x0face042.
//
// Solidity: function totalBuyAmountPerDay(uint256 ) view returns(uint256)
func (_Earth *EarthCaller) TotalBuyAmountPerDay(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "totalBuyAmountPerDay", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBuyAmountPerDay is a free data retrieval call binding the contract method 0x0face042.
//
// Solidity: function totalBuyAmountPerDay(uint256 ) view returns(uint256)
func (_Earth *EarthSession) TotalBuyAmountPerDay(arg0 *big.Int) (*big.Int, error) {
	return _Earth.Contract.TotalBuyAmountPerDay(&_Earth.CallOpts, arg0)
}

// TotalBuyAmountPerDay is a free data retrieval call binding the contract method 0x0face042.
//
// Solidity: function totalBuyAmountPerDay(uint256 ) view returns(uint256)
func (_Earth *EarthCallerSession) TotalBuyAmountPerDay(arg0 *big.Int) (*big.Int, error) {
	return _Earth.Contract.TotalBuyAmountPerDay(&_Earth.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Earth *EarthCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Earth *EarthSession) TotalSupply() (*big.Int, error) {
	return _Earth.Contract.TotalSupply(&_Earth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Earth *EarthCallerSession) TotalSupply() (*big.Int, error) {
	return _Earth.Contract.TotalSupply(&_Earth.CallOpts)
}

// TwoBase is a free data retrieval call binding the contract method 0x1f923b56.
//
// Solidity: function twoBase() view returns(uint256)
func (_Earth *EarthCaller) TwoBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "twoBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TwoBase is a free data retrieval call binding the contract method 0x1f923b56.
//
// Solidity: function twoBase() view returns(uint256)
func (_Earth *EarthSession) TwoBase() (*big.Int, error) {
	return _Earth.Contract.TwoBase(&_Earth.CallOpts)
}

// TwoBase is a free data retrieval call binding the contract method 0x1f923b56.
//
// Solidity: function twoBase() view returns(uint256)
func (_Earth *EarthCallerSession) TwoBase() (*big.Int, error) {
	return _Earth.Contract.TwoBase(&_Earth.CallOpts)
}

// TwoRate is a free data retrieval call binding the contract method 0xda587ee6.
//
// Solidity: function twoRate() view returns(uint256)
func (_Earth *EarthCaller) TwoRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "twoRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TwoRate is a free data retrieval call binding the contract method 0xda587ee6.
//
// Solidity: function twoRate() view returns(uint256)
func (_Earth *EarthSession) TwoRate() (*big.Int, error) {
	return _Earth.Contract.TwoRate(&_Earth.CallOpts)
}

// TwoRate is a free data retrieval call binding the contract method 0xda587ee6.
//
// Solidity: function twoRate() view returns(uint256)
func (_Earth *EarthCallerSession) TwoRate() (*big.Int, error) {
	return _Earth.Contract.TwoRate(&_Earth.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Earth *EarthCaller) UniswapV2Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "uniswapV2Pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Earth *EarthSession) UniswapV2Pair() (common.Address, error) {
	return _Earth.Contract.UniswapV2Pair(&_Earth.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Earth *EarthCallerSession) UniswapV2Pair() (common.Address, error) {
	return _Earth.Contract.UniswapV2Pair(&_Earth.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Earth *EarthCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Earth *EarthSession) UniswapV2Router() (common.Address, error) {
	return _Earth.Contract.UniswapV2Router(&_Earth.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Earth *EarthCallerSession) UniswapV2Router() (common.Address, error) {
	return _Earth.Contract.UniswapV2Router(&_Earth.CallOpts)
}

// UserTotalBuy is a free data retrieval call binding the contract method 0x13dfadd3.
//
// Solidity: function userTotalBuy(address ) view returns(uint256)
func (_Earth *EarthCaller) UserTotalBuy(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "userTotalBuy", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserTotalBuy is a free data retrieval call binding the contract method 0x13dfadd3.
//
// Solidity: function userTotalBuy(address ) view returns(uint256)
func (_Earth *EarthSession) UserTotalBuy(arg0 common.Address) (*big.Int, error) {
	return _Earth.Contract.UserTotalBuy(&_Earth.CallOpts, arg0)
}

// UserTotalBuy is a free data retrieval call binding the contract method 0x13dfadd3.
//
// Solidity: function userTotalBuy(address ) view returns(uint256)
func (_Earth *EarthCallerSession) UserTotalBuy(arg0 common.Address) (*big.Int, error) {
	return _Earth.Contract.UserTotalBuy(&_Earth.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Earth *EarthTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Earth *EarthSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.Approve(&_Earth.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Earth *EarthTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.Approve(&_Earth.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Earth *EarthTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Earth *EarthSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.DecreaseAllowance(&_Earth.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Earth *EarthTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.DecreaseAllowance(&_Earth.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Earth *EarthTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Earth *EarthSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.IncreaseAllowance(&_Earth.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Earth *EarthTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.IncreaseAllowance(&_Earth.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Earth *EarthTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Earth *EarthSession) RenounceOwnership() (*types.Transaction, error) {
	return _Earth.Contract.RenounceOwnership(&_Earth.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Earth *EarthTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Earth.Contract.RenounceOwnership(&_Earth.TransactOpts)
}

// SetBackRate is a paid mutator transaction binding the contract method 0x0a736830.
//
// Solidity: function setBackRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthTransactor) SetBackRate(opts *bind.TransactOpts, _rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "setBackRate", _rate, _base)
}

// SetBackRate is a paid mutator transaction binding the contract method 0x0a736830.
//
// Solidity: function setBackRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthSession) SetBackRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetBackRate(&_Earth.TransactOpts, _rate, _base)
}

// SetBackRate is a paid mutator transaction binding the contract method 0x0a736830.
//
// Solidity: function setBackRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthTransactorSession) SetBackRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetBackRate(&_Earth.TransactOpts, _rate, _base)
}

// SetBuyRate is a paid mutator transaction binding the contract method 0x9a79c3b6.
//
// Solidity: function setBuyRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthTransactor) SetBuyRate(opts *bind.TransactOpts, _rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "setBuyRate", _rate, _base)
}

// SetBuyRate is a paid mutator transaction binding the contract method 0x9a79c3b6.
//
// Solidity: function setBuyRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthSession) SetBuyRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetBuyRate(&_Earth.TransactOpts, _rate, _base)
}

// SetBuyRate is a paid mutator transaction binding the contract method 0x9a79c3b6.
//
// Solidity: function setBuyRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthTransactorSession) SetBuyRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetBuyRate(&_Earth.TransactOpts, _rate, _base)
}

// SetDestroyRate is a paid mutator transaction binding the contract method 0x73700771.
//
// Solidity: function setDestroyRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthTransactor) SetDestroyRate(opts *bind.TransactOpts, _rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "setDestroyRate", _rate, _base)
}

// SetDestroyRate is a paid mutator transaction binding the contract method 0x73700771.
//
// Solidity: function setDestroyRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthSession) SetDestroyRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetDestroyRate(&_Earth.TransactOpts, _rate, _base)
}

// SetDestroyRate is a paid mutator transaction binding the contract method 0x73700771.
//
// Solidity: function setDestroyRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthTransactorSession) SetDestroyRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetDestroyRate(&_Earth.TransactOpts, _rate, _base)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address _addr) returns()
func (_Earth *EarthTransactor) SetFeeAccount(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "setFeeAccount", _addr)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address _addr) returns()
func (_Earth *EarthSession) SetFeeAccount(_addr common.Address) (*types.Transaction, error) {
	return _Earth.Contract.SetFeeAccount(&_Earth.TransactOpts, _addr)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address _addr) returns()
func (_Earth *EarthTransactorSession) SetFeeAccount(_addr common.Address) (*types.Transaction, error) {
	return _Earth.Contract.SetFeeAccount(&_Earth.TransactOpts, _addr)
}

// SetFeeAccountFive is a paid mutator transaction binding the contract method 0xf47851d2.
//
// Solidity: function setFeeAccountFive(address _addr) returns()
func (_Earth *EarthTransactor) SetFeeAccountFive(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "setFeeAccountFive", _addr)
}

// SetFeeAccountFive is a paid mutator transaction binding the contract method 0xf47851d2.
//
// Solidity: function setFeeAccountFive(address _addr) returns()
func (_Earth *EarthSession) SetFeeAccountFive(_addr common.Address) (*types.Transaction, error) {
	return _Earth.Contract.SetFeeAccountFive(&_Earth.TransactOpts, _addr)
}

// SetFeeAccountFive is a paid mutator transaction binding the contract method 0xf47851d2.
//
// Solidity: function setFeeAccountFive(address _addr) returns()
func (_Earth *EarthTransactorSession) SetFeeAccountFive(_addr common.Address) (*types.Transaction, error) {
	return _Earth.Contract.SetFeeAccountFive(&_Earth.TransactOpts, _addr)
}

// SetFeeAccountOne is a paid mutator transaction binding the contract method 0x4be22490.
//
// Solidity: function setFeeAccountOne(address _addr) returns()
func (_Earth *EarthTransactor) SetFeeAccountOne(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "setFeeAccountOne", _addr)
}

// SetFeeAccountOne is a paid mutator transaction binding the contract method 0x4be22490.
//
// Solidity: function setFeeAccountOne(address _addr) returns()
func (_Earth *EarthSession) SetFeeAccountOne(_addr common.Address) (*types.Transaction, error) {
	return _Earth.Contract.SetFeeAccountOne(&_Earth.TransactOpts, _addr)
}

// SetFeeAccountOne is a paid mutator transaction binding the contract method 0x4be22490.
//
// Solidity: function setFeeAccountOne(address _addr) returns()
func (_Earth *EarthTransactorSession) SetFeeAccountOne(_addr common.Address) (*types.Transaction, error) {
	return _Earth.Contract.SetFeeAccountOne(&_Earth.TransactOpts, _addr)
}

// SetFeeAccountTwo is a paid mutator transaction binding the contract method 0x84a5629b.
//
// Solidity: function setFeeAccountTwo(address _addr) returns()
func (_Earth *EarthTransactor) SetFeeAccountTwo(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "setFeeAccountTwo", _addr)
}

// SetFeeAccountTwo is a paid mutator transaction binding the contract method 0x84a5629b.
//
// Solidity: function setFeeAccountTwo(address _addr) returns()
func (_Earth *EarthSession) SetFeeAccountTwo(_addr common.Address) (*types.Transaction, error) {
	return _Earth.Contract.SetFeeAccountTwo(&_Earth.TransactOpts, _addr)
}

// SetFeeAccountTwo is a paid mutator transaction binding the contract method 0x84a5629b.
//
// Solidity: function setFeeAccountTwo(address _addr) returns()
func (_Earth *EarthTransactorSession) SetFeeAccountTwo(_addr common.Address) (*types.Transaction, error) {
	return _Earth.Contract.SetFeeAccountTwo(&_Earth.TransactOpts, _addr)
}

// SetFiveRate is a paid mutator transaction binding the contract method 0xc6832bec.
//
// Solidity: function setFiveRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthTransactor) SetFiveRate(opts *bind.TransactOpts, _rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "setFiveRate", _rate, _base)
}

// SetFiveRate is a paid mutator transaction binding the contract method 0xc6832bec.
//
// Solidity: function setFiveRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthSession) SetFiveRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetFiveRate(&_Earth.TransactOpts, _rate, _base)
}

// SetFiveRate is a paid mutator transaction binding the contract method 0xc6832bec.
//
// Solidity: function setFiveRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthTransactorSession) SetFiveRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetFiveRate(&_Earth.TransactOpts, _rate, _base)
}

// SetFourRate is a paid mutator transaction binding the contract method 0x117535f8.
//
// Solidity: function setFourRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthTransactor) SetFourRate(opts *bind.TransactOpts, _rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "setFourRate", _rate, _base)
}

// SetFourRate is a paid mutator transaction binding the contract method 0x117535f8.
//
// Solidity: function setFourRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthSession) SetFourRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetFourRate(&_Earth.TransactOpts, _rate, _base)
}

// SetFourRate is a paid mutator transaction binding the contract method 0x117535f8.
//
// Solidity: function setFourRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthTransactorSession) SetFourRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetFourRate(&_Earth.TransactOpts, _rate, _base)
}

// SetMaxBuyPer is a paid mutator transaction binding the contract method 0xfd82917e.
//
// Solidity: function setMaxBuyPer(uint256 _maxBuyPer) returns()
func (_Earth *EarthTransactor) SetMaxBuyPer(opts *bind.TransactOpts, _maxBuyPer *big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "setMaxBuyPer", _maxBuyPer)
}

// SetMaxBuyPer is a paid mutator transaction binding the contract method 0xfd82917e.
//
// Solidity: function setMaxBuyPer(uint256 _maxBuyPer) returns()
func (_Earth *EarthSession) SetMaxBuyPer(_maxBuyPer *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetMaxBuyPer(&_Earth.TransactOpts, _maxBuyPer)
}

// SetMaxBuyPer is a paid mutator transaction binding the contract method 0xfd82917e.
//
// Solidity: function setMaxBuyPer(uint256 _maxBuyPer) returns()
func (_Earth *EarthTransactorSession) SetMaxBuyPer(_maxBuyPer *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetMaxBuyPer(&_Earth.TransactOpts, _maxBuyPer)
}

// SetMaxBuyPerDay is a paid mutator transaction binding the contract method 0x04ace00e.
//
// Solidity: function setMaxBuyPerDay(uint256 _maxBuyPerDay) returns()
func (_Earth *EarthTransactor) SetMaxBuyPerDay(opts *bind.TransactOpts, _maxBuyPerDay *big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "setMaxBuyPerDay", _maxBuyPerDay)
}

// SetMaxBuyPerDay is a paid mutator transaction binding the contract method 0x04ace00e.
//
// Solidity: function setMaxBuyPerDay(uint256 _maxBuyPerDay) returns()
func (_Earth *EarthSession) SetMaxBuyPerDay(_maxBuyPerDay *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetMaxBuyPerDay(&_Earth.TransactOpts, _maxBuyPerDay)
}

// SetMaxBuyPerDay is a paid mutator transaction binding the contract method 0x04ace00e.
//
// Solidity: function setMaxBuyPerDay(uint256 _maxBuyPerDay) returns()
func (_Earth *EarthTransactorSession) SetMaxBuyPerDay(_maxBuyPerDay *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetMaxBuyPerDay(&_Earth.TransactOpts, _maxBuyPerDay)
}

// SetOneRate is a paid mutator transaction binding the contract method 0xe85a0d0a.
//
// Solidity: function setOneRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthTransactor) SetOneRate(opts *bind.TransactOpts, _rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "setOneRate", _rate, _base)
}

// SetOneRate is a paid mutator transaction binding the contract method 0xe85a0d0a.
//
// Solidity: function setOneRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthSession) SetOneRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetOneRate(&_Earth.TransactOpts, _rate, _base)
}

// SetOneRate is a paid mutator transaction binding the contract method 0xe85a0d0a.
//
// Solidity: function setOneRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthTransactorSession) SetOneRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetOneRate(&_Earth.TransactOpts, _rate, _base)
}

// SetTwoRate is a paid mutator transaction binding the contract method 0x9f7b8c47.
//
// Solidity: function setTwoRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthTransactor) SetTwoRate(opts *bind.TransactOpts, _rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "setTwoRate", _rate, _base)
}

// SetTwoRate is a paid mutator transaction binding the contract method 0x9f7b8c47.
//
// Solidity: function setTwoRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthSession) SetTwoRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetTwoRate(&_Earth.TransactOpts, _rate, _base)
}

// SetTwoRate is a paid mutator transaction binding the contract method 0x9f7b8c47.
//
// Solidity: function setTwoRate(uint256 _rate, uint256 _base) returns()
func (_Earth *EarthTransactorSession) SetTwoRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.SetTwoRate(&_Earth.TransactOpts, _rate, _base)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Earth *EarthTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Earth *EarthSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.Transfer(&_Earth.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Earth *EarthTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.Transfer(&_Earth.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Earth *EarthTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Earth *EarthSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.TransferFrom(&_Earth.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Earth *EarthTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Earth.Contract.TransferFrom(&_Earth.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Earth *EarthTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Earth *EarthSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Earth.Contract.TransferOwnership(&_Earth.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Earth *EarthTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Earth.Contract.TransferOwnership(&_Earth.TransactOpts, newOwner)
}

// EarthApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Earth contract.
type EarthApprovalIterator struct {
	Event *EarthApproval // Event containing the contract specifics and raw log

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
func (it *EarthApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EarthApproval)
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
		it.Event = new(EarthApproval)
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
func (it *EarthApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EarthApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EarthApproval represents a Approval event raised by the Earth contract.
type EarthApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Earth *EarthFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EarthApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Earth.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EarthApprovalIterator{contract: _Earth.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Earth *EarthFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EarthApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Earth.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EarthApproval)
				if err := _Earth.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Earth *EarthFilterer) ParseApproval(log types.Log) (*EarthApproval, error) {
	event := new(EarthApproval)
	if err := _Earth.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EarthTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Earth contract.
type EarthTransferIterator struct {
	Event *EarthTransfer // Event containing the contract specifics and raw log

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
func (it *EarthTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EarthTransfer)
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
		it.Event = new(EarthTransfer)
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
func (it *EarthTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EarthTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EarthTransfer represents a Transfer event raised by the Earth contract.
type EarthTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Earth *EarthFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EarthTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Earth.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EarthTransferIterator{contract: _Earth.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Earth *EarthFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EarthTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Earth.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EarthTransfer)
				if err := _Earth.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Earth *EarthFilterer) ParseTransfer(log types.Log) (*EarthTransfer, error) {
	event := new(EarthTransfer)
	if err := _Earth.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
