package accounts

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/utils"
)

func TestAccountsManager(t *testing.T) {
	tAssert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get the current working directory
	cwd, err := os.Getwd()
	tAssert.NoError(err)
	tAssert.NotEmpty(cwd)

	// Navigate up one level
	keystorePath := filepath.Join(filepath.Dir(cwd), "data", "faucets")

	testCases := []struct {
		name              string
		clientOptions     *clients.Options
		keystorePath      string
		supportedNetworks []utils.Network
	}{
		{
			name: "Test with Ethereum And Anvil - Mainnet",
			clientOptions: &clients.Options{
				Nodes: []clients.Node{
					{
						Group:                   string(utils.Ethereum),
						Type:                    "mainnet",
						Endpoint:                "https://ethereum.publicnode.com",
						NetworkId:               1,
						ConcurrentClientsNumber: 1,
					},
				},
			},
			keystorePath:      keystorePath,
			supportedNetworks: []utils.Network{utils.Ethereum, utils.AnvilNetwork},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tAssert := assert.New(t)

			pool, err := clients.NewClientPool(ctx, tc.clientOptions)
			tAssert.NoError(err)
			tAssert.NotNil(pool)

			manager, err := NewManager(ctx, pool, &Options{KeystorePath: tc.keystorePath, SupportedNetworks: tc.supportedNetworks})
			require.NoError(t, err)
			tAssert.NotNil(manager)

			err = manager.Load()
			require.NoError(t, err)

			ethereumAccounts := manager.List(utils.Ethereum)
			tAssert.NotEmpty(ethereumAccounts)

			anvilAccounts := manager.List(utils.AnvilNetwork)
			tAssert.NotEmpty(anvilAccounts)

			anvilSimulatorAccounts := manager.List(utils.AnvilNetwork, utils.SimulatorAccountType.String())
			tAssert.NotEmpty(anvilSimulatorAccounts)

		})
	}
}
