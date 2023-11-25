package simulator

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Faucet struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
	Balance    *big.Int
}

func NewFaucet() (*Faucet, error) {
	sk, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	address := crypto.PubkeyToAddress(sk.PublicKey)
	return &Faucet{
		Address:    address,
		PrivateKey: sk,
		Balance:    big.NewInt(0),
	}, nil
}

func (f *Faucet) AddFunds(amount *big.Int) {
	f.Balance.Add(f.Balance, amount)
}

func (f *Faucet) RemoveFunds(amount *big.Int) error {
	if f.Balance.Cmp(amount) < 0 {
		return errors.New("insufficient funds")
	}
	f.Balance.Sub(f.Balance, amount)
	return nil
}

func (f *Faucet) GetBalance() *big.Int {
	return new(big.Int).Set(f.Balance)
}

// Method to create bind.TransactOpts from the Faucet
func (f *Faucet) TransactOpts(simulator *Simulator, chainID *big.Int) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(f.PrivateKey, chainID)
	if err != nil {
		return nil, err
	}

	backend := simulator.GetBackend()

	// Set the nonce, gas price, and gas limit as needed
	nonce, err := backend.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))

	gasPrice, err := backend.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000) // Set your desired gas limit

	return auth, nil
}
