package accounts

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/utils"
)

// Manager handles the account operations across various Ethereum networks.
// It maintains a map of keystores and accounts for each supported network.
type Manager struct {
	ctx      context.Context                      // Context for managing async operations
	cfg      *Options                             // Configuration options for the Manager
	client   *clients.ClientPool                  // Ethereum client pool
	ks       map[utils.Network]*keystore.KeyStore // Keystores for different networks
	accounts map[utils.Network][]*Account         // Accounts mapped by their network
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

	var keystores = make(map[utils.Network]*keystore.KeyStore)

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
		accounts: make(map[utils.Network][]*Account),
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

			if client := m.client.GetClientByGroup(network.String()); client != nil {
				acc.SetClient(client)
			}

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
func (m *Manager) GetNetworkPath(network utils.Network) string {
	return path.Join(m.cfg.KeystorePath, strings.ToLower(string(network)))
}

// GetKeystore retrieves the keystore for a given network.
// Returns an error if the network is not supported.
func (m *Manager) GetKeystore(network utils.Network) (*keystore.KeyStore, error) {
	if _, ok := m.ks[network]; !ok {
		return nil, fmt.Errorf("network %s is not supported", network)
	}

	return m.ks[network], nil
}

// Create creates a new account for a given network with a specified password and optional tags.
// It saves the account to the network's keystore path and adds it to the accounts map.
func (m *Manager) Create(network utils.Network, password string, pin bool, tags ...string) (*Account, error) {
	ks, err := m.GetKeystore(network)
	if err != nil {
		return nil, err
	}

	// In case that we should not pin the account, we need to generate a new private key and avoid saving it to the keystore.
	// This is basically a one time use account.
	if !pin {
		// Generate a new private key
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			return nil, fmt.Errorf("failed to generate private key: %s", err)
		}

		// Obtain the public key from the private key
		publicKey := privateKey.Public()

		// Cast the public key to *ecdsa.PublicKey
		ecdsaPublicKey, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("failed to cast public key to *ecdsa.PublicKey")
		}

		// Obtain the address from the public key
		address := crypto.PubkeyToAddress(*ecdsaPublicKey)

		acc := &Account{
			KeyStore:   ks,
			Address:    address,
			Type:       utils.SimpleAccountType,
			PrivateKey: fmt.Sprintf("%x", privateKey.D),
			PublicKey:  fmt.Sprintf("%x", crypto.FromECDSAPub(ecdsaPublicKey)),
			Password:   base64.StdEncoding.EncodeToString([]byte(password)),
			Network:    network,
			Tags:       tags,
		}

		// Now we need to add the account to the accounts map.
		m.accounts[network] = append(m.accounts[network], acc)

		acc.SetClient(m.client.GetClientByGroup(network.String()))
		return acc, nil
	}

	kacc, err := ks.NewAccount(password)
	if err != nil {
		return nil, err
	}

	acc := &Account{
		KeyStore:        ks,
		Address:         kacc.Address,
		Type:            utils.KeystoreAccountType,
		KeystoreAccount: kacc,
		Password:        base64.StdEncoding.EncodeToString([]byte(password)),
		Network:         network,
		Tags:            tags,
	}

	acc.SetClient(m.client.GetClientByGroup(network.String()))

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
func (m *Manager) List(network utils.Network, tags ...string) []*Account {
	var toReturn []*Account

	if accounts, ok := m.accounts[network]; ok {
		if len(tags) == 0 {
			return accounts
		}

		for _, acc := range accounts {
			for _, tag := range tags {
				if len(tags) > 0 {
					if acc.HasTag(tag) {
						toReturn = append(toReturn, acc)
					}
				} else {
					toReturn = append(toReturn, acc)
				}
			}
		}
	}

	return toReturn
}

// Get retrieves a specific account by its address for a given network.
// Returns an error if the account is not found.
func (m *Manager) Get(network utils.Network, address common.Address) (*Account, error) {
	if accounts, ok := m.accounts[network]; ok {
		for _, acc := range accounts {
			if acc.Address.Hex() == address.Hex() {
				return acc, nil
			}
		}
	}

	return nil, fmt.Errorf("account for network: %s not found: %s", network.String(), address.Hex())
}

// GetRandomAccount returns a random account from all the available accounts across networks.
// Returns an error if there are no accounts available.
func (m *Manager) GetRandomAccount() (*Account, error) {
	var allAccounts []*Account

	// Collect all accounts from all networks
	for _, accs := range m.accounts {
		allAccounts = append(allAccounts, accs...)
	}

	if len(allAccounts) == 0 {
		return nil, fmt.Errorf("no accounts available")
	}

	// Seed the random number generator to ensure different results on each call
	rand.Seed(time.Now().UnixNano())

	// Select a random account
	randomIndex := rand.Intn(len(allAccounts))
	return allAccounts[randomIndex], nil
}

// Delete removes an account for a given network by its address.
// It deletes the account from both the keystore and the Manager's accounts map.
// Returns an error if the account is not found or if there's an issue with deletion.
func (m *Manager) Delete(network utils.Network, address common.Address) error {
	// Check if the account exists in the Manager's accounts map.
	if accounts, ok := m.accounts[network]; ok {
		for i, acc := range accounts {
			if acc.Address == address {
				if acc.PrivateKey == "" && acc.PublicKey == "" {
					password, err := acc.DecodePassword()
					if err != nil {
						return fmt.Errorf("failed to decode password: %s for account: %s", err, address.Hex())
					}

					// Delete the account from the keystore, if applicable.
					ks, err := m.GetKeystore(network)
					if err != nil {
						return err
					}

					err = ks.Delete(acc.KeystoreAccount, password)
					if err != nil {
						return fmt.Errorf("failed to delete account from keystore: %s", err)
					}
				}

				// Remove the account from the Manager's accounts map.
				// Using append to remove the element at index i from accounts.
				m.accounts[network] = append(accounts[:i], accounts[i+1:]...)

				return nil
			}
		}
	}

	return fmt.Errorf("account for network: %s not found: %s", network.String(), address.Hex())
}
