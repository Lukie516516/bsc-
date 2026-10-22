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

// AdminNewMetaData contains all meta data concerning the AdminNew contract.
var AdminNewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aicat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"buyAICAT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reqOrderIdArray\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sendToDead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wbnb\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AdminNewABI is the input ABI used to generate the binding from.
// Deprecated: Use AdminNewMetaData.ABI instead.
var AdminNewABI = AdminNewMetaData.ABI

// AdminNew is an auto generated Go binding around an Ethereum contract.
type AdminNew struct {
	AdminNewCaller     // Read-only binding to the contract
	AdminNewTransactor // Write-only binding to the contract
	AdminNewFilterer   // Log filterer for contract events
}

// AdminNewCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminNewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminNewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminNewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminNewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminNewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminNewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminNewSession struct {
	Contract     *AdminNew         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminNewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminNewCallerSession struct {
	Contract *AdminNewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AdminNewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminNewTransactorSession struct {
	Contract     *AdminNewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AdminNewRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminNewRaw struct {
	Contract *AdminNew // Generic contract binding to access the raw methods on
}

// AdminNewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminNewCallerRaw struct {
	Contract *AdminNewCaller // Generic read-only contract binding to access the raw methods on
}

// AdminNewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminNewTransactorRaw struct {
	Contract *AdminNewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdminNew creates a new instance of AdminNew, bound to a specific deployed contract.
func NewAdminNew(address common.Address, backend bind.ContractBackend) (*AdminNew, error) {
	contract, err := bindAdminNew(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdminNew{AdminNewCaller: AdminNewCaller{contract: contract}, AdminNewTransactor: AdminNewTransactor{contract: contract}, AdminNewFilterer: AdminNewFilterer{contract: contract}}, nil
}

// NewAdminNewCaller creates a new read-only instance of AdminNew, bound to a specific deployed contract.
func NewAdminNewCaller(address common.Address, caller bind.ContractCaller) (*AdminNewCaller, error) {
	contract, err := bindAdminNew(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminNewCaller{contract: contract}, nil
}

// NewAdminNewTransactor creates a new write-only instance of AdminNew, bound to a specific deployed contract.
func NewAdminNewTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminNewTransactor, error) {
	contract, err := bindAdminNew(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminNewTransactor{contract: contract}, nil
}

// NewAdminNewFilterer creates a new log filterer instance of AdminNew, bound to a specific deployed contract.
func NewAdminNewFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminNewFilterer, error) {
	contract, err := bindAdminNew(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminNewFilterer{contract: contract}, nil
}

// bindAdminNew binds a generic wrapper to an already deployed contract.
func bindAdminNew(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminNewABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminNew *AdminNewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdminNew.Contract.AdminNewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminNew *AdminNewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminNew.Contract.AdminNewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminNew *AdminNewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminNew.Contract.AdminNewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminNew *AdminNewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdminNew.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminNew *AdminNewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminNew.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminNew *AdminNewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminNew.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AdminNew *AdminNewCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdminNew.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AdminNew *AdminNewSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AdminNew.Contract.DEFAULTADMINROLE(&_AdminNew.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AdminNew *AdminNewCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AdminNew.Contract.DEFAULTADMINROLE(&_AdminNew.CallOpts)
}

// Aicat is a free data retrieval call binding the contract method 0x5c767d9a.
//
// Solidity: function aicat() view returns(address)
func (_AdminNew *AdminNewCaller) Aicat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdminNew.contract.Call(opts, &out, "aicat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aicat is a free data retrieval call binding the contract method 0x5c767d9a.
//
// Solidity: function aicat() view returns(address)
func (_AdminNew *AdminNewSession) Aicat() (common.Address, error) {
	return _AdminNew.Contract.Aicat(&_AdminNew.CallOpts)
}

// Aicat is a free data retrieval call binding the contract method 0x5c767d9a.
//
// Solidity: function aicat() view returns(address)
func (_AdminNew *AdminNewCallerSession) Aicat() (common.Address, error) {
	return _AdminNew.Contract.Aicat(&_AdminNew.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AdminNew *AdminNewCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AdminNew.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AdminNew *AdminNewSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AdminNew.Contract.GetRoleAdmin(&_AdminNew.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AdminNew *AdminNewCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AdminNew.Contract.GetRoleAdmin(&_AdminNew.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AdminNew *AdminNewCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AdminNew.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AdminNew *AdminNewSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AdminNew.Contract.HasRole(&_AdminNew.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AdminNew *AdminNewCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AdminNew.Contract.HasRole(&_AdminNew.CallOpts, role, account)
}

// ReqOrderIdArray is a free data retrieval call binding the contract method 0xb14819d1.
//
// Solidity: function reqOrderIdArray(uint256 ) view returns(uint256)
func (_AdminNew *AdminNewCaller) ReqOrderIdArray(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdminNew.contract.Call(opts, &out, "reqOrderIdArray", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReqOrderIdArray is a free data retrieval call binding the contract method 0xb14819d1.
//
// Solidity: function reqOrderIdArray(uint256 ) view returns(uint256)
func (_AdminNew *AdminNewSession) ReqOrderIdArray(arg0 *big.Int) (*big.Int, error) {
	return _AdminNew.Contract.ReqOrderIdArray(&_AdminNew.CallOpts, arg0)
}

// ReqOrderIdArray is a free data retrieval call binding the contract method 0xb14819d1.
//
// Solidity: function reqOrderIdArray(uint256 ) view returns(uint256)
func (_AdminNew *AdminNewCallerSession) ReqOrderIdArray(arg0 *big.Int) (*big.Int, error) {
	return _AdminNew.Contract.ReqOrderIdArray(&_AdminNew.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AdminNew *AdminNewCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AdminNew.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AdminNew *AdminNewSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AdminNew.Contract.SupportsInterface(&_AdminNew.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AdminNew *AdminNewCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AdminNew.Contract.SupportsInterface(&_AdminNew.CallOpts, interfaceId)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_AdminNew *AdminNewCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdminNew.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_AdminNew *AdminNewSession) UniswapV2Router() (common.Address, error) {
	return _AdminNew.Contract.UniswapV2Router(&_AdminNew.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_AdminNew *AdminNewCallerSession) UniswapV2Router() (common.Address, error) {
	return _AdminNew.Contract.UniswapV2Router(&_AdminNew.CallOpts)
}

// Wbnb is a free data retrieval call binding the contract method 0x8d72647e.
//
// Solidity: function wbnb() view returns(address)
func (_AdminNew *AdminNewCaller) Wbnb(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdminNew.contract.Call(opts, &out, "wbnb")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wbnb is a free data retrieval call binding the contract method 0x8d72647e.
//
// Solidity: function wbnb() view returns(address)
func (_AdminNew *AdminNewSession) Wbnb() (common.Address, error) {
	return _AdminNew.Contract.Wbnb(&_AdminNew.CallOpts)
}

// Wbnb is a free data retrieval call binding the contract method 0x8d72647e.
//
// Solidity: function wbnb() view returns(address)
func (_AdminNew *AdminNewCallerSession) Wbnb() (common.Address, error) {
	return _AdminNew.Contract.Wbnb(&_AdminNew.CallOpts)
}

// BuyAICAT is a paid mutator transaction binding the contract method 0x02a49bf8.
//
// Solidity: function buyAICAT(uint256 orderId) payable returns()
func (_AdminNew *AdminNewTransactor) BuyAICAT(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _AdminNew.contract.Transact(opts, "buyAICAT", orderId)
}

// BuyAICAT is a paid mutator transaction binding the contract method 0x02a49bf8.
//
// Solidity: function buyAICAT(uint256 orderId) payable returns()
func (_AdminNew *AdminNewSession) BuyAICAT(orderId *big.Int) (*types.Transaction, error) {
	return _AdminNew.Contract.BuyAICAT(&_AdminNew.TransactOpts, orderId)
}

// BuyAICAT is a paid mutator transaction binding the contract method 0x02a49bf8.
//
// Solidity: function buyAICAT(uint256 orderId) payable returns()
func (_AdminNew *AdminNewTransactorSession) BuyAICAT(orderId *big.Int) (*types.Transaction, error) {
	return _AdminNew.Contract.BuyAICAT(&_AdminNew.TransactOpts, orderId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AdminNew *AdminNewTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdminNew.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AdminNew *AdminNewSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdminNew.Contract.GrantRole(&_AdminNew.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AdminNew *AdminNewTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdminNew.Contract.GrantRole(&_AdminNew.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AdminNew *AdminNewTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AdminNew.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AdminNew *AdminNewSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AdminNew.Contract.RenounceRole(&_AdminNew.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AdminNew *AdminNewTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AdminNew.Contract.RenounceRole(&_AdminNew.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AdminNew *AdminNewTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdminNew.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AdminNew *AdminNewSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdminNew.Contract.RevokeRole(&_AdminNew.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AdminNew *AdminNewTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdminNew.Contract.RevokeRole(&_AdminNew.TransactOpts, role, account)
}

// SendToDead is a paid mutator transaction binding the contract method 0xeba735a5.
//
// Solidity: function sendToDead() returns()
func (_AdminNew *AdminNewTransactor) SendToDead(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminNew.contract.Transact(opts, "sendToDead")
}

// SendToDead is a paid mutator transaction binding the contract method 0xeba735a5.
//
// Solidity: function sendToDead() returns()
func (_AdminNew *AdminNewSession) SendToDead() (*types.Transaction, error) {
	return _AdminNew.Contract.SendToDead(&_AdminNew.TransactOpts)
}

// SendToDead is a paid mutator transaction binding the contract method 0xeba735a5.
//
// Solidity: function sendToDead() returns()
func (_AdminNew *AdminNewTransactorSession) SendToDead() (*types.Transaction, error) {
	return _AdminNew.Contract.SendToDead(&_AdminNew.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AdminNew *AdminNewTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminNew.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AdminNew *AdminNewSession) Receive() (*types.Transaction, error) {
	return _AdminNew.Contract.Receive(&_AdminNew.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AdminNew *AdminNewTransactorSession) Receive() (*types.Transaction, error) {
	return _AdminNew.Contract.Receive(&_AdminNew.TransactOpts)
}

// AdminNewRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AdminNew contract.
type AdminNewRoleAdminChangedIterator struct {
	Event *AdminNewRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AdminNewRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminNewRoleAdminChanged)
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
		it.Event = new(AdminNewRoleAdminChanged)
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
func (it *AdminNewRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminNewRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminNewRoleAdminChanged represents a RoleAdminChanged event raised by the AdminNew contract.
type AdminNewRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AdminNew *AdminNewFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AdminNewRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AdminNew.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AdminNewRoleAdminChangedIterator{contract: _AdminNew.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AdminNew *AdminNewFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AdminNewRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AdminNew.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminNewRoleAdminChanged)
				if err := _AdminNew.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AdminNew *AdminNewFilterer) ParseRoleAdminChanged(log types.Log) (*AdminNewRoleAdminChanged, error) {
	event := new(AdminNewRoleAdminChanged)
	if err := _AdminNew.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminNewRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AdminNew contract.
type AdminNewRoleGrantedIterator struct {
	Event *AdminNewRoleGranted // Event containing the contract specifics and raw log

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
func (it *AdminNewRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminNewRoleGranted)
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
		it.Event = new(AdminNewRoleGranted)
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
func (it *AdminNewRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminNewRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminNewRoleGranted represents a RoleGranted event raised by the AdminNew contract.
type AdminNewRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AdminNew *AdminNewFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AdminNewRoleGrantedIterator, error) {

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

	logs, sub, err := _AdminNew.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AdminNewRoleGrantedIterator{contract: _AdminNew.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AdminNew *AdminNewFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AdminNewRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AdminNew.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminNewRoleGranted)
				if err := _AdminNew.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AdminNew *AdminNewFilterer) ParseRoleGranted(log types.Log) (*AdminNewRoleGranted, error) {
	event := new(AdminNewRoleGranted)
	if err := _AdminNew.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminNewRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AdminNew contract.
type AdminNewRoleRevokedIterator struct {
	Event *AdminNewRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AdminNewRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminNewRoleRevoked)
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
		it.Event = new(AdminNewRoleRevoked)
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
func (it *AdminNewRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminNewRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminNewRoleRevoked represents a RoleRevoked event raised by the AdminNew contract.
type AdminNewRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AdminNew *AdminNewFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AdminNewRoleRevokedIterator, error) {

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

	logs, sub, err := _AdminNew.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AdminNewRoleRevokedIterator{contract: _AdminNew.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AdminNew *AdminNewFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AdminNewRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AdminNew.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminNewRoleRevoked)
				if err := _AdminNew.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AdminNew *AdminNewFilterer) ParseRoleRevoked(log types.Log) (*AdminNewRoleRevoked, error) {
	event := new(AdminNewRoleRevoked)
	if err := _AdminNew.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
