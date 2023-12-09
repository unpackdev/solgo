package inspector

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/bytecode"
	"github.com/unpackdev/solgo/utils"
)

type AuditApprovalResults struct {
	Detected          bool            `json:"detected"`
	ApprovalRequested bool            `json:"approval_requested"`
	Approved          bool            `json:"approved"`
	TxHash            common.Hash     `json:"transaction_hash"`
	Receipt           bool            `json:"receipt"`
	ReceiptStatus     uint64          `json:"receipt_status"`
	Logs              []*bytecode.Log `json:"logs"`
	RequestedAmount   *big.Int        `json:"requested_amount"`
}

type AuditSwapResults struct {
	Detected        bool             `json:"detected"`
	SwapRequested   bool             `json:"swap_requested"`
	PairDetails     []common.Address `json:"pair_details"`
	TxHash          common.Hash      `json:"transaction_hash"`
	Receipt         bool             `json:"receipt"`
	ReceiptStatus   uint64           `json:"receipt_status"`
	Logs            []*bytecode.Log  `json:"logs"`
	RequestedAmount *big.Int         `json:"requested_amount"`
	ReceivedAmount  *big.Int         `json:"received_amount"`
}

type AuditBuyOrSellResults struct {
	Detected bool                  `json:"detected"`
	Approval *AuditApprovalResults `json:"approval"`
	Results  *AuditSwapResults     `json:"results"`
}

type Exchange struct {
	ExchangeType utils.ExchangeType `json:"exchange_type"`
	Address      common.Address     `json:"address"`
	PairAddress  common.Address     `json:"pair_address"`
	Balance      *big.Int           `json:"balance"`
}

type AuditTaxResults struct {
	// Fields for Buy transaction tax calculation
	BuyTradeLossAmountRaw  *big.Int `json:"buy_trade_loss_amount_raw"`
	BuyTradeLossAmount     string   `json:"buy_trade_loss_amount"`
	BuyTradeLossPercentage string   `json:"buy_trade_loss_percentage"`

	// Fields for Sell transaction tax calculation
	SellTradeLossAmountRaw  *big.Int `json:"sell_trade_loss_amount_raw"`
	SellTradeLossAmount     string   `json:"sell_trade_loss_amount"`
	SellTradeLossPercentage string   `json:"sell_trade_loss_percentage"`

	// Fields for Total transaction tax calculation
	TotalTradeLossAmountRaw  *big.Int `json:"total_trade_loss_amount_raw"`
	TotalTradeLossAmount     string   `json:"total_trade_loss_amount"`
	TotalTradeLossPercentage string   `json:"total_trade_loss_percentage"`
}

type AuditResults struct {
	Detected                    bool                   `json:"detected"`
	BlockNumber                 *big.Int               `json:"block_number"`
	BlockHash                   common.Hash            `json:"block_hash"`
	FaucetAddress               common.Address         `json:"faucet_address"`
	FaucetAccountEthBalance     *big.Int               `json:"faucet_account_eth_balance"`
	FaucetAccountInitialBalance *big.Int               `json:"faucet_account_initial_balance"`
	HoneyPot                    bool                   `json:"honey_pot"`
	Exchange                    *Exchange              `json:"exchange"`
	Buy                         *AuditBuyOrSellResults `json:"buy"`
	Sell                        *AuditBuyOrSellResults `json:"sell"`
	Tax                         *AuditTaxResults       `json:"tax"`
}

type AuditDetector struct {
	ctx context.Context
	*Inspector
	results *AuditResults
}

func NewAuditDetector(ctx context.Context, inspector *Inspector) Detector {
	return &AuditDetector{
		ctx:       ctx,
		Inspector: inspector,
		results:   &AuditResults{},
	}
}

func (m *AuditDetector) Name() string {
	return "State Variable Detector"
}

func (m *AuditDetector) Type() DetectorType {
	return AuditDetectorType
}

func (m *AuditDetector) Enter(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *AuditDetector) Detect(ctx context.Context) (DetectorFn, error) {
	/* 	if m.GetDetector() != nil && m.GetDetector().GetIR() != nil && m.GetDetector().GetIR().GetRoot() != nil {
		report := m.GetReport()
		if report.HasDetector(TokenDetectorType) {
			if tokenDetector, ok := report.GetDetector(TokenDetectorType).(*TokenResults); ok && tokenDetector.Detected {
				if err := m.handleTokenDetection(ctx); err != nil {
					return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
				}
			}
		}
	} */
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *AuditDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *AuditDetector) Results() any {
	return m.results
}

func (m *AuditDetector) IsReady() bool {
	return m.GetDetector() != nil && m.GetDetector().GetIR() != nil && m.GetDetector().GetIR().GetRoot() != nil && m.GetReport().HasDetector(TokenDetectorType)
}

/* func (m *AuditDetector) GetTokenDetector() (*TokenResults, bool) {
	if m.GetReport().HasDetector(TokenDetectorType) {
		if tokenDetector, ok := m.GetReport().GetDetector(TokenDetectorType).(*TokenResults); ok {
			return tokenDetector, true
		}
	}

	return nil, false
}

func (m *AuditDetector) GetCurrency(token *TokenResults) *entities.Token {
	return entities.NewToken(
		uint(utils.GetNetworkID(m.report.Network)),
		m.GetAddress(),
		uint(token.Decimals),
		token.Symbol,
		token.Name,
	)
} */

func (m *AuditDetector) handleTokenDetection(ctx context.Context) error {
	if m.GetDetector() != nil && m.GetDetector().GetIR() != nil && m.GetDetector().GetIR().GetRoot() != nil {
		/* 		report := m.GetReport()
		   		if report.HasDetector(TokenDetectorType) {
		   			if tokenDetector, ok := report.GetDetector(TokenDetectorType).(*TokenResults); ok && tokenDetector.Detected {
		   				if client := m.GetBindingManager().GetClient().GetClientByGroup(string(utils.Ethereum)); client != nil {
		   					latestBlock, err := client.HeaderByNumber(ctx, nil)
		   					if err != nil {
		   						return fmt.Errorf("failed to get latest block for network %s", utils.Ethereum)
		   					}

		   					// Following function will return client by block. In case that block is not yet ready it will spawn new anvil node
		   					// and wait for it to be ready. Once it's ready it will return the client.
		   					client, err := m.sim.GetClient(ctx, utils.AnvilSimulator, latestBlock.Number)
		   					if err != nil {
		   						return fmt.Errorf("failed to get simulated client for network %s", utils.AnvilNetwork)
		   					}

		   					bindingManager, err := bindings.NewManager(ctx, m.sim.GetProvider(utils.AnvilSimulator).GetClientPool())
		   					if err != nil {
		   						return fmt.Errorf("failed to create binding manager for network %s", utils.AnvilNetwork)
		   					}

		   					uniswapBind, err := bindings.NewUniswap(ctx, utils.AnvilNetwork, bindingManager, bindings.DefaultUniswapBindOptions())
		   					if err != nil {
		   						return fmt.Errorf("failed to create uniswap bindings for network %s", utils.AnvilNetwork)
		   					}

		   					tokenBind, err := bindings.NewToken(ctx, utils.AnvilNetwork, bindingManager, bindings.DefaultTokenBindOptions(m.GetAddress()))
		   					if err != nil {
		   						return fmt.Errorf("failed to create token bindings for network %s", utils.AnvilNetwork)
		   					}

		   					ethAddr, err := uniswapBind.WETH()
		   					if err != nil {
		   						return fmt.Errorf("failed to get WETH address for network %s", utils.AnvilNetwork)
		   					}

		   					account := m.sim.GetFaucet().List(utils.AnvilNetwork)[0]
		   					m.results.FaucetAddress = account.GetAddress()

		   					balance, err := account.Balance(ctx, nil)
		   					if err != nil {
		   						return fmt.Errorf("failed to get faucet account balance for network %s", utils.AnvilNetwork)
		   					}
		   					m.results.FaucetAccountEthBalance = balance

		   					zap.L().Info(
		   						"Faucet account balance",
		   						zap.Any("simulator", utils.AnvilSimulator),
		   						zap.Any("network", utils.AnvilNetwork),
		   						zap.Any("contract_address", m.GetAddress().Hex()),
		   						zap.Any("eth_address", ethAddr.Hex()),
		   						zap.Any("faucet_address", account.Address.Hex()),
		   						zap.Any("balance", balance),
		   					)

		   					m.results.FaucetAccountInitialBalance, err = tokenBind.BalanceOf(account.Address)
		   					if err != nil {
		   						return fmt.Errorf("failed to get faucet account initial balance for network %s", utils.AnvilNetwork)
		   					}

		   					uniswapAddr, err := uniswapBind.GetAddress(bindings.UniswapV2Router)
		   					if err != nil {
		   						zap.L().Error(
		   							"failed to get uniswap address",
		   							zap.Error(err),
		   							zap.Any("simulator", utils.AnvilSimulator),
		   							zap.Any("network", utils.AnvilNetwork),
		   							zap.Any("address", m.GetAddress().Hex()),
		   							zap.Any("eth_address", ethAddr.Hex()),
		   							zap.Any("faucet_address", account.Address.Hex()),
		   						)
		   						return err
		   					}

		   					m.results.Exchange = &Exchange{
		   						ExchangeType: utils.UniswapV2,
		   						Address:      uniswapAddr,
		   					}

		   					// Lets figure out what the pair address is...
		   					pairAddr, err := uniswapBind.GetPair(m.GetAddress(), ethAddr)
		   					if err != nil {
		   						zap.L().Error(
		   							"failed to get pair address",
		   							zap.Error(err),
		   							zap.Any("simulator", utils.AnvilSimulator),
		   							zap.Any("network", utils.AnvilNetwork),
		   							zap.Any("address", m.GetAddress().Hex()),
		   							zap.Any("eth_address", ethAddr.Hex()),
		   							zap.Any("faucet_address", account.Address.Hex()),
		   						)
		   					} else {
		   						m.results.Exchange.PairAddress = pairAddr

		   						// Getting the balance from the pair to be able calculate later on taxes and what not...
		   						balance, err := tokenBind.BalanceOf(pairAddr)
		   						if err != nil {
		   							zap.L().Error(
		   								"failed to get pair token balance",
		   								zap.Error(err),
		   								zap.Any("simulator", utils.AnvilSimulator),
		   								zap.Any("network", utils.AnvilNetwork),
		   								zap.Any("contract_address", m.GetAddress().Hex()),
		   								zap.Any("pair_address", pairAddr.Hex()),
		   							)
		   						} else {
		   							m.results.Exchange.Balance = balance
		   						}
		   					}

		   					purchaseAmount := big.NewInt(1000000000)

		   					tokenBinding, _ := tokenBind.GetBinding(utils.AnvilNetwork, bindings.Erc20)
		   					uniswapBinding, _ := uniswapBind.GetBinding(utils.AnvilNetwork, bindings.UniswapV2Router)
		   					uniswapPairBinding, _ := uniswapBind.GetBinding(utils.AnvilNetwork, bindings.UniswapV2Pair)

		   					authApprove, err := account.TransactOpts(client, nil, false) // Approval cannot take value as value is in the approve call
		   					if err != nil {
		   						zap.L().Error(
		   							"failed to create transaction options",
		   							zap.Error(err),
		   							zap.Any("simulator", utils.AnvilSimulator),
		   							zap.Any("network", utils.AnvilNetwork),
		   							zap.Any("address", m.GetAddress().Hex()),
		   							zap.Any("eth_address", ethAddr.Hex()),
		   							zap.Any("faucet_address", account.Address.Hex()),
		   							zap.Any("purchase_amount", purchaseAmount),
		   						)
		   						return err
		   					}

		   					_, approveReceiptTx, err := tokenBind.Approve(authApprove, uniswapAddr, purchaseAmount, false)
		   					if err != nil {
		   						zap.L().Error(
		   							"failed to approve tokens",
		   							zap.Error(err),
		   							zap.Any("simulator", utils.AnvilSimulator),
		   							zap.Any("network", utils.AnvilNetwork),
		   							zap.Any("address", m.GetAddress().Hex()),
		   							zap.Any("eth_address", ethAddr.Hex()),
		   							zap.Any("faucet_address", account.Address.Hex()),
		   							zap.Any("purchase_amount", purchaseAmount),
		   						)
		   						//return err
		   					}

		   					buyRequest := &AuditBuyOrSellResults{
		   						Approval: &AuditApprovalResults{
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
		   							Logs:            make([]*bytecode.Log, 0),
		   							RequestedAmount: purchaseAmount,
		   						},
		   					}

		   					m.results.Buy = buyRequest

		   					if approveReceiptTx != nil {
		   						if len(approveReceiptTx.Logs) > 0 {
		   							if decodedApprovalLog, err := bytecode.DecodeLogFromAbi(approveReceiptTx.Logs[0], []byte(tokenBinding.RawABI)); err == nil {
		   								m.results.Buy.Approval.Logs = append(m.results.Buy.Approval.Logs, decodedApprovalLog)
		   							}
		   						}
		   					}

		   					buyResults := &AuditSwapResults{
		   						Detected:      false,
		   						SwapRequested: true,
		   						PairDetails: []common.Address{
		   							currencies.WETH.Addresses[utils.Ethereum],
		   							m.GetAddress(),
		   						},
		   					}

		   					buyOpts, err := account.TransactOpts(client, purchaseAmount, false)
		   					if err != nil {
		   						zap.L().Error(
		   							"failed to create transfer transact information",
		   							zap.Error(err),
		   							zap.Any("simulator", utils.AnvilSimulator),
		   							zap.Any("network", utils.AnvilNetwork),
		   							zap.Any("address", m.GetAddress().Hex()),
		   							zap.Any("eth_address", ethAddr.Hex()),
		   							zap.Any("faucet_address", account.Address.Hex()),
		   							zap.Any("purchase_amount", purchaseAmount),
		   						)
		   						return err
		   					}

		   					deadline := big.NewInt(time.Now().Add(time.Minute).Unix())
		   					_, buyReceipt, err := uniswapBind.Buy(buyOpts, big.NewInt(0), buyResults.PairDetails, account.Address, deadline, false)
		   					if err != nil {
		   						zap.L().Error(
		   							"failed to transfer tokens",
		   							zap.Error(err),
		   							zap.Any("simulator", utils.AnvilSimulator),
		   							zap.Any("network", utils.AnvilNetwork),
		   							zap.Any("address", m.GetAddress().Hex()),
		   							zap.Any("eth_address", ethAddr.Hex()),
		   							zap.Any("faucet_address", account.Address.Hex()),
		   							zap.Any("purchase_amount", purchaseAmount),
		   						)
		   					} else {
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
		   							if len(buyReceipt.Logs) > 0 {
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

		   								buyResults.Logs = buyLogs

		   								for _, log := range buyLogs {
		   									if log.Type == "swap" {
		   										if amountOut, ok := log.Data["amount0Out"]; ok {
		   											if amountOut, ok := amountOut.(*big.Int); ok {
		   												buyResults.ReceivedAmount = amountOut
		   											}
		   										}

		   										if amountIn, ok := log.Data["amount1In"]; ok {
		   											if amountIn, ok := amountIn.(*big.Int); ok {
		   												buyResults.RequestedAmount = amountIn
		   											}
		   										}
		   									}
		   								}
		   							}
		   						}
		   					}

		   					m.results.Buy.Results = buyResults */

		/* 					// SELL RESULTS

		   					if buyResults.Detected && buyResults.Receipt && buyResults.ReceivedAmount != nil {
		   						sellApprove, err := account.TransactOpts(client, nil, false) // Approval cannot take value as value is in the approve call
		   						if err != nil {
		   							zap.L().Error(
		   								"failed to create sell transaction options",
		   								zap.Error(err),
		   								zap.Any("simulator", utils.AnvilSimulator),
		   								zap.Any("network", utils.AnvilNetwork),
		   								zap.Any("address", m.GetAddress().Hex()),
		   								zap.Any("eth_address", ethAddr.Hex()),
		   								zap.Any("faucet_address", account.Address.Hex()),
		   								zap.Any("sell_amount", buyResults.ReceivedAmount),
		   							)
		   							return err
		   						}

		   						_, approveReceiptTx, err := tokenBind.Approve(sellApprove, uniswapAddr, buyResults.ReceivedAmount, false)
		   						if err != nil {
		   							zap.L().Error(
		   								"failed to approve sell tokens",
		   								zap.Error(err),
		   								zap.Any("simulator", utils.AnvilSimulator),
		   								zap.Any("network", utils.AnvilNetwork),
		   								zap.Any("address", m.GetAddress().Hex()),
		   								zap.Any("eth_address", ethAddr.Hex()),
		   								zap.Any("faucet_address", account.Address.Hex()),
		   								zap.Any("sell_amount", buyResults.ReceivedAmount),
		   							)
		   						}

		   						sellRequest := &AuditBuyOrSellResults{
		   							Approval: &AuditApprovalResults{
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
		   								Logs:            make([]*bytecode.Log, 0),
		   								RequestedAmount: buyResults.ReceivedAmount,
		   							},
		   						}

		   						m.results.Sell = sellRequest

		   						if approveReceiptTx != nil {
		   							if len(approveReceiptTx.Logs) > 0 {
		   								if decodedApprovalLog, err := bytecode.DecodeLogFromAbi(approveReceiptTx.Logs[0], []byte(tokenBinding.RawABI)); err == nil {
		   									m.results.Buy.Approval.Logs = append(m.results.Buy.Approval.Logs, decodedApprovalLog)
		   								}
		   							}
		   						}

		   						sellResults := &AuditSwapResults{
		   							Detected:      false,
		   							SwapRequested: true,
		   							PairDetails: []common.Address{
		   								m.GetAddress(),
		   								currencies.WETH.Addresses[utils.Ethereum],
		   							},
		   						}

		   						sellOpts, err := account.TransactOpts(client, nil, false)
		   						if err != nil {
		   							zap.L().Error(
		   								"failed to create transfer transact information",
		   								zap.Error(err),
		   								zap.Any("simulator", utils.AnvilSimulator),
		   								zap.Any("network", utils.AnvilNetwork),
		   								zap.Any("address", m.GetAddress().Hex()),
		   								zap.Any("eth_address", ethAddr.Hex()),
		   								zap.Any("faucet_address", account.Address.Hex()),
		   								zap.Any("sell_amount", buyResults.ReceivedAmount),
		   							)
		   							return err
		   						}

		   						deadline := big.NewInt(time.Now().Add(time.Minute).Unix())
		   						_, sellReceipt, err := uniswapBind.Sell(sellOpts, buyResults.ReceivedAmount, big.NewInt(0), sellResults.PairDetails, account.Address, deadline, false)
		   						if err != nil {
		   							zap.L().Error(
		   								"failed to sell tokens",
		   								zap.Error(err),
		   								zap.Any("simulator", utils.AnvilSimulator),
		   								zap.Any("network", utils.AnvilNetwork),
		   								zap.Any("address", m.GetAddress().Hex()),
		   								zap.Any("eth_address", ethAddr.Hex()),
		   								zap.Any("faucet_address", account.Address.Hex()),
		   								zap.Any("purchase_amount", purchaseAmount),
		   							)
		   						} else {
		   							sellResults.Detected = true
		   							sellResults.TxHash = sellReceipt.TxHash
		   							sellResults.Receipt = sellReceipt != nil
		   							sellResults.ReceiptStatus = func() uint64 {
		   								if sellReceipt != nil {
		   									return sellReceipt.Status
		   								}

		   								return 0
		   							}()

		   							if sellReceipt != nil {
		   								if len(sellReceipt.Logs) > 0 {
		   									var buyLogs []*bytecode.Log
		   									for _, log := range sellReceipt.Logs {
		   										if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(tokenBinding.RawABI)); err == nil {
		   											buyLogs = append(buyLogs, decodedBuyLog)
		   										} else if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(uniswapBinding.RawABI)); err == nil {
		   											buyLogs = append(buyLogs, decodedBuyLog)
		   										} else if decodedBuyLog, err := bytecode.DecodeLogFromAbi(log, []byte(uniswapPairBinding.RawABI)); err == nil {
		   											buyLogs = append(buyLogs, decodedBuyLog)
		   										}
		   									}

		   									sellResults.Logs = buyLogs

		   									for _, log := range buyLogs {
		   										if log.Type == "swap" {
		   											spew.Dump(log.Data)
		   											if amountOut, ok := log.Data["amount1Out"]; ok {
		   												if amountOut, ok := amountOut.(*big.Int); ok {
		   													sellResults.ReceivedAmount = amountOut
		   												}
		   											}

		   											if amountIn, ok := log.Data["amount0In"]; ok {
		   												if amountIn, ok := amountIn.(*big.Int); ok {
		   													sellResults.RequestedAmount = amountIn
		   												}
		   											}
		   										}
		   									}
		   								}
		   							}
		   						}

		   						m.results.Sell.Results = sellResults
		   					}

		   					// TAX CALCULATION
		   					m.calculateTaxes(ctx) */

		/* 				}

			}
		} */
	}

	return nil
}

/* func (m *AuditDetector) calculateTaxes(ctx context.Context) error {
	tax := &AuditTaxResults{}

	if m.results.Buy.Results.Receipt && m.results.Sell.Results.Receipt {
		// Corrected calculation for raw trade loss amount
		tax.TotalTradeLossAmountRaw = big.NewInt(0).Sub(m.results.Buy.Results.RequestedAmount, m.results.Sell.Results.ReceivedAmount)
		tax.TotalTradeLossAmount = entities.FromRawAmount(entities.WETH9[1], tax.TotalTradeLossAmountRaw).ToFixed(18)

		// Calculate trade loss percentage
		if m.results.Buy.Results.RequestedAmount.Cmp(big.NewInt(0)) != 0 { // Avoid division by zero
			lossPercentage := new(big.Float).Quo(
				new(big.Float).SetInt(tax.TotalTradeLossAmountRaw),
				new(big.Float).SetInt(m.results.Buy.Results.RequestedAmount),
			)
			lossPercentage = lossPercentage.Mul(lossPercentage, big.NewFloat(100)) // Convert to percentage
			tax.TotalTradeLossPercentage = lossPercentage.Text('f', 2)             // Format as string with 2 decimal places
		} else {
			tax.TotalTradeLossPercentage = "0.00" // Default to 0.00% if buy amount is zero
		}
	}

	m.results.Tax = tax
	return nil
} */
