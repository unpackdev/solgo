package storage

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
	"github.com/unpackdev/solgo/cfg"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/detector"
	"github.com/unpackdev/solgo/providers/etherscan"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestStorage(t *testing.T) {
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
	require.NoError(t, err)
	require.NotNil(t, pool)

	etherscanApiKeys := os.Getenv("ETHERSCAN_API_KEYS")
	etherscanProvider, err := etherscan.NewProvider(ctx, nil, &etherscan.Options{
		Provider: etherscan.EtherScan,
		Endpoint: "https://api.etherscan.io/api",
		Keys:     strings.Split(etherscanApiKeys, ","),
	})
	require.NoError(t, err)

	storage, err := NewStorage(ctx, utils.Ethereum, pool, NewDefaultOptions())
	tAssert.NoError(err)
	require.NoError(t, err)

	testCases := []struct {
		name               string
		address            common.Address
		atBlock            *big.Int
		expectError        bool
		expectedSlotsCount int
		expectedSlots      map[int]*SlotDescriptor
	}{
		{
			name:               "Valid GROK (ETH) Contract: 0x8390a1da07e376ef7add4be859ba74fb83aa02d5",
			address:            common.HexToAddress("0x8390a1da07e376ef7add4be859ba74fb83aa02d5"),
			atBlock:            nil,
			expectError:        false,
			expectedSlotsCount: 24,
			expectedSlots:      map[int]*SlotDescriptor{},
		}, {
			name:               "Operation Black Rock: 0x01e99288ea767084cdabb1542aaa017425525f5b",
			address:            common.HexToAddress("0x01e99288ea767084cdabb1542aaa017425525f5b"),
			atBlock:            nil,
			expectError:        false,
			expectedSlotsCount: 25,
			expectedSlots:      map[int]*SlotDescriptor{},
		},
		{
			name:               "NonfungiblePositionManager: 0xC36442b4a4522E871399CD717aBDD847Ab11FE88",
			address:            common.HexToAddress("0xC36442b4a4522E871399CD717aBDD847Ab11FE88"),
			atBlock:            nil,
			expectError:        false,
			expectedSlotsCount: 25,
			expectedSlots:      map[int]*SlotDescriptor{},
		}, {
			name:               "UniswapV3Pool: 0xE67b950F4b84c5b06Ee36DEd6727a17443fE7493",
			address:            common.HexToAddress("0xE67b950F4b84c5b06Ee36DEd6727a17443fE7493"),
			atBlock:            nil,
			expectError:        false,
			expectedSlotsCount: 16,
			expectedSlots:      map[int]*SlotDescriptor{},
		},
		{
			name:               "SpareBytes: 0xc2F78739074b5dDCA2aDB85DE63826Cc92cE792e",
			address:            common.HexToAddress("0xc2F78739074b5dDCA2aDB85DE63826Cc92cE792e"),
			atBlock:            nil,
			expectError:        false,
			expectedSlotsCount: 37,
			expectedSlots:      map[int]*SlotDescriptor{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tAssert := assert.New(t)

			response, err := etherscanProvider.ScanContract(tc.address)
			tAssert.NoError(err)
			tAssert.NotNil(response)

			sources, err := solgo.NewSourcesFromEtherScan(response.Name, response.SourceCode)
			tAssert.NoError(err)
			tAssert.NotNil(sources)
			require.True(t, sources.HasUnits())

			parser, err := detector.NewDetectorFromSources(ctx, nil, sources)
			tAssert.NoError(err)
			tAssert.NotNil(parser)

			// So far, contracts bellow 0.6.0 are doing some weird stuff, so we are disabling it for now...
			require.False(t, utils.IsSemanticVersionLowerOrEqualTo(response.CompilerVersion, utils.SemanticVersion{Major: 0, Minor: 6, Patch: 0}))

			errs := parser.Parse()
			tAssert.Equal(len(errs), 0)

			err = parser.Build()
			tAssert.NoError(err)

			builder, err := cfg.NewBuilder(context.Background(), parser.GetIR())
			assert.NoError(t, err)
			assert.NotNil(t, builder)

			assert.NoError(t, builder.Build())

			// Use the test case data to run the test
			reader, err := storage.Describe(ctx, tc.address, parser, builder, tc.atBlock)

			if tc.expectError {
				tAssert.Error(err)
				tAssert.Nil(reader)
			} else {
				tAssert.NoError(err)
				tAssert.NotNil(reader)
			}

			require.NotNil(t, reader, "reader should not be nil")
			require.NotNil(t, reader.GetDescriptor(), "reader descriptor should not be nil")
			sortedSlots := reader.GetDescriptor().GetSortedSlots()
			tAssert.NotNil(sortedSlots)
			tAssert.Equal(tc.expectedSlotsCount, len(sortedSlots))

			storageVars, err := builder.GetStorageStateVariables()
			require.NoError(t, err)
			require.NotNil(t, storageVars)
		})
	}

}
