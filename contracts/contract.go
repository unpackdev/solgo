package contracts

import (
	"context"
	"fmt"

	"github.com/0x19/solc-switch"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/exchanges"
	"github.com/unpackdev/solgo/liquidity"
	"github.com/unpackdev/solgo/providers/bitquery"
	"github.com/unpackdev/solgo/providers/etherscan"
	"github.com/unpackdev/solgo/simulator"
	"github.com/unpackdev/solgo/tokens"
	"github.com/unpackdev/solgo/utils"
)

// Metadata holds essential data related to an Ethereum smart contract.
// It includes information about the contract's bytecode, associated transactions, and blockchain context.
type Metadata struct {
	RuntimeBytecode  []byte
	DeployedBytecode []byte
	Block            *types.Block
	Transaction      *types.Transaction
	Receipt          *types.Receipt
}

// Contract represents an Ethereum smart contract within the context of a specific network.
// It encapsulates the contract's address, network information, and associated metadata,
// and provides methods to interact with the contract on the blockchain.
type Contract struct {
	ctx             context.Context
	clientPool      *clients.ClientPool
	client          *clients.Client
	addr            common.Address
	network         utils.Network
	descriptor      *Descriptor
	liq             *liquidity.Liquidity
	sim             *simulator.Simulator
	token           *tokens.Token
	bqp             *bitquery.BitQueryProvider
	etherscan       *etherscan.EtherScanProvider
	compiler        *solc.Solc
	bindings        *bindings.Manager
	exchangeManager *exchanges.Manager
	tokenBind       *bindings.Token
}

// NewContract creates a new instance of Contract for a given Ethereum address and network.
// It initializes the contract's context, metadata, and associated blockchain clients.
// The function validates the contract's existence and its bytecode before creation.
func NewContract(ctx context.Context, network utils.Network, clientPool *clients.ClientPool, sim *simulator.Simulator, liq *liquidity.Liquidity, bqp *bitquery.BitQueryProvider, etherscan *etherscan.EtherScanProvider, compiler *solc.Solc, bindManager *bindings.Manager, exchange *exchanges.Manager, addr common.Address) (*Contract, error) {
	if clientPool == nil {
		return nil, fmt.Errorf("client pool is nil")
	}

	client := clientPool.GetClientByGroup(network.String())
	if client == nil {
		return nil, fmt.Errorf("client for network %s is nil", network.String())
	}

	if !common.IsHexAddress(addr.Hex()) {
		return nil, fmt.Errorf("invalid address provided: %s", addr.Hex())
	}

	tokenBind, err := bindings.NewToken(ctx, network, bindManager, bindings.DefaultTokenBindOptions(addr))
	if err != nil {
		return nil, fmt.Errorf("failed to create new token %s bindings: %w", addr, err)
	}

	token, err := tokens.NewToken(
		ctx,
		network,
		addr,
		utils.AnvilSimulator,
		bindManager,
		exchange,
		sim,
		clientPool,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create new token %s instance: %w", addr, err)
	}

	toReturn := &Contract{
		ctx:        ctx,
		network:    network,
		clientPool: clientPool,
		client:     client,
		addr:       addr,
		liq:        liq,
		bqp:        bqp,
		etherscan:  etherscan,
		compiler:   compiler,
		sim:        sim,
		descriptor: &Descriptor{
			Network:        network,
			NetworkID:      utils.GetNetworkID(network),
			Address:        addr,
			LiquidityPairs: make(map[utils.ExchangeType]common.Address),
			Safety:         &SafetyDescriptor{},
		},
		bindings:        bindManager,
		exchangeManager: exchange,
		token:           token,
		tokenBind:       tokenBind,
	}

	/* 	inspect, err := inspector.NewInspector(ctx, clientPool, toReturn)
	   	if err != nil {
	   		return nil, fmt.Errorf("failed to create new inspector: %w", err)
	   	}
	   	toReturn.inspector = inspect
	*/

	if valid, err := toReturn.IsValid(); err != nil {
		return nil, fmt.Errorf("failure to check for contract validity: %s", err)
	} else if !valid {
		return nil, fmt.Errorf("requested contract '%s' is not valid", addr.Hex())
	}

	return toReturn, nil
}

// GetAddress returns the Ethereum address of the contract.
func (c *Contract) GetAddress() common.Address {
	return c.addr
}

// GetNetwork returns the network (e.g., Mainnet, Ropsten) on which the contract is deployed.
func (c *Contract) GetNetwork() utils.Network {
	return c.network
}

// GetDeployedBytecode returns the deployed bytecode of the contract.
// This bytecode is the compiled contract code that exists on the Ethereum blockchain.
func (c *Contract) GetDeployedBytecode() []byte {
	return c.descriptor.DeployedBytecode
}

// GetRuntimeBytecode returns the runtime bytecode of the contract.
// This bytecode is used during the execution of contract calls and transactions.
func (c *Contract) GetRuntimeBytecode() []byte {
	return c.descriptor.RuntimeBytecode
}

// GetBlock returns the blockchain block in which the contract was deployed or involved.
func (c *Contract) GetBlock() *types.Block {
	return c.descriptor.Block
}

// GetTransaction returns the Ethereum transaction associated with the contract's deployment or a specific operation.
func (c *Contract) GetTransaction() *types.Transaction {
	return c.descriptor.Transaction
}

// GetReceipt returns the receipt of the transaction in which the contract was involved,
// providing details such as gas used and logs generated.
func (c *Contract) GetReceipt() *types.Receipt {
	return c.descriptor.Receipt
}

// GetSender returns the Ethereum address of the sender of the contract's transaction.
// It extracts the sender's address using the transaction's signature.
func (c *Contract) GetSender() (common.Address, error) {
	from, err := types.Sender(types.LatestSignerForChainID(c.descriptor.Transaction.ChainId()), c.descriptor.Transaction)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get sender: %s", err)
	}

	return from, nil
}

func (c *Contract) GetToken() *tokens.Token {
	return c.token
}

// IsValid checks if the contract is valid by verifying its deployed bytecode.
// A contract is considered valid if it has non-empty deployed bytecode on the blockchain.
func (c *Contract) IsValid() (bool, error) {
	if err := c.discoverDeployedBytecode(); err != nil {
		return false, err
	}

	return len(c.descriptor.DeployedBytecode) > 0, nil
}

func (c *Contract) GetDescriptor() *Descriptor {
	return c.descriptor
}
