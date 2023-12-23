// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapuniversalrouter

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
	Seaport                     common.Address
	NftxZap                     common.Address
	X2y2                        common.Address
	Foundation                  common.Address
	Sudoswap                    common.Address
	Nft20Zap                    common.Address
	Cryptopunks                 common.Address
	LooksRare                   common.Address
	RouterRewardsDistributor    common.Address
	LooksRareRewardsDistributor common.Address
	LooksRareToken              common.Address
	V2Factory                   common.Address
	V3Factory                   common.Address
	PairInitCodeHash            [32]byte
	PoolInitCodeHash            [32]byte
}

// UniswapUniversalRouterMetaData contains all meta data concerning the UniswapUniversalRouter contract.
var UniswapUniversalRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"permit2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth9\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seaport\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftxZap\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"x2y2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"foundation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sudoswap\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft20Zap\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cryptopunks\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"looksRare\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"routerRewardsDistributor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"looksRareRewardsDistributor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"looksRareToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v2Factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v3Factory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pairInitCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"poolInitCodeHash\",\"type\":\"bytes32\"}],\"internalType\":\"structRouterParameters\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ContractLocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHNotAccepted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commandIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"ExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FromAddressIsNotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBips\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commandType\",\"type\":\"uint256\"}],\"name\":\"InvalidCommandType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwnerERC1155\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwnerERC721\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReserves\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSlice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SliceOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SliceOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ToAddressOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ToAddressOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ToUint24OutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ToUint24Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionDeadlinePassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnableToClaim\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsafeCast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2TooLittleReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2TooMuchRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidAmountOut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidSwap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3TooLittleReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3TooMuchRequested\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsSent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"looksRareClaim\",\"type\":\"bytes\"}],\"name\":\"collectRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UniswapUniversalRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapUniversalRouterMetaData.ABI instead.
var UniswapUniversalRouterABI = UniswapUniversalRouterMetaData.ABI

// UniswapUniversalRouter is an auto generated Go binding around an Ethereum contract.
type UniswapUniversalRouter struct {
	UniswapUniversalRouterCaller     // Read-only binding to the contract
	UniswapUniversalRouterTransactor // Write-only binding to the contract
	UniswapUniversalRouterFilterer   // Log filterer for contract events
}

// UniswapUniversalRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapUniversalRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapUniversalRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapUniversalRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapUniversalRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapUniversalRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapUniversalRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapUniversalRouterSession struct {
	Contract     *UniswapUniversalRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UniswapUniversalRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapUniversalRouterCallerSession struct {
	Contract *UniswapUniversalRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// UniswapUniversalRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapUniversalRouterTransactorSession struct {
	Contract     *UniswapUniversalRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// UniswapUniversalRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapUniversalRouterRaw struct {
	Contract *UniswapUniversalRouter // Generic contract binding to access the raw methods on
}

// UniswapUniversalRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapUniversalRouterCallerRaw struct {
	Contract *UniswapUniversalRouterCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapUniversalRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapUniversalRouterTransactorRaw struct {
	Contract *UniswapUniversalRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapUniversalRouter creates a new instance of UniswapUniversalRouter, bound to a specific deployed contract.
func NewUniswapUniversalRouter(address common.Address, backend bind.ContractBackend) (*UniswapUniversalRouter, error) {
	contract, err := bindUniswapUniversalRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapUniversalRouter{UniswapUniversalRouterCaller: UniswapUniversalRouterCaller{contract: contract}, UniswapUniversalRouterTransactor: UniswapUniversalRouterTransactor{contract: contract}, UniswapUniversalRouterFilterer: UniswapUniversalRouterFilterer{contract: contract}}, nil
}

// NewUniswapUniversalRouterCaller creates a new read-only instance of UniswapUniversalRouter, bound to a specific deployed contract.
func NewUniswapUniversalRouterCaller(address common.Address, caller bind.ContractCaller) (*UniswapUniversalRouterCaller, error) {
	contract, err := bindUniswapUniversalRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapUniversalRouterCaller{contract: contract}, nil
}

// NewUniswapUniversalRouterTransactor creates a new write-only instance of UniswapUniversalRouter, bound to a specific deployed contract.
func NewUniswapUniversalRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapUniversalRouterTransactor, error) {
	contract, err := bindUniswapUniversalRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapUniversalRouterTransactor{contract: contract}, nil
}

// NewUniswapUniversalRouterFilterer creates a new log filterer instance of UniswapUniversalRouter, bound to a specific deployed contract.
func NewUniswapUniversalRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapUniversalRouterFilterer, error) {
	contract, err := bindUniswapUniversalRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapUniversalRouterFilterer{contract: contract}, nil
}

// bindUniswapUniversalRouter binds a generic wrapper to an already deployed contract.
func bindUniswapUniversalRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapUniversalRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapUniversalRouter *UniswapUniversalRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapUniversalRouter.Contract.UniswapUniversalRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapUniversalRouter *UniswapUniversalRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapUniversalRouter.Contract.UniswapUniversalRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapUniversalRouter *UniswapUniversalRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapUniversalRouter.Contract.UniswapUniversalRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapUniversalRouter *UniswapUniversalRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapUniversalRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapUniversalRouter *UniswapUniversalRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapUniversalRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapUniversalRouter *UniswapUniversalRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapUniversalRouter.Contract.contract.Transact(opts, method, params...)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouter *UniswapUniversalRouterCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _UniswapUniversalRouter.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouter *UniswapUniversalRouterSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _UniswapUniversalRouter.Contract.OnERC1155BatchReceived(&_UniswapUniversalRouter.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouter *UniswapUniversalRouterCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _UniswapUniversalRouter.Contract.OnERC1155BatchReceived(&_UniswapUniversalRouter.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouter *UniswapUniversalRouterCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _UniswapUniversalRouter.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouter *UniswapUniversalRouterSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _UniswapUniversalRouter.Contract.OnERC1155Received(&_UniswapUniversalRouter.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouter *UniswapUniversalRouterCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _UniswapUniversalRouter.Contract.OnERC1155Received(&_UniswapUniversalRouter.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouter *UniswapUniversalRouterCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _UniswapUniversalRouter.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouter *UniswapUniversalRouterSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _UniswapUniversalRouter.Contract.OnERC721Received(&_UniswapUniversalRouter.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_UniswapUniversalRouter *UniswapUniversalRouterCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _UniswapUniversalRouter.Contract.OnERC721Received(&_UniswapUniversalRouter.CallOpts, arg0, arg1, arg2, arg3)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_UniswapUniversalRouter *UniswapUniversalRouterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _UniswapUniversalRouter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_UniswapUniversalRouter *UniswapUniversalRouterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UniswapUniversalRouter.Contract.SupportsInterface(&_UniswapUniversalRouter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_UniswapUniversalRouter *UniswapUniversalRouterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UniswapUniversalRouter.Contract.SupportsInterface(&_UniswapUniversalRouter.CallOpts, interfaceId)
}

// CollectRewards is a paid mutator transaction binding the contract method 0x709a1cc2.
//
// Solidity: function collectRewards(bytes looksRareClaim) returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterTransactor) CollectRewards(opts *bind.TransactOpts, looksRareClaim []byte) (*types.Transaction, error) {
	return _UniswapUniversalRouter.contract.Transact(opts, "collectRewards", looksRareClaim)
}

// CollectRewards is a paid mutator transaction binding the contract method 0x709a1cc2.
//
// Solidity: function collectRewards(bytes looksRareClaim) returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterSession) CollectRewards(looksRareClaim []byte) (*types.Transaction, error) {
	return _UniswapUniversalRouter.Contract.CollectRewards(&_UniswapUniversalRouter.TransactOpts, looksRareClaim)
}

// CollectRewards is a paid mutator transaction binding the contract method 0x709a1cc2.
//
// Solidity: function collectRewards(bytes looksRareClaim) returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterTransactorSession) CollectRewards(looksRareClaim []byte) (*types.Transaction, error) {
	return _UniswapUniversalRouter.Contract.CollectRewards(&_UniswapUniversalRouter.TransactOpts, looksRareClaim)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterTransactor) Execute(opts *bind.TransactOpts, commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _UniswapUniversalRouter.contract.Transact(opts, "execute", commands, inputs)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterSession) Execute(commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _UniswapUniversalRouter.Contract.Execute(&_UniswapUniversalRouter.TransactOpts, commands, inputs)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterTransactorSession) Execute(commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _UniswapUniversalRouter.Contract.Execute(&_UniswapUniversalRouter.TransactOpts, commands, inputs)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterTransactor) Execute0(opts *bind.TransactOpts, commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapUniversalRouter.contract.Transact(opts, "execute0", commands, inputs, deadline)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterSession) Execute0(commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapUniversalRouter.Contract.Execute0(&_UniswapUniversalRouter.TransactOpts, commands, inputs, deadline)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterTransactorSession) Execute0(commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapUniversalRouter.Contract.Execute0(&_UniswapUniversalRouter.TransactOpts, commands, inputs, deadline)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapUniversalRouter.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapUniversalRouter.Contract.UniswapV3SwapCallback(&_UniswapUniversalRouter.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapUniversalRouter.Contract.UniswapV3SwapCallback(&_UniswapUniversalRouter.TransactOpts, amount0Delta, amount1Delta, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapUniversalRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterSession) Receive() (*types.Transaction, error) {
	return _UniswapUniversalRouter.Contract.Receive(&_UniswapUniversalRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapUniversalRouter *UniswapUniversalRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _UniswapUniversalRouter.Contract.Receive(&_UniswapUniversalRouter.TransactOpts)
}

// UniswapUniversalRouterRewardsSentIterator is returned from FilterRewardsSent and is used to iterate over the raw logs and unpacked data for RewardsSent events raised by the UniswapUniversalRouter contract.
type UniswapUniversalRouterRewardsSentIterator struct {
	Event *UniswapUniversalRouterRewardsSent // Event containing the contract specifics and raw log

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
func (it *UniswapUniversalRouterRewardsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapUniversalRouterRewardsSent)
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
		it.Event = new(UniswapUniversalRouterRewardsSent)
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
func (it *UniswapUniversalRouterRewardsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapUniversalRouterRewardsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapUniversalRouterRewardsSent represents a RewardsSent event raised by the UniswapUniversalRouter contract.
type UniswapUniversalRouterRewardsSent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsSent is a free log retrieval operation binding the contract event 0x1e8f03f716bc104bf7d728131967a0c771e85ab54d09c1e2d6ed9e0bc4e2a16c.
//
// Solidity: event RewardsSent(uint256 amount)
func (_UniswapUniversalRouter *UniswapUniversalRouterFilterer) FilterRewardsSent(opts *bind.FilterOpts) (*UniswapUniversalRouterRewardsSentIterator, error) {

	logs, sub, err := _UniswapUniversalRouter.contract.FilterLogs(opts, "RewardsSent")
	if err != nil {
		return nil, err
	}
	return &UniswapUniversalRouterRewardsSentIterator{contract: _UniswapUniversalRouter.contract, event: "RewardsSent", logs: logs, sub: sub}, nil
}

// WatchRewardsSent is a free log subscription operation binding the contract event 0x1e8f03f716bc104bf7d728131967a0c771e85ab54d09c1e2d6ed9e0bc4e2a16c.
//
// Solidity: event RewardsSent(uint256 amount)
func (_UniswapUniversalRouter *UniswapUniversalRouterFilterer) WatchRewardsSent(opts *bind.WatchOpts, sink chan<- *UniswapUniversalRouterRewardsSent) (event.Subscription, error) {

	logs, sub, err := _UniswapUniversalRouter.contract.WatchLogs(opts, "RewardsSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapUniversalRouterRewardsSent)
				if err := _UniswapUniversalRouter.contract.UnpackLog(event, "RewardsSent", log); err != nil {
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
func (_UniswapUniversalRouter *UniswapUniversalRouterFilterer) ParseRewardsSent(log types.Log) (*UniswapUniversalRouterRewardsSent, error) {
	event := new(UniswapUniversalRouterRewardsSent)
	if err := _UniswapUniversalRouter.contract.UnpackLog(event, "RewardsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
