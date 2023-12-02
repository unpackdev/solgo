package simulator

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestAnvilSimulator(t *testing.T) {
	tAssert := assert.New(t)

	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	tAssert.NoError(err)
	zap.ReplaceGlobals(logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	simulator, err := CreateNewTestSimulator(ctx, t)
	require.NoError(t, err)
	require.NotNil(t, simulator)
	defer simulator.Close()

	testCases := []struct {
		name      string
		provider  utils.SimulatorType
		expectErr bool
	}{
		{
			name:      "Anvil simulator start and stop with periodic status checks",
			provider:  utils.AnvilSimulator,
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			err := simulator.Start(ctx)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for i := 0; i < 10; i++ {
				statuses, err := simulator.Status(ctx)
				if tc.expectErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					tAssert.NotNil(statuses)
				}

				anvilStatuses, found := statuses.GetNodesByType(utils.AnvilSimulator)
				tAssert.NotNil(anvilStatuses)
				tAssert.True(found)
				tAssert.Exactly(1, len(anvilStatuses))

				time.Sleep(300 * time.Millisecond)
			}

			err = simulator.Stop(ctx)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

/* func TestSimulator(t *testing.T) {
	tAssert := assert.New(t)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	etherscanApiKeys := os.Getenv("ETHERSCAN_API_KEYS")
	etherscanProvider := etherscan.NewEtherScanProvider(ctx, nil, &etherscan.Options{
		Provider: etherscan.EtherScan,
		Endpoint: "https://api.etherscan.io/api",
		Keys:     strings.Split(etherscanApiKeys, ","),
	})

	simulator, err := NewSimulator(ctx, etherscanProvider, &Options{})
	tAssert.NoError(err)
	tAssert.NotNil(simulator)

	anvilProvider, err := NewAnvilProvider(ctx, &AnvilProviderOptions{
		SimulatedClients:        10,
		MaxSimulatedConnections: 1,
		Endpoint:                "wss://ethereum.rpc.thirdweb.com:8545",
		Addr:                    "http://localhost:3030",
		SecurityKey:             "96429b17-56fd-4355-a43f-5b8a0225575d",
	})
	tAssert.NoError(err)
	tAssert.NotNil(anvilProvider)

	simulator.RegisterProvider(utils.AnvilSimulator, anvilProvider)

	for _, node := range anvilProvider.GetNodes() {
		tAssert.NotNil(node)
		tAssert.NotNil(node.GetID())
		tAssert.NotNil(node.GetHost())
		tAssert.NotNil(node.GetPort())
		tAssert.NotNil(node.GetIpcPath())
		tAssert.NotNil(node.GetAutoImpersonate())
		tAssert.NotNil(node.GetBlockNumber())
		tAssert.NotNil(node.GetAddresses())
		tAssert.NotNil(node.GetPrivateKeys())

		accounts := node.GetAccounts()
		tAssert.NotNil(accounts)

		for _, account := range accounts {
			tAssert.NotNil(account.Address)
			tAssert.NotNil(account.PrivateKey)
		}
	}


	// Alright now we have all of the nodes and pools available. What now?!

	// Following function will return client by block. In case that block is not yet ready it will spawn new anvil node
	// and wait for it to be ready. Once it's ready it will return the client.
	client, err := simulator.GetClient(ctx, utils.AnvilSimulator, big.NewInt(18686381))
	tAssert.NoError(err)
	tAssert.NotNil(client)

	// ^ This is the shit above that allows us to pretty much spoof the entire mainnet at near real-time and opens up
	// possibilities to do any type of classification that we can possibly phantom!
	// I'm just excited... Now it would be great if I knew how to simulate but heck, lets deal with this here....

	// Now lets get bindings...
	// Gonna go harder route as we want to test all the methods we can...
	bindingManager, err := bindings.NewManager(ctx, simulator.GetProvider(utils.AnvilSimulator).GetClientPool())
	tAssert.NoError(err)
	tAssert.NotNil(bindingManager)

	uniswapBind, err := bindings.NewUniswap(ctx, utils.AnvilNetwork, bindingManager, bindings.DefaultUniswapBindOptions())
	tAssert.NoError(err)
	tAssert.NotNil(uniswapBind)

	// Lets play with this address...
	// It's been at the time of writing a ghost that goes away and comes back somehow...
	// https://etherscan.io/address/0x8390a1da07e376ef7add4be859ba74fb83aa02d5
	contractAddr := common.HexToAddress("0x8390a1da07e376ef7add4be859ba74fb83aa02d5")

	tokenBind, err := bindings.NewToken(ctx, utils.AnvilNetwork, bindingManager, bindings.DefaultTokenBindOptions(contractAddr))
	tAssert.NoError(err)
	tAssert.NotNil(tokenBind)

	// Lets fetch the name of the token...
	name, err := tokenBind.GetName()
	tAssert.NoError(err)
	tAssert.Equal("GROK", name)

	// Lets fetch the symbol of the token...
	symbol, err := tokenBind.GetSymbol()
	tAssert.NoError(err)
	tAssert.Equal("GROK", symbol)

	// Lets fetch the decimals of the token...
	decimals, err := tokenBind.GetDecimals()
	tAssert.NoError(err)
	tAssert.Equal(uint8(0x9), decimals)

	// Lets fetch the total supply of the token...
	totalSupply, err := tokenBind.GetTotalSupply()
	tAssert.NoError(err)
	tAssert.Equal("6900000000000000000", totalSupply.String())

	// Lets fetch ethereum address from the uniswap contract :)
	ethAddr, err := uniswapBind.WETH()
	tAssert.NoError(err)
	tAssert.Equal(currencies.WETH.Addresses[utils.Ethereum], ethAddr)

	// Lets figure out what the pair address is...
	pairAddr, err := uniswapBind.GetPair(contractAddr, ethAddr)
	tAssert.NoError(err)
	tAssert.Equal("0x69c66BeAfB06674Db41b22CFC50c34A93b8d82a2", pairAddr.String())

	// Lets fetch the balance of the token...
	balance, err := tokenBind.BalanceOf(contractAddr)
	tAssert.NoError(err)

	// Lets fetch the balance of the pair...
	pairBalance, err := tokenBind.BalanceOf(pairAddr)
	tAssert.NoError(err)
	// Lets fetch the balance of the pair...
	burnBalance, err := tokenBind.BalanceOf(utils.ZeroAddress)
	tAssert.NoError(err)

	node, found := anvilProvider.GetNodeByBlockNumber(big.NewInt(18686381))
	tAssert.True(found)
	tAssert.NotNil(node)

	account, found := node.GetAccount(node.GetAddresses()[0])
	tAssert.True(found)
	tAssert.NotNil(account)

	// Lets fetch the balance of the pair...
	accountBalance, err := tokenBind.BalanceOf(account.Address)
	tAssert.NoError(err)

	contractPercentage := new(big.Float).Quo(new(big.Float).SetInt(balance), new(big.Float).SetInt(totalSupply))
	pairPercentage := new(big.Float).Quo(new(big.Float).SetInt(pairBalance), new(big.Float).SetInt(totalSupply))
	burnPercentage := new(big.Float).Quo(new(big.Float).SetInt(burnBalance), new(big.Float).SetInt(totalSupply))
	accountPercentage := new(big.Float).Quo(new(big.Float).SetInt(accountBalance), new(big.Float).SetInt(totalSupply))

	// Convert percentages to human-readable format
	cP, _ := contractPercentage.Mul(contractPercentage, big.NewFloat(100)).Float64()
	pP, _ := pairPercentage.Mul(pairPercentage, big.NewFloat(100)).Float64()
	bP, _ := burnPercentage.Mul(burnPercentage, big.NewFloat(100)).Float64()
	aP, _ := accountPercentage.Mul(accountPercentage, big.NewFloat(100)).Float64()
	contractPercentageStr := fmt.Sprintf("%.9f%%", cP)
	pairPercentageStr := fmt.Sprintf("%.9f%%", pP)
	burnPercentageStr := fmt.Sprintf("%.9f%%", bP)
	accountPercentageStr := fmt.Sprintf("%.9f%%", aP)


	// An example of impersonating account....
	whaleAddress := common.HexToAddress("0x8390a1DA07E376ef7aDd4Be859BA74Fb83aA02D5")
	stealAddress, err := bindingManager.ImpersonateAccount(utils.AnvilNetwork, whaleAddress)
	tAssert.NoError(err)
	tAssert.Equal(stealAddress, whaleAddress)

	// An example of stop impersonating account....
	stealAddress, err = bindingManager.StopImpersonateAccount(utils.AnvilNetwork, whaleAddress)
	tAssert.NoError(err)
	tAssert.Equal(stealAddress, whaleAddress)

	uniswapAddr, err := uniswapBind.GetAddress(bindings.UniswapV2Router)
	tAssert.NoError(err)
	tAssert.NotNil(uniswapAddr)

	tradeAmount := big.NewInt(10000000000000000)

	authApprove, err := account.TransactOpts(client, nil)
	tAssert.NoError(err)
	tAssert.NotNil(authApprove)

	approveTx, approveReceiptTx, err := tokenBind.Approve(authApprove, uniswapAddr, tradeAmount, false)
	tAssert.NoError(err)
	tAssert.NotNil(approveTx)
	tAssert.NotNil(approveReceiptTx)

	tokenBinding, _ := tokenBind.GetBinding(utils.AnvilNetwork, bindings.Erc20)
	uniswapBinding, _ := uniswapBind.GetBinding(utils.AnvilNetwork, bindings.UniswapV2Router)
	uniswapPairBinding, _ := uniswapBind.GetBinding(utils.AnvilNetwork, bindings.UniswapV2Pair)

	decodedApprovalLog, err := bytecode.DecodeLogFromAbi(approveReceiptTx.Logs[0], []byte(tokenBinding.RawABI))
	tAssert.NoError(err)
	tAssert.NotNil(decodedApprovalLog)

	swapApprove, err := account.TransactOpts(client, tradeAmount)
	tAssert.NoError(err)
	tAssert.NotNil(authApprove)

	path := []common.Address{
		currencies.WETH.Addresses[utils.Ethereum],
		contractAddr,
	}

	buyTx, buyReceipt, err := uniswapBind.Buy(
		swapApprove,
		big.NewInt(0),
		path,
		account.Address,
		big.NewInt(time.Now().Add(time.Minute).Unix()),
	)
	tAssert.NoError(err)
	tAssert.NotNil(buyTx)
	tAssert.NotNil(buyReceipt)

	var buyLogs []*bytecode.Log
	for _, log := range buyReceipt.Logs {
		if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(tokenBinding.RawABI)); err == nil {
			buyLogs = append(buyLogs, decodedBuyLog)
		} else if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(uniswapBinding.RawABI)); err == nil {
			buyLogs = append(buyLogs, decodedBuyLog)
		} else if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(uniswapPairBinding.RawABI)); err == nil {
			buyLogs = append(buyLogs, decodedBuyLog)
		}
	}

	spew.Dump("Token information", map[string]interface{}{
		"name":                         name,
		"symbol":                       symbol,
		"decimals":                     decimals,
		"totalSupply":                  totalSupply,
		"WETH":                         ethAddr,
		"contract_balance":             balance,
		"pair_balance":                 pairBalance,
		"burn_balance":                 burnBalance,
		"account_balance":              accountBalance,
		"contract_percentage":          contractPercentageStr,
		"pair_percentage":              pairPercentageStr,
		"burn_percentage":              burnPercentageStr,
		"account_percentage":           accountPercentageStr,
		"impersonated_account_address": stealAddress,
		"approve_tx_hash":              approveReceiptTx.TxHash.Hex(),
		"approve_receipt_status":       approveReceiptTx.Status,
		"approve_log":                  decodedApprovalLog,
		"buy_tx_hash":                  buyReceipt.TxHash.Hex(),
		"buy_receipt_status":           buyReceipt.Status,
		"buy_log_count":                len(buyReceipt.Logs),
		"buy_logs":                     buyLogs,
	})

	const uniswapTaxRate = 0.0003 // 0.03% Uniswap tax rate

	for _, log := range buyLogs {
		if log.Type == "swap" {
			if amount1In, ok := log.Data["amount1In"].(*big.Int); ok {
				if amount0Out, ok := log.Data["amount0Out"].(*big.Int); ok {
					// Calculate Uniswap tax
					uniswapTax := new(big.Int).Mul(amount0Out, big.NewInt(int64(uniswapTaxRate*1000000)))
					uniswapTax = uniswapTax.Div(uniswapTax, big.NewInt(1000000))

					// Calculate the difference between requested and received amounts
					amountDiff := new(big.Int).Sub(amount1In, amount0Out)

					// Calculate net received amount after tax
					netReceived := new(big.Int).Sub(amount0Out, uniswapTax)

					// Calculate percentages
					netReceivedTaxPercent := new(big.Float).Quo(new(big.Float).SetInt(uniswapTax), new(big.Float).SetInt(amountDiff)).Mul(new(big.Float), big.NewFloat(100))
					totalTaxPercent := new(big.Float).Quo(new(big.Float).SetInt(uniswapTax), new(big.Float).SetInt(amountDiff)).Mul(new(big.Float), big.NewFloat(100))

					spew.Dump(
						"Swap information",
						map[string]interface{}{
							"requested_amount": amount1In,
							"received_amount":  amount0Out,
							"uniswap_tax":      uniswapTax,
							"net_received_tax": netReceivedTaxPercent.Text('f', 2) + "%",
							"total_tax":        totalTaxPercent.Text('f', 2) + "%",
							"net_received":     netReceived,
							"uniswap_tax_rate": fmt.Sprintf("%.2f%%", uniswapTaxRate*100),
						},
					)
				}
			}
		}
	}

}
*/
