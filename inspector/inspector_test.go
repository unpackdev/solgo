package inspector

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/detector"
	"github.com/unpackdev/solgo/providers/etherscan"
	"github.com/unpackdev/solgo/simulator"
	"github.com/unpackdev/solgo/standards"
	"github.com/unpackdev/solgo/storage"
	"github.com/unpackdev/solgo/utils"
)

func TestInspector(t *testing.T) {
	tAssert := assert.New(t)

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

	etherscanApiKeys := os.Getenv("ETHERSCAN_API_KEYS")
	etherscanProvider := etherscan.NewEtherScanProvider(ctx, nil, &etherscan.Options{
		Provider: etherscan.EtherScan,
		Endpoint: "https://api.etherscan.io/api",
		Keys:     strings.Split(etherscanApiKeys, ","),
	})

	storage, err := storage.NewStorage(ctx, utils.Ethereum, pool, nil, storage.NewDefaultOptions())
	tAssert.NoError(err)
	tAssert.NotNil(storage)

	bindManager, err := bindings.NewManager(ctx, pool)
	tAssert.NoError(err)
	tAssert.NotNil(bindManager)

	sim, err := simulator.CreateNewTestSimulator(ctx, t)
	require.NoError(t, err)
	require.NotNil(t, sim)
	defer sim.Close()

	err = sim.Start(ctx)
	require.NoError(t, err)

	defer func() {
		err := sim.Stop(ctx)
		require.NoError(t, err)
	}()

	testCases := []struct {
		name          string
		contractAddr  common.Address
		expectedError bool
	}{
		{
			name:          "HoneyPot 1",
			contractAddr:  common.HexToAddress("0x0c65b5d43f2c2252897ce04d86d7fa46b83ed514"),
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response, err := etherscanProvider.ScanContract(tc.contractAddr)
			if tc.expectedError {
				tAssert.Error(err)
			} else {
				tAssert.NoError(err)
				tAssert.NotNil(response)

				sources, err := solgo.NewSourcesFromEtherScan(response.Name, response.SourceCode)
				tAssert.NoError(err)
				tAssert.NotNil(sources)
				require.True(t, sources.HasUnits())

				parser, err := detector.NewDetectorFromSources(ctx, nil, sources)
				tAssert.NoError(err)
				tAssert.NotNil(parser)

				// So far contracts bellow 0.6.0 are doing some weird shit so we are disabling it for now...
				require.False(t, utils.IsSemanticVersionLowerOrEqualTo(response.CompilerVersion, utils.SemanticVersion{Major: 0, Minor: 6, Patch: 0}))

				parser.Parse()

				err = parser.Build()
				tAssert.NoError(err)

				inspector, err := NewInspector(ctx, parser, sim, storage, bindManager, tc.contractAddr)
				tAssert.NoError(err)
				tAssert.NotNil(inspector)

				// Register default detectors...
				inspector.RegisterDetectors()

				// If contract does not have any source code available we don't want to check it here.
				// In that case we will in the future go towards the opcodes...
				require.True(t, inspector.IsReady())

				// First we don't want to do any type of inspections if contract is not ERC20
				require.True(t, inspector.HasStandard(standards.ERC20))

				// Now we are going to do use transfers check as we don't want to continue with this contract if there are
				// no transfers involved...
				require.True(t, inspector.UsesTransfers())

				// Alright now we're at the point that we know contract should be checked for any type of malicious activity
				tAssert.NoError(inspector.Inspect())

				utils.DumpNodeNoExit(inspector.GetReport().Detectors[StandardsDetectorType])

				/* 				md := GetDetector(MintDetectorType)
				   				require.NotNil(t, md)

				   				mdd := ToDetector[*MintDetector](md)
				   				require.NotNil(t, mdd)
				   				utils.DumpNodeNoExit(mdd.GetFunctionNames()) */

			}
		})
	}
}
