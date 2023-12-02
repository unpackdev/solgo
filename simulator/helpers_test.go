package simulator

import (
	"context"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unpackdev/solgo/utils"
)

func CreateNewTestSimulator(ctx context.Context, t *testing.T) (*Simulator, error) {
	tAssert := assert.New(t)

	// Get the current working directory
	cwd, err := os.Getwd()
	tAssert.NoError(err)
	tAssert.NotEmpty(cwd)

	// Navigate up one level
	keystorePath := filepath.Join(filepath.Dir(cwd), "data", "faucets")

	// Establish base simulator
	// It acts as a faucet provider and manager for all the simulation providers.
	// It also provides a way to manage the simulation providers.
	simulator, err := NewSimulator(ctx, &Options{
		KeystorePath:                keystorePath,
		SupportedNetworks:           []utils.Network{utils.Ethereum, utils.AnvilNetwork},
		FaucetsEnabled:              true,
		FaucetAccountCount:          10,
		FaucetAccountDefaultBalance: new(big.Int).Mul(big.NewInt(1000), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)),
		DefaultPassword:             "wearegoingtogethacked",
	})

	require.NoError(t, err)
	tAssert.NotNil(simulator)

	anvilProvider, err := NewAnvilProvider(ctx, simulator, &AnvilProviderOptions{
		Network:             utils.AnvilNetwork,
		NetworkID:           utils.EthereumNetworkID,
		ClientCount:         1,
		MaxClientCount:      10,
		AutoImpersonate:     false,
		PidPath:             filepath.Join("/", "tmp", ".solgo", "/", "simulator", "/", "anvil"),
		AnvilExecutablePath: "/home/cortex/.cargo/bin/anvil",
		Fork:                true,
		ForkEndpoint:        os.Getenv("SOLGO_SIMULATOR_FORK_ENDPOINT"),
		IPAddr:              net.ParseIP("127.0.0.1"),
		StartPort:           5400,
		EndPort:             5500,
	})

	require.NoError(t, err)
	tAssert.NotNil(anvilProvider)

	simulator.RegisterProvider(utils.AnvilSimulator, anvilProvider)

	return simulator, nil
}
