package abi

import (
	"context"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/ir"
	"github.com/unpackdev/solgo/providers/etherscan"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestOnchainContracts(t *testing.T) {
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
				Endpoint:                "ws://localhost:8545",
				NetworkId:               1,
				ConcurrentClientsNumber: 1,
			},
		},
	}

	pool, err := clients.NewClientPool(ctx, clientOptions)
	tAssert.NoError(err)
	tAssert.NotNil(pool)

	etherscanApiKeys := os.Getenv("ETHERSCAN_API_KEYS")
	etherscanProvider := etherscan.NewEtherScanProvider(ctx, nil, &etherscan.Options{
		Provider: etherscan.EtherScan,
		Endpoint: "https://api.etherscan.io/api",
		Keys:     strings.Split(etherscanApiKeys, ","),
	})
	tAssert.NotNil(etherscanProvider)

	testCases := []struct {
		name        string
		address     common.Address
		isEmpty     bool
		atBlock     *big.Int
		expectError bool
	}{
		/* 		{
			name:        "ERC20 - Enumerable Set Test - 0xC36442b4a4522E871399CD717aBDD847Ab11FE88",
			address:     common.HexToAddress("0xC36442b4a4522E871399CD717aBDD847Ab11FE88"),
			atBlock:     nil,
			expectError: false,
		}, */
		/* 		{
			name:        "ERC20 - 0x881D40237659C251811CEC9c364ef91dC08D300C",
			address:     common.HexToAddress("0x881D40237659C251811CEC9c364ef91dC08D300C"),
			atBlock:     nil,
			expectError: false,
		}, */
		/* 		{
			name:        "ERC20 - 0xee2eBCcB7CDb34a8A822b589F9E8427C24351bfc",
			address:     common.HexToAddress("0xee2eBCcB7CDb34a8A822b589F9E8427C24351bfc"),
			atBlock:     nil,
			expectError: false,
		}, */
		/* 		{
			name:        "ERC20 - 0x772c44b5166647B135BB4836AbC4E06c28E94978",
			address:     common.HexToAddress("0x772c44b5166647B135BB4836AbC4E06c28E94978"),
			atBlock:     nil,
			expectError: false,
		}, */
		{
			name:        "ERC20 - 0x1F98431c8aD98523631AE4a59f267346ea31F984",
			address:     common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984"),
			atBlock:     nil,
			expectError: false,
		},
	}

	for _, tc := range testCases {
		//time.Sleep(1000 * time.Second)
		t.Run(tc.name, func(t *testing.T) {
			tAssert := assert.New(t)

			response, err := etherscanProvider.ScanContract(tc.address)
			tAssert.NoError(err)
			tAssert.NotNil(response)

			sources, err := solgo.NewSourcesFromEtherScan(response.Name, response.SourceCode)
			tAssert.NoError(err)
			tAssert.NotNil(sources)
			require.True(t, sources.HasUnits())

			builder, err := NewBuilderFromSources(context.TODO(), sources)
			assert.NoError(t, err)
			assert.NotNil(t, builder)

			assert.IsType(t, builder.GetAstBuilder(), &ast.ASTBuilder{})
			assert.IsType(t, builder.GetParser(), &ir.Builder{})

			// Important step which will parse the entire AST and build the IR including check for
			// reference errors and syntax errors.
			assert.Empty(t, builder.Parse())

			// Now we can get into the business of building the ABIs...
			assert.NoError(t, builder.Build())

			// Get the root node of the ABI
			root := builder.GetRoot()
			assert.NotNil(t, root)

			assert.NotNil(t, builder.GetSources())
			assert.NotNil(t, builder.GetTypeResolver())

			if !tc.isEmpty {
				assert.NotNil(t, root.GetIR())
				assert.NotNil(t, root.GetContractsAsSlice())
				assert.NotNil(t, root.GetEntryContract())
			}

			pretty, err := builder.ToJSONPretty()
			assert.NoError(t, err)
			assert.NotNil(t, pretty)

			// Check for entry contract name equivalence...
			if root.HasContracts() {
				assert.Equal(t, sources.EntrySourceUnitName, root.GetEntryName())
			}

			// Test of the ABI builder is to load generated ABI and parse it with
			// go-ethereum ABI parser and compare the results. On this way we know that
			// what users see will be literally identical and useful to wider community.
			for contractName, contract := range root.GetContracts() {
				// Assert GetContractByName() method to know that it's working...
				assert.Equal(t, contract, root.GetContractByName(contractName))

				jsonAbi, err := builder.ToJSON(contract)
				assert.NoError(t, err)
				assert.NotNil(t, jsonAbi)

				etherAbi, err := builder.ToABI(contract)
				if err != nil {
					utils.DumpNodeNoExit(contract)
				}

				require.NoError(t, err)
				require.NotNil(t, etherAbi)
			}

		})
	}

}
