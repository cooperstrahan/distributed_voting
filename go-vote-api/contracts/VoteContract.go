// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractsABI is the input ABI used to generate the binding from.
const ContractsABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_v\",\"type\":\"bytes32\"}],\"name\":\"checkCastVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_v\",\"type\":\"bytes32\"}],\"name\":\"checkVoterRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_v\",\"type\":\"bytes32\"}],\"name\":\"recordVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_v\",\"type\":\"bytes32\"}],\"name\":\"registerVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractsBin is the compiled bytecode used for deploying new contracts.
var ContractsBin = "0x608060405234801561001057600080fd5b506101ae806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806310103b7214610051578063793c5d2114610088578063c27f59011461009b578063eed50256146100be575b600080fd5b61007461005f366004610160565b60009081526020819052604090205460ff1690565b604051901515815260200160405180910390f35b610074610096366004610160565b6100d1565b6100746100a9366004610160565b60009081526001602052604090205460ff1690565b6100746100cc366004610160565b61010f565b60008181526020819052604081205460ff1661010757506000908152602081905260409020805460ff1916600190811790915590565b506000919050565b60008181526020819052604081205460ff16801561013c575060008281526001602052604090205460ff16155b1561010757506000908152600160208190526040909120805460ff19168217905590565b600060208284031215610171578081fd5b503591905056fea26469706673582212208030ad7310bf8b8471983c2931bef2c802b969a41b03f4ca00882dc22e26708064736f6c63430008040033"

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// CheckCastVote is a free data retrieval call binding the contract method 0xc27f5901.
//
// Solidity: function checkCastVote(bytes32 _v) view returns(bool)
func (_Contracts *ContractsCaller) CheckCastVote(opts *bind.CallOpts, _v [32]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "checkCastVote", _v)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckCastVote is a free data retrieval call binding the contract method 0xc27f5901.
//
// Solidity: function checkCastVote(bytes32 _v) view returns(bool)
func (_Contracts *ContractsSession) CheckCastVote(_v [32]byte) (bool, error) {
	return _Contracts.Contract.CheckCastVote(&_Contracts.CallOpts, _v)
}

// CheckCastVote is a free data retrieval call binding the contract method 0xc27f5901.
//
// Solidity: function checkCastVote(bytes32 _v) view returns(bool)
func (_Contracts *ContractsCallerSession) CheckCastVote(_v [32]byte) (bool, error) {
	return _Contracts.Contract.CheckCastVote(&_Contracts.CallOpts, _v)
}

// CheckVoterRegistered is a free data retrieval call binding the contract method 0x10103b72.
//
// Solidity: function checkVoterRegistered(bytes32 _v) view returns(bool)
func (_Contracts *ContractsCaller) CheckVoterRegistered(opts *bind.CallOpts, _v [32]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "checkVoterRegistered", _v)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckVoterRegistered is a free data retrieval call binding the contract method 0x10103b72.
//
// Solidity: function checkVoterRegistered(bytes32 _v) view returns(bool)
func (_Contracts *ContractsSession) CheckVoterRegistered(_v [32]byte) (bool, error) {
	return _Contracts.Contract.CheckVoterRegistered(&_Contracts.CallOpts, _v)
}

// CheckVoterRegistered is a free data retrieval call binding the contract method 0x10103b72.
//
// Solidity: function checkVoterRegistered(bytes32 _v) view returns(bool)
func (_Contracts *ContractsCallerSession) CheckVoterRegistered(_v [32]byte) (bool, error) {
	return _Contracts.Contract.CheckVoterRegistered(&_Contracts.CallOpts, _v)
}

// RecordVote is a paid mutator transaction binding the contract method 0xeed50256.
//
// Solidity: function recordVote(bytes32 _v) returns(bool)
func (_Contracts *ContractsTransactor) RecordVote(opts *bind.TransactOpts, _v [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "recordVote", _v)
}

// RecordVote is a paid mutator transaction binding the contract method 0xeed50256.
//
// Solidity: function recordVote(bytes32 _v) returns(bool)
func (_Contracts *ContractsSession) RecordVote(_v [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.RecordVote(&_Contracts.TransactOpts, _v)
}

// RecordVote is a paid mutator transaction binding the contract method 0xeed50256.
//
// Solidity: function recordVote(bytes32 _v) returns(bool)
func (_Contracts *ContractsTransactorSession) RecordVote(_v [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.RecordVote(&_Contracts.TransactOpts, _v)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x793c5d21.
//
// Solidity: function registerVoter(bytes32 _v) returns(bool)
func (_Contracts *ContractsTransactor) RegisterVoter(opts *bind.TransactOpts, _v [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerVoter", _v)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x793c5d21.
//
// Solidity: function registerVoter(bytes32 _v) returns(bool)
func (_Contracts *ContractsSession) RegisterVoter(_v [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterVoter(&_Contracts.TransactOpts, _v)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x793c5d21.
//
// Solidity: function registerVoter(bytes32 _v) returns(bool)
func (_Contracts *ContractsTransactorSession) RegisterVoter(_v [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterVoter(&_Contracts.TransactOpts, _v)
}
