package bindings

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/standards"
	"github.com/unpackdev/solgo/utils"
)

type PairDetails struct {
	Address  common.Address // Address of the Uniswap pair contract
	Token0   common.Address // Address of the first token in the pair
	Token1   common.Address // Address of the second token in the pair
	Reserve0 *big.Int       // Reserve amount of the first token
	Reserve1 *big.Int       // Reserve amount of the second token
}

type Manager struct {
	ctx        context.Context
	clientPool *clients.ClientPool
	bindings   map[utils.Network]map[BindingType]*Binding
	mu         sync.RWMutex // Mutex for thread-safe operations
}

func NewManager(ctx context.Context, clientPool *clients.ClientPool) (*Manager, error) {
	if !standards.StandardsLoaded() {
		if err := standards.LoadStandards(); err != nil {
			return nil, fmt.Errorf("failed to load standards: %w", err)
		}
	}

	return &Manager{
		ctx:        ctx,
		clientPool: clientPool,
		bindings:   make(map[utils.Network]map[BindingType]*Binding),
	}, nil
}

func NewSimulatedManager(ctx context.Context) (*Manager, error) {
	if !standards.StandardsLoaded() {
		if err := standards.LoadStandards(); err != nil {
			return nil, fmt.Errorf("failed to load standards: %w", err)
		}
	}

	return &Manager{
		ctx:      ctx,
		bindings: make(map[utils.Network]map[BindingType]*Binding),
	}, nil
}

func (m *Manager) RegisterBinding(network utils.Network, networkID utils.NetworkID, name BindingType, address common.Address, rawABI string) (*Binding, error) {
	parsedABI, err := abi.JSON(strings.NewReader(rawABI))
	if err != nil {
		return nil, err
	}

	binding := &Binding{
		network:   network,
		networkID: networkID,
		Type:      name,
		Address:   address,
		RawABI:    rawABI,
		ABI:       &parsedABI,
	}

	m.mu.RLock()
	if _, ok := m.bindings[network]; !ok {
		m.bindings[network] = make(map[BindingType]*Binding)
	}
	m.mu.RUnlock()

	// We don't want to overwrite existing bindings and we don't want to register the same binding twice
	if !m.BindingExist(network, name) {
		m.bindings[network][name] = binding
	}

	return binding, nil
}

func (m *Manager) GetClient() *clients.ClientPool {
	return m.clientPool
}

func (m *Manager) GetBinding(network utils.Network, name BindingType) (*Binding, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if networkBindings, ok := m.bindings[network]; ok {
		if binding, ok := networkBindings[name]; ok {
			return binding, nil
		}
	}
	return nil, fmt.Errorf("binding %s not found", name)
}

func (m *Manager) GetBindings(network utils.Network) map[BindingType]*Binding {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.bindings[network]
}

func (m *Manager) BindingExist(network utils.Network, name BindingType) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if networkBindings, ok := m.bindings[network]; ok {
		if _, ok := networkBindings[name]; ok {
			return true
		}
	}
	return false
}

// WatchEvents sets up a subscription to watch events from a specific contract.
func (m *Manager) WatchEvents(network utils.Network, bindingType BindingType, eventName string, ch chan<- types.Log) (ethereum.Subscription, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	binding, ok := m.bindings[network][bindingType]
	if !ok {
		return nil, fmt.Errorf("binding %s not found for network %s", bindingType, network)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{binding.GetAddress()},
	}

	event, ok := binding.ABI.Events[eventName]
	if !ok {
		return nil, fmt.Errorf("event %s not found in ABI", eventName)
	}

	query.Topics = [][]common.Hash{{event.ID}}

	client := m.clientPool.GetClientByGroup(network.String())
	if client == nil {
		return nil, fmt.Errorf("client not found for network %s", network)
	}

	return client.SubscribeFilterLogs(context.Background(), query, ch)
}

// CallContractMethod calls a method on a registered contract.
func (m *Manager) CallContractMethod(ctx context.Context, network utils.Network, bindingType BindingType, toAddr common.Address, methodName string, params ...interface{}) (any, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	binding, ok := m.bindings[network][bindingType]
	if !ok {
		return nil, fmt.Errorf("binding %s not found for network %s", bindingType, network)
	}

	method, ok := binding.ABI.Methods[methodName]
	if !ok {
		return nil, fmt.Errorf("binding %s method %s not found in ABI", bindingType, methodName)
	}

	data, err := method.Inputs.Pack(params...)
	if err != nil {
		return nil, err
	}

	destinationAddr := toAddr
	if destinationAddr == utils.ZeroAddress {
		destinationAddr = binding.Address
	}

	callMsg := ethereum.CallMsg{
		To:   &destinationAddr,
		Data: append(method.ID, data...),
	}

	var result []byte

	client := m.clientPool.GetClientByGroup(network.String())
	if client == nil {
		return nil, fmt.Errorf("client not found for network %s", network)
	}

	result, err = client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	var unpackedResults any
	err = binding.ABI.UnpackIntoInterface(&unpackedResults, methodName, result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack results: %w", err)
	}

	return unpackedResults, nil
}

func (m *Manager) CallContractMethodUnpackMap(ctx context.Context, network utils.Network, bindingType BindingType, toAddr common.Address, methodName string, params ...interface{}) (map[string]any, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	binding, ok := m.bindings[network][bindingType]
	if !ok {
		return nil, fmt.Errorf("binding %s not found for network %s", bindingType, network)
	}

	method, ok := binding.ABI.Methods[methodName]
	if !ok {
		return nil, fmt.Errorf("binding %s method %s not found in ABI", bindingType, methodName)
	}

	data, err := method.Inputs.Pack(params...)
	if err != nil {
		return nil, err
	}

	destinationAddr := toAddr
	if destinationAddr == utils.ZeroAddress {
		destinationAddr = binding.Address
	}

	callMsg := ethereum.CallMsg{
		To:   &destinationAddr,
		Data: append(method.ID, data...),
	}

	var result []byte

	client := m.clientPool.GetClientByGroup(network.String())
	if client == nil {
		return nil, fmt.Errorf("client not found for network %s", network)
	}

	result, err = client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	unpackedResults := map[string]any{}
	err = binding.ABI.UnpackIntoMap(unpackedResults, methodName, result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack results: %w", err)
	}

	return unpackedResults, nil
}
