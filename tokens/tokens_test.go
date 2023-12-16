package tokens

import (
	"context"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/exchanges"
	"github.com/unpackdev/solgo/simulator"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestToken(t *testing.T) {
	tAssert := assert.New(t)

	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	tAssert.NoError(err)
	zap.ReplaceGlobals(logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientOptions := &clients.Options{
		Nodes: []clients.Node{
			{
				Group:                   string(utils.Ethereum),
				Type:                    "mainnet",
				Endpoint:                "https://ethereum.publicnode.com",
				NetworkId:               1,
				ConcurrentClientsNumber: 1,
			},
		},
	}

	pool, err := clients.NewClientPool(ctx, clientOptions)
	tAssert.NoError(err)
	tAssert.NotNil(pool)

	simulatorOPts := &simulator.AnvilProviderOptions{
		Network:             utils.AnvilNetwork,
		NetworkID:           utils.EthereumNetworkID,
		ClientCount:         0, // We do not want any clients as they will be fetched when necessary...
		MaxClientCount:      10,
		AutoImpersonate:     true,
		PidPath:             filepath.Join("/", "tmp", ".solgo", "/", "simulator", "/", "anvil"),
		AnvilExecutablePath: os.Getenv("SOLGO_ANVIL_PATH"),
		Fork:                true,
		ForkEndpoint:        os.Getenv("SOLGO_SIMULATOR_FORK_ENDPOINT"),
		IPAddr:              net.ParseIP("127.0.0.1"),
		StartPort:           5400,
		EndPort:             5500,
	}

	exchangeManager, err := exchanges.NewManager(ctx, pool, exchanges.DefaultOptions())
	require.NoError(t, err)
	require.NotNil(t, exchangeManager)

	bindManager, err := bindings.NewManager(ctx, pool)
	require.NoError(t, err)
	require.NotNil(t, bindManager)

	sim, err := simulator.CreateNewTestSimulator(ctx, pool, t, simulatorOPts)
	require.NoError(t, err)
	require.NotNil(t, sim)
	defer sim.Close()

	err = sim.Start(ctx)
	require.NoError(t, err)

	testCases := []struct {
		enabled       bool
		name          string
		address       common.Address
		network       utils.Network
		simulate      bool
		simulatorType utils.SimulatorType
		exchangeType  utils.ExchangeType
		atBlock       *big.Int
		amount        *big.Int
		simulations   []string
		expectError   bool
	}{
		/* 		{
		   			enabled:       true,
		   			name:          "Grok Token - No Simulation",
		   			address:       common.HexToAddress("0x8390a1da07e376ef7add4be859ba74fb83aa02d5"),
		   			network:       utils.Ethereum,
		   			simulate:      false,
		   			simulatorType: utils.NoSimulator,
		   			exchangeType:  utils.UniswapV2,
		   			amount:        big.NewInt(1),
		   			simulations:   []string{"get_pair"},
		   			expectError:   false,
		   		},
		   		{
		   			enabled:       false,
		   			name:          "Grok Token - Simulated",
		   			address:       common.HexToAddress("0x8390a1da07e376ef7add4be859ba74fb83aa02d5"),
		   			network:       utils.AnvilNetwork,
		   			simulate:      true,
		   			simulatorType: utils.AnvilSimulator,
		   			exchangeType:  utils.UniswapV2,
		   			amount:        big.NewInt(1),
		   			simulations:   []string{"get_pair", "buy_token", "sell_token"},
		   			expectError:   false,
		   		}, */
		{
			enabled:       false,
			name:          "Grok Token - Simulated",
			address:       common.HexToAddress("0x1E241521f4767853B376C2Fe795a222a07D588eE"),
			network:       utils.AnvilNetwork,
			simulate:      true,
			simulatorType: utils.AnvilSimulator,
			exchangeType:  utils.UniswapV2,
			amount:        big.NewInt(100000000000000000),
			simulations:   []string{"get_pair", "buy_token", "sell_token"},
			expectError:   false,
		},
	}
	// // 0xb777d386a9f6bf14ff85d92b27dc70209141e787

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//tAssert := assert.New(t)

			token, err := NewToken(ctx, tc.network, tc.address, tc.simulatorType, bindManager, exchangeManager, sim, pool)
			if tc.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, token)
			require.NotNil(t, token.GetContext())
			require.NotNil(t, token.GetClientPool())
			require.NotNil(t, token.GetBindManager())
			require.NotNil(t, token.GetSimulator())
			require.NotNil(t, token.GetDescriptor())
			require.NotNil(t, token.GetExchangeManager())
			require.Equal(t, tc.network, token.GetNetwork())
			require.Equal(t, tc.address, token.GetDescriptor().Address)
			require.Equal(t, tc.simulatorType, token.GetSimulatorType())

			descriptor, err := token.Unpack(ctx, tc.atBlock, tc.simulate, tc.simulatorType)
			require.NoError(t, err)
			require.NotNil(t, descriptor)

			// Modify block from this point on as we have one...
			tc.atBlock = descriptor.BlockNumber

			require.Equal(t, tc.simulate, token.IsInSimulation())

			utils.DumpNodeNoExit(descriptor)

			/* if tc.simulate {
				account, err := sim.GetFaucet().GetRandomAccount()
				require.NoError(t, err)
				require.NotNil(t, account)

				for _, simulation := range tc.simulations {
					switch simulation {
					case "buy_token":

						// We are going to test the get pair function based on the WETH9
						weth := entities.EtherOnChain(uint(utils.GetNetworkID(tc.network))).Wrapped()
						amount := utils.FromWei(tc.amount, weth)

						tradeResponse, err := token.Buy(ctx, tc.exchangeType, tc.simulatorType, account, weth, amount, tc.atBlock)
						tAssert.NoError(err)
						tAssert.NotNil(tradeResponse)

						utils.DumpNodeNoExit(tradeResponse)

					case "sell_token":
						continue
					}
				}
			} */

		})
	}
}
