package solgo

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// ListenerName represents the name of a listener.
type ListenerName string

// Predefined listener names.
const (
	ListenerAbi          ListenerName = "abi"
	ListenerContractInfo ListenerName = "contract_info"
	ListenerAst          ListenerName = "ast"
	ListenerSyntaxErrors ListenerName = "syntax_errors"
)

func (l ListenerName) String() string {
	return string(l)
}

type listeners map[ListenerName]antlr.ParseTreeListener

func (s *Parser) RegisterListener(name ListenerName, listener antlr.ParseTreeListener) error {
	if s.IsListenerRegistered(name) {
		return fmt.Errorf("listener %s already registered", name)
	}

	s.listeners[name] = listener
	return nil
}

func (s *Parser) GetAllListeners() map[ListenerName]antlr.ParseTreeListener {
	return s.listeners
}

func (s *Parser) GetListener(name ListenerName) (antlr.ParseTreeListener, error) {
	if !s.IsListenerRegistered(name) {
		return nil, fmt.Errorf("listener %s not registered", name)
	}

	return s.listeners[name], nil
}

func (s *Parser) IsListenerRegistered(name ListenerName) bool {
	_, ok := s.listeners[name]
	return ok
}
