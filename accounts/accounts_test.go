package accounts

import (
	"context"
	"math/big"
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
						Group:             string(utils.Ethereum),
						Type:              "mainnet",
						Endpoint:          "https://ethereum.publicnode.com",
						NetworkId:         1,
						ConcurrentClients: 1,
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

			require.NotNil(t, manager.GetConfig())

			found, err := manager.GetKeystore(utils.AnvilNetwork)
			require.Nil(t, err)
			require.NotNil(t, found)

			path := manager.GetNetworkPath(utils.AnvilNetwork)
			require.Equal(t, path, filepath.Join(keystorePath, "anvil"))

			err = manager.Load()
			require.NoError(t, err)

			ethereumAccounts := manager.List(utils.Ethereum)
			tAssert.NotEmpty(ethereumAccounts)

			anvilAccounts := manager.List(utils.AnvilNetwork)
			tAssert.NotEmpty(anvilAccounts)

			anvilSimulatorAccounts := manager.List(utils.AnvilNetwork, utils.SimulatorAccountType.String())
			tAssert.NotEmpty(anvilSimulatorAccounts)

			// Create new account but do not pin it as we don't want to overbloat the keystore
			newAccount, err := manager.Create(utils.Ethereum, utils.SimulatorAccountType.String(), false, "test")
			require.NoError(t, err)
			require.NotNil(t, newAccount)
			require.NotNil(t, newAccount.GetAddress())

			require.True(t, newAccount.HasTag("test"))
			require.False(t, newAccount.HasTag("test2"))

			newAccountRead, err := manager.Get(utils.Ethereum, newAccount.GetAddress())
			require.NoError(t, err)
			require.NotNil(t, newAccountRead)
			require.NotNil(t, newAccountRead.GetAddress())

			require.Equal(t, newAccount.GetAddress(), newAccountRead.GetAddress())

			currentBalance, err := newAccount.Balance(ctx, nil)
			require.NoError(t, err)
			require.NotNil(t, currentBalance)
			require.Equal(t, "0", currentBalance.String())

			client := pool.GetClientByGroup(utils.Ethereum.String())

			simulatedTopts, err := newAccount.TransactOpts(client, big.NewInt(1000000000000000000), true)
			tAssert.NoError(err)
			tAssert.NotNil(simulatedTopts)

			err = manager.Delete(utils.Ethereum, newAccount.GetAddress())
			require.NoError(t, err)

			// Create new account with pin
			newAccount, err = manager.Create(utils.Ethereum, utils.SimulatorAccountType.String(), true, "test")
			require.NoError(t, err)
			require.NotNil(t, newAccount)
			require.NotNil(t, newAccount.GetAddress())

			newAccountRead, err = manager.Get(utils.Ethereum, newAccount.GetAddress())
			tAssert.NoError(err)
			tAssert.NotNil(newAccountRead)
			tAssert.NotNil(newAccountRead.GetAddress())

			simulatedTopts, err = newAccount.TransactOpts(client, big.NewInt(1000000000000000000), false)
			tAssert.NoError(err)
			tAssert.NotNil(simulatedTopts)

			tAssert.True(newAccountRead.HasAddress(newAccount.GetAddress()))
			tAssert.Equal(newAccount.GetAddress().Hex(), newAccountRead.GetAddress().Hex())

			err = manager.Delete(utils.Ethereum, newAccount.GetAddress())
			require.NoError(t, err)

		})
	}
}
