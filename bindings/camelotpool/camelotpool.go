// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package camelotpool

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

// CamelotPoolMetaData contains all meta data concerning the CamelotPool contract.
var CamelotPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"DrainWrongToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"token0FeePercent\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"token1FeePercent\",\"type\":\"uint16\"}],\"name\":\"FeePercentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SetPairTypeImmutable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prevStableSwap\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"stableSwap\",\"type\":\"bool\"}],\"name\":\"SetStableSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Skim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_FEE_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"drainWrongToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint16\",\"name\":\"_token0FeePercent\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_token1FeePercent\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pairTypeImmutable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"precisionMultiplier0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"precisionMultiplier1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newToken0FeePercent\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newToken1FeePercent\",\"type\":\"uint16\"}],\"name\":\"setFeePercent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setPairTypeImmutable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"expectedReserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"expectedReserve1\",\"type\":\"uint112\"}],\"name\":\"setStableSwap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stableSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0FeePercent\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token1FeePercent\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CamelotPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use CamelotPoolMetaData.ABI instead.
var CamelotPoolABI = CamelotPoolMetaData.ABI

// CamelotPool is an auto generated Go binding around an Ethereum contract.
type CamelotPool struct {
	CamelotPoolCaller     // Read-only binding to the contract
	CamelotPoolTransactor // Write-only binding to the contract
	CamelotPoolFilterer   // Log filterer for contract events
}

// CamelotPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type CamelotPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CamelotPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CamelotPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CamelotPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CamelotPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CamelotPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CamelotPoolSession struct {
	Contract     *CamelotPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CamelotPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CamelotPoolCallerSession struct {
	Contract *CamelotPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CamelotPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CamelotPoolTransactorSession struct {
	Contract     *CamelotPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CamelotPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type CamelotPoolRaw struct {
	Contract *CamelotPool // Generic contract binding to access the raw methods on
}

// CamelotPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CamelotPoolCallerRaw struct {
	Contract *CamelotPoolCaller // Generic read-only contract binding to access the raw methods on
}

// CamelotPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CamelotPoolTransactorRaw struct {
	Contract *CamelotPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCamelotPool creates a new instance of CamelotPool, bound to a specific deployed contract.
func NewCamelotPool(address common.Address, backend bind.ContractBackend) (*CamelotPool, error) {
	contract, err := bindCamelotPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CamelotPool{CamelotPoolCaller: CamelotPoolCaller{contract: contract}, CamelotPoolTransactor: CamelotPoolTransactor{contract: contract}, CamelotPoolFilterer: CamelotPoolFilterer{contract: contract}}, nil
}

// NewCamelotPoolCaller creates a new read-only instance of CamelotPool, bound to a specific deployed contract.
func NewCamelotPoolCaller(address common.Address, caller bind.ContractCaller) (*CamelotPoolCaller, error) {
	contract, err := bindCamelotPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CamelotPoolCaller{contract: contract}, nil
}

// NewCamelotPoolTransactor creates a new write-only instance of CamelotPool, bound to a specific deployed contract.
func NewCamelotPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*CamelotPoolTransactor, error) {
	contract, err := bindCamelotPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CamelotPoolTransactor{contract: contract}, nil
}

// NewCamelotPoolFilterer creates a new log filterer instance of CamelotPool, bound to a specific deployed contract.
func NewCamelotPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*CamelotPoolFilterer, error) {
	contract, err := bindCamelotPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CamelotPoolFilterer{contract: contract}, nil
}

// bindCamelotPool binds a generic wrapper to an already deployed contract.
func bindCamelotPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CamelotPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CamelotPool *CamelotPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CamelotPool.Contract.CamelotPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CamelotPool *CamelotPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CamelotPool.Contract.CamelotPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CamelotPool *CamelotPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CamelotPool.Contract.CamelotPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CamelotPool *CamelotPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CamelotPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CamelotPool *CamelotPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CamelotPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CamelotPool *CamelotPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CamelotPool.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CamelotPool *CamelotPoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CamelotPool *CamelotPoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CamelotPool.Contract.DOMAINSEPARATOR(&_CamelotPool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CamelotPool *CamelotPoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CamelotPool.Contract.DOMAINSEPARATOR(&_CamelotPool.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_CamelotPool *CamelotPoolCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "FEE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_CamelotPool *CamelotPoolSession) FEEDENOMINATOR() (*big.Int, error) {
	return _CamelotPool.Contract.FEEDENOMINATOR(&_CamelotPool.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_CamelotPool *CamelotPoolCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _CamelotPool.Contract.FEEDENOMINATOR(&_CamelotPool.CallOpts)
}

// MAXFEEPERCENT is a free data retrieval call binding the contract method 0x67d81740.
//
// Solidity: function MAX_FEE_PERCENT() view returns(uint256)
func (_CamelotPool *CamelotPoolCaller) MAXFEEPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "MAX_FEE_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEEPERCENT is a free data retrieval call binding the contract method 0x67d81740.
//
// Solidity: function MAX_FEE_PERCENT() view returns(uint256)
func (_CamelotPool *CamelotPoolSession) MAXFEEPERCENT() (*big.Int, error) {
	return _CamelotPool.Contract.MAXFEEPERCENT(&_CamelotPool.CallOpts)
}

// MAXFEEPERCENT is a free data retrieval call binding the contract method 0x67d81740.
//
// Solidity: function MAX_FEE_PERCENT() view returns(uint256)
func (_CamelotPool *CamelotPoolCallerSession) MAXFEEPERCENT() (*big.Int, error) {
	return _CamelotPool.Contract.MAXFEEPERCENT(&_CamelotPool.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_CamelotPool *CamelotPoolCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_CamelotPool *CamelotPoolSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _CamelotPool.Contract.MINIMUMLIQUIDITY(&_CamelotPool.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_CamelotPool *CamelotPoolCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _CamelotPool.Contract.MINIMUMLIQUIDITY(&_CamelotPool.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_CamelotPool *CamelotPoolCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_CamelotPool *CamelotPoolSession) PERMITTYPEHASH() ([32]byte, error) {
	return _CamelotPool.Contract.PERMITTYPEHASH(&_CamelotPool.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_CamelotPool *CamelotPoolCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _CamelotPool.Contract.PERMITTYPEHASH(&_CamelotPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_CamelotPool *CamelotPoolCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_CamelotPool *CamelotPoolSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CamelotPool.Contract.Allowance(&_CamelotPool.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_CamelotPool *CamelotPoolCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CamelotPool.Contract.Allowance(&_CamelotPool.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CamelotPool *CamelotPoolCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CamelotPool *CamelotPoolSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CamelotPool.Contract.BalanceOf(&_CamelotPool.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CamelotPool *CamelotPoolCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CamelotPool.Contract.BalanceOf(&_CamelotPool.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CamelotPool *CamelotPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CamelotPool *CamelotPoolSession) Decimals() (uint8, error) {
	return _CamelotPool.Contract.Decimals(&_CamelotPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CamelotPool *CamelotPoolCallerSession) Decimals() (uint8, error) {
	return _CamelotPool.Contract.Decimals(&_CamelotPool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CamelotPool *CamelotPoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CamelotPool *CamelotPoolSession) Factory() (common.Address, error) {
	return _CamelotPool.Contract.Factory(&_CamelotPool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CamelotPool *CamelotPoolCallerSession) Factory() (common.Address, error) {
	return _CamelotPool.Contract.Factory(&_CamelotPool.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn) view returns(uint256)
func (_CamelotPool *CamelotPoolCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, tokenIn common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "getAmountOut", amountIn, tokenIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn) view returns(uint256)
func (_CamelotPool *CamelotPoolSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address) (*big.Int, error) {
	return _CamelotPool.Contract.GetAmountOut(&_CamelotPool.CallOpts, amountIn, tokenIn)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn) view returns(uint256)
func (_CamelotPool *CamelotPoolCallerSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address) (*big.Int, error) {
	return _CamelotPool.Contract.GetAmountOut(&_CamelotPool.CallOpts, amountIn, tokenIn)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint16 _token0FeePercent, uint16 _token1FeePercent)
func (_CamelotPool *CamelotPoolCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0         *big.Int
	Reserve1         *big.Int
	Token0FeePercent uint16
	Token1FeePercent uint16
}, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0         *big.Int
		Reserve1         *big.Int
		Token0FeePercent uint16
		Token1FeePercent uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Token0FeePercent = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.Token1FeePercent = *abi.ConvertType(out[3], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint16 _token0FeePercent, uint16 _token1FeePercent)
func (_CamelotPool *CamelotPoolSession) GetReserves() (struct {
	Reserve0         *big.Int
	Reserve1         *big.Int
	Token0FeePercent uint16
	Token1FeePercent uint16
}, error) {
	return _CamelotPool.Contract.GetReserves(&_CamelotPool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint16 _token0FeePercent, uint16 _token1FeePercent)
func (_CamelotPool *CamelotPoolCallerSession) GetReserves() (struct {
	Reserve0         *big.Int
	Reserve1         *big.Int
	Token0FeePercent uint16
	Token1FeePercent uint16
}, error) {
	return _CamelotPool.Contract.GetReserves(&_CamelotPool.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_CamelotPool *CamelotPoolCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_CamelotPool *CamelotPoolSession) Initialized() (bool, error) {
	return _CamelotPool.Contract.Initialized(&_CamelotPool.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_CamelotPool *CamelotPoolCallerSession) Initialized() (bool, error) {
	return _CamelotPool.Contract.Initialized(&_CamelotPool.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_CamelotPool *CamelotPoolCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_CamelotPool *CamelotPoolSession) KLast() (*big.Int, error) {
	return _CamelotPool.Contract.KLast(&_CamelotPool.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_CamelotPool *CamelotPoolCallerSession) KLast() (*big.Int, error) {
	return _CamelotPool.Contract.KLast(&_CamelotPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CamelotPool *CamelotPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CamelotPool *CamelotPoolSession) Name() (string, error) {
	return _CamelotPool.Contract.Name(&_CamelotPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CamelotPool *CamelotPoolCallerSession) Name() (string, error) {
	return _CamelotPool.Contract.Name(&_CamelotPool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_CamelotPool *CamelotPoolCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_CamelotPool *CamelotPoolSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _CamelotPool.Contract.Nonces(&_CamelotPool.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_CamelotPool *CamelotPoolCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _CamelotPool.Contract.Nonces(&_CamelotPool.CallOpts, arg0)
}

// PairTypeImmutable is a free data retrieval call binding the contract method 0xb6200b07.
//
// Solidity: function pairTypeImmutable() view returns(bool)
func (_CamelotPool *CamelotPoolCaller) PairTypeImmutable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "pairTypeImmutable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PairTypeImmutable is a free data retrieval call binding the contract method 0xb6200b07.
//
// Solidity: function pairTypeImmutable() view returns(bool)
func (_CamelotPool *CamelotPoolSession) PairTypeImmutable() (bool, error) {
	return _CamelotPool.Contract.PairTypeImmutable(&_CamelotPool.CallOpts)
}

// PairTypeImmutable is a free data retrieval call binding the contract method 0xb6200b07.
//
// Solidity: function pairTypeImmutable() view returns(bool)
func (_CamelotPool *CamelotPoolCallerSession) PairTypeImmutable() (bool, error) {
	return _CamelotPool.Contract.PairTypeImmutable(&_CamelotPool.CallOpts)
}

// PrecisionMultiplier0 is a free data retrieval call binding the contract method 0x3b9f1dc0.
//
// Solidity: function precisionMultiplier0() view returns(uint256)
func (_CamelotPool *CamelotPoolCaller) PrecisionMultiplier0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "precisionMultiplier0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrecisionMultiplier0 is a free data retrieval call binding the contract method 0x3b9f1dc0.
//
// Solidity: function precisionMultiplier0() view returns(uint256)
func (_CamelotPool *CamelotPoolSession) PrecisionMultiplier0() (*big.Int, error) {
	return _CamelotPool.Contract.PrecisionMultiplier0(&_CamelotPool.CallOpts)
}

// PrecisionMultiplier0 is a free data retrieval call binding the contract method 0x3b9f1dc0.
//
// Solidity: function precisionMultiplier0() view returns(uint256)
func (_CamelotPool *CamelotPoolCallerSession) PrecisionMultiplier0() (*big.Int, error) {
	return _CamelotPool.Contract.PrecisionMultiplier0(&_CamelotPool.CallOpts)
}

// PrecisionMultiplier1 is a free data retrieval call binding the contract method 0x288e5d02.
//
// Solidity: function precisionMultiplier1() view returns(uint256)
func (_CamelotPool *CamelotPoolCaller) PrecisionMultiplier1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "precisionMultiplier1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrecisionMultiplier1 is a free data retrieval call binding the contract method 0x288e5d02.
//
// Solidity: function precisionMultiplier1() view returns(uint256)
func (_CamelotPool *CamelotPoolSession) PrecisionMultiplier1() (*big.Int, error) {
	return _CamelotPool.Contract.PrecisionMultiplier1(&_CamelotPool.CallOpts)
}

// PrecisionMultiplier1 is a free data retrieval call binding the contract method 0x288e5d02.
//
// Solidity: function precisionMultiplier1() view returns(uint256)
func (_CamelotPool *CamelotPoolCallerSession) PrecisionMultiplier1() (*big.Int, error) {
	return _CamelotPool.Contract.PrecisionMultiplier1(&_CamelotPool.CallOpts)
}

// StableSwap is a free data retrieval call binding the contract method 0x9e548b7f.
//
// Solidity: function stableSwap() view returns(bool)
func (_CamelotPool *CamelotPoolCaller) StableSwap(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "stableSwap")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StableSwap is a free data retrieval call binding the contract method 0x9e548b7f.
//
// Solidity: function stableSwap() view returns(bool)
func (_CamelotPool *CamelotPoolSession) StableSwap() (bool, error) {
	return _CamelotPool.Contract.StableSwap(&_CamelotPool.CallOpts)
}

// StableSwap is a free data retrieval call binding the contract method 0x9e548b7f.
//
// Solidity: function stableSwap() view returns(bool)
func (_CamelotPool *CamelotPoolCallerSession) StableSwap() (bool, error) {
	return _CamelotPool.Contract.StableSwap(&_CamelotPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CamelotPool *CamelotPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CamelotPool *CamelotPoolSession) Symbol() (string, error) {
	return _CamelotPool.Contract.Symbol(&_CamelotPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CamelotPool *CamelotPoolCallerSession) Symbol() (string, error) {
	return _CamelotPool.Contract.Symbol(&_CamelotPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CamelotPool *CamelotPoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CamelotPool *CamelotPoolSession) Token0() (common.Address, error) {
	return _CamelotPool.Contract.Token0(&_CamelotPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CamelotPool *CamelotPoolCallerSession) Token0() (common.Address, error) {
	return _CamelotPool.Contract.Token0(&_CamelotPool.CallOpts)
}

// Token0FeePercent is a free data retrieval call binding the contract method 0x62ecec03.
//
// Solidity: function token0FeePercent() view returns(uint16)
func (_CamelotPool *CamelotPoolCaller) Token0FeePercent(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "token0FeePercent")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Token0FeePercent is a free data retrieval call binding the contract method 0x62ecec03.
//
// Solidity: function token0FeePercent() view returns(uint16)
func (_CamelotPool *CamelotPoolSession) Token0FeePercent() (uint16, error) {
	return _CamelotPool.Contract.Token0FeePercent(&_CamelotPool.CallOpts)
}

// Token0FeePercent is a free data retrieval call binding the contract method 0x62ecec03.
//
// Solidity: function token0FeePercent() view returns(uint16)
func (_CamelotPool *CamelotPoolCallerSession) Token0FeePercent() (uint16, error) {
	return _CamelotPool.Contract.Token0FeePercent(&_CamelotPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CamelotPool *CamelotPoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CamelotPool *CamelotPoolSession) Token1() (common.Address, error) {
	return _CamelotPool.Contract.Token1(&_CamelotPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CamelotPool *CamelotPoolCallerSession) Token1() (common.Address, error) {
	return _CamelotPool.Contract.Token1(&_CamelotPool.CallOpts)
}

// Token1FeePercent is a free data retrieval call binding the contract method 0x2fcd1692.
//
// Solidity: function token1FeePercent() view returns(uint16)
func (_CamelotPool *CamelotPoolCaller) Token1FeePercent(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "token1FeePercent")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Token1FeePercent is a free data retrieval call binding the contract method 0x2fcd1692.
//
// Solidity: function token1FeePercent() view returns(uint16)
func (_CamelotPool *CamelotPoolSession) Token1FeePercent() (uint16, error) {
	return _CamelotPool.Contract.Token1FeePercent(&_CamelotPool.CallOpts)
}

// Token1FeePercent is a free data retrieval call binding the contract method 0x2fcd1692.
//
// Solidity: function token1FeePercent() view returns(uint16)
func (_CamelotPool *CamelotPoolCallerSession) Token1FeePercent() (uint16, error) {
	return _CamelotPool.Contract.Token1FeePercent(&_CamelotPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CamelotPool *CamelotPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CamelotPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CamelotPool *CamelotPoolSession) TotalSupply() (*big.Int, error) {
	return _CamelotPool.Contract.TotalSupply(&_CamelotPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CamelotPool *CamelotPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _CamelotPool.Contract.TotalSupply(&_CamelotPool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CamelotPool *CamelotPoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CamelotPool *CamelotPoolSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CamelotPool.Contract.Approve(&_CamelotPool.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CamelotPool *CamelotPoolTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CamelotPool.Contract.Approve(&_CamelotPool.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_CamelotPool *CamelotPoolTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_CamelotPool *CamelotPoolSession) Burn(to common.Address) (*types.Transaction, error) {
	return _CamelotPool.Contract.Burn(&_CamelotPool.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_CamelotPool *CamelotPoolTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _CamelotPool.Contract.Burn(&_CamelotPool.TransactOpts, to)
}

// DrainWrongToken is a paid mutator transaction binding the contract method 0xf39ac11f.
//
// Solidity: function drainWrongToken(address token, address to) returns()
func (_CamelotPool *CamelotPoolTransactor) DrainWrongToken(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "drainWrongToken", token, to)
}

// DrainWrongToken is a paid mutator transaction binding the contract method 0xf39ac11f.
//
// Solidity: function drainWrongToken(address token, address to) returns()
func (_CamelotPool *CamelotPoolSession) DrainWrongToken(token common.Address, to common.Address) (*types.Transaction, error) {
	return _CamelotPool.Contract.DrainWrongToken(&_CamelotPool.TransactOpts, token, to)
}

// DrainWrongToken is a paid mutator transaction binding the contract method 0xf39ac11f.
//
// Solidity: function drainWrongToken(address token, address to) returns()
func (_CamelotPool *CamelotPoolTransactorSession) DrainWrongToken(token common.Address, to common.Address) (*types.Transaction, error) {
	return _CamelotPool.Contract.DrainWrongToken(&_CamelotPool.TransactOpts, token, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_CamelotPool *CamelotPoolTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "initialize", _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_CamelotPool *CamelotPoolSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _CamelotPool.Contract.Initialize(&_CamelotPool.TransactOpts, _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_CamelotPool *CamelotPoolTransactorSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _CamelotPool.Contract.Initialize(&_CamelotPool.TransactOpts, _token0, _token1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_CamelotPool *CamelotPoolTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_CamelotPool *CamelotPoolSession) Mint(to common.Address) (*types.Transaction, error) {
	return _CamelotPool.Contract.Mint(&_CamelotPool.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_CamelotPool *CamelotPoolTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _CamelotPool.Contract.Mint(&_CamelotPool.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CamelotPool *CamelotPoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CamelotPool *CamelotPoolSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CamelotPool.Contract.Permit(&_CamelotPool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CamelotPool *CamelotPoolTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CamelotPool.Contract.Permit(&_CamelotPool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0x48e5d260.
//
// Solidity: function setFeePercent(uint16 newToken0FeePercent, uint16 newToken1FeePercent) returns()
func (_CamelotPool *CamelotPoolTransactor) SetFeePercent(opts *bind.TransactOpts, newToken0FeePercent uint16, newToken1FeePercent uint16) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "setFeePercent", newToken0FeePercent, newToken1FeePercent)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0x48e5d260.
//
// Solidity: function setFeePercent(uint16 newToken0FeePercent, uint16 newToken1FeePercent) returns()
func (_CamelotPool *CamelotPoolSession) SetFeePercent(newToken0FeePercent uint16, newToken1FeePercent uint16) (*types.Transaction, error) {
	return _CamelotPool.Contract.SetFeePercent(&_CamelotPool.TransactOpts, newToken0FeePercent, newToken1FeePercent)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0x48e5d260.
//
// Solidity: function setFeePercent(uint16 newToken0FeePercent, uint16 newToken1FeePercent) returns()
func (_CamelotPool *CamelotPoolTransactorSession) SetFeePercent(newToken0FeePercent uint16, newToken1FeePercent uint16) (*types.Transaction, error) {
	return _CamelotPool.Contract.SetFeePercent(&_CamelotPool.TransactOpts, newToken0FeePercent, newToken1FeePercent)
}

// SetPairTypeImmutable is a paid mutator transaction binding the contract method 0x3ba17077.
//
// Solidity: function setPairTypeImmutable() returns()
func (_CamelotPool *CamelotPoolTransactor) SetPairTypeImmutable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "setPairTypeImmutable")
}

// SetPairTypeImmutable is a paid mutator transaction binding the contract method 0x3ba17077.
//
// Solidity: function setPairTypeImmutable() returns()
func (_CamelotPool *CamelotPoolSession) SetPairTypeImmutable() (*types.Transaction, error) {
	return _CamelotPool.Contract.SetPairTypeImmutable(&_CamelotPool.TransactOpts)
}

// SetPairTypeImmutable is a paid mutator transaction binding the contract method 0x3ba17077.
//
// Solidity: function setPairTypeImmutable() returns()
func (_CamelotPool *CamelotPoolTransactorSession) SetPairTypeImmutable() (*types.Transaction, error) {
	return _CamelotPool.Contract.SetPairTypeImmutable(&_CamelotPool.TransactOpts)
}

// SetStableSwap is a paid mutator transaction binding the contract method 0x3029e5d4.
//
// Solidity: function setStableSwap(bool stable, uint112 expectedReserve0, uint112 expectedReserve1) returns()
func (_CamelotPool *CamelotPoolTransactor) SetStableSwap(opts *bind.TransactOpts, stable bool, expectedReserve0 *big.Int, expectedReserve1 *big.Int) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "setStableSwap", stable, expectedReserve0, expectedReserve1)
}

// SetStableSwap is a paid mutator transaction binding the contract method 0x3029e5d4.
//
// Solidity: function setStableSwap(bool stable, uint112 expectedReserve0, uint112 expectedReserve1) returns()
func (_CamelotPool *CamelotPoolSession) SetStableSwap(stable bool, expectedReserve0 *big.Int, expectedReserve1 *big.Int) (*types.Transaction, error) {
	return _CamelotPool.Contract.SetStableSwap(&_CamelotPool.TransactOpts, stable, expectedReserve0, expectedReserve1)
}

// SetStableSwap is a paid mutator transaction binding the contract method 0x3029e5d4.
//
// Solidity: function setStableSwap(bool stable, uint112 expectedReserve0, uint112 expectedReserve1) returns()
func (_CamelotPool *CamelotPoolTransactorSession) SetStableSwap(stable bool, expectedReserve0 *big.Int, expectedReserve1 *big.Int) (*types.Transaction, error) {
	return _CamelotPool.Contract.SetStableSwap(&_CamelotPool.TransactOpts, stable, expectedReserve0, expectedReserve1)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_CamelotPool *CamelotPoolTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_CamelotPool *CamelotPoolSession) Skim(to common.Address) (*types.Transaction, error) {
	return _CamelotPool.Contract.Skim(&_CamelotPool.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_CamelotPool *CamelotPoolTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _CamelotPool.Contract.Skim(&_CamelotPool.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_CamelotPool *CamelotPoolTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_CamelotPool *CamelotPoolSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _CamelotPool.Contract.Swap(&_CamelotPool.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_CamelotPool *CamelotPoolTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _CamelotPool.Contract.Swap(&_CamelotPool.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap0 is a paid mutator transaction binding the contract method 0x6e1fdd7f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data, address referrer) returns()
func (_CamelotPool *CamelotPoolTransactor) Swap0(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte, referrer common.Address) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "swap0", amount0Out, amount1Out, to, data, referrer)
}

// Swap0 is a paid mutator transaction binding the contract method 0x6e1fdd7f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data, address referrer) returns()
func (_CamelotPool *CamelotPoolSession) Swap0(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte, referrer common.Address) (*types.Transaction, error) {
	return _CamelotPool.Contract.Swap0(&_CamelotPool.TransactOpts, amount0Out, amount1Out, to, data, referrer)
}

// Swap0 is a paid mutator transaction binding the contract method 0x6e1fdd7f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data, address referrer) returns()
func (_CamelotPool *CamelotPoolTransactorSession) Swap0(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte, referrer common.Address) (*types.Transaction, error) {
	return _CamelotPool.Contract.Swap0(&_CamelotPool.TransactOpts, amount0Out, amount1Out, to, data, referrer)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_CamelotPool *CamelotPoolTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_CamelotPool *CamelotPoolSession) Sync() (*types.Transaction, error) {
	return _CamelotPool.Contract.Sync(&_CamelotPool.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_CamelotPool *CamelotPoolTransactorSession) Sync() (*types.Transaction, error) {
	return _CamelotPool.Contract.Sync(&_CamelotPool.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CamelotPool *CamelotPoolTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CamelotPool *CamelotPoolSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CamelotPool.Contract.Transfer(&_CamelotPool.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CamelotPool *CamelotPoolTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CamelotPool.Contract.Transfer(&_CamelotPool.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CamelotPool *CamelotPoolTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CamelotPool.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CamelotPool *CamelotPoolSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CamelotPool.Contract.TransferFrom(&_CamelotPool.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CamelotPool *CamelotPoolTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CamelotPool.Contract.TransferFrom(&_CamelotPool.TransactOpts, from, to, value)
}

// CamelotPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CamelotPool contract.
type CamelotPoolApprovalIterator struct {
	Event *CamelotPoolApproval // Event containing the contract specifics and raw log

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
func (it *CamelotPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CamelotPoolApproval)
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
		it.Event = new(CamelotPoolApproval)
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
func (it *CamelotPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CamelotPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CamelotPoolApproval represents a Approval event raised by the CamelotPool contract.
type CamelotPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CamelotPool *CamelotPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CamelotPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CamelotPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CamelotPoolApprovalIterator{contract: _CamelotPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CamelotPool *CamelotPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CamelotPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CamelotPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CamelotPoolApproval)
				if err := _CamelotPool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CamelotPool *CamelotPoolFilterer) ParseApproval(log types.Log) (*CamelotPoolApproval, error) {
	event := new(CamelotPoolApproval)
	if err := _CamelotPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CamelotPoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the CamelotPool contract.
type CamelotPoolBurnIterator struct {
	Event *CamelotPoolBurn // Event containing the contract specifics and raw log

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
func (it *CamelotPoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CamelotPoolBurn)
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
		it.Event = new(CamelotPoolBurn)
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
func (it *CamelotPoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CamelotPoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CamelotPoolBurn represents a Burn event raised by the CamelotPool contract.
type CamelotPoolBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_CamelotPool *CamelotPoolFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*CamelotPoolBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CamelotPool.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CamelotPoolBurnIterator{contract: _CamelotPool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_CamelotPool *CamelotPoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *CamelotPoolBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CamelotPool.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CamelotPoolBurn)
				if err := _CamelotPool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_CamelotPool *CamelotPoolFilterer) ParseBurn(log types.Log) (*CamelotPoolBurn, error) {
	event := new(CamelotPoolBurn)
	if err := _CamelotPool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CamelotPoolDrainWrongTokenIterator is returned from FilterDrainWrongToken and is used to iterate over the raw logs and unpacked data for DrainWrongToken events raised by the CamelotPool contract.
type CamelotPoolDrainWrongTokenIterator struct {
	Event *CamelotPoolDrainWrongToken // Event containing the contract specifics and raw log

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
func (it *CamelotPoolDrainWrongTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CamelotPoolDrainWrongToken)
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
		it.Event = new(CamelotPoolDrainWrongToken)
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
func (it *CamelotPoolDrainWrongTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CamelotPoolDrainWrongTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CamelotPoolDrainWrongToken represents a DrainWrongToken event raised by the CamelotPool contract.
type CamelotPoolDrainWrongToken struct {
	Token common.Address
	To    common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDrainWrongToken is a free log retrieval operation binding the contract event 0x368a9dc863ecb94b5ba32a682e26295b10d9c2666fad7d785ebdf262c3c52413.
//
// Solidity: event DrainWrongToken(address indexed token, address to)
func (_CamelotPool *CamelotPoolFilterer) FilterDrainWrongToken(opts *bind.FilterOpts, token []common.Address) (*CamelotPoolDrainWrongTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CamelotPool.contract.FilterLogs(opts, "DrainWrongToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CamelotPoolDrainWrongTokenIterator{contract: _CamelotPool.contract, event: "DrainWrongToken", logs: logs, sub: sub}, nil
}

// WatchDrainWrongToken is a free log subscription operation binding the contract event 0x368a9dc863ecb94b5ba32a682e26295b10d9c2666fad7d785ebdf262c3c52413.
//
// Solidity: event DrainWrongToken(address indexed token, address to)
func (_CamelotPool *CamelotPoolFilterer) WatchDrainWrongToken(opts *bind.WatchOpts, sink chan<- *CamelotPoolDrainWrongToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CamelotPool.contract.WatchLogs(opts, "DrainWrongToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CamelotPoolDrainWrongToken)
				if err := _CamelotPool.contract.UnpackLog(event, "DrainWrongToken", log); err != nil {
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

// ParseDrainWrongToken is a log parse operation binding the contract event 0x368a9dc863ecb94b5ba32a682e26295b10d9c2666fad7d785ebdf262c3c52413.
//
// Solidity: event DrainWrongToken(address indexed token, address to)
func (_CamelotPool *CamelotPoolFilterer) ParseDrainWrongToken(log types.Log) (*CamelotPoolDrainWrongToken, error) {
	event := new(CamelotPoolDrainWrongToken)
	if err := _CamelotPool.contract.UnpackLog(event, "DrainWrongToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CamelotPoolFeePercentUpdatedIterator is returned from FilterFeePercentUpdated and is used to iterate over the raw logs and unpacked data for FeePercentUpdated events raised by the CamelotPool contract.
type CamelotPoolFeePercentUpdatedIterator struct {
	Event *CamelotPoolFeePercentUpdated // Event containing the contract specifics and raw log

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
func (it *CamelotPoolFeePercentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CamelotPoolFeePercentUpdated)
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
		it.Event = new(CamelotPoolFeePercentUpdated)
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
func (it *CamelotPoolFeePercentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CamelotPoolFeePercentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CamelotPoolFeePercentUpdated represents a FeePercentUpdated event raised by the CamelotPool contract.
type CamelotPoolFeePercentUpdated struct {
	Token0FeePercent uint16
	Token1FeePercent uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFeePercentUpdated is a free log retrieval operation binding the contract event 0xa4877b8ecb5a00ba277e4bceeeb187a669e7113649774dfbea05c259ce27f17b.
//
// Solidity: event FeePercentUpdated(uint16 token0FeePercent, uint16 token1FeePercent)
func (_CamelotPool *CamelotPoolFilterer) FilterFeePercentUpdated(opts *bind.FilterOpts) (*CamelotPoolFeePercentUpdatedIterator, error) {

	logs, sub, err := _CamelotPool.contract.FilterLogs(opts, "FeePercentUpdated")
	if err != nil {
		return nil, err
	}
	return &CamelotPoolFeePercentUpdatedIterator{contract: _CamelotPool.contract, event: "FeePercentUpdated", logs: logs, sub: sub}, nil
}

// WatchFeePercentUpdated is a free log subscription operation binding the contract event 0xa4877b8ecb5a00ba277e4bceeeb187a669e7113649774dfbea05c259ce27f17b.
//
// Solidity: event FeePercentUpdated(uint16 token0FeePercent, uint16 token1FeePercent)
func (_CamelotPool *CamelotPoolFilterer) WatchFeePercentUpdated(opts *bind.WatchOpts, sink chan<- *CamelotPoolFeePercentUpdated) (event.Subscription, error) {

	logs, sub, err := _CamelotPool.contract.WatchLogs(opts, "FeePercentUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CamelotPoolFeePercentUpdated)
				if err := _CamelotPool.contract.UnpackLog(event, "FeePercentUpdated", log); err != nil {
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

// ParseFeePercentUpdated is a log parse operation binding the contract event 0xa4877b8ecb5a00ba277e4bceeeb187a669e7113649774dfbea05c259ce27f17b.
//
// Solidity: event FeePercentUpdated(uint16 token0FeePercent, uint16 token1FeePercent)
func (_CamelotPool *CamelotPoolFilterer) ParseFeePercentUpdated(log types.Log) (*CamelotPoolFeePercentUpdated, error) {
	event := new(CamelotPoolFeePercentUpdated)
	if err := _CamelotPool.contract.UnpackLog(event, "FeePercentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CamelotPoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the CamelotPool contract.
type CamelotPoolMintIterator struct {
	Event *CamelotPoolMint // Event containing the contract specifics and raw log

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
func (it *CamelotPoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CamelotPoolMint)
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
		it.Event = new(CamelotPoolMint)
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
func (it *CamelotPoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CamelotPoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CamelotPoolMint represents a Mint event raised by the CamelotPool contract.
type CamelotPoolMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_CamelotPool *CamelotPoolFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*CamelotPoolMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CamelotPool.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &CamelotPoolMintIterator{contract: _CamelotPool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_CamelotPool *CamelotPoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *CamelotPoolMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CamelotPool.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CamelotPoolMint)
				if err := _CamelotPool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_CamelotPool *CamelotPoolFilterer) ParseMint(log types.Log) (*CamelotPoolMint, error) {
	event := new(CamelotPoolMint)
	if err := _CamelotPool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CamelotPoolSetPairTypeImmutableIterator is returned from FilterSetPairTypeImmutable and is used to iterate over the raw logs and unpacked data for SetPairTypeImmutable events raised by the CamelotPool contract.
type CamelotPoolSetPairTypeImmutableIterator struct {
	Event *CamelotPoolSetPairTypeImmutable // Event containing the contract specifics and raw log

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
func (it *CamelotPoolSetPairTypeImmutableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CamelotPoolSetPairTypeImmutable)
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
		it.Event = new(CamelotPoolSetPairTypeImmutable)
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
func (it *CamelotPoolSetPairTypeImmutableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CamelotPoolSetPairTypeImmutableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CamelotPoolSetPairTypeImmutable represents a SetPairTypeImmutable event raised by the CamelotPool contract.
type CamelotPoolSetPairTypeImmutable struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetPairTypeImmutable is a free log retrieval operation binding the contract event 0x09122c41ae733a4d7740324d50e35fbd6ee85be3c1312a45596d8045150ab2f2.
//
// Solidity: event SetPairTypeImmutable()
func (_CamelotPool *CamelotPoolFilterer) FilterSetPairTypeImmutable(opts *bind.FilterOpts) (*CamelotPoolSetPairTypeImmutableIterator, error) {

	logs, sub, err := _CamelotPool.contract.FilterLogs(opts, "SetPairTypeImmutable")
	if err != nil {
		return nil, err
	}
	return &CamelotPoolSetPairTypeImmutableIterator{contract: _CamelotPool.contract, event: "SetPairTypeImmutable", logs: logs, sub: sub}, nil
}

// WatchSetPairTypeImmutable is a free log subscription operation binding the contract event 0x09122c41ae733a4d7740324d50e35fbd6ee85be3c1312a45596d8045150ab2f2.
//
// Solidity: event SetPairTypeImmutable()
func (_CamelotPool *CamelotPoolFilterer) WatchSetPairTypeImmutable(opts *bind.WatchOpts, sink chan<- *CamelotPoolSetPairTypeImmutable) (event.Subscription, error) {

	logs, sub, err := _CamelotPool.contract.WatchLogs(opts, "SetPairTypeImmutable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CamelotPoolSetPairTypeImmutable)
				if err := _CamelotPool.contract.UnpackLog(event, "SetPairTypeImmutable", log); err != nil {
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

// ParseSetPairTypeImmutable is a log parse operation binding the contract event 0x09122c41ae733a4d7740324d50e35fbd6ee85be3c1312a45596d8045150ab2f2.
//
// Solidity: event SetPairTypeImmutable()
func (_CamelotPool *CamelotPoolFilterer) ParseSetPairTypeImmutable(log types.Log) (*CamelotPoolSetPairTypeImmutable, error) {
	event := new(CamelotPoolSetPairTypeImmutable)
	if err := _CamelotPool.contract.UnpackLog(event, "SetPairTypeImmutable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CamelotPoolSetStableSwapIterator is returned from FilterSetStableSwap and is used to iterate over the raw logs and unpacked data for SetStableSwap events raised by the CamelotPool contract.
type CamelotPoolSetStableSwapIterator struct {
	Event *CamelotPoolSetStableSwap // Event containing the contract specifics and raw log

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
func (it *CamelotPoolSetStableSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CamelotPoolSetStableSwap)
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
		it.Event = new(CamelotPoolSetStableSwap)
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
func (it *CamelotPoolSetStableSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CamelotPoolSetStableSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CamelotPoolSetStableSwap represents a SetStableSwap event raised by the CamelotPool contract.
type CamelotPoolSetStableSwap struct {
	PrevStableSwap bool
	StableSwap     bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetStableSwap is a free log retrieval operation binding the contract event 0xb6a86710bde53aa7fb1b3856279e2af5b476d53e2dd0902cf17a0911b5a43a8b.
//
// Solidity: event SetStableSwap(bool prevStableSwap, bool stableSwap)
func (_CamelotPool *CamelotPoolFilterer) FilterSetStableSwap(opts *bind.FilterOpts) (*CamelotPoolSetStableSwapIterator, error) {

	logs, sub, err := _CamelotPool.contract.FilterLogs(opts, "SetStableSwap")
	if err != nil {
		return nil, err
	}
	return &CamelotPoolSetStableSwapIterator{contract: _CamelotPool.contract, event: "SetStableSwap", logs: logs, sub: sub}, nil
}

// WatchSetStableSwap is a free log subscription operation binding the contract event 0xb6a86710bde53aa7fb1b3856279e2af5b476d53e2dd0902cf17a0911b5a43a8b.
//
// Solidity: event SetStableSwap(bool prevStableSwap, bool stableSwap)
func (_CamelotPool *CamelotPoolFilterer) WatchSetStableSwap(opts *bind.WatchOpts, sink chan<- *CamelotPoolSetStableSwap) (event.Subscription, error) {

	logs, sub, err := _CamelotPool.contract.WatchLogs(opts, "SetStableSwap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CamelotPoolSetStableSwap)
				if err := _CamelotPool.contract.UnpackLog(event, "SetStableSwap", log); err != nil {
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

// ParseSetStableSwap is a log parse operation binding the contract event 0xb6a86710bde53aa7fb1b3856279e2af5b476d53e2dd0902cf17a0911b5a43a8b.
//
// Solidity: event SetStableSwap(bool prevStableSwap, bool stableSwap)
func (_CamelotPool *CamelotPoolFilterer) ParseSetStableSwap(log types.Log) (*CamelotPoolSetStableSwap, error) {
	event := new(CamelotPoolSetStableSwap)
	if err := _CamelotPool.contract.UnpackLog(event, "SetStableSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CamelotPoolSkimIterator is returned from FilterSkim and is used to iterate over the raw logs and unpacked data for Skim events raised by the CamelotPool contract.
type CamelotPoolSkimIterator struct {
	Event *CamelotPoolSkim // Event containing the contract specifics and raw log

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
func (it *CamelotPoolSkimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CamelotPoolSkim)
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
		it.Event = new(CamelotPoolSkim)
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
func (it *CamelotPoolSkimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CamelotPoolSkimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CamelotPoolSkim represents a Skim event raised by the CamelotPool contract.
type CamelotPoolSkim struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSkim is a free log retrieval operation binding the contract event 0x21ad22495c9c75cd1c94756f91824e779c0c8a8e168b267c790df464fe056b79.
//
// Solidity: event Skim()
func (_CamelotPool *CamelotPoolFilterer) FilterSkim(opts *bind.FilterOpts) (*CamelotPoolSkimIterator, error) {

	logs, sub, err := _CamelotPool.contract.FilterLogs(opts, "Skim")
	if err != nil {
		return nil, err
	}
	return &CamelotPoolSkimIterator{contract: _CamelotPool.contract, event: "Skim", logs: logs, sub: sub}, nil
}

// WatchSkim is a free log subscription operation binding the contract event 0x21ad22495c9c75cd1c94756f91824e779c0c8a8e168b267c790df464fe056b79.
//
// Solidity: event Skim()
func (_CamelotPool *CamelotPoolFilterer) WatchSkim(opts *bind.WatchOpts, sink chan<- *CamelotPoolSkim) (event.Subscription, error) {

	logs, sub, err := _CamelotPool.contract.WatchLogs(opts, "Skim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CamelotPoolSkim)
				if err := _CamelotPool.contract.UnpackLog(event, "Skim", log); err != nil {
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

// ParseSkim is a log parse operation binding the contract event 0x21ad22495c9c75cd1c94756f91824e779c0c8a8e168b267c790df464fe056b79.
//
// Solidity: event Skim()
func (_CamelotPool *CamelotPoolFilterer) ParseSkim(log types.Log) (*CamelotPoolSkim, error) {
	event := new(CamelotPoolSkim)
	if err := _CamelotPool.contract.UnpackLog(event, "Skim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CamelotPoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the CamelotPool contract.
type CamelotPoolSwapIterator struct {
	Event *CamelotPoolSwap // Event containing the contract specifics and raw log

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
func (it *CamelotPoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CamelotPoolSwap)
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
		it.Event = new(CamelotPoolSwap)
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
func (it *CamelotPoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CamelotPoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CamelotPoolSwap represents a Swap event raised by the CamelotPool contract.
type CamelotPoolSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_CamelotPool *CamelotPoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*CamelotPoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CamelotPool.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CamelotPoolSwapIterator{contract: _CamelotPool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_CamelotPool *CamelotPoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *CamelotPoolSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CamelotPool.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CamelotPoolSwap)
				if err := _CamelotPool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_CamelotPool *CamelotPoolFilterer) ParseSwap(log types.Log) (*CamelotPoolSwap, error) {
	event := new(CamelotPoolSwap)
	if err := _CamelotPool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CamelotPoolSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the CamelotPool contract.
type CamelotPoolSyncIterator struct {
	Event *CamelotPoolSync // Event containing the contract specifics and raw log

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
func (it *CamelotPoolSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CamelotPoolSync)
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
		it.Event = new(CamelotPoolSync)
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
func (it *CamelotPoolSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CamelotPoolSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CamelotPoolSync represents a Sync event raised by the CamelotPool contract.
type CamelotPoolSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_CamelotPool *CamelotPoolFilterer) FilterSync(opts *bind.FilterOpts) (*CamelotPoolSyncIterator, error) {

	logs, sub, err := _CamelotPool.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &CamelotPoolSyncIterator{contract: _CamelotPool.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_CamelotPool *CamelotPoolFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *CamelotPoolSync) (event.Subscription, error) {

	logs, sub, err := _CamelotPool.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CamelotPoolSync)
				if err := _CamelotPool.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_CamelotPool *CamelotPoolFilterer) ParseSync(log types.Log) (*CamelotPoolSync, error) {
	event := new(CamelotPoolSync)
	if err := _CamelotPool.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CamelotPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CamelotPool contract.
type CamelotPoolTransferIterator struct {
	Event *CamelotPoolTransfer // Event containing the contract specifics and raw log

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
func (it *CamelotPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CamelotPoolTransfer)
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
		it.Event = new(CamelotPoolTransfer)
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
func (it *CamelotPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CamelotPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CamelotPoolTransfer represents a Transfer event raised by the CamelotPool contract.
type CamelotPoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CamelotPool *CamelotPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CamelotPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CamelotPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CamelotPoolTransferIterator{contract: _CamelotPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CamelotPool *CamelotPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CamelotPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CamelotPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CamelotPoolTransfer)
				if err := _CamelotPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CamelotPool *CamelotPoolFilterer) ParseTransfer(log types.Log) (*CamelotPoolTransfer, error) {
	event := new(CamelotPoolTransfer)
	if err := _CamelotPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
