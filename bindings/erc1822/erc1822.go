// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc1822

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

// ERC1822MetaData contains all meta data concerning the ERC1822 contract.
var ERC1822MetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"setProxyOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"\",\"type\":\"address\"}],\"name\":\"ProxyOwnershipTransferred\",\"type\":\"event\"}]",
}

// ERC1822ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1822MetaData.ABI instead.
var ERC1822ABI = ERC1822MetaData.ABI

// ERC1822 is an auto generated Go binding around an Ethereum contract.
type ERC1822 struct {
	ERC1822Caller     // Read-only binding to the contract
	ERC1822Transactor // Write-only binding to the contract
	ERC1822Filterer   // Log filterer for contract events
}

// ERC1822Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1822Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1822Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1822Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1822Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1822Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1822Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1822Session struct {
	Contract     *ERC1822          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1822CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1822CallerSession struct {
	Contract *ERC1822Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC1822TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1822TransactorSession struct {
	Contract     *ERC1822Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC1822Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1822Raw struct {
	Contract *ERC1822 // Generic contract binding to access the raw methods on
}

// ERC1822CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1822CallerRaw struct {
	Contract *ERC1822Caller // Generic read-only contract binding to access the raw methods on
}

// ERC1822TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1822TransactorRaw struct {
	Contract *ERC1822Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1822 creates a new instance of ERC1822, bound to a specific deployed contract.
func NewERC1822(address common.Address, backend bind.ContractBackend) (*ERC1822, error) {
	contract, err := bindERC1822(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1822{ERC1822Caller: ERC1822Caller{contract: contract}, ERC1822Transactor: ERC1822Transactor{contract: contract}, ERC1822Filterer: ERC1822Filterer{contract: contract}}, nil
}

// NewERC1822Caller creates a new read-only instance of ERC1822, bound to a specific deployed contract.
func NewERC1822Caller(address common.Address, caller bind.ContractCaller) (*ERC1822Caller, error) {
	contract, err := bindERC1822(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1822Caller{contract: contract}, nil
}

// NewERC1822Transactor creates a new write-only instance of ERC1822, bound to a specific deployed contract.
func NewERC1822Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC1822Transactor, error) {
	contract, err := bindERC1822(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1822Transactor{contract: contract}, nil
}

// NewERC1822Filterer creates a new log filterer instance of ERC1822, bound to a specific deployed contract.
func NewERC1822Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC1822Filterer, error) {
	contract, err := bindERC1822(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1822Filterer{contract: contract}, nil
}

// bindERC1822 binds a generic wrapper to an already deployed contract.
func bindERC1822(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1822ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1822 *ERC1822Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1822.Contract.ERC1822Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1822 *ERC1822Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1822.Contract.ERC1822Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1822 *ERC1822Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1822.Contract.ERC1822Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1822 *ERC1822CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1822.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1822 *ERC1822TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1822.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1822 *ERC1822TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1822.Contract.contract.Transact(opts, method, params...)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_ERC1822 *ERC1822Caller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC1822.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_ERC1822 *ERC1822Session) GetImplementation() (common.Address, error) {
	return _ERC1822.Contract.GetImplementation(&_ERC1822.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_ERC1822 *ERC1822CallerSession) GetImplementation() (common.Address, error) {
	return _ERC1822.Contract.GetImplementation(&_ERC1822.CallOpts)
}

// SetProxyOwner is a paid mutator transaction binding the contract method 0xcaaee91c.
//
// Solidity: function setProxyOwner(address ) returns()
func (_ERC1822 *ERC1822Transactor) SetProxyOwner(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _ERC1822.contract.Transact(opts, "setProxyOwner", arg0)
}

// SetProxyOwner is a paid mutator transaction binding the contract method 0xcaaee91c.
//
// Solidity: function setProxyOwner(address ) returns()
func (_ERC1822 *ERC1822Session) SetProxyOwner(arg0 common.Address) (*types.Transaction, error) {
	return _ERC1822.Contract.SetProxyOwner(&_ERC1822.TransactOpts, arg0)
}

// SetProxyOwner is a paid mutator transaction binding the contract method 0xcaaee91c.
//
// Solidity: function setProxyOwner(address ) returns()
func (_ERC1822 *ERC1822TransactorSession) SetProxyOwner(arg0 common.Address) (*types.Transaction, error) {
	return _ERC1822.Contract.SetProxyOwner(&_ERC1822.TransactOpts, arg0)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address ) returns()
func (_ERC1822 *ERC1822Transactor) UpgradeTo(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _ERC1822.contract.Transact(opts, "upgradeTo", arg0)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address ) returns()
func (_ERC1822 *ERC1822Session) UpgradeTo(arg0 common.Address) (*types.Transaction, error) {
	return _ERC1822.Contract.UpgradeTo(&_ERC1822.TransactOpts, arg0)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address ) returns()
func (_ERC1822 *ERC1822TransactorSession) UpgradeTo(arg0 common.Address) (*types.Transaction, error) {
	return _ERC1822.Contract.UpgradeTo(&_ERC1822.TransactOpts, arg0)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x57975c29.
//
// Solidity: function upgradeToAndCall(address , string ) returns()
func (_ERC1822 *ERC1822Transactor) UpgradeToAndCall(opts *bind.TransactOpts, arg0 common.Address, arg1 string) (*types.Transaction, error) {
	return _ERC1822.contract.Transact(opts, "upgradeToAndCall", arg0, arg1)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x57975c29.
//
// Solidity: function upgradeToAndCall(address , string ) returns()
func (_ERC1822 *ERC1822Session) UpgradeToAndCall(arg0 common.Address, arg1 string) (*types.Transaction, error) {
	return _ERC1822.Contract.UpgradeToAndCall(&_ERC1822.TransactOpts, arg0, arg1)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x57975c29.
//
// Solidity: function upgradeToAndCall(address , string ) returns()
func (_ERC1822 *ERC1822TransactorSession) UpgradeToAndCall(arg0 common.Address, arg1 string) (*types.Transaction, error) {
	return _ERC1822.Contract.UpgradeToAndCall(&_ERC1822.TransactOpts, arg0, arg1)
}

// ERC1822ProxyOwnershipTransferredIterator is returned from FilterProxyOwnershipTransferred and is used to iterate over the raw logs and unpacked data for ProxyOwnershipTransferred events raised by the ERC1822 contract.
type ERC1822ProxyOwnershipTransferredIterator struct {
	Event *ERC1822ProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC1822ProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1822ProxyOwnershipTransferred)
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
		it.Event = new(ERC1822ProxyOwnershipTransferred)
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
func (it *ERC1822ProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1822ProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1822ProxyOwnershipTransferred represents a ProxyOwnershipTransferred event raised by the ERC1822 contract.
type ERC1822ProxyOwnershipTransferred struct {
	Arg0 common.Address
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterProxyOwnershipTransferred is a free log retrieval operation binding the contract event 0x5a3e66efaa1e445ebd894728a69d6959842ea1e97bd79b892797106e270efcd9.
//
// Solidity: event ProxyOwnershipTransferred(address indexed arg0, address indexed arg1)
func (_ERC1822 *ERC1822Filterer) FilterProxyOwnershipTransferred(opts *bind.FilterOpts, arg0 []common.Address, arg1 []common.Address) (*ERC1822ProxyOwnershipTransferredIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _ERC1822.contract.FilterLogs(opts, "ProxyOwnershipTransferred", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return &ERC1822ProxyOwnershipTransferredIterator{contract: _ERC1822.contract, event: "ProxyOwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchProxyOwnershipTransferred is a free log subscription operation binding the contract event 0x5a3e66efaa1e445ebd894728a69d6959842ea1e97bd79b892797106e270efcd9.
//
// Solidity: event ProxyOwnershipTransferred(address indexed arg0, address indexed arg1)
func (_ERC1822 *ERC1822Filterer) WatchProxyOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC1822ProxyOwnershipTransferred, arg0 []common.Address, arg1 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _ERC1822.contract.WatchLogs(opts, "ProxyOwnershipTransferred", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1822ProxyOwnershipTransferred)
				if err := _ERC1822.contract.UnpackLog(event, "ProxyOwnershipTransferred", log); err != nil {
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

// ParseProxyOwnershipTransferred is a log parse operation binding the contract event 0x5a3e66efaa1e445ebd894728a69d6959842ea1e97bd79b892797106e270efcd9.
//
// Solidity: event ProxyOwnershipTransferred(address indexed arg0, address indexed arg1)
func (_ERC1822 *ERC1822Filterer) ParseProxyOwnershipTransferred(log types.Log) (*ERC1822ProxyOwnershipTransferred, error) {
	event := new(ERC1822ProxyOwnershipTransferred)
	if err := _ERC1822.contract.UnpackLog(event, "ProxyOwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1822UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC1822 contract.
type ERC1822UpgradedIterator struct {
	Event *ERC1822Upgraded // Event containing the contract specifics and raw log

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
func (it *ERC1822UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1822Upgraded)
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
		it.Event = new(ERC1822Upgraded)
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
func (it *ERC1822UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1822UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1822Upgraded represents a Upgraded event raised by the ERC1822 contract.
type ERC1822Upgraded struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed arg0)
func (_ERC1822 *ERC1822Filterer) FilterUpgraded(opts *bind.FilterOpts, arg0 []common.Address) (*ERC1822UpgradedIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _ERC1822.contract.FilterLogs(opts, "Upgraded", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &ERC1822UpgradedIterator{contract: _ERC1822.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed arg0)
func (_ERC1822 *ERC1822Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1822Upgraded, arg0 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _ERC1822.contract.WatchLogs(opts, "Upgraded", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1822Upgraded)
				if err := _ERC1822.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed arg0)
func (_ERC1822 *ERC1822Filterer) ParseUpgraded(log types.Log) (*ERC1822Upgraded, error) {
	event := new(ERC1822Upgraded)
	if err := _ERC1822.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
