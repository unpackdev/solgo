// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv3ticklens

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

// ITickLensPopulatedTick is an auto generated low-level Go binding around an user-defined struct.
type ITickLensPopulatedTick struct {
	Tick           *big.Int
	LiquidityNet   *big.Int
	LiquidityGross *big.Int
}

// UniswapV3TicklensMetaData contains all meta data concerning the UniswapV3Ticklens contract.
var UniswapV3TicklensMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"int16\",\"name\":\"tickBitmapIndex\",\"type\":\"int16\"}],\"name\":\"getPopulatedTicksInWord\",\"outputs\":[{\"components\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"int128\",\"name\":\"liquidityNet\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityGross\",\"type\":\"uint128\"}],\"internalType\":\"structITickLens.PopulatedTick[]\",\"name\":\"populatedTicks\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapV3TicklensABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3TicklensMetaData.ABI instead.
var UniswapV3TicklensABI = UniswapV3TicklensMetaData.ABI

// UniswapV3Ticklens is an auto generated Go binding around an Ethereum contract.
type UniswapV3Ticklens struct {
	UniswapV3TicklensCaller     // Read-only binding to the contract
	UniswapV3TicklensTransactor // Write-only binding to the contract
	UniswapV3TicklensFilterer   // Log filterer for contract events
}

// UniswapV3TicklensCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3TicklensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3TicklensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3TicklensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3TicklensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3TicklensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3TicklensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3TicklensSession struct {
	Contract     *UniswapV3Ticklens // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UniswapV3TicklensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3TicklensCallerSession struct {
	Contract *UniswapV3TicklensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UniswapV3TicklensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3TicklensTransactorSession struct {
	Contract     *UniswapV3TicklensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UniswapV3TicklensRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3TicklensRaw struct {
	Contract *UniswapV3Ticklens // Generic contract binding to access the raw methods on
}

// UniswapV3TicklensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3TicklensCallerRaw struct {
	Contract *UniswapV3TicklensCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3TicklensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3TicklensTransactorRaw struct {
	Contract *UniswapV3TicklensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3Ticklens creates a new instance of UniswapV3Ticklens, bound to a specific deployed contract.
func NewUniswapV3Ticklens(address common.Address, backend bind.ContractBackend) (*UniswapV3Ticklens, error) {
	contract, err := bindUniswapV3Ticklens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3Ticklens{UniswapV3TicklensCaller: UniswapV3TicklensCaller{contract: contract}, UniswapV3TicklensTransactor: UniswapV3TicklensTransactor{contract: contract}, UniswapV3TicklensFilterer: UniswapV3TicklensFilterer{contract: contract}}, nil
}

// NewUniswapV3TicklensCaller creates a new read-only instance of UniswapV3Ticklens, bound to a specific deployed contract.
func NewUniswapV3TicklensCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3TicklensCaller, error) {
	contract, err := bindUniswapV3Ticklens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3TicklensCaller{contract: contract}, nil
}

// NewUniswapV3TicklensTransactor creates a new write-only instance of UniswapV3Ticklens, bound to a specific deployed contract.
func NewUniswapV3TicklensTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3TicklensTransactor, error) {
	contract, err := bindUniswapV3Ticklens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3TicklensTransactor{contract: contract}, nil
}

// NewUniswapV3TicklensFilterer creates a new log filterer instance of UniswapV3Ticklens, bound to a specific deployed contract.
func NewUniswapV3TicklensFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3TicklensFilterer, error) {
	contract, err := bindUniswapV3Ticklens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3TicklensFilterer{contract: contract}, nil
}

// bindUniswapV3Ticklens binds a generic wrapper to an already deployed contract.
func bindUniswapV3Ticklens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapV3TicklensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Ticklens *UniswapV3TicklensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Ticklens.Contract.UniswapV3TicklensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Ticklens *UniswapV3TicklensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Ticklens.Contract.UniswapV3TicklensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Ticklens *UniswapV3TicklensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Ticklens.Contract.UniswapV3TicklensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Ticklens *UniswapV3TicklensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Ticklens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Ticklens *UniswapV3TicklensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Ticklens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Ticklens *UniswapV3TicklensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Ticklens.Contract.contract.Transact(opts, method, params...)
}

// GetPopulatedTicksInWord is a free data retrieval call binding the contract method 0x351fb478.
//
// Solidity: function getPopulatedTicksInWord(address pool, int16 tickBitmapIndex) view returns((int24,int128,uint128)[] populatedTicks)
func (_UniswapV3Ticklens *UniswapV3TicklensCaller) GetPopulatedTicksInWord(opts *bind.CallOpts, pool common.Address, tickBitmapIndex int16) ([]ITickLensPopulatedTick, error) {
	var out []interface{}
	err := _UniswapV3Ticklens.contract.Call(opts, &out, "getPopulatedTicksInWord", pool, tickBitmapIndex)

	if err != nil {
		return *new([]ITickLensPopulatedTick), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITickLensPopulatedTick)).(*[]ITickLensPopulatedTick)

	return out0, err

}

// GetPopulatedTicksInWord is a free data retrieval call binding the contract method 0x351fb478.
//
// Solidity: function getPopulatedTicksInWord(address pool, int16 tickBitmapIndex) view returns((int24,int128,uint128)[] populatedTicks)
func (_UniswapV3Ticklens *UniswapV3TicklensSession) GetPopulatedTicksInWord(pool common.Address, tickBitmapIndex int16) ([]ITickLensPopulatedTick, error) {
	return _UniswapV3Ticklens.Contract.GetPopulatedTicksInWord(&_UniswapV3Ticklens.CallOpts, pool, tickBitmapIndex)
}

// GetPopulatedTicksInWord is a free data retrieval call binding the contract method 0x351fb478.
//
// Solidity: function getPopulatedTicksInWord(address pool, int16 tickBitmapIndex) view returns((int24,int128,uint128)[] populatedTicks)
func (_UniswapV3Ticklens *UniswapV3TicklensCallerSession) GetPopulatedTicksInWord(pool common.Address, tickBitmapIndex int16) ([]ITickLensPopulatedTick, error) {
	return _UniswapV3Ticklens.Contract.GetPopulatedTicksInWord(&_UniswapV3Ticklens.CallOpts, pool, tickBitmapIndex)
}
