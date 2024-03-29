package observers

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/unpackdev/solgo/clients"
)

func TestNewBlockHeaderSubscriber(t *testing.T) {
	t.Skip("Skipping New Block Subscriber tests as they require a node to be running...")

	client, err := clients.NewClientPool(context.Background(), &clients.Options{
		Nodes: []clients.Node{
			{
				Group:             "bsc",
				Type:              "fullnode",
				FailoverGroup:     "bsc",
				FailoverType:      "archive",
				NetworkId:         56,
				Endpoint:          os.Getenv("FULL_NODE_TEST_URL"),
				ConcurrentClients: 2,
			},
			{
				Group:             "bsc",
				Type:              "archive",
				FailoverGroup:     "bsc",
				FailoverType:      "fullnode",
				NetworkId:         56,
				Endpoint:          os.Getenv("ARCHIVE_NODE_TEST_URL"),
				ConcurrentClients: 1,
			},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, client)

	tests := []struct {
		name       string
		opts       *BlockSubscriberOptions
		wantErr    bool
		wantBlocks int64
	}{
		{
			name: "Block Header Subscriber",
			opts: &BlockSubscriberOptions{
				NetworkID: 1,
				Group:     "bsc",
				Type:      "fullnode",
				Head:      true,
			},
			wantErr:    false,
			wantBlocks: 2,
		},
		{
			name: "Block Header Subscriber Start - End",
			opts: &BlockSubscriberOptions{
				NetworkID:        1,
				Group:            "bsc",
				Type:             "fullnode",
				Head:             false,
				StartBlockNumber: big.NewInt(31913866),
				EndBlockNumber:   big.NewInt(31913868),
			},
			wantErr:    false,
			wantBlocks: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subscriber, err := NewBlockSubscriber(context.Background(), client)
			if tt.wantErr {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
			}
			assert.NotNil(t, subscriber)

			blockCh := make(chan *types.Header)
			go func(blockCh chan *types.Header) {
				err = subscriber.SubscribeHeader(tt.opts, blockCh)
				assert.NoError(t, err)
			}(blockCh)

			receivedBlocks := int64(0)

		lookupTo:
			for {
				select {
				case block := <-blockCh:
					assert.True(t, subscriber.IsActive())
					t.Logf("Received block header number: %d", block.Number.Int64())
					receivedBlocks++
					if receivedBlocks >= tt.wantBlocks {
						break lookupTo
					}
				case <-time.After(10 * time.Second):
					assert.FailNow(t, "timeout")
					break lookupTo
				}
			}

			subscriber.Close()
			assert.False(t, subscriber.IsActive())
		})
	}
}

func TestNewBlockSubscriber(t *testing.T) {
	t.Skip("Skipping New Block Subscriber tests as they require a node to be running...")

	client, err := clients.NewClientPool(context.Background(), &clients.Options{
		Nodes: []clients.Node{
			{
				Group:             "bsc",
				Type:              "fullnode",
				FailoverGroup:     "bsc",
				FailoverType:      "archive",
				NetworkId:         56,
				Endpoint:          os.Getenv("FULL_NODE_TEST_URL"),
				ConcurrentClients: 2,
			},
			{
				Group:             "bsc",
				Type:              "archive",
				FailoverGroup:     "bsc",
				FailoverType:      "fullnode",
				NetworkId:         56,
				Endpoint:          os.Getenv("ARCHIVE_NODE_TEST_URL"),
				ConcurrentClients: 1,
			},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, client)

	tests := []struct {
		name       string
		opts       *BlockSubscriberOptions
		wantErr    bool
		wantBlocks int64
	}{
		{
			name: "Block Subscriber",
			opts: &BlockSubscriberOptions{
				NetworkID: 1,
				Group:     "bsc",
				Type:      "fullnode",
				Head:      true,
			},
			wantErr:    false,
			wantBlocks: 2,
		},
		{
			name: "Block Subscriber Start - End",
			opts: &BlockSubscriberOptions{
				NetworkID:        1,
				Group:            "bsc",
				Type:             "fullnode",
				Head:             false,
				StartBlockNumber: big.NewInt(31913866),
				EndBlockNumber:   big.NewInt(31913868),
			},
			wantErr:    false,
			wantBlocks: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subscriber, err := NewBlockSubscriber(context.Background(), client)
			if tt.wantErr {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
			}
			assert.NotNil(t, subscriber)

			blockCh := make(chan *types.Block)
			go func(blockCh chan *types.Block) {
				err = subscriber.Subscribe(tt.opts, blockCh)
				assert.NoError(t, err)
			}(blockCh)

			receivedBlocks := int64(0)

		lookupTo:
			for {
				select {
				case block := <-blockCh:
					assert.True(t, subscriber.IsActive())
					t.Logf("Received block header number: %d", block.NumberU64())
					receivedBlocks++
					if receivedBlocks >= tt.wantBlocks {
						break lookupTo
					}
				case <-time.After(10 * time.Second):
					assert.FailNow(t, "timeout")
					break lookupTo
				}
			}

			subscriber.Close()
			assert.False(t, subscriber.IsActive())
		})
	}
}
