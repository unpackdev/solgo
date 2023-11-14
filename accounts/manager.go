package accounts

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/utils"
)

// Manager handles the account operations across various Ethereum networks.
// It maintains a map of keystores and accounts for each supported network.
type Manager struct {
	ctx      context.Context                // Context for managing async operations
	cfg      *Options                       // Configuration options for the Manager
	client   *clients.ClientPool            // Ethereum client pool
	ks       map[Network]*keystore.KeyStore // Keystores for different networks
	accounts map[Network][]*Account         // Accounts mapped by their network
}

// NewManager initializes a new Manager instance.
// It checks for the existence of keystore paths and supported networks,
// and loads existing accounts from the keystore.
func NewManager(ctx context.Context, client *clients.ClientPool, cfg *Options) (*Manager, error) {
	if !utils.PathExists(cfg.KeystorePath) {
		return nil, fmt.Errorf("keystore path does not exist: %s", cfg.KeystorePath)
	}

	if len(cfg.SupportedNetworks) == 0 {
		return nil, fmt.Errorf("no supported networks provided. You must provide at least one network")
	}

	var keystores = make(map[Network]*keystore.KeyStore)

	// Now for each supported network, we need to create a subdirectory in the keystore path if it does not exist.
	// Be sure that write permissions are set correctly.
	for _, network := range cfg.SupportedNetworks {
		networkPath := path.Join(cfg.KeystorePath, strings.ToLower(string(network)))
		if !utils.PathExists(networkPath) {
			if err := os.MkdirAll(networkPath, 0700); err != nil {
				return nil, err
			}
		}

		keystores[network] = keystore.NewKeyStore(
			networkPath,
			keystore.StandardScryptN,
			keystore.StandardScryptP,
		)
	}

	toReturn := &Manager{
		ctx:      ctx,
		cfg:      cfg,
		ks:       keystores,
		client:   client,
		accounts: make(map[Network][]*Account),
	}

	if err := toReturn.Load(); err != nil {
		return nil, err
	}

	return toReturn, nil
}

// Load loads accounts from the keystore for each network.
// Returns an error if it fails to load accounts for any network.
func (m *Manager) Load() error {
	for _, network := range m.cfg.SupportedNetworks {
		ks, err := m.GetKeystore(network)
		if err != nil {
			return err
		}

		for _, kAcc := range ks.Accounts() {
			path := path.Join(m.GetNetworkPath(network), kAcc.Address.Hex()+".json")
			acc, err := LoadAccount(path)
			if err != nil {
				return err
			}
			acc.ClientPool = m.client
			acc.KeyStore = ks
			m.accounts[network] = append(m.accounts[network], acc)
		}
	}

	return nil
}

// GetConfig returns the configuration options of the Manager.
func (m *Manager) GetConfig() *Options {
	return m.cfg
}

// GetNetworkPath returns the file path for a given network's keystore.
func (m *Manager) GetNetworkPath(network Network) string {
	return path.Join(m.cfg.KeystorePath, strings.ToLower(string(network)))
}

// GetKeystore retrieves the keystore for a given network.
// Returns an error if the network is not supported.
func (m *Manager) GetKeystore(network Network) (*keystore.KeyStore, error) {
	if _, ok := m.ks[network]; !ok {
		return nil, fmt.Errorf("network %s is not supported", network)
	}

	return m.ks[network], nil
}

// Import imports an account with a given private key and password into the keystore of a specified network.
// TODO: Currently, the function body is empty and needs implementation.
func (m *Manager) Import(network Network, privateKey string, password string) error {
	return nil
}

// Create creates a new account for a given network with a specified password and optional tags.
// It saves the account to the network's keystore path and adds it to the accounts map.
func (m *Manager) Create(network Network, password string, tags ...string) (*Account, error) {
	ks, err := m.GetKeystore(network)
	if err != nil {
		return nil, err
	}

	kacc, err := ks.NewAccount(password)
	if err != nil {
		return nil, err
	}

	acc := &Account{
		ClientPool:      m.client,
		KeyStore:        ks,
		KeystoreAccount: kacc,
		Password:        base64.StdEncoding.EncodeToString([]byte(password)),
		Network:         network,
		Tags:            tags,
	}

	// Now we need to save the account to the keystore path.
	path := path.Join(m.GetNetworkPath(network), kacc.Address.Hex()+".json")

	if err := acc.SaveToPath(path); err != nil {
		return nil, err
	}

	// Now we need to add the account to the accounts map.
	m.accounts[network] = append(m.accounts[network], acc)

	return acc, nil
}

// List lists all accounts for a given network, optionally filtered by tags.
func (m *Manager) List(network Network, tags ...string) []*Account {
	var toReturn []*Account

	if accounts, ok := m.accounts[network]; ok {
		if len(tags) == 0 {
			return accounts
		}

		for _, acc := range accounts {
			for _, tag := range tags {
				if acc.HasTag(tag) {
					toReturn = append(toReturn, acc)
				}
			}
		}
	}

	return toReturn
}

// Get retrieves a specific account by its address for a given network.
// Returns an error if the account is not found.
func (m *Manager) Get(network Network, address common.Address) (*Account, error) {
	if accounts, ok := m.accounts[network]; ok {
		for _, acc := range accounts {
			if acc.KeystoreAccount.Address.Hex() == address.Hex() {
				return acc, nil
			}
		}
	}

	return nil, fmt.Errorf("account not found")
}
