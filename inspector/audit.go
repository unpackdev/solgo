package inspector

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/bytecode"
	"github.com/unpackdev/solgo/utils"
	"github.com/unpackdev/solgo/utils/currencies"
	"go.uber.org/zap"
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
	Detected          bool             `json:"detected"`
	TransferRequested bool             `json:"transfer_requested"`
	PairDetails       []common.Address `json:"pair_details"`
	TxHash            common.Hash      `json:"transaction_hash"`
	Receipt           bool             `json:"receipt"`
	ReceiptStatus     uint64           `json:"receipt_status"`
	Logs              []*bytecode.Log  `json:"logs"`
}

type AuditBuyResults struct {
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

type AuditResults struct {
	Detected                    bool             `json:"detected"`
	FaucetAddress               common.Address   `json:"faucet_address"`
	FaucetAccountEthBalance     *big.Int         `json:"faucet_account_eth_balance"`
	FaucetAccountInitialBalance *big.Int         `json:"faucet_account_initial_balance"`
	HoneyPot                    bool             `json:"honey_pot"`
	Exchange                    *Exchange        `json:"exchange"`
	Buy                         *AuditBuyResults `json:"buy"`
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
	if m.GetDetector() != nil && m.GetDetector().GetIR() != nil && m.GetDetector().GetIR().GetRoot() != nil {
		report := m.GetReport()
		if report.HasDetector(TokenDetectorType) {
			if tokenDetector, ok := report.GetDetector(TokenDetectorType).(*TokenResults); ok && tokenDetector.Detected {
				if client := m.GetBindingManager().GetClient().GetClientByGroup(string(utils.Ethereum)); client != nil {
					latestBlock, err := client.HeaderByNumber(ctx, nil)
					if err != nil {
						zap.L().Error(
							"failed to get latest block",
							zap.Error(err),
							zap.Any("network", utils.Ethereum),
							zap.Any("address", m.GetAddress().Hex()),
						)
						return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
					}

					// Following function will return client by block. In case that block is not yet ready it will spawn new anvil node
					// and wait for it to be ready. Once it's ready it will return the client.
					client, err := m.sim.GetClient(ctx, utils.AnvilSimulator, latestBlock.Number)
					if err != nil {
						zap.L().Error(
							"failed to get simulated client",
							zap.Error(err),
							zap.Any("simulator", utils.AnvilSimulator),
							zap.Any("network", utils.AnvilNetwork),
							zap.Any("address", m.GetAddress().Hex()),
						)
						return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
					}

					bindingManager, err := bindings.NewManager(ctx, m.sim.GetProvider(utils.AnvilSimulator).GetClientPool())
					if err != nil {
						zap.L().Error(
							"failed to create binding manager",
							zap.Error(err),
							zap.Any("simulator", utils.AnvilSimulator),
							zap.Any("network", utils.AnvilNetwork),
							zap.Any("address", m.GetAddress().Hex()),
						)
						return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
					}

					uniswapBind, err := bindings.NewUniswap(ctx, utils.AnvilNetwork, bindingManager, bindings.DefaultUniswapBindOptions())
					if err != nil {
						zap.L().Error(
							"failed to create uniswap bindings",
							zap.Error(err),
							zap.Any("simulator", utils.AnvilSimulator),
							zap.Any("network", utils.AnvilNetwork),
							zap.Any("address", m.GetAddress().Hex()),
						)
						return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
					}

					tokenBind, err := bindings.NewToken(ctx, utils.AnvilNetwork, bindingManager, bindings.DefaultTokenBindOptions(m.GetAddress()))
					if err != nil {
						zap.L().Error(
							"failed to create token bindings",
							zap.Error(err),
							zap.Any("simulator", utils.AnvilSimulator),
							zap.Any("network", utils.AnvilNetwork),
							zap.Any("address", m.GetAddress().Hex()),
						)
						return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
					}

					ethAddr, err := uniswapBind.WETH()
					if err != nil {
						zap.L().Error(
							"failed to get WETH address",
							zap.Error(err),
							zap.Any("simulator", utils.AnvilSimulator),
							zap.Any("network", utils.AnvilNetwork),
							zap.Any("address", m.GetAddress().Hex()),
						)
						return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
					}

					// Lets figure out what the pair address is...
					_, err = uniswapBind.GetPair(m.GetAddress(), ethAddr)
					if err != nil {
						zap.L().Error(
							"failed to get pair address",
							zap.Error(err),
							zap.Any("simulator", utils.AnvilSimulator),
							zap.Any("network", utils.AnvilNetwork),
							zap.Any("address", m.GetAddress().Hex()),
							zap.Any("eth_address", ethAddr.Hex()),
						)
						return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
					}

					//anvilProvider := m.sim.GetProvider(utils.AnvilSimulator)

					account := m.sim.GetFaucet().List(utils.AnvilNetwork)[0]
					m.results.FaucetAddress = account.GetAddress()

					balance, err := account.Balance(ctx, nil)
					if err != nil {
						zap.L().Error(
							"failed to get faucet account balance",
							zap.Error(err),
							zap.Any("simulator", utils.AnvilSimulator),
							zap.Any("network", utils.AnvilNetwork),
							zap.Any("address", m.GetAddress().Hex()),
							zap.Any("eth_address", ethAddr.Hex()),
							zap.Any("faucet_address", account.Address.Hex()),
						)
						return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
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

					faucetInitialBalance, err := tokenBind.BalanceOf(account.Address)
					if err != nil {
						zap.L().Error(
							"failed to get faucet account balance",
							zap.Error(err),
							zap.Any("simulator", utils.AnvilSimulator),
							zap.Any("network", utils.AnvilNetwork),
							zap.Any("address", m.GetAddress().Hex()),
							zap.Any("eth_address", ethAddr.Hex()),
							zap.Any("faucet_address", account.Address.Hex()),
						)
						return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
					}
					m.results.FaucetAccountInitialBalance = faucetInitialBalance

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
						return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
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

					// TODO: Prior we can go into the transact to approve we need to know the amount to approve.

					purchaseAmount := big.NewInt(1000000000)
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
						return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
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
						//return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
					}

					buyRequest := &AuditBuyResults{
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
						tokenBinding, _ := tokenBind.GetBinding(utils.AnvilNetwork, bindings.Erc20)
						if len(approveReceiptTx.Logs) > 0 {
							if decodedApprovalLog, err := bytecode.DecodeLogFromAbi(approveReceiptTx.Logs[0], []byte(tokenBinding.RawABI)); err == nil {
								m.results.Buy.Approval.Logs = append(m.results.Buy.Approval.Logs, decodedApprovalLog)
							}
						}
					}

					buyResults := &AuditSwapResults{
						Detected:          false,
						TransferRequested: true,
						PairDetails: []common.Address{
							currencies.WETH.Addresses[utils.Ethereum],
							m.GetAddress(),
						},
					}

					transferOpts, err := account.TransactOpts(client, purchaseAmount, true)
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
						return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
					}

					deadline := big.NewInt(time.Now().Add(time.Minute).Unix())
					_, buyReceipt, err := uniswapBind.Buy(transferOpts, big.NewInt(0), buyResults.PairDetails, account.Address, deadline, true)
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
							}
						}
					}

					m.results.Buy.Results = buyResults
				}

			}
		}
	}
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *AuditDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *AuditDetector) Results() any {
	return m.results
}
