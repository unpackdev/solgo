// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapuniversalrouterv2

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

// RouterParameters is an auto generated low-level Go binding around an user-defined struct.
type RouterParameters struct {
	Permit2                     common.Address
	Weth9                       common.Address
	SeaportV15                  common.Address
	SeaportV14                  common.Address
	OpenseaConduit              common.Address
	NftxZap                     common.Address
	X2y2                        common.Address
	Foundation                  common.Address
	Sudoswap                    common.Address
	ElementMarket               common.Address
	Nft20Zap                    common.Address
	Cryptopunks                 common.Address
	LooksRareV2                 common.Address
	RouterRewardsDistributor    common.Address
	LooksRareRewardsDistributor common.Address
	LooksRareToken              common.Address
	V2Factory                   common.Address
	V3Factory                   common.Address
	PairInitCodeHash            [32]byte
	PoolInitCodeHash            [32]byte
}

// UniswapUniversalRouterV2MetaData contains all meta data concerning the UniswapUniversalRouterV2 contract.
var UniswapUniversalRouterV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"permit2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth9\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seaportV1_5\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seaportV1_4\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"openseaConduit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftxZap\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"x2y2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"foundation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sudoswap\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"elementMarket\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft20Zap\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cryptopunks\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"looksRareV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"routerRewardsDistributor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"looksRareRewardsDistributor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"looksRareToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v2Factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v3Factory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pairInitCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"poolInitCodeHash\",\"type\":\"bytes32\"}],\"internalType\":\"structRouterParameters\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyPunkFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractLocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHNotAccepted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commandIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"ExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FromAddressIsNotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBips\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commandType\",\"type\":\"uint256\"}],\"name\":\"InvalidCommandType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwnerERC1155\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwnerERC721\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReserves\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SliceOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionDeadlinePassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnableToClaim\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsafeCast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2TooLittleReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2TooMuchRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidAmountOut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidSwap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3TooLittleReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3TooMuchRequested\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsSent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"looksRareClaim\",\"type\":\"bytes\"}],\"name\":\"collectRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UniswapUniversalRouterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapUniversalRouterV2MetaData.ABI instead.
var UniswapUniversalRouterV2ABI = UniswapUniversalRouterV2MetaData.ABI

// UniswapUniversalRouterV2 is an auto generated Go binding around an Ethereum contract.
type UniswapUniversalRouterV2 struct {
	UniswapUniversalRouterV2Caller     // Read-only binding to the contract
	UniswapUniversalRouterV2Transactor // Write-only binding to the contract
	UniswapUniversalRouterV2Filterer   // Log filterer for contract events
}

// UniswapUniversalRouterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapUniversalRouterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapUniversalRouterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapUniversalRouterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapUniversalRouterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapUniversalRouterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapUniversalRouterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapUniversalRouterV2Session struct {
	Contract     *UniswapUniversalRouterV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UniswapUniversalRouterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapUniversalRouterV2CallerSession struct {
	Contract *UniswapUniversalRouterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// UniswapUniversalRouterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapUniversalRouterV2TransactorSession struct {
	Contract     *UniswapUniversalRouterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// UniswapUniversalRouterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapUniversalRouterV2Raw struct {
	Contract *UniswapUniversalRouterV2 // Generic contract binding to access the raw methods on
}

// UniswapUniversalRouterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapUniversalRouterV2CallerRaw struct {
	Contract *UniswapUniversalRouterV2Caller // Generic read-only contract binding to access the raw methods on
}

// UniswapUniversalRouterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapUniversalRouterV2TransactorRaw struct {
	Contract *UniswapUniversalRouterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapUniversalRouterV2 creates a new instance of UniswapUniversalRouterV2, bound to a specific deployed contract.
func NewUniswapUniversalRouterV2(address common.Address, backend bind.ContractBackend) (*UniswapUniversalRouterV2, error) {
	contract, err := bindUniswapUniversalRouterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapUniversalRouterV2{UniswapUniversalRouterV2Caller: UniswapUniversalRouterV2Caller{contract: contract}, UniswapUniversalRouterV2Transactor: UniswapUniversalRouterV2Transactor{contract: contract}, UniswapUniversalRouterV2Filterer: UniswapUniversalRouterV2Filterer{contract: contract}}, nil
}

// NewUniswapUniversalRouterV2Caller creates a new read-only instance of UniswapUniversalRouterV2, bound to a specific deployed contract.
func NewUniswapUniversalRouterV2Caller(address common.Address, caller bind.ContractCaller) (*UniswapUniversalRouterV2Caller, error) {
	contract, err := bindUniswapUniversalRouterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapUniversalRouterV2Caller{contract: contract}, nil
}

// NewUniswapUniversalRouterV2Transactor creates a new write-only instance of UniswapUniversalRouterV2, bound to a specific deployed contract.
func NewUniswapUniversalRouterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*UniswapUniversalRouterV2Transactor, error) {
	contract, err := bindUniswapUniversalRouterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapUniversalRouterV2Transactor{contract: contract}, nil
}

// NewUniswapUniversalRouterV2Filterer creates a new log filterer instance of UniswapUniversalRouterV2, bound to a specific deployed contract.
func NewUniswapUniversalRouterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*UniswapUniversalRouterV2Filterer, error) {
	contract, err := bindUniswapUniversalRouterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapUniversalRouterV2Filterer{contract: contract}, nil
}

// bindUniswapUniversalRouterV2 binds a generic wrapper to an already deployed contract.
func bindUniswapUniversalRouterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapUniversalRouterV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapUniversalRouterV2.Contract.UniswapUniversalRouterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.Contract.UniswapUniversalRouterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.Contract.UniswapUniversalRouterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapUniversalRouterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.Contract.contract.Transact(opts, method, params...)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Caller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _UniswapUniversalRouterV2.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Session) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _UniswapUniversalRouterV2.Contract.OnERC1155BatchReceived(&_UniswapUniversalRouterV2.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2CallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _UniswapUniversalRouterV2.Contract.OnERC1155BatchReceived(&_UniswapUniversalRouterV2.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Caller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _UniswapUniversalRouterV2.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Session) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _UniswapUniversalRouterV2.Contract.OnERC1155Received(&_UniswapUniversalRouterV2.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2CallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _UniswapUniversalRouterV2.Contract.OnERC1155Received(&_UniswapUniversalRouterV2.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Caller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _UniswapUniversalRouterV2.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Session) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _UniswapUniversalRouterV2.Contract.OnERC721Received(&_UniswapUniversalRouterV2.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2CallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _UniswapUniversalRouterV2.Contract.OnERC721Received(&_UniswapUniversalRouterV2.CallOpts, arg0, arg1, arg2, arg3)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _UniswapUniversalRouterV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UniswapUniversalRouterV2.Contract.SupportsInterface(&_UniswapUniversalRouterV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UniswapUniversalRouterV2.Contract.SupportsInterface(&_UniswapUniversalRouterV2.CallOpts, interfaceId)
}

// CollectRewards is a paid mutator transaction binding the contract method 0x709a1cc2.
//
// Solidity: function collectRewards(bytes looksRareClaim) returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Transactor) CollectRewards(opts *bind.TransactOpts, looksRareClaim []byte) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.contract.Transact(opts, "collectRewards", looksRareClaim)
}

// CollectRewards is a paid mutator transaction binding the contract method 0x709a1cc2.
//
// Solidity: function collectRewards(bytes looksRareClaim) returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Session) CollectRewards(looksRareClaim []byte) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.Contract.CollectRewards(&_UniswapUniversalRouterV2.TransactOpts, looksRareClaim)
}

// CollectRewards is a paid mutator transaction binding the contract method 0x709a1cc2.
//
// Solidity: function collectRewards(bytes looksRareClaim) returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2TransactorSession) CollectRewards(looksRareClaim []byte) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.Contract.CollectRewards(&_UniswapUniversalRouterV2.TransactOpts, looksRareClaim)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Transactor) Execute(opts *bind.TransactOpts, commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.contract.Transact(opts, "execute", commands, inputs)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Session) Execute(commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.Contract.Execute(&_UniswapUniversalRouterV2.TransactOpts, commands, inputs)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2TransactorSession) Execute(commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.Contract.Execute(&_UniswapUniversalRouterV2.TransactOpts, commands, inputs)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Transactor) Execute0(opts *bind.TransactOpts, commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.contract.Transact(opts, "execute0", commands, inputs, deadline)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Session) Execute0(commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.Contract.Execute0(&_UniswapUniversalRouterV2.TransactOpts, commands, inputs, deadline)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2TransactorSession) Execute0(commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.Contract.Execute0(&_UniswapUniversalRouterV2.TransactOpts, commands, inputs, deadline)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Transactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Session) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.Contract.UniswapV3SwapCallback(&_UniswapUniversalRouterV2.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2TransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.Contract.UniswapV3SwapCallback(&_UniswapUniversalRouterV2.TransactOpts, amount0Delta, amount1Delta, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Session) Receive() (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.Contract.Receive(&_UniswapUniversalRouterV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2TransactorSession) Receive() (*types.Transaction, error) {
	return _UniswapUniversalRouterV2.Contract.Receive(&_UniswapUniversalRouterV2.TransactOpts)
}

// UniswapUniversalRouterV2RewardsSentIterator is returned from FilterRewardsSent and is used to iterate over the raw logs and unpacked data for RewardsSent events raised by the UniswapUniversalRouterV2 contract.
type UniswapUniversalRouterV2RewardsSentIterator struct {
	Event *UniswapUniversalRouterV2RewardsSent // Event containing the contract specifics and raw log

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
func (it *UniswapUniversalRouterV2RewardsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapUniversalRouterV2RewardsSent)
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
		it.Event = new(UniswapUniversalRouterV2RewardsSent)
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
func (it *UniswapUniversalRouterV2RewardsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapUniversalRouterV2RewardsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapUniversalRouterV2RewardsSent represents a RewardsSent event raised by the UniswapUniversalRouterV2 contract.
type UniswapUniversalRouterV2RewardsSent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsSent is a free log retrieval operation binding the contract event 0x1e8f03f716bc104bf7d728131967a0c771e85ab54d09c1e2d6ed9e0bc4e2a16c.
//
// Solidity: event RewardsSent(uint256 amount)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Filterer) FilterRewardsSent(opts *bind.FilterOpts) (*UniswapUniversalRouterV2RewardsSentIterator, error) {

	logs, sub, err := _UniswapUniversalRouterV2.contract.FilterLogs(opts, "RewardsSent")
	if err != nil {
		return nil, err
	}
	return &UniswapUniversalRouterV2RewardsSentIterator{contract: _UniswapUniversalRouterV2.contract, event: "RewardsSent", logs: logs, sub: sub}, nil
}

// WatchRewardsSent is a free log subscription operation binding the contract event 0x1e8f03f716bc104bf7d728131967a0c771e85ab54d09c1e2d6ed9e0bc4e2a16c.
//
// Solidity: event RewardsSent(uint256 amount)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Filterer) WatchRewardsSent(opts *bind.WatchOpts, sink chan<- *UniswapUniversalRouterV2RewardsSent) (event.Subscription, error) {

	logs, sub, err := _UniswapUniversalRouterV2.contract.WatchLogs(opts, "RewardsSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapUniversalRouterV2RewardsSent)
				if err := _UniswapUniversalRouterV2.contract.UnpackLog(event, "RewardsSent", log); err != nil {
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

// ParseRewardsSent is a log parse operation binding the contract event 0x1e8f03f716bc104bf7d728131967a0c771e85ab54d09c1e2d6ed9e0bc4e2a16c.
//
// Solidity: event RewardsSent(uint256 amount)
func (_UniswapUniversalRouterV2 *UniswapUniversalRouterV2Filterer) ParseRewardsSent(log types.Log) (*UniswapUniversalRouterV2RewardsSent, error) {
	event := new(UniswapUniversalRouterV2RewardsSent)
	if err := _UniswapUniversalRouterV2.contract.UnpackLog(event, "RewardsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
