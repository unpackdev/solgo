package detector

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/providers/etherscan"
	"github.com/unpackdev/solgo/utils"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/0x19/solc-switch"
	"github.com/stretchr/testify/assert"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/abi"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/audit"
	"github.com/unpackdev/solgo/ir"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestOnchainEthereumDetectorFromSources(t *testing.T) {
	fullNodeUrl := os.Getenv("FULL_NODE_TEST_URL")
	etherscanApiKeys := os.Getenv("ETHERSCAN_API_KEYS")

	if fullNodeUrl == "" || etherscanApiKeys == "" {
		t.Skip("Skipping onchain ethereum (detector) tests as keys are not properly set")
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

	solcConfig, err := solc.NewDefaultConfig()
	assert.NoError(t, err)
	assert.NotNil(t, solcConfig)

	// Preparation of solc repository. In the tests, this is required as we need to due to CI/CD permissions
	// have ability to set the releases path to the local repository.
	// @TODO: in the future investigate permissions between different go modules.
	cwd, err := os.Getwd()
	assert.NoError(t, err)
	assert.NotEmpty(t, cwd)

	releasesPath := filepath.Join(cwd, "..", "data", "solc", "releases")
	err = solcConfig.SetReleasesPath(releasesPath)
	assert.NoError(t, err)

	compiler, err := solc.New(context.Background(), solcConfig)
	assert.NoError(t, err)
	assert.NotNil(t, compiler)

	// Make sure to sync the releases...
	err = compiler.Sync()
	assert.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientOptions := &clients.Options{
		Nodes: []clients.Node{
			{
				Group:             string(utils.Ethereum),
				Type:              "mainnet",
				Endpoint:          fullNodeUrl,
				NetworkId:         1,
				ConcurrentClients: 1,
			},
		},
	}

	pool, err := clients.NewClientPool(ctx, clientOptions)
	require.NoError(t, err)
	require.NotNil(t, pool)

	etherscanProvider, err := etherscan.NewProvider(ctx, nil, &etherscan.Options{
		Provider:  etherscan.EtherScan,
		Endpoint:  "https://api.etherscan.io/api",
		RateLimit: 3,
		Keys:      strings.Split(etherscanApiKeys, ","),
	})
	require.NoError(t, err)

	// Define multiple test cases
	testCases := []struct {
		name                 string
		address              common.Address
		expectsErrors        bool
		unresolvedReferences int
	}{
		{
			name:                 "FiatTokenV2_2",
			address:              common.HexToAddress("0x43506849D7C04F9138D1A2050bbF3A0c054402dd"),
			expectsErrors:        false,
			unresolvedReferences: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			tctx, tcancel := context.WithTimeout(ctx, 10*time.Second)
			defer tcancel()

			response, err := etherscanProvider.ScanContract(tctx, testCase.address)
			assert.NoError(t, err)
			assert.NotNil(t, response)

			sources, err := solgo.NewSourcesFromEtherScan(response.Name, response.SourceCode)
			assert.NoError(t, err)
			assert.NotNil(t, sources)
			require.True(t, sources.HasUnits())

			detector, err := NewDetectorFromSources(ctx, compiler, sources)
			require.NoError(t, err)
			assert.Equal(t, ctx, detector.GetContext())
			assert.IsType(t, &Detector{}, detector)
			assert.IsType(t, &abi.Builder{}, detector.GetABI())
			assert.IsType(t, &ast.ASTBuilder{}, detector.GetAST())
			assert.IsType(t, &ir.Builder{}, detector.GetIR())
			assert.IsType(t, &solgo.Parser{}, detector.GetParser())
			assert.IsType(t, &solgo.Sources{}, detector.GetSources())
			assert.IsType(t, &solc.Solc{}, detector.GetSolc())
			assert.IsType(t, &audit.Auditor{}, detector.GetAuditor())

			syntaxErrs := detector.Parse()
			require.Equal(t, len(syntaxErrs), 0)

			err = detector.Build()
			require.NoError(t, err)

			astBuilder := detector.GetAST()
			errs := astBuilder.ResolveReferences()
			if testCase.expectsErrors {
				var errsExpected []error
				assert.Equal(t, errsExpected, errs)
			}
			assert.Equal(t, int(testCase.unresolvedReferences), astBuilder.GetResolver().GetUnprocessedCount())
			assert.Equal(t, len(astBuilder.GetResolver().GetUnprocessedNodes()), astBuilder.GetResolver().GetUnprocessedCount())

		})
	}
}
