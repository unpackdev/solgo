// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc1967

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

// ERC1967MetaData contains all meta data concerning the ERC1967 contract.
var ERC1967MetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"setInterfaceImplementer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getInterfaceImplementer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"string\"}],\"name\":\"interfaceHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"updateERC165Cache\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"implementsERC165InterfaceNoCache\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"implementsERC165Interface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"\",\"type\":\"address\"}],\"name\":\"InterfaceImplementerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"}]",
}

// ERC1967ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1967MetaData.ABI instead.
var ERC1967ABI = ERC1967MetaData.ABI

// ERC1967 is an auto generated Go binding around an Ethereum contract.
type ERC1967 struct {
	ERC1967Caller     // Read-only binding to the contract
	ERC1967Transactor // Write-only binding to the contract
	ERC1967Filterer   // Log filterer for contract events
}

// ERC1967Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1967Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1967Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1967Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1967Session struct {
	Contract     *ERC1967          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1967CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1967CallerSession struct {
	Contract *ERC1967Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC1967TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1967TransactorSession struct {
	Contract     *ERC1967Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC1967Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1967Raw struct {
	Contract *ERC1967 // Generic contract binding to access the raw methods on
}

// ERC1967CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1967CallerRaw struct {
	Contract *ERC1967Caller // Generic read-only contract binding to access the raw methods on
}

// ERC1967TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1967TransactorRaw struct {
	Contract *ERC1967Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1967 creates a new instance of ERC1967, bound to a specific deployed contract.
func NewERC1967(address common.Address, backend bind.ContractBackend) (*ERC1967, error) {
	contract, err := bindERC1967(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1967{ERC1967Caller: ERC1967Caller{contract: contract}, ERC1967Transactor: ERC1967Transactor{contract: contract}, ERC1967Filterer: ERC1967Filterer{contract: contract}}, nil
}

// NewERC1967Caller creates a new read-only instance of ERC1967, bound to a specific deployed contract.
func NewERC1967Caller(address common.Address, caller bind.ContractCaller) (*ERC1967Caller, error) {
	contract, err := bindERC1967(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967Caller{contract: contract}, nil
}

// NewERC1967Transactor creates a new write-only instance of ERC1967, bound to a specific deployed contract.
func NewERC1967Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC1967Transactor, error) {
	contract, err := bindERC1967(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967Transactor{contract: contract}, nil
}

// NewERC1967Filterer creates a new log filterer instance of ERC1967, bound to a specific deployed contract.
func NewERC1967Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC1967Filterer, error) {
	contract, err := bindERC1967(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1967Filterer{contract: contract}, nil
}

// bindERC1967 binds a generic wrapper to an already deployed contract.
func bindERC1967(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1967ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967 *ERC1967Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967.Contract.ERC1967Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967 *ERC1967Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967.Contract.ERC1967Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967 *ERC1967Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967.Contract.ERC1967Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967 *ERC1967CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967 *ERC1967TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967 *ERC1967TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967.Contract.contract.Transact(opts, method, params...)
}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address , bytes32 ) view returns(address)
func (_ERC1967 *ERC1967Caller) GetInterfaceImplementer(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC1967.contract.Call(opts, &out, "getInterfaceImplementer", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address , bytes32 ) view returns(address)
func (_ERC1967 *ERC1967Session) GetInterfaceImplementer(arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	return _ERC1967.Contract.GetInterfaceImplementer(&_ERC1967.CallOpts, arg0, arg1)
}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address , bytes32 ) view returns(address)
func (_ERC1967 *ERC1967CallerSession) GetInterfaceImplementer(arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	return _ERC1967.Contract.GetInterfaceImplementer(&_ERC1967.CallOpts, arg0, arg1)
}

// ImplementsERC165Interface is a free data retrieval call binding the contract method 0xc03c5d2e.
//
// Solidity: function implementsERC165Interface(address , bytes32 ) view returns(bool)
func (_ERC1967 *ERC1967Caller) ImplementsERC165Interface(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _ERC1967.contract.Call(opts, &out, "implementsERC165Interface", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ImplementsERC165Interface is a free data retrieval call binding the contract method 0xc03c5d2e.
//
// Solidity: function implementsERC165Interface(address , bytes32 ) view returns(bool)
func (_ERC1967 *ERC1967Session) ImplementsERC165Interface(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _ERC1967.Contract.ImplementsERC165Interface(&_ERC1967.CallOpts, arg0, arg1)
}

// ImplementsERC165Interface is a free data retrieval call binding the contract method 0xc03c5d2e.
//
// Solidity: function implementsERC165Interface(address , bytes32 ) view returns(bool)
func (_ERC1967 *ERC1967CallerSession) ImplementsERC165Interface(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _ERC1967.Contract.ImplementsERC165Interface(&_ERC1967.CallOpts, arg0, arg1)
}

// ImplementsERC165InterfaceNoCache is a free data retrieval call binding the contract method 0x49aa8eba.
//
// Solidity: function implementsERC165InterfaceNoCache(address , bytes32 ) view returns(bool)
func (_ERC1967 *ERC1967Caller) ImplementsERC165InterfaceNoCache(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _ERC1967.contract.Call(opts, &out, "implementsERC165InterfaceNoCache", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ImplementsERC165InterfaceNoCache is a free data retrieval call binding the contract method 0x49aa8eba.
//
// Solidity: function implementsERC165InterfaceNoCache(address , bytes32 ) view returns(bool)
func (_ERC1967 *ERC1967Session) ImplementsERC165InterfaceNoCache(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _ERC1967.Contract.ImplementsERC165InterfaceNoCache(&_ERC1967.CallOpts, arg0, arg1)
}

// ImplementsERC165InterfaceNoCache is a free data retrieval call binding the contract method 0x49aa8eba.
//
// Solidity: function implementsERC165InterfaceNoCache(address , bytes32 ) view returns(bool)
func (_ERC1967 *ERC1967CallerSession) ImplementsERC165InterfaceNoCache(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _ERC1967.Contract.ImplementsERC165InterfaceNoCache(&_ERC1967.CallOpts, arg0, arg1)
}

// InterfaceHash is a free data retrieval call binding the contract method 0x65ba36c1.
//
// Solidity: function interfaceHash(string ) view returns(bytes32)
func (_ERC1967 *ERC1967Caller) InterfaceHash(opts *bind.CallOpts, arg0 string) ([32]byte, error) {
	var out []interface{}
	err := _ERC1967.contract.Call(opts, &out, "interfaceHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InterfaceHash is a free data retrieval call binding the contract method 0x65ba36c1.
//
// Solidity: function interfaceHash(string ) view returns(bytes32)
func (_ERC1967 *ERC1967Session) InterfaceHash(arg0 string) ([32]byte, error) {
	return _ERC1967.Contract.InterfaceHash(&_ERC1967.CallOpts, arg0)
}

// InterfaceHash is a free data retrieval call binding the contract method 0x65ba36c1.
//
// Solidity: function interfaceHash(string ) view returns(bytes32)
func (_ERC1967 *ERC1967CallerSession) InterfaceHash(arg0 string) ([32]byte, error) {
	return _ERC1967.Contract.InterfaceHash(&_ERC1967.CallOpts, arg0)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x29965a1d.
//
// Solidity: function setInterfaceImplementer(address , bytes32 , address ) returns()
func (_ERC1967 *ERC1967Transactor) SetInterfaceImplementer(opts *bind.TransactOpts, arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*types.Transaction, error) {
	return _ERC1967.contract.Transact(opts, "setInterfaceImplementer", arg0, arg1, arg2)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x29965a1d.
//
// Solidity: function setInterfaceImplementer(address , bytes32 , address ) returns()
func (_ERC1967 *ERC1967Session) SetInterfaceImplementer(arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*types.Transaction, error) {
	return _ERC1967.Contract.SetInterfaceImplementer(&_ERC1967.TransactOpts, arg0, arg1, arg2)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x29965a1d.
//
// Solidity: function setInterfaceImplementer(address , bytes32 , address ) returns()
func (_ERC1967 *ERC1967TransactorSession) SetInterfaceImplementer(arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*types.Transaction, error) {
	return _ERC1967.Contract.SetInterfaceImplementer(&_ERC1967.TransactOpts, arg0, arg1, arg2)
}

// UpdateERC165Cache is a paid mutator transaction binding the contract method 0x6a73781d.
//
// Solidity: function updateERC165Cache(address , bytes32 ) returns()
func (_ERC1967 *ERC1967Transactor) UpdateERC165Cache(opts *bind.TransactOpts, arg0 common.Address, arg1 [32]byte) (*types.Transaction, error) {
	return _ERC1967.contract.Transact(opts, "updateERC165Cache", arg0, arg1)
}

// UpdateERC165Cache is a paid mutator transaction binding the contract method 0x6a73781d.
//
// Solidity: function updateERC165Cache(address , bytes32 ) returns()
func (_ERC1967 *ERC1967Session) UpdateERC165Cache(arg0 common.Address, arg1 [32]byte) (*types.Transaction, error) {
	return _ERC1967.Contract.UpdateERC165Cache(&_ERC1967.TransactOpts, arg0, arg1)
}

// UpdateERC165Cache is a paid mutator transaction binding the contract method 0x6a73781d.
//
// Solidity: function updateERC165Cache(address , bytes32 ) returns()
func (_ERC1967 *ERC1967TransactorSession) UpdateERC165Cache(arg0 common.Address, arg1 [32]byte) (*types.Transaction, error) {
	return _ERC1967.Contract.UpdateERC165Cache(&_ERC1967.TransactOpts, arg0, arg1)
}

// ERC1967AdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ERC1967 contract.
type ERC1967AdminChangedIterator struct {
	Event *ERC1967AdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC1967AdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967AdminChanged)
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
		it.Event = new(ERC1967AdminChanged)
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
func (it *ERC1967AdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967AdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967AdminChanged represents a AdminChanged event raised by the ERC1967 contract.
type ERC1967AdminChanged struct {
	Arg0 common.Address
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed arg0, address indexed arg1)
func (_ERC1967 *ERC1967Filterer) FilterAdminChanged(opts *bind.FilterOpts, arg0 []common.Address, arg1 []common.Address) (*ERC1967AdminChangedIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _ERC1967.contract.FilterLogs(opts, "AdminChanged", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return &ERC1967AdminChangedIterator{contract: _ERC1967.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed arg0, address indexed arg1)
func (_ERC1967 *ERC1967Filterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC1967AdminChanged, arg0 []common.Address, arg1 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _ERC1967.contract.WatchLogs(opts, "AdminChanged", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967AdminChanged)
				if err := _ERC1967.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed arg0, address indexed arg1)
func (_ERC1967 *ERC1967Filterer) ParseAdminChanged(log types.Log) (*ERC1967AdminChanged, error) {
	event := new(ERC1967AdminChanged)
	if err := _ERC1967.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967InterfaceImplementerSetIterator is returned from FilterInterfaceImplementerSet and is used to iterate over the raw logs and unpacked data for InterfaceImplementerSet events raised by the ERC1967 contract.
type ERC1967InterfaceImplementerSetIterator struct {
	Event *ERC1967InterfaceImplementerSet // Event containing the contract specifics and raw log

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
func (it *ERC1967InterfaceImplementerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967InterfaceImplementerSet)
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
		it.Event = new(ERC1967InterfaceImplementerSet)
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
func (it *ERC1967InterfaceImplementerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967InterfaceImplementerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967InterfaceImplementerSet represents a InterfaceImplementerSet event raised by the ERC1967 contract.
type ERC1967InterfaceImplementerSet struct {
	Arg0 common.Address
	Arg1 [32]byte
	Arg2 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterInterfaceImplementerSet is a free log retrieval operation binding the contract event 0x93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db153.
//
// Solidity: event InterfaceImplementerSet(address indexed arg0, bytes32 indexed arg1, address indexed arg2)
func (_ERC1967 *ERC1967Filterer) FilterInterfaceImplementerSet(opts *bind.FilterOpts, arg0 []common.Address, arg1 [][32]byte, arg2 []common.Address) (*ERC1967InterfaceImplementerSetIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}
	var arg2Rule []interface{}
	for _, arg2Item := range arg2 {
		arg2Rule = append(arg2Rule, arg2Item)
	}

	logs, sub, err := _ERC1967.contract.FilterLogs(opts, "InterfaceImplementerSet", arg0Rule, arg1Rule, arg2Rule)
	if err != nil {
		return nil, err
	}
	return &ERC1967InterfaceImplementerSetIterator{contract: _ERC1967.contract, event: "InterfaceImplementerSet", logs: logs, sub: sub}, nil
}

// WatchInterfaceImplementerSet is a free log subscription operation binding the contract event 0x93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db153.
//
// Solidity: event InterfaceImplementerSet(address indexed arg0, bytes32 indexed arg1, address indexed arg2)
func (_ERC1967 *ERC1967Filterer) WatchInterfaceImplementerSet(opts *bind.WatchOpts, sink chan<- *ERC1967InterfaceImplementerSet, arg0 []common.Address, arg1 [][32]byte, arg2 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}
	var arg2Rule []interface{}
	for _, arg2Item := range arg2 {
		arg2Rule = append(arg2Rule, arg2Item)
	}

	logs, sub, err := _ERC1967.contract.WatchLogs(opts, "InterfaceImplementerSet", arg0Rule, arg1Rule, arg2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967InterfaceImplementerSet)
				if err := _ERC1967.contract.UnpackLog(event, "InterfaceImplementerSet", log); err != nil {
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

// ParseInterfaceImplementerSet is a log parse operation binding the contract event 0x93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db153.
//
// Solidity: event InterfaceImplementerSet(address indexed arg0, bytes32 indexed arg1, address indexed arg2)
func (_ERC1967 *ERC1967Filterer) ParseInterfaceImplementerSet(log types.Log) (*ERC1967InterfaceImplementerSet, error) {
	event := new(ERC1967InterfaceImplementerSet)
	if err := _ERC1967.contract.UnpackLog(event, "InterfaceImplementerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
