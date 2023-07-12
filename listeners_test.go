package solgo

import (
	"context"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
)

type mockListener struct {
	antlr.ParseTreeListener
}

func TestListenerGetterAndSetter(t *testing.T) {
	ctx := context.TODO()
	input := strings.NewReader("contract Test {}")

	// Test New
	s, err := NewParser(ctx, input)
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, input, s.GetInput())
	assert.NotNil(t, s.GetInputStream())
	assert.NotNil(t, s.GetLexer())
	assert.NotNil(t, s.GetTokenStream())
	assert.NotNil(t, s.GetParser())
	assert.NotNil(t, s.GetTree())

	// Test RegisterListener
	listener := &mockListener{}
	err = s.RegisterListener(ListenerAbi, listener)
	assert.NoError(t, err)

	// Test IsListenerRegistered
	assert.True(t, s.IsListenerRegistered(ListenerAbi))

	gotListener, err := s.GetListener(ListenerAbi)
	assert.NoError(t, err)
	assert.Equal(t, listener, gotListener)

	gotListeners := s.GetAllListeners()
	assert.Equal(t, 1, len(gotListeners))
	assert.Equal(t, listener, gotListeners[ListenerAbi])

	err = s.RegisterListener(ListenerAbi, listener)
	assert.Error(t, err)

	_, err = s.GetListener(ListenerAst)
	assert.Error(t, err)
}
