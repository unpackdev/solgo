package clients

import (
	"context"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientPool(t *testing.T) {
	tests := []struct {
		name        string
		opts        *Options
		group       string
		typ         string
		expectError bool
		errorMsg    string
	}{
		{
			name: "Invalid Client",
			opts: &Options{
				Nodes: []Node{
					{
						Group:                   "testGroup",
						Type:                    "testType",
						Endpoint:                "http://localhost:2222",
						ConcurrentClientsNumber: 1,
						NetworkId:               1,
					},
				},
			},
			group:       "testGroup",
			typ:         "testType",
			expectError: true,
			errorMsg:    "failed to initialize Ethereum client: Post \"http://localhost:2222\": dial tcp [::1]:2222: connect: connection refused",
		},
		{
			name: "Bsc Clients",
			opts: &Options{
				Nodes: []Node{
					{
						Group:                   "bsc",
						Type:                    "mainnet",
						Endpoint:                "https://bsc-dataseed.binance.org/",
						ConcurrentClientsNumber: 1,
						NetworkId:               56,
					},
					{
						Group:                   "bsc",
						Type:                    "archive",
						Endpoint:                "https://bsc-dataseed.binance.org/",
						ConcurrentClientsNumber: 1,
						NetworkId:               56,
					},
				},
			},
			group:       "bsc",
			typ:         "mainnet",
			expectError: false,
		},
		{
			name: "Empty Endpoint",
			opts: &Options{
				Nodes: []Node{
					{
						Group:                   "testGroup",
						Type:                    "testType",
						Endpoint:                "",
						ConcurrentClientsNumber: 1,
						NetworkId:               1,
					},
				},
			},
			group:       "testGroup",
			typ:         "testType",
			expectError: true,
			errorMsg:    "configuration client URL not set",
		},
		{
			name: "Zero Concurrent Clients",
			opts: &Options{
				Nodes: []Node{
					{
						Group:                   "testGroup",
						Type:                    "testType",
						Endpoint:                "http://localhost:8545",
						ConcurrentClientsNumber: 0,
						NetworkId:               1,
					},
				},
			},
			group:       "testGroup",
			typ:         "testType",
			expectError: true,
			errorMsg:    "configuration amount of concurrent clients is not set",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pool, err := NewClientPool(context.Background(), tt.opts)
			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMsg, err.Error())
				return
			} else {
				assert.NoError(t, err)
			}

			client := pool.GetClient(tt.group, tt.typ)
			if tt.expectError {
				assert.Nil(t, client)
			} else {
				assert.NotNil(t, client)
			}

			assert.Equal(t, len(tt.opts.GetNodes()), pool.Len())

			clientByGroup := pool.GetClientByGroup(tt.group)
			if tt.expectError {
				assert.Nil(t, clientByGroup)
			} else {
				assert.NotNil(t, clientByGroup)
			}

			clientByGroupAndType := pool.GetClientByGroupAndType(tt.group, tt.typ)
			if tt.expectError {
				assert.Nil(t, clientByGroupAndType)
			} else {
				assert.NotNil(t, clientByGroupAndType)
			}

			assert.Equal(t, tt.group, clientByGroupAndType.GetGroup())
			assert.Equal(t, tt.typ, clientByGroupAndType.GetType())
			assert.Equal(t, tt.opts.GetNodes()[0].Endpoint, clientByGroupAndType.GetEndpoint())
			assert.Equal(t, tt.opts.GetNodes()[0].NetworkId, clientByGroupAndType.GetNetworkID())

			nodeGroup, nodeType := pool.GetClientDescriptionByNetworkId(big.NewInt(clientByGroupAndType.GetNetworkID()))
			assert.Equal(t, tt.group, nodeGroup)
			assert.Equal(t, tt.typ, nodeType)

			pool.Close()
		})
	}
}
