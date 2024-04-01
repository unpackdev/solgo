package cfg

import (
	"context"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

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
				Group:             string(utils.Ethereum),
				Type:              "mainnet",
				Endpoint:          os.Getenv("FULL_NODE_TEST_URL"),
				NetworkId:         1,
				ConcurrentClients: 1,
			},
		},
	}

	pool, err := clients.NewClientPool(ctx, clientOptions)
	tAssert.NoError(err)
	tAssert.NotNil(pool)

	etherscanApiKeys := os.Getenv("ETHERSCAN_API_KEYS")
	etherscanProvider, err := etherscan.NewProvider(ctx, nil, &etherscan.Options{
		Provider: etherscan.EtherScan,
		Endpoint: "https://api.etherscan.io/api",
		Keys:     strings.Split(etherscanApiKeys, ","),
	})
	require.NoError(t, err)
	tAssert.NotNil(etherscanProvider)

	testCases := []struct {
		name        string
		address     common.Address
		isEmpty     bool
		atBlock     *big.Int
		length      int
		expectError bool
	}{
		{
			name:        "IgnoreFudETH - 0x8dB4beACcd1698892821a9a0Dc367792c0cB9940",
			address:     common.HexToAddress("0x8dB4beACcd1698892821a9a0Dc367792c0cB9940"),
			atBlock:     nil,
			length:      27,
			expectError: false,
		},
		{
			name:        "GROK - 0x8390a1da07e376ef7add4be859ba74fb83aa02d5",
			address:     common.HexToAddress("0x8390a1da07e376ef7add4be859ba74fb83aa02d5"),
			atBlock:     nil,
			length:      28,
			expectError: false,
		},
		{
			name:        "OperationBlackRock - 0x01e99288ea767084cdabb1542aaa017425525f5b",
			address:     common.HexToAddress("0x01e99288ea767084cdabb1542aaa017425525f5b"),
			atBlock:     nil,
			length:      26,
			expectError: false,
		},
		{
			name:        "NomiswapStableFactory - 0x818339b4E536E707f14980219037c5046b049dD4",
			address:     common.HexToAddress("0x818339b4E536E707f14980219037c5046b049dD4"),
			atBlock:     nil,
			length:      6,
			expectError: false,
		},
	}

	for _, tc := range testCases {
		//time.Sleep(1000 * time.Second)
		t.Run(tc.name, func(t *testing.T) {
			tAssert := assert.New(t)

			tctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			response, err := etherscanProvider.ScanContract(tctx, tc.address)
			tAssert.NoError(err)
			require.NotNil(t, response)

			sources, err := solgo.NewSourcesFromEtherScan(response.Name, response.SourceCode)
			tAssert.NoError(err)
			tAssert.NotNil(sources)
			require.True(t, sources.HasUnits())

			parser, err := ir.NewBuilderFromSources(context.TODO(), sources)
			if tc.expectError {
				require.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, parser)
			assert.IsType(t, &ir.Builder{}, parser)
			assert.IsType(t, &ast.ASTBuilder{}, parser.GetAstBuilder())
			assert.IsType(t, &solgo.Parser{}, parser.GetParser())
			assert.IsType(t, &solgo.Sources{}, parser.GetSources())

			// Important step which will parse the sources and build the AST including check for
			// reference errors and syntax errors.
			// If you wish to only parse the sources without checking for errors, use
			// parser.GetParser().Parse()
			assert.Empty(t, parser.Parse())

			// Now we can get into the business of building the intermediate representation
			assert.NoError(t, parser.Build())

			// Now we can get into the business of building the control flow graph
			builder, err := NewBuilder(context.Background(), parser)
			assert.NoError(t, err)
			assert.NotNil(t, builder)
			assert.IsType(t, &Builder{}, builder)

			err = builder.Build()
			assert.NoError(t, err)

			jsonBytes, err := builder.ToJSON("")
			assert.NoError(t, err)
			assert.NotNil(t, jsonBytes)

			jsonPretty, err := utils.ToJSONPretty(builder.GetGraph().GetNodes())
			assert.NoError(t, err)
			assert.NotNil(t, jsonPretty)

			mermaid := builder.ToMermaid()
			assert.NotEmpty(t, mermaid)

			vars, err := builder.GetStorageStateVariables()
			assert.NoError(t, err)
			assert.NotNil(t, vars)
			require.Equal(t, tc.length, len(vars))
		})
	}

}
