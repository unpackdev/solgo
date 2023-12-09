package exchanges

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/accounts"
	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/bytecode"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/simulator"
	"github.com/unpackdev/solgo/utils"
	"github.com/unpackdev/solgo/utils/entities"
)

type UniswapV2Exchange struct {
	ctx         context.Context
	clientPool  *clients.ClientPool
	opts        *ExchangeOptions
	uniswapBind *bindings.Uniswap
	sim         *simulator.Simulator
	bindings    map[bindings.BindingType]*bindings.Binding
}

func NewUniswapV2(ctx context.Context, clientsPool *clients.ClientPool, sim *simulator.Simulator, uniswapBind *bindings.Uniswap, opts *ExchangeOptions) (*UniswapV2Exchange, error) {
	if clientsPool == nil {
		return nil, fmt.Errorf("uniswapv2 exchange: clients pool is nil")
	}

	if opts == nil {
		return nil, fmt.Errorf("uniswapv2 exchange: options are nil")
	}

	wethBind, err := bindings.NewWETH(ctx, uniswapBind.Manager, bindings.DefaultWETHBindOptions())
	if err != nil {
		return nil, fmt.Errorf("failed to create weth binding: %s", err)
	}

	wethBinding, err := wethBind.GetBinding(utils.AnvilNetwork, bindings.WETH9)
	if err != nil {
		return nil, fmt.Errorf("failed to get weth binding: %s", err)
	}

	pairBind, err := uniswapBind.GetBinding(utils.AnvilNetwork, bindings.UniswapV2Pair)
	if err != nil {
		return nil, fmt.Errorf("failed to get uniswap v2 pair binding: %s", err)
	}

	uniswapv3Bind, err := bindings.NewUniswapV3(ctx, uniswapBind.Manager, bindings.DefaultUniswapV3BindOptions())
	if err != nil {
		return nil, fmt.Errorf("failed to create uniswap v3 binding: %s", err)
	}

	uniswapv3Binding, err := uniswapv3Bind.GetBinding(utils.AnvilNetwork, bindings.UniswapV3Pool)
	if err != nil {
		return nil, fmt.Errorf("failed to get uniswap v3 pool binding: %s", err)
	}

	return &UniswapV2Exchange{
		ctx:         ctx,
		uniswapBind: uniswapBind,
		clientPool:  clientsPool,
		opts:        opts,
		sim:         sim,
		bindings: map[bindings.BindingType]*bindings.Binding{
			bindings.WETH9:         wethBinding,
			bindings.UniswapV2Pair: pairBind,
			bindings.UniswapV3Pool: uniswapv3Binding,
		},
	}, nil
}

// ToUniswapV2 converts an Exchange to a UniswapV2Exchange. This is a helper function that you can use to
// access interface methods that are not part of the Exchange interface.
func ToUniswapV2(exchange Exchange) *UniswapV2Exchange {
	return exchange.(*UniswapV2Exchange)
}

func (u *UniswapV2Exchange) GetType() utils.ExchangeType {
	return utils.UniswapV2
}

func (u *UniswapV2Exchange) GetRouterAddress() common.Address {
	return u.opts.RouterAddress
}

func (u *UniswapV2Exchange) GetFactoryAddress() common.Address {
	return u.opts.FactoryAddress
}

func (u *UniswapV2Exchange) GetOptions() *ExchangeOptions {
	return u.opts
}

func (u *UniswapV2Exchange) GetClient(ctx context.Context, network utils.Network, simulatorType utils.SimulatorType, atBlock *big.Int) (*clients.Client, error) {
	if u.sim != nil && simulatorType != utils.NoSimulator {
		client, _, err := u.sim.GetClient(ctx, simulatorType, atBlock)
		if err != nil {
			return nil, fmt.Errorf("failed to get client from simulator: %s", err)
		}
		return client, nil
	}

	// This is going to return one of the normal clients (not simulated) from the pool.
	return u.clientPool.GetClientByGroup(network.String()), nil
}

func (u *UniswapV2Exchange) Buy(ctx context.Context, network utils.Network, simulatorType utils.SimulatorType, tokenBind *bindings.Token, spender *accounts.Account, baseToken *entities.Token, quoteToken *entities.Token, amount *entities.CurrencyAmount, atBlock *big.Int) (*UniswapV2TradeDescriptor, error) {
	networkID := utils.GetNetworkID(network)
	toReturn := &UniswapV2TradeDescriptor{
		ExchangeType:   utils.UniswapV2,
		Network:        network,
		Simulation:     network == utils.AnvilNetwork,
		NetworkID:      networkID,
		RouterAddress:  u.opts.RouterAddress,
		FactoryAddress: u.opts.FactoryAddress,
		WETHAddress:    entities.WETH9[uint(networkID)].Address,
		SpenderAddress: spender.Address,
		AmountRaw:      amount.Quotient(),
		Amount:         amount.ToExact(),
	}

	client, err := u.GetClient(ctx, network, simulatorType, atBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to get client: %s", err)
	}

	usdtToken := entities.USDT[uint(networkID)]

	currentBalance, err := spender.Balance(ctx, atBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to get current balance: %s", err)
	}
	toReturn.SpenderBalanceBefore = currentBalance

	pairAddr, err := u.uniswapBind.GetPair(ctx, quoteToken.Address, baseToken.Address)
	if err != nil {
		return nil, err
	}
	toReturn.PairAddress = pairAddr

	reserves, err := u.uniswapBind.GetReserves(ctx, pairAddr)
	if err != nil {
		return nil, err
	}

	unixReserveTime := time.Unix(int64(reserves.BlockTimestampLast), 0)
	toReturn.PairReserves = &UniswapV2PairReserves{
		Token0:    quoteToken.Address,
		Token1:    baseToken.Address,
		Reserve0:  reserves.Reserve0,
		Reserve1:  reserves.Reserve1,
		BlockTime: unixReserveTime,
	}

	var inverted bool
	var reserveIn *big.Int
	var reserveOut *big.Int
	var tokenIn *entities.Token
	var tokenOut *entities.Token

	if toReturn.PairReserves.Reserve1.Uint64() < toReturn.PairReserves.Reserve0.Uint64() {
		inverted = true
		reserveIn = toReturn.PairReserves.Reserve1
		reserveOut = toReturn.PairReserves.Reserve0
		tokenIn = quoteToken
		tokenOut = baseToken
	} else {
		reserveIn = toReturn.PairReserves.Reserve0
		reserveOut = toReturn.PairReserves.Reserve1
		tokenIn = baseToken
		tokenOut = quoteToken
	}

	toReturn.Price = entities.NewPrice(tokenIn, tokenOut, reserveIn, reserveOut)

	if inverted {
		toReturn.PricePerToken = toReturn.Price.Invert().ToSignificant(9)
	} else {
		if tokenIn.Decimals() > tokenOut.Decimals() {
			toReturn.Price = entities.NewPrice(tokenIn, tokenOut, reserveOut, reserveIn)
			toReturn.PricePerToken = toReturn.Price.Invert().ToSignificant(9)
		} else {
			toReturn.PricePerToken = toReturn.Price.ToSignificant(9)
		}
	}

	yes, _ := tokenIn.SortsBefore(tokenOut)
	spew.Dump(
		toReturn.Price.Invert().Quotient(),
		toReturn.Price.Quotient(),
		tokenIn.Decimals(),
		tokenOut.Decimals(),
		inverted,
		reserveIn,
		reserveOut,
		yes,
	)

	/* 	if inverted {
		toReturn.PricePerToken = toReturn.Price.Invert().ToSignificant(9)
	} else { */
	//toReturn.PricePerToken = toReturn.Price.ToSignificant(9)
	/* 	} */

	usdtPairAddr, err := u.uniswapBind.GetPair(ctx, usdtToken.Address, entities.WETH9[1].Address)
	if err != nil {
		return nil, err
	}

	usdtReserves, err := u.uniswapBind.GetReserves(ctx, usdtPairAddr)
	if err != nil {
		return nil, err
	}

	usdtEthPrice := entities.NewPrice(usdtToken, entities.WETH9[1], usdtReserves.Reserve1, usdtReserves.Reserve0)
	toReturn.UsdToEthPriceRaw = usdtEthPrice
	toReturn.UsdToEthPrice = usdtEthPrice.ToSignificant(9)
	toReturn.EthToUsdPriceRaw = usdtEthPrice.Invert()
	toReturn.EthToUsdPrice = toReturn.EthToUsdPriceRaw.ToSignificant(9)

	// Calculate Token Price in USD
	// First, ensure both prices are in the same scale (adjust decimals if needed)
	//tokenPriceInUsdRaw := new(big.Int).Mul(toReturn.Price.Quotient(), toReturn.UsdToEthPriceRaw.Quotient())

	// Adjust for the decimals to get the final price in USD
	// Assuming 18 decimals for ETH and your token
	//tokenPriceInUsd := new(big.Float).Quo(new(big.Float).SetInt(tokenPriceInUsdRaw), big.NewFloat(math.Pow10(18)))

	//toReturn.PricePerTokenUsd = entities.FromRawAmount(entities.WETH9[1], tokenPriceInUsdRaw).Invert().ToFixed(9)

	amountOut, err := u.uniswapBind.GetAmountOut(ctx, amount.Quotient(), reserves.Reserve1, reserves.Reserve0)
	if err != nil {
		return nil, err
	}
	toReturn.MaxAmountRaw = amountOut
	toReturn.MaxAmount = entities.FromRawAmount(tokenOut, amountOut).ToFixed(2)

	tokenBinding, _ := tokenBind.GetBinding(utils.AnvilNetwork, bindings.Erc20)
	uniswapBinding, _ := u.uniswapBind.GetBinding(utils.AnvilNetwork, bindings.UniswapV2Router)

	authApprove, err := spender.TransactOpts(client, nil, false) // Approval cannot take value as value is in the approve call
	if err != nil {
		return nil, fmt.Errorf("failed to create approve transact opts: %s", err)
	}

	_, approveReceiptTx, err := tokenBind.Approve(ctx, network, simulatorType, client, authApprove, u.opts.RouterAddress, amount.Quotient(), atBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to approve token: %s", err)
	}

	approvalResults := &AuditApprovalResults{
		Detected:          true,
		ApprovalRequested: true,
		Approved: func() bool {
			if approveReceiptTx != nil {
				return approveReceiptTx.Status == 1
			}

			return false
		}(),
		TxHash:  approveReceiptTx.TxHash,
		Receipt: approveReceiptTx != nil,
		ReceiptStatus: func() uint64 {
			if approveReceiptTx != nil {
				return approveReceiptTx.Status
			}

			return 0
		}(),
		Logs: make([]*bytecode.Log, 0),
	}

	tradeRequest := &AuditBuyOrSellResults{
		Approval: approvalResults,
	}

	if approveReceiptTx != nil {
		approvalResults.GasUsedRaw = approveReceiptTx.GasUsed
		approvalResults.GasUsed = entities.FromRawAmount(quoteToken, big.NewInt(int64(approveReceiptTx.GasUsed))).ToFixed(int32(quoteToken.Decimals()))

		if len(approveReceiptTx.Logs) > 0 {
			if decodedApprovalLog, err := bytecode.DecodeLogFromAbi(approveReceiptTx.Logs[0], []byte(tokenBinding.RawABI)); err == nil {
				tradeRequest.Approval.Logs = append(tradeRequest.Approval.Logs, decodedApprovalLog)
			}
		}
	}

	buyResults := &AuditSwapResults{
		Detected:      false,
		SwapRequested: true,
		PairDetails:   []common.Address{baseToken.Address, quoteToken.Address},
	}

	authBuy, err := spender.TransactOpts(client, amount.Quotient(), false) // Approval cannot take value as value is in the approve call
	if err != nil {
		return nil, fmt.Errorf("failed to create buy transact opts: %s", err)
	}

	// We are using hack to pretend sending normal transaction while using simulated client...
	// Therefore, instead of passing simulatorType we pass utils.NoSimulator
	deadline := big.NewInt(time.Now().Add(time.Minute).Unix())
	_, buyReceipt, err := u.uniswapBind.Buy(authBuy, network, simulatorType, client, big.NewInt(0), buyResults.PairDetails, spender.Address, deadline)
	if err != nil {
		return nil, fmt.Errorf("failed to buy token: %s", err)

	}

	buyResults.Detected = true
	buyResults.TxHash = buyReceipt.TxHash
	buyResults.Receipt = buyReceipt != nil
	buyResults.ReceiptStatus = func() uint64 {
		if buyReceipt != nil {
			return buyReceipt.Status
		}

		return 0
	}()

	if buyReceipt != nil {
		buyResults.GasUsedRaw = buyReceipt.GasUsed
		buyResults.GasUsed = entities.FromRawAmount(quoteToken, big.NewInt(int64(buyReceipt.GasUsed))).ToFixed(int32(quoteToken.Decimals()))

		if len(buyReceipt.Logs) > 0 {
			var buyLogs []*bytecode.Log
			for _, log := range buyReceipt.Logs {
				if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(tokenBinding.RawABI)); err == nil {
					buyLogs = append(buyLogs, decodedBuyLog)
				} else if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(uniswapBinding.RawABI)); err == nil {
					buyLogs = append(buyLogs, decodedBuyLog)
				} else if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(u.bindings[bindings.UniswapV2Pair].RawABI)); err == nil {
					buyLogs = append(buyLogs, decodedBuyLog)
				} else if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(u.bindings[bindings.WETH9].RawABI)); err == nil {
					buyLogs = append(buyLogs, decodedBuyLog)
				} else if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(u.bindings[bindings.UniswapV3Pool].RawABI)); err == nil {
					buyLogs = append(buyLogs, decodedBuyLog)
				}
			}

			buyResults.Logs = buyLogs

			for _, log := range buyLogs {
				if log.Type == utils.SwapLogEventType {
					if amountOut, ok := log.Data["amount0Out"]; ok {
						if amountOut, ok := amountOut.(*big.Int); ok {
							buyResults.SwapReceivedAmountRaw = amountOut
							buyResults.SwapReceivedAmount = entities.FromRawAmount(tokenOut, amountOut).ToFixed(2)
						}
					}

					if buyResults.SwapReceivedAmountRaw == nil || buyResults.SwapReceivedAmountRaw.Uint64() == 0 {
						if amountOut, ok := log.Data["amount1Out"]; ok {
							if amountOut, ok := amountOut.(*big.Int); ok {
								buyResults.SwapReceivedAmountRaw = amountOut
								buyResults.SwapReceivedAmount = entities.FromRawAmount(tokenOut, amountOut).ToFixed(2)
							}
						}
					}
				}

				if log.Type == utils.TransferLogEventType {
					for _, topic := range log.Topics {
						if topic.Name == "to" && topic.Value == spender.Address {
							if amountOut, ok := log.Data["value"]; ok {
								if amountOut, ok := amountOut.(*big.Int); ok {
									buyResults.ReceivedAmountRaw = amountOut
									buyResults.ReceivedAmount = entities.FromRawAmount(tokenOut, amountOut).ToFixed(2)
								}
							}
						}
					}
				}
			}
		}
	}

	if buyResults.ReceivedAmountRaw != nil && buyResults.SwapReceivedAmountRaw != nil {
		tax := CalculatePercentageDifference(buyResults.SwapReceivedAmountRaw, buyResults.ReceivedAmountRaw, quoteToken.Decimals())
		buyResults.TaxRaw = tax
		buyResults.Tax = entities.FromRawAmount(quoteToken, tax).ToFixed(2)
	}

	tradeRequest.Results = buyResults
	toReturn.Trade = tradeRequest

	if approvalResults.Detected || buyResults.Detected {
		toReturn.Trade.Detected = true
	}

	afterBalance, err := spender.Balance(ctx, atBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to get current balance: %s", err)
	}

	toReturn.SpenderBalanceAfter = afterBalance
	utils.DumpNodeNoExit(toReturn)
	return toReturn, nil
}

func (u *UniswapV2Exchange) Sell(ctx context.Context, network utils.Network, simulatorType utils.SimulatorType, tokenBind *bindings.Token, spender *accounts.Account, baseToken *entities.Token, quoteToken *entities.Token, amount *entities.CurrencyAmount, atBlock *big.Int) (*UniswapV2TradeDescriptor, error) {
	networkID := utils.GetNetworkID(network)
	toReturn := &UniswapV2TradeDescriptor{
		ExchangeType:   utils.UniswapV2,
		Network:        network,
		Simulation:     network == utils.AnvilNetwork,
		NetworkID:      networkID,
		RouterAddress:  u.opts.RouterAddress,
		FactoryAddress: u.opts.FactoryAddress,
		WETHAddress:    entities.WETH9[uint(networkID)].Address,
		SpenderAddress: spender.Address,
		AmountRaw:      amount.Quotient(),
		Amount:         amount.ToExact(),
	}

	client, err := u.GetClient(ctx, network, simulatorType, atBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to get client: %s", err)
	}

	wethBind, err := bindings.NewWETH(ctx, tokenBind.Manager, bindings.DefaultWETHBindOptions())
	if err != nil {
		return nil, fmt.Errorf("failed to create weth binding: %s", err)
	}
	wethBinding, _ := wethBind.GetBinding(utils.AnvilNetwork, bindings.WETH9)

	fiatBind, err := bindings.NewFiat(ctx, tokenBind.Manager, bindings.DefaultFiatBindOptions())
	if err != nil {
		return nil, fmt.Errorf("failed to create usdc binding: %s", err)
	}
	usdcBinding, _ := fiatBind.GetBinding(utils.AnvilNetwork, bindings.USDC)
	usdtToken := entities.USDT[uint(networkID)]
	_ = usdcBinding

	currentBalance, err := spender.Balance(ctx, atBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to get current balance: %s", err)
	}
	toReturn.SpenderBalanceBefore = currentBalance

	pairAddr, err := u.uniswapBind.GetPair(ctx, quoteToken.Address, baseToken.Address)
	if err != nil {
		return nil, err
	}
	toReturn.PairAddress = pairAddr

	reserves, err := u.uniswapBind.GetReserves(ctx, pairAddr)
	if err != nil {
		return nil, err
	}

	unixReserveTime := time.Unix(int64(reserves.BlockTimestampLast), 0)
	toReturn.PairReserves = &UniswapV2PairReserves{
		Token0:    quoteToken.Address,
		Token1:    baseToken.Address,
		Reserve0:  reserves.Reserve0,
		Reserve1:  reserves.Reserve1,
		BlockTime: unixReserveTime,
	}

	var inverted bool
	var reserveIn *big.Int
	var reserveOut *big.Int
	var tokenIn *entities.Token
	var tokenOut *entities.Token

	if toReturn.PairReserves.Reserve1.Uint64() < toReturn.PairReserves.Reserve0.Uint64() {
		inverted = true
		reserveIn = toReturn.PairReserves.Reserve1
		reserveOut = toReturn.PairReserves.Reserve0
		tokenIn = quoteToken
		tokenOut = baseToken
	} else {
		reserveIn = toReturn.PairReserves.Reserve0
		reserveOut = toReturn.PairReserves.Reserve1
		tokenIn = baseToken
		tokenOut = quoteToken
	}

	toReturn.Price = entities.NewPrice(tokenIn, tokenOut, reserveIn, reserveOut)

	if inverted {
		toReturn.PricePerToken = toReturn.Price.Invert().ToSignificant(9)
	} else {
		if tokenIn.Decimals() > tokenOut.Decimals() {
			toReturn.PricePerToken = toReturn.Price.Invert().ToSignificant(9)
		} else {
			toReturn.PricePerToken = toReturn.Price.ToSignificant(9)
		}
	}

	usdtPairAddr, err := u.uniswapBind.GetPair(ctx, usdtToken.Address, entities.WETH9[1].Address)
	if err != nil {
		return nil, err
	}

	usdtReserves, err := u.uniswapBind.GetReserves(ctx, usdtPairAddr)
	if err != nil {
		return nil, err
	}

	usdtEthPrice := entities.NewPrice(usdtToken, entities.WETH9[1], usdtReserves.Reserve1, usdtReserves.Reserve0)
	toReturn.UsdToEthPriceRaw = usdtEthPrice
	toReturn.UsdToEthPrice = usdtEthPrice.ToSignificant(9)
	toReturn.EthToUsdPriceRaw = usdtEthPrice.Invert()
	toReturn.EthToUsdPrice = toReturn.EthToUsdPriceRaw.ToSignificant(9)

	// Calculate Token Price in USD
	// First, ensure both prices are in the same scale (adjust decimals if needed)
	//tokenPriceInUsdRaw := new(big.Int).Mul(toReturn.Price.Quotient(), toReturn.UsdToEthPriceRaw.Quotient())

	// Adjust for the decimals to get the final price in USD
	// Assuming 18 decimals for ETH and your token
	//tokenPriceInUsd := new(big.Float).Quo(new(big.Float).SetInt(tokenPriceInUsdRaw), big.NewFloat(math.Pow10(18)))

	//toReturn.PricePerTokenUsd = entities.FromRawAmount(entities.WETH9[1], tokenPriceInUsdRaw).Invert().ToFixed(9)

	amountOut, err := u.uniswapBind.GetAmountOut(ctx, amount.Quotient(), reserves.Reserve0, reserves.Reserve1)
	if err != nil {
		return nil, err
	}
	toReturn.MaxAmountRaw = amountOut
	toReturn.MaxAmount = entities.FromRawAmount(tokenOut, amountOut).ToFixed(2)

	tokenBinding, _ := tokenBind.GetBinding(utils.AnvilNetwork, bindings.Erc20)
	uniswapBinding, _ := u.uniswapBind.GetBinding(utils.AnvilNetwork, bindings.UniswapV2Router)
	uniswapPairBinding, _ := u.uniswapBind.GetBinding(utils.AnvilNetwork, bindings.UniswapV2Pair)

	authApprove, err := spender.TransactOpts(client, amount.Quotient(), false) // Approval cannot take value as value is in the approve call
	if err != nil {
		return nil, fmt.Errorf("failed to create sell approve transact opts: %s", err)
	}
	authApprove.Value = big.NewInt(0)

	_, approveReceiptTx, err := tokenBind.Approve(ctx, network, simulatorType, client, authApprove, u.opts.RouterAddress, amount.Quotient(), atBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to approve sell token: %s", err)
	}

	approvalResults := &AuditApprovalResults{
		Detected:          true,
		ApprovalRequested: true,
		Approved: func() bool {
			if approveReceiptTx != nil {
				return approveReceiptTx.Status == 1
			}

			return false
		}(),
		TxHash:  approveReceiptTx.TxHash,
		Receipt: approveReceiptTx != nil,
		ReceiptStatus: func() uint64 {
			if approveReceiptTx != nil {
				return approveReceiptTx.Status
			}

			return 0
		}(),
		Logs: make([]*bytecode.Log, 0),
	}

	tradeRequest := &AuditBuyOrSellResults{
		Approval: approvalResults,
	}

	if approveReceiptTx != nil {
		approvalResults.GasUsedRaw = approveReceiptTx.GasUsed
		approvalResults.GasUsed = entities.FromRawAmount(quoteToken, big.NewInt(int64(approveReceiptTx.GasUsed))).ToFixed(int32(quoteToken.Decimals()))

		if len(approveReceiptTx.Logs) > 0 {
			if decodedApprovalLog, err := bytecode.DecodeLogFromAbi(approveReceiptTx.Logs[0], []byte(tokenBinding.RawABI)); err == nil {
				tradeRequest.Approval.Logs = append(tradeRequest.Approval.Logs, decodedApprovalLog)
			}
		}
	}

	sellResults := &AuditSwapResults{
		Detected:      true,
		SwapRequested: true,
		Failure:       false,
		PairDetails:   []common.Address{baseToken.Address, quoteToken.Address},
	}

	authSell, err := spender.TransactOpts(client, amount.Quotient(), false) // Approval cannot take value as value is in the approve call
	if err != nil {
		sellResults.Failure = true
		tradeRequest.Results = sellResults
		toReturn.Trade = tradeRequest
		return toReturn, fmt.Errorf("failed to create sell transact opts: %s", err)
	}
	authSell.Value = big.NewInt(0)

	// We are using hack to pretend sending normal transaction while using simulated client...
	// Therefore, instead of passing simulatorType we pass utils.NoSimulator
	deadline := big.NewInt(time.Now().Add(time.Minute).Unix())
	_, buyReceipt, err := u.uniswapBind.Sell(authSell, network, simulatorType, client, amount.Quotient(), big.NewInt(0), sellResults.PairDetails, spender.Address, deadline)
	if err != nil {
		sellResults.Failure = true
		tradeRequest.Results = sellResults
		toReturn.Trade = tradeRequest
		return nil, fmt.Errorf("failed to sell token: %s", err)
	}

	sellResults.Detected = true
	sellResults.TxHash = buyReceipt.TxHash
	sellResults.Receipt = buyReceipt != nil
	sellResults.ReceiptStatus = func() uint64 {
		if buyReceipt != nil {
			return buyReceipt.Status
		}

		return 0
	}()

	if buyReceipt != nil {
		sellResults.GasUsedRaw = buyReceipt.GasUsed
		sellResults.GasUsed = entities.FromRawAmount(quoteToken, big.NewInt(int64(buyReceipt.GasUsed))).ToFixed(int32(quoteToken.Decimals()))

		if len(buyReceipt.Logs) > 0 {
			var buyLogs []*bytecode.Log
			for _, log := range buyReceipt.Logs {
				if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(tokenBinding.RawABI)); err == nil {
					buyLogs = append(buyLogs, decodedBuyLog)
				} else if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(uniswapBinding.RawABI)); err == nil {
					buyLogs = append(buyLogs, decodedBuyLog)
				} else if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(uniswapPairBinding.RawABI)); err == nil {
					buyLogs = append(buyLogs, decodedBuyLog)
				} else if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(wethBinding.RawABI)); err == nil {
					buyLogs = append(buyLogs, decodedBuyLog)
				}
			}

			sellResults.Logs = buyLogs

			for _, log := range buyLogs {
				if log.Type == utils.SwapLogEventType {
					if amountOut, ok := log.Data["amount1Out"]; ok {
						if amountOut, ok := amountOut.(*big.Int); ok {
							sellResults.SwapReceivedAmountRaw = amountOut
							sellResults.SwapReceivedAmount = entities.FromRawAmount(tokenOut, amountOut).ToSignificant(2)
						}
					}

					if sellResults.SwapReceivedAmountRaw == nil || sellResults.SwapReceivedAmountRaw.Uint64() == 0 {
						if amountOut, ok := log.Data["amount1Out"]; ok {
							if amountOut, ok := amountOut.(*big.Int); ok {
								sellResults.SwapReceivedAmountRaw = amountOut
								sellResults.SwapReceivedAmount = entities.FromRawAmount(tokenOut, amountOut).ToSignificant(2)
							}
						}
					}

					if amountOut, ok := log.Data["amount0In"]; ok {
						if amountOut, ok := amountOut.(*big.Int); ok {
							sellResults.ReceivedAmountRaw = amountOut
							sellResults.ReceivedAmount = entities.FromRawAmount(tokenOut, amountOut).ToSignificant(2)
						}
					}

					if sellResults.ReceivedAmountRaw == nil || sellResults.ReceivedAmountRaw.Uint64() == 0 {
						if amountOut, ok := log.Data["amount1In"]; ok {
							if amountOut, ok := amountOut.(*big.Int); ok {
								sellResults.ReceivedAmountRaw = amountOut
								sellResults.ReceivedAmount = entities.FromRawAmount(tokenOut, amountOut).ToSignificant(2)
							}
						}
					}
				}
			}
		}
	}

	if sellResults.ReceivedAmountRaw != nil && sellResults.SwapReceivedAmountRaw != nil {
		tax := CalculatePercentageDifference(amount.Quotient(), sellResults.ReceivedAmountRaw, quoteToken.Decimals())
		sellResults.TaxRaw = tax
		sellResults.Tax = entities.FromRawAmount(quoteToken, tax).ToFixed(2)
	}

	tradeRequest.Results = sellResults
	toReturn.Trade = tradeRequest

	if approvalResults.Detected || sellResults.Detected {
		toReturn.Trade.Detected = true
	}

	afterBalance, err := spender.Balance(ctx, atBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to get current balance: %s", err)
	}

	toReturn.SpenderBalanceAfter = afterBalance
	utils.DumpNodeNoExit(toReturn)
	return toReturn, nil
}

// CalculatePercentageDifference calculates the percentage difference between two big.Int values and returns a big.Int
// scaled up by the specified number of decimals.
func CalculatePercentageDifference(value1, value2 *big.Int, decimals uint) *big.Int {
	// Step 1: Find the absolute difference
	difference := new(big.Int).Sub(value1, value2)
	difference.Abs(difference)

	// Step 2: Calculate the average of the two values
	sum := new(big.Int).Add(value1, value2)
	average := new(big.Int).Div(sum, big.NewInt(2))

	// Step 3: Calculate the percentage difference
	// Multiply the difference by 10000 (to move the decimal place four places to the right)
	difference.Mul(difference, big.NewInt(10000))

	// Divide by the average
	percentageDifference := new(big.Int).Div(difference, average)

	// Scale up by 10^decimals to retain the specified number of decimal places
	scale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals-2)), nil) // Subtracting 2 because we already scaled by 100
	percentageDifference.Mul(percentageDifference, scale)

	return percentageDifference
}
