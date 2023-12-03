package bindings

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/unpackdev/solgo/utils"
)

const (
	UniswapV2Factory BindingType = "UniswapV2Factory"
	UniswapV2Pair    BindingType = "UniswapV2Pair"
	UniswapV2Router  BindingType = "UniswapV2Router"
)

type Uniswap struct {
	*Manager
	network utils.Network
	ctx     context.Context
	opts    []*BindOptions
}

func NewUniswap(ctx context.Context, network utils.Network, manager *Manager, opts []*BindOptions) (*Uniswap, error) {
	if opts == nil {
		opts = DefaultUniswapBindOptions()
	}

	for _, opt := range opts {
		if err := opt.Validate(); err != nil {
			return nil, err
		}
	}

	// Now lets register all the bindings with the manager
	for _, opt := range opts {
		for _, network := range opt.Networks {
			if _, err := manager.RegisterBinding(network, opt.NetworkID, opt.Type, opt.Address, opt.ABI); err != nil {
				return nil, err
			}
		}
	}

	return &Uniswap{
		network: network,
		Manager: manager,
		ctx:     ctx,
		opts:    opts,
	}, nil
}

func (u *Uniswap) GetAddress(bindingType BindingType) (common.Address, error) {
	for _, opt := range u.opts {
		if opt.Type == bindingType {
			return opt.Address, nil
		}
	}

	return common.Address{}, fmt.Errorf("binding not found for type %s", bindingType)
}

func (u *Uniswap) WETH() (common.Address, error) {
	result, err := u.Manager.CallContractMethod(u.network, UniswapV2Router, "WETH")
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get WETH address: %w", err)
	}

	return result.(common.Address), nil
}

func (u *Uniswap) GetPair(tokenA, tokenB common.Address) (common.Address, error) {
	// Ensure tokenA is less than tokenB to conform with Uniswap's sorting
	if tokenA.Hex() > tokenB.Hex() {
		tokenA, tokenB = tokenB, tokenA
	}

	// Call the 'getPair' method on the Uniswap V2 Factory contract
	result, err := u.Manager.CallContractMethod(u.network, UniswapV2Factory, "getPair", tokenA, tokenB)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get pair: %w", err)
	}

	// The result is expected to be the address of the pair
	pairAddress, ok := result.(common.Address)
	if !ok {
		return common.Address{}, fmt.Errorf("failed to assert result as common.Address - pair address")
	}

	return pairAddress, nil
}

func (u *Uniswap) Buy(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, simulate bool) (*types.Transaction, *types.Receipt, error) {
	bind, err := u.GetBinding(utils.Ethereum, UniswapV2Router)
	if err != nil {
		return nil, nil, err
	}
	bindAbi := bind.GetABI()

	method, exists := bindAbi.Methods["swapExactETHForTokensSupportingFeeOnTransferTokens"]
	if !exists {
		return nil, nil, errors.New("swap method not found")
	}

	input, err := bindAbi.Pack(method.Name, amountOutMin, path, to, deadline)
	if err != nil {
		return nil, nil, err
	}

	if simulate {
		txHash, err := u.Manager.SendSimulatedTransaction(opts, u.network, &bind.Address, method, input)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to send swapExactETHForTokensSupportingFeeOnTransferTokens transaction: %w", err)
		}

		spew.Dump(opts)

		receipt, err := u.Manager.WaitForReceipt(u.ctx, u.network, *txHash)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get swapExactETHForTokensSupportingFeeOnTransferTokens transaction receipt: %w", err)
		}

		tx, _, err := u.Manager.GetTransactionByHash(u.ctx, u.network, receipt.TxHash)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get swapExactETHForTokensSupportingFeeOnTransferTokens transaction by hash: %w", err)
		}

		return tx, receipt, nil
	}

	tx, err := u.Manager.SendTransaction(opts, u.network, &bind.Address, input)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to send swapExactETHForTokensSupportingFeeOnTransferTokens transaction: %w", err)
	}

	receipt, err := u.Manager.WaitForReceipt(u.ctx, u.network, tx.Hash())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get swapExactETHForTokensSupportingFeeOnTransferTokens transaction receipt: %w", err)
	}

	return tx, receipt, nil
}

func (u *Uniswap) Sell(opts *bind.TransactOpts, amountIn *big.Int, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int, simulate bool) (*types.Transaction, *types.Receipt, error) {
	bind, err := u.GetBinding(utils.Ethereum, UniswapV2Router)
	if err != nil {
		return nil, nil, err
	}
	bindAbi := bind.GetABI()

	method, exists := bindAbi.Methods["swapExactTokensForETHSupportingFeeOnTransferTokens"]
	if !exists {
		return nil, nil, errors.New("swap method not found")
	}

	input, err := bindAbi.Pack(method.Name, amountIn, amountOut, path, to, deadline)
	if err != nil {
		return nil, nil, err
	}

	if simulate {
		txHash, err := u.Manager.SendSimulatedTransaction(opts, u.network, &bind.Address, method, input)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to send swapExactTokensForETHSupportingFeeOnTransferTokens transaction: %w", err)
		}

		spew.Dump(opts)

		receipt, err := u.Manager.WaitForReceipt(u.ctx, u.network, *txHash)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get swapExactTokensForETHSupportingFeeOnTransferTokens transaction receipt: %w", err)
		}

		tx, _, err := u.Manager.GetTransactionByHash(u.ctx, u.network, receipt.TxHash)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get swapExactTokensForETHSupportingFeeOnTransferTokens transaction by hash: %w", err)
		}

		return tx, receipt, nil
	}

	tx, err := u.Manager.SendTransaction(opts, u.network, &bind.Address, input)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to send swapExactTokensForETHSupportingFeeOnTransferTokens transaction: %w", err)
	}

	receipt, err := u.Manager.WaitForReceipt(u.ctx, u.network, tx.Hash())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get swapExactTokensForETHSupportingFeeOnTransferTokens transaction receipt: %w", err)
	}

	return tx, receipt, nil
}

func (u *Uniswap) EstimateTaxesForToken(tokenAddress common.Address) (*big.Int, error) {
	/* 	wEthAddr, err := u.WETH()
	   	if err != nil {
	   		return nil, err
	   	}

	   	// Figure out the pair address...
	   	pairAddress, err := u.GetPair(tokenAddress, wEthAddr)
	   	if err != nil {
	   		return nil, err
	   	}

	   	fmt.Println("Pair Address:", pairAddress.Hex())

	   	currencyEth, _ := currencies.NewEther(big.NewInt(100000000000000000))
	   	fmt.Println(currencyEth.Addresses[utils.Ethereum].Hex())

	   	amounts, err := u.CalculateAmounts(currencyEth.Raw(), tokenAddress, wEthAddr)
	   	if err != nil {
	   		panic(err)
	   		return nil, err
	   	}

	   	fmt.Println("Amounts:", amounts)

	   	utils.DumpNodeWithExit(1) */
	return nil, nil
}

func (u *Uniswap) CalculateAmounts(amount *big.Int, baseAddress, quoteAddress common.Address) ([]common.Address, error) {
	/*
		bind, err := u.GetBinding(utils.Ethereum, UniswapV2Router)
		if err != nil {
			return nil, err
		}

		 	bindAbi := bind.GetABI()

		   	createPairMethod, exists := bindAbi.Methods["swapExactTokensForTokensSupportingFeeOnTransferTokens"]
		   	if !exists {
		   		return nil, errors.New("createPair method not found")
		   	}

		   	_ = createPairMethod

		   	getAmountsOut, exists := bindAbi.Methods["getAmountsOut"]
		   	if !exists {
		   		return nil, errors.New("getAmountsOut method not found")
		   	}

		   	packedInput, err := getAmountsOut.Inputs.Pack(amount, baseAddress, quoteAddress)
		   	if err != nil {
		   		return nil, err
		   	} */

	amountsOut, err := u.Manager.CallContractMethod(utils.Ethereum, UniswapV2Router, "getAmountsOut", amount, []common.Address{baseAddress, quoteAddress})
	if err != nil {
		return nil, err
	}

	amountsIn, err := u.Manager.CallContractMethod(utils.Ethereum, UniswapV2Router, "getAmountsIn", amount, []common.Address{baseAddress, quoteAddress})
	if err != nil {
		return nil, err
	}

	fmt.Println(amountsOut, amountsIn)

	return nil, nil
}

func (u *Uniswap) CreatePair(opts *bind.TransactOpts, tokenA, tokenB common.Address) (*types.Transaction, error) {
	// Ensure tokenA is less than tokenB
	if tokenA.Hex() > tokenB.Hex() {
		tokenA, tokenB = tokenB, tokenA
	}

	bind, err := u.GetBinding(utils.Ethereum, UniswapV2Factory)
	if err != nil {
		return nil, err
	}

	bindAbi := bind.GetABI()

	// Find the `createPair` method in the ABI
	createPairMethod, exists := bindAbi.Methods["createPair"]
	if !exists {
		return nil, errors.New("createPair method not found")
	}

	// ABI encode the input parameters for the `createPair` method
	input, err := createPairMethod.Inputs.Pack(tokenA, tokenB)
	if err != nil {
		return nil, err
	}

	// Send the createPair transaction
	tx, err := u.Manager.SendTransaction(opts, u.network, &bind.Address, input)
	if err != nil {
		return nil, fmt.Errorf("failed to send createPair transaction: %w", err)
	}

	return tx, nil
}

func (u *Uniswap) FetchPairs(pairAddresses []common.Address) ([]*PairDetails, error) {
	var pairs []*PairDetails

	for _, pairAddress := range pairAddresses {
		pair, err := u.FetchPair(pairAddress)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, pair)
	}

	return pairs, nil
}

func (u *Uniswap) FetchPair(pairAddress common.Address) (*PairDetails, error) {
	// Use Manager to call the 'token0', 'token1', and 'getReserves' methods on the pair contract
	token0, err := u.Manager.CallContractMethod(utils.Ethereum, UniswapV2Pair, "token0", pairAddress)
	if err != nil {
		return nil, err
	}
	token1, err := u.Manager.CallContractMethod(utils.Ethereum, UniswapV2Pair, "token1", pairAddress)
	if err != nil {
		return nil, err
	}

	reserves, err := u.Manager.CallContractMethod(utils.Ethereum, UniswapV2Pair, "getReserves", pairAddress)
	if err != nil {
		return nil, err
	}

	// Parse the result to get the reserve values
	reserveA, reserveB, err := parseReserves(reserves)
	if err != nil {
		return nil, err
	}

	return &PairDetails{
		Address:  pairAddress,
		Token0:   common.HexToAddress(token0.(string)),
		Token1:   common.HexToAddress(token1.(string)),
		Reserve0: reserveA,
		Reserve1: reserveB,
	}, nil
}

/*
func (u *Uniswap) AddLiquidity(pairAddress common.Address, amountTokenA, amountTokenB *big.Int) error {
	// Implement logic to add liquidity to the specified pair
	// This might involve interacting with the pair contract directly
	// or through a router contract if necessary
}

func (u *Uniswap) RemoveLiquidity(pairAddress common.Address, liquidity *big.Int) error {
	// Implement logic to remove liquidity from the specified pair
}

func (u *Uniswap) SwapTokens(pairAddress common.Address, amountIn *big.Int, tokenIn, tokenOut common.Address) (*big.Int, error) {
	// Implement logic to simulate a token swap
	// This is useful to analyze the pair's behavior in a trade
}

func (u *Uniswap) GetReserves(pairAddress common.Address) (*big.Int, *big.Int, error) {
	// Implement logic to interact with the Uniswap Pair contract
	// to fetch the reserves for each token in the pair
} */

func GetPairAddress(factoryAddress, token0, token1 common.Address) (common.Address, error) {
	// Ensure token0 is less than token1
	if token0.Hex() > token1.Hex() {
		token0, token1 = token1, token0
	}

	// Create the hash of the token addresses
	tokenHash := crypto.Keccak256Hash(abiEncodePacked(token0, token1))

	// Prepend the "ff" prefix and append the Uniswap-specific hash
	data := append([]byte{0xff}, factoryAddress.Bytes()...)
	data = append(data, tokenHash.Bytes()...)
	data = append(data, common.Hex2Bytes("96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f")...)

	// Calculate the final address hash
	finalHash := crypto.Keccak256Hash(data)

	// Convert the hash to an address
	return common.BytesToAddress(finalHash.Bytes()[12:]), nil
}

func abiEncodePacked(addresses ...common.Address) []byte {
	var data []byte
	for _, address := range addresses {
		data = append(data, address.Bytes()...)
	}
	return data
}

func parseReserves(data any) (*big.Int, *big.Int, error) {
	// Assuming data is a byte slice returned from the contract call
	if dataBytes, ok := data.([]byte); ok {
		// The first 32 bytes are the length of the array (skip it)
		// Next 32 bytes each for the two reserve values
		if len(dataBytes) >= 96 { // 32 bytes length + 2 * 32 bytes values
			reserve0 := new(big.Int).SetBytes(dataBytes[32:64])
			reserve1 := new(big.Int).SetBytes(dataBytes[64:96])
			return reserve0, reserve1, nil
		}
	}
	return nil, nil, fmt.Errorf("invalid data format for reserves")
}

func DefaultUniswapBindOptions() []*BindOptions {
	return []*BindOptions{
		{
			Networks:  []utils.Network{utils.Ethereum, utils.AnvilNetwork},
			NetworkID: utils.EthereumNetworkID,
			Name:      "UniswapV2: Router",
			Type:      UniswapV2Router,
			Address:   common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"),
			ABI:       `[{"inputs":[{"internalType":"address","name":"_factory","type":"address"},{"internalType":"address","name":"_WETH","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"WETH","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"},{"internalType":"uint256","name":"amountADesired","type":"uint256"},{"internalType":"uint256","name":"amountBDesired","type":"uint256"},{"internalType":"uint256","name":"amountAMin","type":"uint256"},{"internalType":"uint256","name":"amountBMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"addLiquidity","outputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"amountB","type":"uint256"},{"internalType":"uint256","name":"liquidity","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"amountTokenDesired","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"addLiquidityETH","outputs":[{"internalType":"uint256","name":"amountToken","type":"uint256"},{"internalType":"uint256","name":"amountETH","type":"uint256"},{"internalType":"uint256","name":"liquidity","type":"uint256"}],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"factory","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"uint256","name":"reserveIn","type":"uint256"},{"internalType":"uint256","name":"reserveOut","type":"uint256"}],"name":"getAmountIn","outputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"reserveIn","type":"uint256"},{"internalType":"uint256","name":"reserveOut","type":"uint256"}],"name":"getAmountOut","outputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"}],"name":"getAmountsIn","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"}],"name":"getAmountsOut","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"reserveA","type":"uint256"},{"internalType":"uint256","name":"reserveB","type":"uint256"}],"name":"quote","outputs":[{"internalType":"uint256","name":"amountB","type":"uint256"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountAMin","type":"uint256"},{"internalType":"uint256","name":"amountBMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"removeLiquidity","outputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"amountB","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"removeLiquidityETH","outputs":[{"internalType":"uint256","name":"amountToken","type":"uint256"},{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"removeLiquidityETHSupportingFeeOnTransferTokens","outputs":[{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"bool","name":"approveMax","type":"bool"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"removeLiquidityETHWithPermit","outputs":[{"internalType":"uint256","name":"amountToken","type":"uint256"},{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"bool","name":"approveMax","type":"bool"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens","outputs":[{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountAMin","type":"uint256"},{"internalType":"uint256","name":"amountBMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"bool","name":"approveMax","type":"bool"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"removeLiquidityWithPermit","outputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"amountB","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapETHForExactTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactETHForTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactETHForTokensSupportingFeeOnTransferTokens","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForETH","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForETHSupportingFeeOnTransferTokens","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForTokensSupportingFeeOnTransferTokens","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"uint256","name":"amountInMax","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapTokensForExactETH","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"uint256","name":"amountInMax","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapTokensForExactTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"},{"stateMutability":"payable","type":"receive"}]`,
		},
		{
			Networks:  []utils.Network{utils.Ethereum, utils.AnvilNetwork},
			NetworkID: utils.EthereumNetworkID,
			Name:      "UniswapV2: Factory Contract",
			Type:      UniswapV2Factory,
			Address:   common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"),
			ABI:       `[{"inputs":[{"internalType":"address","name":"_feeToSetter","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"token0","type":"address"},{"indexed":true,"internalType":"address","name":"token1","type":"address"},{"indexed":false,"internalType":"address","name":"pair","type":"address"},{"indexed":false,"internalType":"uint256","name":"","type":"uint256"}],"name":"PairCreated","type":"event"},{"constant":true,"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"allPairs","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"allPairsLength","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"}],"name":"createPair","outputs":[{"internalType":"address","name":"pair","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"feeTo","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"feeToSetter","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"","type":"address"},{"internalType":"address","name":"","type":"address"}],"name":"getPair","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"_feeTo","type":"address"}],"name":"setFeeTo","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"_feeToSetter","type":"address"}],"name":"setFeeToSetter","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`,
		},
		{
			Networks:  []utils.Network{utils.Ethereum, utils.AnvilNetwork},
			NetworkID: utils.EthereumNetworkID,
			Name:      "UniswapV2: Pair",
			Type:      UniswapV2Pair,
			Address:   common.HexToAddress("0x3356c9a8f40f8e9c1d192a4347a76d18243fabc5"),
			ABI:       `[{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"spender","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"sender","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount0","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount1","type":"uint256"},{"indexed":true,"internalType":"address","name":"to","type":"address"}],"name":"Burn","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"sender","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount0","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount1","type":"uint256"}],"name":"Mint","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"sender","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount0In","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount1In","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount0Out","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount1Out","type":"uint256"},{"indexed":true,"internalType":"address","name":"to","type":"address"}],"name":"Swap","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint112","name":"reserve0","type":"uint112"},{"indexed":false,"internalType":"uint112","name":"reserve1","type":"uint112"}],"name":"Sync","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"constant":true,"inputs":[],"name":"DOMAIN_SEPARATOR","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"MINIMUM_LIQUIDITY","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"PERMIT_TYPEHASH","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"","type":"address"},{"internalType":"address","name":"","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"to","type":"address"}],"name":"burn","outputs":[{"internalType":"uint256","name":"amount0","type":"uint256"},{"internalType":"uint256","name":"amount1","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"factory","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getReserves","outputs":[{"internalType":"uint112","name":"_reserve0","type":"uint112"},{"internalType":"uint112","name":"_reserve1","type":"uint112"},{"internalType":"uint32","name":"_blockTimestampLast","type":"uint32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"_token0","type":"address"},{"internalType":"address","name":"_token1","type":"address"}],"name":"initialize","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"kLast","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"to","type":"address"}],"name":"mint","outputs":[{"internalType":"uint256","name":"liquidity","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"nonces","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"permit","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"price0CumulativeLast","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"price1CumulativeLast","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"to","type":"address"}],"name":"skim","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"uint256","name":"amount0Out","type":"uint256"},{"internalType":"uint256","name":"amount1Out","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"swap","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"sync","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"token0","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"token1","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"transferFrom","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]`,
		},
	}
}

func GetBindByType(opts []*BindOptions, t BindingType) *BindOptions {
	for _, opt := range opts {
		if opt.Type == t {
			return opt
		}
	}
	return nil
}
