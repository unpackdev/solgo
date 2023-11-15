// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc1820

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

// ERC1820MetaData contains all meta data concerning the ERC1820 contract.
var ERC1820MetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"interfaceHash\",\"type\":\"bytes32\"}],\"name\":\"getInterfaceImplementer\",\"outputs\":[{\"name\":\"implementer\",\"type\":\"address\"}],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"interfaceHash\",\"type\":\"bytes32\"},{\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"setInterfaceImplementer\",\"outputs\":[],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getManager\",\"outputs\":[{\"name\":\"manager\",\"type\":\"address\"}],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setInterfaceImplementer\",\"outputs\":[],\"type\":\"function\"},{\"inputs\":[],\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"interfaceHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"InterfaceImplementerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"ManagerChanged\",\"type\":\"event\"}]",
}

// ERC1820ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1820MetaData.ABI instead.
var ERC1820ABI = ERC1820MetaData.ABI

// ERC1820 is an auto generated Go binding around an Ethereum contract.
type ERC1820 struct {
	ERC1820Caller     // Read-only binding to the contract
	ERC1820Transactor // Write-only binding to the contract
	ERC1820Filterer   // Log filterer for contract events
}

// ERC1820Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1820Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1820Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1820Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1820Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1820Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1820Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1820Session struct {
	Contract     *ERC1820          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1820CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1820CallerSession struct {
	Contract *ERC1820Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC1820TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1820TransactorSession struct {
	Contract     *ERC1820Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC1820Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1820Raw struct {
	Contract *ERC1820 // Generic contract binding to access the raw methods on
}

// ERC1820CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1820CallerRaw struct {
	Contract *ERC1820Caller // Generic read-only contract binding to access the raw methods on
}

// ERC1820TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1820TransactorRaw struct {
	Contract *ERC1820Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1820 creates a new instance of ERC1820, bound to a specific deployed contract.
func NewERC1820(address common.Address, backend bind.ContractBackend) (*ERC1820, error) {
	contract, err := bindERC1820(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1820{ERC1820Caller: ERC1820Caller{contract: contract}, ERC1820Transactor: ERC1820Transactor{contract: contract}, ERC1820Filterer: ERC1820Filterer{contract: contract}}, nil
}

// NewERC1820Caller creates a new read-only instance of ERC1820, bound to a specific deployed contract.
func NewERC1820Caller(address common.Address, caller bind.ContractCaller) (*ERC1820Caller, error) {
	contract, err := bindERC1820(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1820Caller{contract: contract}, nil
}

// NewERC1820Transactor creates a new write-only instance of ERC1820, bound to a specific deployed contract.
func NewERC1820Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC1820Transactor, error) {
	contract, err := bindERC1820(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1820Transactor{contract: contract}, nil
}

// NewERC1820Filterer creates a new log filterer instance of ERC1820, bound to a specific deployed contract.
func NewERC1820Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC1820Filterer, error) {
	contract, err := bindERC1820(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1820Filterer{contract: contract}, nil
}

// bindERC1820 binds a generic wrapper to an already deployed contract.
func bindERC1820(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1820ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1820 *ERC1820Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1820.Contract.ERC1820Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1820 *ERC1820Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1820.Contract.ERC1820Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1820 *ERC1820Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1820.Contract.ERC1820Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1820 *ERC1820CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1820.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1820 *ERC1820TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1820.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1820 *ERC1820TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1820.Contract.contract.Transact(opts, method, params...)
}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address account, bytes32 interfaceHash) returns(address implementer)
func (_ERC1820 *ERC1820Caller) GetInterfaceImplementer(opts *bind.CallOpts, account common.Address, interfaceHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC1820.contract.Call(opts, &out, "getInterfaceImplementer", account, interfaceHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address account, bytes32 interfaceHash) returns(address implementer)
func (_ERC1820 *ERC1820Session) GetInterfaceImplementer(account common.Address, interfaceHash [32]byte) (common.Address, error) {
	return _ERC1820.Contract.GetInterfaceImplementer(&_ERC1820.CallOpts, account, interfaceHash)
}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address account, bytes32 interfaceHash) returns(address implementer)
func (_ERC1820 *ERC1820CallerSession) GetInterfaceImplementer(account common.Address, interfaceHash [32]byte) (common.Address, error) {
	return _ERC1820.Contract.GetInterfaceImplementer(&_ERC1820.CallOpts, account, interfaceHash)
}

// GetManager is a free data retrieval call binding the contract method 0x3d584063.
//
// Solidity: function getManager(address account) returns(address manager)
func (_ERC1820 *ERC1820Caller) GetManager(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC1820.contract.Call(opts, &out, "getManager", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetManager is a free data retrieval call binding the contract method 0x3d584063.
//
// Solidity: function getManager(address account) returns(address manager)
func (_ERC1820 *ERC1820Session) GetManager(account common.Address) (common.Address, error) {
	return _ERC1820.Contract.GetManager(&_ERC1820.CallOpts, account)
}

// GetManager is a free data retrieval call binding the contract method 0x3d584063.
//
// Solidity: function getManager(address account) returns(address manager)
func (_ERC1820 *ERC1820CallerSession) GetManager(account common.Address) (common.Address, error) {
	return _ERC1820.Contract.GetManager(&_ERC1820.CallOpts, account)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x989f4baf.
//
// Solidity: function setInterfaceImplementer(bytes32 interfaceHash, address implementer) returns()
func (_ERC1820 *ERC1820Transactor) SetInterfaceImplementer(opts *bind.TransactOpts, interfaceHash [32]byte, implementer common.Address) (*types.Transaction, error) {
	return _ERC1820.contract.Transact(opts, "setInterfaceImplementer", interfaceHash, implementer)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x989f4baf.
//
// Solidity: function setInterfaceImplementer(bytes32 interfaceHash, address implementer) returns()
func (_ERC1820 *ERC1820Session) SetInterfaceImplementer(interfaceHash [32]byte, implementer common.Address) (*types.Transaction, error) {
	return _ERC1820.Contract.SetInterfaceImplementer(&_ERC1820.TransactOpts, interfaceHash, implementer)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x989f4baf.
//
// Solidity: function setInterfaceImplementer(bytes32 interfaceHash, address implementer) returns()
func (_ERC1820 *ERC1820TransactorSession) SetInterfaceImplementer(interfaceHash [32]byte, implementer common.Address) (*types.Transaction, error) {
	return _ERC1820.Contract.SetInterfaceImplementer(&_ERC1820.TransactOpts, interfaceHash, implementer)
}

// SetInterfaceImplementer0 is a paid mutator transaction binding the contract method 0x73922911.
//
// Solidity: function setInterfaceImplementer(address account) returns()
func (_ERC1820 *ERC1820Transactor) SetInterfaceImplementer0(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ERC1820.contract.Transact(opts, "setInterfaceImplementer0", account)
}

// SetInterfaceImplementer0 is a paid mutator transaction binding the contract method 0x73922911.
//
// Solidity: function setInterfaceImplementer(address account) returns()
func (_ERC1820 *ERC1820Session) SetInterfaceImplementer0(account common.Address) (*types.Transaction, error) {
	return _ERC1820.Contract.SetInterfaceImplementer0(&_ERC1820.TransactOpts, account)
}

// SetInterfaceImplementer0 is a paid mutator transaction binding the contract method 0x73922911.
//
// Solidity: function setInterfaceImplementer(address account) returns()
func (_ERC1820 *ERC1820TransactorSession) SetInterfaceImplementer0(account common.Address) (*types.Transaction, error) {
	return _ERC1820.Contract.SetInterfaceImplementer0(&_ERC1820.TransactOpts, account)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address newManager) returns()
func (_ERC1820 *ERC1820Transactor) SetManager(opts *bind.TransactOpts, newManager common.Address) (*types.Transaction, error) {
	return _ERC1820.contract.Transact(opts, "setManager", newManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address newManager) returns()
func (_ERC1820 *ERC1820Session) SetManager(newManager common.Address) (*types.Transaction, error) {
	return _ERC1820.Contract.SetManager(&_ERC1820.TransactOpts, newManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address newManager) returns()
func (_ERC1820 *ERC1820TransactorSession) SetManager(newManager common.Address) (*types.Transaction, error) {
	return _ERC1820.Contract.SetManager(&_ERC1820.TransactOpts, newManager)
}

// ERC1820InterfaceImplementerSetIterator is returned from FilterInterfaceImplementerSet and is used to iterate over the raw logs and unpacked data for InterfaceImplementerSet events raised by the ERC1820 contract.
type ERC1820InterfaceImplementerSetIterator struct {
	Event *ERC1820InterfaceImplementerSet // Event containing the contract specifics and raw log

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
func (it *ERC1820InterfaceImplementerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1820InterfaceImplementerSet)
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
		it.Event = new(ERC1820InterfaceImplementerSet)
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
func (it *ERC1820InterfaceImplementerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1820InterfaceImplementerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1820InterfaceImplementerSet represents a InterfaceImplementerSet event raised by the ERC1820 contract.
type ERC1820InterfaceImplementerSet struct {
	Account       common.Address
	InterfaceHash [32]byte
	Implementer   common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInterfaceImplementerSet is a free log retrieval operation binding the contract event 0x93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db153.
//
// Solidity: event InterfaceImplementerSet(address indexed account, bytes32 indexed interfaceHash, address indexed implementer)
func (_ERC1820 *ERC1820Filterer) FilterInterfaceImplementerSet(opts *bind.FilterOpts, account []common.Address, interfaceHash [][32]byte, implementer []common.Address) (*ERC1820InterfaceImplementerSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var interfaceHashRule []interface{}
	for _, interfaceHashItem := range interfaceHash {
		interfaceHashRule = append(interfaceHashRule, interfaceHashItem)
	}
	var implementerRule []interface{}
	for _, implementerItem := range implementer {
		implementerRule = append(implementerRule, implementerItem)
	}

	logs, sub, err := _ERC1820.contract.FilterLogs(opts, "InterfaceImplementerSet", accountRule, interfaceHashRule, implementerRule)
	if err != nil {
		return nil, err
	}
	return &ERC1820InterfaceImplementerSetIterator{contract: _ERC1820.contract, event: "InterfaceImplementerSet", logs: logs, sub: sub}, nil
}

// WatchInterfaceImplementerSet is a free log subscription operation binding the contract event 0x93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db153.
//
// Solidity: event InterfaceImplementerSet(address indexed account, bytes32 indexed interfaceHash, address indexed implementer)
func (_ERC1820 *ERC1820Filterer) WatchInterfaceImplementerSet(opts *bind.WatchOpts, sink chan<- *ERC1820InterfaceImplementerSet, account []common.Address, interfaceHash [][32]byte, implementer []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var interfaceHashRule []interface{}
	for _, interfaceHashItem := range interfaceHash {
		interfaceHashRule = append(interfaceHashRule, interfaceHashItem)
	}
	var implementerRule []interface{}
	for _, implementerItem := range implementer {
		implementerRule = append(implementerRule, implementerItem)
	}

	logs, sub, err := _ERC1820.contract.WatchLogs(opts, "InterfaceImplementerSet", accountRule, interfaceHashRule, implementerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1820InterfaceImplementerSet)
				if err := _ERC1820.contract.UnpackLog(event, "InterfaceImplementerSet", log); err != nil {
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
// Solidity: event InterfaceImplementerSet(address indexed account, bytes32 indexed interfaceHash, address indexed implementer)
func (_ERC1820 *ERC1820Filterer) ParseInterfaceImplementerSet(log types.Log) (*ERC1820InterfaceImplementerSet, error) {
	event := new(ERC1820InterfaceImplementerSet)
	if err := _ERC1820.contract.UnpackLog(event, "InterfaceImplementerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1820ManagerChangedIterator is returned from FilterManagerChanged and is used to iterate over the raw logs and unpacked data for ManagerChanged events raised by the ERC1820 contract.
type ERC1820ManagerChangedIterator struct {
	Event *ERC1820ManagerChanged // Event containing the contract specifics and raw log

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
func (it *ERC1820ManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1820ManagerChanged)
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
		it.Event = new(ERC1820ManagerChanged)
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
func (it *ERC1820ManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1820ManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1820ManagerChanged represents a ManagerChanged event raised by the ERC1820 contract.
type ERC1820ManagerChanged struct {
	Account    common.Address
	NewManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterManagerChanged is a free log retrieval operation binding the contract event 0x605c2dbf762e5f7d60a546d42e7205dcb1b011ebc62a61736a57c9089d3a4350.
//
// Solidity: event ManagerChanged(address indexed account, address indexed newManager)
func (_ERC1820 *ERC1820Filterer) FilterManagerChanged(opts *bind.FilterOpts, account []common.Address, newManager []common.Address) (*ERC1820ManagerChangedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _ERC1820.contract.FilterLogs(opts, "ManagerChanged", accountRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return &ERC1820ManagerChangedIterator{contract: _ERC1820.contract, event: "ManagerChanged", logs: logs, sub: sub}, nil
}

// WatchManagerChanged is a free log subscription operation binding the contract event 0x605c2dbf762e5f7d60a546d42e7205dcb1b011ebc62a61736a57c9089d3a4350.
//
// Solidity: event ManagerChanged(address indexed account, address indexed newManager)
func (_ERC1820 *ERC1820Filterer) WatchManagerChanged(opts *bind.WatchOpts, sink chan<- *ERC1820ManagerChanged, account []common.Address, newManager []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _ERC1820.contract.WatchLogs(opts, "ManagerChanged", accountRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1820ManagerChanged)
				if err := _ERC1820.contract.UnpackLog(event, "ManagerChanged", log); err != nil {
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

// ParseManagerChanged is a log parse operation binding the contract event 0x605c2dbf762e5f7d60a546d42e7205dcb1b011ebc62a61736a57c9089d3a4350.
//
// Solidity: event ManagerChanged(address indexed account, address indexed newManager)
func (_ERC1820 *ERC1820Filterer) ParseManagerChanged(log types.Log) (*ERC1820ManagerChanged, error) {
	event := new(ERC1820ManagerChanged)
	if err := _ERC1820.contract.UnpackLog(event, "ManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
