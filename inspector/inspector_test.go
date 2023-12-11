package inspector

import (
	"context"
	"net"
	"os"
	"path/filepath"
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
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestInspector(t *testing.T) {
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
				Endpoint:                "https://eth-mainnet.g.alchemy.com/v2/Dcctb0Q9tu7V_4FgggehOvoJK4lT1ppG",
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

	simulatorOPts := &simulator.AnvilProviderOptions{
		Network:             utils.AnvilNetwork,
		NetworkID:           utils.EthereumNetworkID,
		ClientCount:         0, // We do not want any clients as they will be fetched when necessary...
		MaxClientCount:      10,
		AutoImpersonate:     true,
		PidPath:             filepath.Join("/", "tmp", ".solgo", "/", "simulator", "/", "anvil"),
		AnvilExecutablePath: "/home/cortex/.cargo/bin/anvil",
		Fork:                true,
		ForkEndpoint:        os.Getenv("SOLGO_SIMULATOR_FORK_ENDPOINT"),
		IPAddr:              net.ParseIP("127.0.0.1"),
		StartPort:           5400,
		EndPort:             5500,
	}

	sim, err := simulator.CreateNewTestSimulator(ctx, pool, t, simulatorOPts)
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
		network       utils.Network
		expectedError bool
	}{
		{
			name: "GROK Token",
			//contractAddr:  common.HexToAddress("0x0c65b5d43f2c2252897ce04d86d7fa46b83ed514"),
			contractAddr:  common.HexToAddress("0x8390a1da07e376ef7add4be859ba74fb83aa02d5"),
			network:       utils.Ethereum,
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

				inspector, err := NewInspector(ctx, tc.network, parser, sim, storage, bindManager, tc.contractAddr, nil)
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

				toJson, _ := utils.ToJSONPretty(inspector.GetReport())
				utils.WriteToFile("report.json", toJson)

				utils.DumpNodeNoExit(inspector.GetReport().Detectors[AuditDetectorType])

				/* 				md := GetDetector(MintDetectorType)
				   				require.NotNil(t, md)

				   				mdd := ToDetector[*MintDetector](md)
				   				require.NotNil(t, mdd)
				   				utils.DumpNodeNoExit(mdd.GetFunctionNames()) */

			}
		})
	}
}
