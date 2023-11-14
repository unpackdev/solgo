package accounts

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	account "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/unpackdev/solgo/clients"
)

// Account represents an Ethereum account with extended functionalities.
// It embeds ClientPool for network interactions and KeyStore for account management.
// It also includes fields for account details, network information, and additional tags.
type Account struct {
	*clients.ClientPool `json:"-" yaml:"-"` // ClientPool for Ethereum client interactions
	*keystore.KeyStore  `json:"-" yaml:"-"` // KeyStore for managing account keys
	KeystoreAccount     account.Account     `json:"account" yaml:"account"`   // Ethereum account information
	Password            string              `json:"password" yaml:"password"` // Account's password
	Network             Network             `json:"network" yaml:"network"`   // Network information
	Tags                []string            `json:"tags" yaml:"tags"`         // Arbitrary tags for the account
}

// HasTag checks if the account has a specific tag.
// Returns true if the tag exists, false otherwise.
func (a *Account) HasTag(tag string) bool {
	for _, t := range a.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

// DecodePassword decodes the base64-encoded password.
// Returns the decoded password or an error if decoding fails.
func (a *Account) DecodePassword() (string, error) {
	passwd, err := base64.StdEncoding.DecodeString(a.Password)
	return string(passwd), err
}

// GetAddress retrieves the Ethereum address of the account.
func (a *Account) GetAddress() common.Address {
	return a.KeystoreAccount.Address
}

// Balance queries the Ethereum network for the account's balance at a specific block number.
// Returns the balance or an error if the query fails.
func (a *Account) Balance(ctx context.Context, blockNum *big.Int) (*big.Int, error) {
	client := a.ClientPool.GetClientByGroup(string(a.Network))
	if client == nil {
		return big.NewInt(0), fmt.Errorf("no client found for network %s", string(a.Network))
	}

	balance, err := client.BalanceAt(ctx, a.KeystoreAccount.Address, blockNum)
	if err != nil {
		return big.NewInt(0), err
	}

	return balance, nil
}

// Transfer handles the transfer of Ethereum from the account to another address.
// Validates sufficient balance and signs the transaction with the account's passphrase.
// Returns the signed transaction or an error if the transfer fails.
func (a *Account) Transfer(ctx context.Context, to common.Address, value *big.Int) (*types.Transaction, error) {
	client := a.ClientPool.GetClientByGroup(string(a.Network))
	if client == nil {
		return nil, fmt.Errorf("no client found for network %s", string(a.Network))
	}

	currentBalance, err := a.Balance(ctx, nil)
	if err != nil {
		return nil, err
	}

	if currentBalance.Cmp(value) < 0 {
		return nil, fmt.Errorf("insufficient balance")
	}

	passwd, err := a.DecodePassword()
	if err != nil {
		return nil, fmt.Errorf("failed to decode password: %s", err.Error())
	}

	nonce, err := client.PendingNonceAt(context.Background(), a.KeystoreAccount.Address)
	if err != nil {
		return nil, err
	}

	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	var data []byte
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)

	signedTx, err := a.KeyStore.SignTxWithPassphrase(a.KeystoreAccount, passwd, tx, big.NewInt(client.GetNetworkID()))
	if err != nil {
		return nil, err
	}

	if err := client.SendTransaction(ctx, signedTx); err != nil {
		return nil, err
	}

	return signedTx, nil
}

// SaveToPath saves the account information to a specified file path in JSON format.
// Returns an error if the saving process fails.
func (a *Account) SaveToPath(path string) error {
	file, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, file, 0644)
}

// LoadAccount is a utility function to load an account from a JSON file at a given path.
// Returns the loaded Account or an error if loading fails.
func LoadAccount(path string) (*Account, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var account Account
	err = json.Unmarshal(file, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
