package observers

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/unpackdev/solgo/clients"
)

func TestNewContractSubscriber(t *testing.T) {
	t.Skip("Skipping Contract Subscriber tests as they require a BSC node to be running...")

	client, err := clients.NewClientPool(context.Background(), &clients.Options{
		Nodes: []clients.Node{
			{
				Group:                   "bsc",
				Type:                    "fullnode",
				FailoverGroup:           "bsc",
				FailoverType:            "archive",
				NetworkId:               56,
				Endpoint:                os.Getenv("FULL_NODE_TEST_URL"),
				ConcurrentClientsNumber: 2,
			},
			{
				Group:                   "bsc",
				Type:                    "archive",
				FailoverGroup:           "bsc",
				FailoverType:            "fullnode",
				NetworkId:               56,
				Endpoint:                os.Getenv("ARCHIVE_NODE_TEST_URL"),
				ConcurrentClientsNumber: 1,
			},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, client)

	tests := []struct {
		name          string
		opts          *ContractSubscriberOptions
		wantErr       bool
		wantContracts int64
		timeout       time.Duration
		timeoutSkip   bool
	}{
		{
			name: "Contracts Subscriber",
			opts: &ContractSubscriberOptions{
				NetworkID: 1,
				Group:     "bsc",
				Type:      "fullnode",
				Head:      true,
			},
			wantErr:       false,
			wantContracts: 1,
			timeout:       10 * time.Second,
			timeoutSkip:   true, // Sometimes heads will have it in 10s, sometimes not, let's test it but skip timeout...
		},
		{
			name: "Contracts Subscriber Start - End",
			opts: &ContractSubscriberOptions{
				NetworkID:        1,
				Group:            "bsc",
				Type:             "fullnode",
				Head:             false,
				StartBlockNumber: big.NewInt(31913866),
				EndBlockNumber:   big.NewInt(31913890),
			},
			wantErr:       false,
			wantContracts: 1,
			timeout:       10 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subscriber, err := NewContractSubscriber(context.Background(), client)
			if tt.wantErr {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
			}
			assert.NotNil(t, subscriber)

			contractsCh := make(chan *Contract)
			go func(contractsCh chan *Contract) {
				err = subscriber.Subscribe(tt.opts, contractsCh)
				assert.NoError(t, err)
			}(contractsCh)

			receivedContracts := int64(0)

		lookupTo:
			for {
				select {
				case contract := <-contractsCh:
					assert.True(t, subscriber.IsActive())
					t.Logf("Received contract: %v", contract)
					receivedContracts++
					if receivedContracts >= tt.wantContracts {
						break lookupTo
					}
				case <-time.After(tt.timeout):
					if !tt.timeoutSkip {
						assert.FailNow(t, "timeout")
					}
					break lookupTo
				}
			}

			subscriber.Close()
			assert.False(t, subscriber.IsActive())
		})
	}
}
