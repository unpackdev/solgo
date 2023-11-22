package observers

import "fmt"

type ProcessorType string
type HookType string

func (h HookType) String() string {
	return string(h)
}

const (
	PreHook  HookType = "pre"
	PostHook HookType = "post"

	BlockProcessor       ProcessorType = "blocks"
	TransactionProcessor ProcessorType = "transactions"
	LogProcessor         ProcessorType = "logs"
	ContractProcessor    ProcessorType = "contracts"
)

type BlockHookFn func(*BlockEntry) (*BlockEntry, error)
type TransactionHookFn func(*TransactionEntry) (*TransactionEntry, error)
type LogHookFn func(*LogEntry) (*LogEntry, error)
type ContractHookFn func(*ContractEntry) (*ContractEntry, error)

func (m *Manager) RegisterHook(processor ProcessorType, hookType HookType, fn interface{}) error {
	if hooks, ok := m.hooks[processor]; ok {
		if _, ok := hooks[hookType]; ok {
			return fmt.Errorf("hook %s already registered for processor %s", hookType, processor)
		}
	}

	if _, ok := m.hooks[processor]; !ok {
		m.hooks[processor] = make(map[HookType]interface{})
	}

	m.hooks[processor][hookType] = fn
	return nil
}

func (m *Manager) GetHooks(processor ProcessorType) map[HookType]interface{} {
	return m.hooks[processor]
}
