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
	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/providers/etherscan"
	"github.com/unpackdev/solgo/simulator"
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

	simulator, err := simulator.NewSimulator(ctx, pool, nil)
	tAssert.NoError(err)
	tAssert.NotNil(simulator)

	etherscanApiKeys := os.Getenv("ETHERSCAN_API_KEYS")
	etherscanProvider := etherscan.NewEtherScanProvider(ctx, nil, &etherscan.Options{
		Provider: etherscan.EtherScan,
		Endpoint: "https://api.etherscan.io/api",
		Keys:     strings.Split(etherscanApiKeys, ","),
	})

	bindManager, err := bindings.NewManager(ctx, pool)
	tAssert.NoError(err)
	tAssert.NotNil(bindManager)

	storage, err := NewStorage(ctx, utils.Ethereum, pool, simulator, etherscanProvider, nil, bindManager, NewDefaultOptions())
	tAssert.NoError(err)
	tAssert.NotNil(storage)

	testCases := []struct {
		name               string
		address            common.Address
		atBlock            *big.Int
		expectError        bool
		expectedSlotsCount int
		expectedSlots      map[int]*SlotDescriptor
	}{
		/* 		{
			name:               "Valid GROK (ETH) Contract: 0x8390a1da07e376ef7add4be859ba74fb83aa02d5",
			address:            common.HexToAddress("0x8390a1da07e376ef7add4be859ba74fb83aa02d5"),
			atBlock:            nil,
			expectError:        false,
			expectedSlotsCount: 24,
			expectedSlots:      map[int]*SlotDescriptor{},
		}, */
		{
			name:               "Operation Black Rock: 0x01e99288ea767084cdabb1542aaa017425525f5b",
			address:            common.HexToAddress("0x01e99288ea767084cdabb1542aaa017425525f5b"),
			atBlock:            nil,
			expectError:        false,
			expectedSlotsCount: 20,
			expectedSlots: map[int]*SlotDescriptor{
				0: {},
			},
		},
		/* 		{
			name:               "Q*: 0x9abfc0f085c82ec1be31d30843965fcc63053ffe",
			address:            common.HexToAddress("0x9abfc0f085c82ec1be31d30843965fcc63053ffe"),
			atBlock:            nil,
			expectError:        false,
			expectedSlotsCount: 24,
			expectedSlots:      map[int]*SlotDescriptor{},
		}, */
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tAssert := assert.New(t)

			// Use the test case data to run the test
			reader, err := storage.Describe(ctx, tc.address, tc.atBlock)

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

			for _, slot := range sortedSlots {
				utils.DumpNodeNoExit(slot)
			}

			/* 			for i, slot := range sortedSlots {
				require.NotNil(t, tc.expectedSlots[i])
				tAssert.Equal(tc.expectedSlots[i].Slot, slot.Slot)
				tAssert.Equal(tc.expectedSlots[i].Offset, slot.Offset)
				tAssert.Equal(tc.expectedSlots[i].Type, slot.Type)
				tAssert.Equal(tc.expectedSlots[i].Name, slot.Name)
				tAssert.Equal(tc.expectedSlots[i].Value, slot.Value)
				tAssert.Equal(tc.expectedSlots[i].Size, slot.Size)
				tAssert.Equal(tc.expectedSlots[i].DeclarationLine, slot.DeclarationLine)
			} */

		})
	}

}
