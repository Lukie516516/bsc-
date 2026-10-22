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

// CatMetaData contains all meta data concerning the Cat contract.
var CatMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dailyFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openBuy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openSell\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setFeeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_open\",\"type\":\"bool\"}],\"name\":\"setOpenBuy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_open\",\"type\":\"bool\"}],\"name\":\"setOpenSell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_base\",\"type\":\"uint256\"}],\"name\":\"setSellRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setWhiteMap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CatABI is the input ABI used to generate the binding from.
// Deprecated: Use CatMetaData.ABI instead.
var CatABI = CatMetaData.ABI

// Cat is an auto generated Go binding around an Ethereum contract.
type Cat struct {
	CatCaller     // Read-only binding to the contract
	CatTransactor // Write-only binding to the contract
	CatFilterer   // Log filterer for contract events
}

// CatCaller is an auto generated read-only Go binding around an Ethereum contract.
type CatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CatSession struct {
	Contract     *Cat              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CatCallerSession struct {
	Contract *CatCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CatTransactorSession struct {
	Contract     *CatTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CatRaw is an auto generated low-level Go binding around an Ethereum contract.
type CatRaw struct {
	Contract *Cat // Generic contract binding to access the raw methods on
}

// CatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CatCallerRaw struct {
	Contract *CatCaller // Generic read-only contract binding to access the raw methods on
}

// CatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CatTransactorRaw struct {
	Contract *CatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCat creates a new instance of Cat, bound to a specific deployed contract.
func NewCat(address common.Address, backend bind.ContractBackend) (*Cat, error) {
	contract, err := bindCat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cat{CatCaller: CatCaller{contract: contract}, CatTransactor: CatTransactor{contract: contract}, CatFilterer: CatFilterer{contract: contract}}, nil
}

// NewCatCaller creates a new read-only instance of Cat, bound to a specific deployed contract.
func NewCatCaller(address common.Address, caller bind.ContractCaller) (*CatCaller, error) {
	contract, err := bindCat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CatCaller{contract: contract}, nil
}

// NewCatTransactor creates a new write-only instance of Cat, bound to a specific deployed contract.
func NewCatTransactor(address common.Address, transactor bind.ContractTransactor) (*CatTransactor, error) {
	contract, err := bindCat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CatTransactor{contract: contract}, nil
}

// NewCatFilterer creates a new log filterer instance of Cat, bound to a specific deployed contract.
func NewCatFilterer(address common.Address, filterer bind.ContractFilterer) (*CatFilterer, error) {
	contract, err := bindCat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CatFilterer{contract: contract}, nil
}

// bindCat binds a generic wrapper to an already deployed contract.
func bindCat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CatABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cat *CatRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cat.Contract.CatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cat *CatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cat.Contract.CatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cat *CatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cat.Contract.CatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cat *CatCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cat *CatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cat *CatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cat.Contract.contract.Transact(opts, method, params...)
}

// DAY is a free data retrieval call binding the contract method 0x27cfe856.
//
// Solidity: function DAY() view returns(uint256)
func (_Cat *CatCaller) DAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "DAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DAY is a free data retrieval call binding the contract method 0x27cfe856.
//
// Solidity: function DAY() view returns(uint256)
func (_Cat *CatSession) DAY() (*big.Int, error) {
	return _Cat.Contract.DAY(&_Cat.CallOpts)
}

// DAY is a free data retrieval call binding the contract method 0x27cfe856.
//
// Solidity: function DAY() view returns(uint256)
func (_Cat *CatCallerSession) DAY() (*big.Int, error) {
	return _Cat.Contract.DAY(&_Cat.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cat *CatCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cat *CatSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Cat.Contract.Allowance(&_Cat.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cat *CatCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Cat.Contract.Allowance(&_Cat.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cat *CatCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cat *CatSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Cat.Contract.BalanceOf(&_Cat.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cat *CatCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Cat.Contract.BalanceOf(&_Cat.CallOpts, account)
}

// DailyFee is a free data retrieval call binding the contract method 0x3bc1b433.
//
// Solidity: function dailyFee(uint256 ) view returns(uint256)
func (_Cat *CatCaller) DailyFee(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "dailyFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyFee is a free data retrieval call binding the contract method 0x3bc1b433.
//
// Solidity: function dailyFee(uint256 ) view returns(uint256)
func (_Cat *CatSession) DailyFee(arg0 *big.Int) (*big.Int, error) {
	return _Cat.Contract.DailyFee(&_Cat.CallOpts, arg0)
}

// DailyFee is a free data retrieval call binding the contract method 0x3bc1b433.
//
// Solidity: function dailyFee(uint256 ) view returns(uint256)
func (_Cat *CatCallerSession) DailyFee(arg0 *big.Int) (*big.Int, error) {
	return _Cat.Contract.DailyFee(&_Cat.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cat *CatCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cat *CatSession) Decimals() (uint8, error) {
	return _Cat.Contract.Decimals(&_Cat.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cat *CatCallerSession) Decimals() (uint8, error) {
	return _Cat.Contract.Decimals(&_Cat.CallOpts)
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() view returns(address)
func (_Cat *CatCaller) FeeAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "feeAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() view returns(address)
func (_Cat *CatSession) FeeAccount() (common.Address, error) {
	return _Cat.Contract.FeeAccount(&_Cat.CallOpts)
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() view returns(address)
func (_Cat *CatCallerSession) FeeAccount() (common.Address, error) {
	return _Cat.Contract.FeeAccount(&_Cat.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cat *CatCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cat *CatSession) Name() (string, error) {
	return _Cat.Contract.Name(&_Cat.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cat *CatCallerSession) Name() (string, error) {
	return _Cat.Contract.Name(&_Cat.CallOpts)
}

// OpenBuy is a free data retrieval call binding the contract method 0xd4a67930.
//
// Solidity: function openBuy() view returns(bool)
func (_Cat *CatCaller) OpenBuy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "openBuy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OpenBuy is a free data retrieval call binding the contract method 0xd4a67930.
//
// Solidity: function openBuy() view returns(bool)
func (_Cat *CatSession) OpenBuy() (bool, error) {
	return _Cat.Contract.OpenBuy(&_Cat.CallOpts)
}

// OpenBuy is a free data retrieval call binding the contract method 0xd4a67930.
//
// Solidity: function openBuy() view returns(bool)
func (_Cat *CatCallerSession) OpenBuy() (bool, error) {
	return _Cat.Contract.OpenBuy(&_Cat.CallOpts)
}

// OpenSell is a free data retrieval call binding the contract method 0xf1eae64a.
//
// Solidity: function openSell() view returns(bool)
func (_Cat *CatCaller) OpenSell(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "openSell")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OpenSell is a free data retrieval call binding the contract method 0xf1eae64a.
//
// Solidity: function openSell() view returns(bool)
func (_Cat *CatSession) OpenSell() (bool, error) {
	return _Cat.Contract.OpenSell(&_Cat.CallOpts)
}

// OpenSell is a free data retrieval call binding the contract method 0xf1eae64a.
//
// Solidity: function openSell() view returns(bool)
func (_Cat *CatCallerSession) OpenSell() (bool, error) {
	return _Cat.Contract.OpenSell(&_Cat.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cat *CatCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cat *CatSession) Owner() (common.Address, error) {
	return _Cat.Contract.Owner(&_Cat.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cat *CatCallerSession) Owner() (common.Address, error) {
	return _Cat.Contract.Owner(&_Cat.CallOpts)
}

// SellBase is a free data retrieval call binding the contract method 0x8b2033bf.
//
// Solidity: function sellBase() view returns(uint256)
func (_Cat *CatCaller) SellBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "sellBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellBase is a free data retrieval call binding the contract method 0x8b2033bf.
//
// Solidity: function sellBase() view returns(uint256)
func (_Cat *CatSession) SellBase() (*big.Int, error) {
	return _Cat.Contract.SellBase(&_Cat.CallOpts)
}

// SellBase is a free data retrieval call binding the contract method 0x8b2033bf.
//
// Solidity: function sellBase() view returns(uint256)
func (_Cat *CatCallerSession) SellBase() (*big.Int, error) {
	return _Cat.Contract.SellBase(&_Cat.CallOpts)
}

// SellRate is a free data retrieval call binding the contract method 0x6217229b.
//
// Solidity: function sellRate() view returns(uint256)
func (_Cat *CatCaller) SellRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "sellRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellRate is a free data retrieval call binding the contract method 0x6217229b.
//
// Solidity: function sellRate() view returns(uint256)
func (_Cat *CatSession) SellRate() (*big.Int, error) {
	return _Cat.Contract.SellRate(&_Cat.CallOpts)
}

// SellRate is a free data retrieval call binding the contract method 0x6217229b.
//
// Solidity: function sellRate() view returns(uint256)
func (_Cat *CatCallerSession) SellRate() (*big.Int, error) {
	return _Cat.Contract.SellRate(&_Cat.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cat *CatCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cat *CatSession) Symbol() (string, error) {
	return _Cat.Contract.Symbol(&_Cat.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cat *CatCallerSession) Symbol() (string, error) {
	return _Cat.Contract.Symbol(&_Cat.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cat *CatCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cat *CatSession) TotalSupply() (*big.Int, error) {
	return _Cat.Contract.TotalSupply(&_Cat.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cat *CatCallerSession) TotalSupply() (*big.Int, error) {
	return _Cat.Contract.TotalSupply(&_Cat.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Cat *CatCaller) UniswapV2Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "uniswapV2Pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Cat *CatSession) UniswapV2Pair() (common.Address, error) {
	return _Cat.Contract.UniswapV2Pair(&_Cat.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Cat *CatCallerSession) UniswapV2Pair() (common.Address, error) {
	return _Cat.Contract.UniswapV2Pair(&_Cat.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Cat *CatCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Cat *CatSession) UniswapV2Router() (common.Address, error) {
	return _Cat.Contract.UniswapV2Router(&_Cat.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Cat *CatCallerSession) UniswapV2Router() (common.Address, error) {
	return _Cat.Contract.UniswapV2Router(&_Cat.CallOpts)
}

// WhiteMap is a free data retrieval call binding the contract method 0x7b895238.
//
// Solidity: function whiteMap(address ) view returns(bool)
func (_Cat *CatCaller) WhiteMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Cat.contract.Call(opts, &out, "whiteMap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteMap is a free data retrieval call binding the contract method 0x7b895238.
//
// Solidity: function whiteMap(address ) view returns(bool)
func (_Cat *CatSession) WhiteMap(arg0 common.Address) (bool, error) {
	return _Cat.Contract.WhiteMap(&_Cat.CallOpts, arg0)
}

// WhiteMap is a free data retrieval call binding the contract method 0x7b895238.
//
// Solidity: function whiteMap(address ) view returns(bool)
func (_Cat *CatCallerSession) WhiteMap(arg0 common.Address) (bool, error) {
	return _Cat.Contract.WhiteMap(&_Cat.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cat *CatTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cat.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cat *CatSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cat.Contract.Approve(&_Cat.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cat *CatTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cat.Contract.Approve(&_Cat.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cat *CatTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cat.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cat *CatSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cat.Contract.DecreaseAllowance(&_Cat.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cat *CatTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cat.Contract.DecreaseAllowance(&_Cat.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cat *CatTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cat.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cat *CatSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cat.Contract.IncreaseAllowance(&_Cat.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cat *CatTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cat.Contract.IncreaseAllowance(&_Cat.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cat *CatTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cat.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cat *CatSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cat.Contract.RenounceOwnership(&_Cat.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cat *CatTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cat.Contract.RenounceOwnership(&_Cat.TransactOpts)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address _addr) returns()
func (_Cat *CatTransactor) SetFeeAccount(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Cat.contract.Transact(opts, "setFeeAccount", _addr)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address _addr) returns()
func (_Cat *CatSession) SetFeeAccount(_addr common.Address) (*types.Transaction, error) {
	return _Cat.Contract.SetFeeAccount(&_Cat.TransactOpts, _addr)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address _addr) returns()
func (_Cat *CatTransactorSession) SetFeeAccount(_addr common.Address) (*types.Transaction, error) {
	return _Cat.Contract.SetFeeAccount(&_Cat.TransactOpts, _addr)
}

// SetOpenBuy is a paid mutator transaction binding the contract method 0x29e9a524.
//
// Solidity: function setOpenBuy(bool _open) returns()
func (_Cat *CatTransactor) SetOpenBuy(opts *bind.TransactOpts, _open bool) (*types.Transaction, error) {
	return _Cat.contract.Transact(opts, "setOpenBuy", _open)
}

// SetOpenBuy is a paid mutator transaction binding the contract method 0x29e9a524.
//
// Solidity: function setOpenBuy(bool _open) returns()
func (_Cat *CatSession) SetOpenBuy(_open bool) (*types.Transaction, error) {
	return _Cat.Contract.SetOpenBuy(&_Cat.TransactOpts, _open)
}

// SetOpenBuy is a paid mutator transaction binding the contract method 0x29e9a524.
//
// Solidity: function setOpenBuy(bool _open) returns()
func (_Cat *CatTransactorSession) SetOpenBuy(_open bool) (*types.Transaction, error) {
	return _Cat.Contract.SetOpenBuy(&_Cat.TransactOpts, _open)
}

// SetOpenSell is a paid mutator transaction binding the contract method 0x3c03b9a4.
//
// Solidity: function setOpenSell(bool _open) returns()
func (_Cat *CatTransactor) SetOpenSell(opts *bind.TransactOpts, _open bool) (*types.Transaction, error) {
	return _Cat.contract.Transact(opts, "setOpenSell", _open)
}

// SetOpenSell is a paid mutator transaction binding the contract method 0x3c03b9a4.
//
// Solidity: function setOpenSell(bool _open) returns()
func (_Cat *CatSession) SetOpenSell(_open bool) (*types.Transaction, error) {
	return _Cat.Contract.SetOpenSell(&_Cat.TransactOpts, _open)
}

// SetOpenSell is a paid mutator transaction binding the contract method 0x3c03b9a4.
//
// Solidity: function setOpenSell(bool _open) returns()
func (_Cat *CatTransactorSession) SetOpenSell(_open bool) (*types.Transaction, error) {
	return _Cat.Contract.SetOpenSell(&_Cat.TransactOpts, _open)
}

// SetSellRate is a paid mutator transaction binding the contract method 0x00f8e9fe.
//
// Solidity: function setSellRate(uint256 _rate, uint256 _base) returns()
func (_Cat *CatTransactor) SetSellRate(opts *bind.TransactOpts, _rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Cat.contract.Transact(opts, "setSellRate", _rate, _base)
}

// SetSellRate is a paid mutator transaction binding the contract method 0x00f8e9fe.
//
// Solidity: function setSellRate(uint256 _rate, uint256 _base) returns()
func (_Cat *CatSession) SetSellRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Cat.Contract.SetSellRate(&_Cat.TransactOpts, _rate, _base)
}

// SetSellRate is a paid mutator transaction binding the contract method 0x00f8e9fe.
//
// Solidity: function setSellRate(uint256 _rate, uint256 _base) returns()
func (_Cat *CatTransactorSession) SetSellRate(_rate *big.Int, _base *big.Int) (*types.Transaction, error) {
	return _Cat.Contract.SetSellRate(&_Cat.TransactOpts, _rate, _base)
}

// SetWhiteMap is a paid mutator transaction binding the contract method 0xa08d970c.
//
// Solidity: function setWhiteMap(address[] users, bool value) returns()
func (_Cat *CatTransactor) SetWhiteMap(opts *bind.TransactOpts, users []common.Address, value bool) (*types.Transaction, error) {
	return _Cat.contract.Transact(opts, "setWhiteMap", users, value)
}

// SetWhiteMap is a paid mutator transaction binding the contract method 0xa08d970c.
//
// Solidity: function setWhiteMap(address[] users, bool value) returns()
func (_Cat *CatSession) SetWhiteMap(users []common.Address, value bool) (*types.Transaction, error) {
	return _Cat.Contract.SetWhiteMap(&_Cat.TransactOpts, users, value)
}

// SetWhiteMap is a paid mutator transaction binding the contract method 0xa08d970c.
//
// Solidity: function setWhiteMap(address[] users, bool value) returns()
func (_Cat *CatTransactorSession) SetWhiteMap(users []common.Address, value bool) (*types.Transaction, error) {
	return _Cat.Contract.SetWhiteMap(&_Cat.TransactOpts, users, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cat *CatTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cat.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cat *CatSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cat.Contract.Transfer(&_Cat.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cat *CatTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cat.Contract.Transfer(&_Cat.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cat *CatTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cat.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cat *CatSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cat.Contract.TransferFrom(&_Cat.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cat *CatTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cat.Contract.TransferFrom(&_Cat.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cat *CatTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cat.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cat *CatSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cat.Contract.TransferOwnership(&_Cat.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cat *CatTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cat.Contract.TransferOwnership(&_Cat.TransactOpts, newOwner)
}

// CatApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Cat contract.
type CatApprovalIterator struct {
	Event *CatApproval // Event containing the contract specifics and raw log

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
func (it *CatApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CatApproval)
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
		it.Event = new(CatApproval)
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
func (it *CatApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CatApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CatApproval represents a Approval event raised by the Cat contract.
type CatApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Cat *CatFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CatApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cat.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CatApprovalIterator{contract: _Cat.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Cat *CatFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CatApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cat.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CatApproval)
				if err := _Cat.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Cat *CatFilterer) ParseApproval(log types.Log) (*CatApproval, error) {
	event := new(CatApproval)
	if err := _Cat.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CatTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Cat contract.
type CatTransferIterator struct {
	Event *CatTransfer // Event containing the contract specifics and raw log

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
func (it *CatTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CatTransfer)
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
		it.Event = new(CatTransfer)
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
func (it *CatTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CatTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CatTransfer represents a Transfer event raised by the Cat contract.
type CatTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cat *CatFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CatTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cat.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CatTransferIterator{contract: _Cat.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cat *CatFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CatTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cat.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CatTransfer)
				if err := _Cat.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Cat *CatFilterer) ParseTransfer(log types.Log) (*CatTransfer, error) {
	event := new(CatTransfer)
	if err := _Cat.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
