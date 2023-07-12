package metadata

import (
	"context"
	"testing"

	ipfs "github.com/ipfs/go-ipfs-api"
	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo/tests"
)

func TestIpfsProvider(t *testing.T) {
	t.Skip("Skipping IPFS tests as they require an IPFS node to be running...")
	tAssert := assert.New(t)

	context, cancel := context.WithCancel(context.TODO())
	defer cancel()

	sh := ipfs.NewShell("4.tcp.eu.ngrok.io:18285")

	provider, err := NewIpfsProvider(context, sh)
	tAssert.NoError(err)
	tAssert.NotNil(provider)

	tests := []struct {
		name    string
		hash    string
		want    []byte
		wantErr bool
	}{
		{
			name:    "Partial Metadata",
			hash:    "ipfs://QmPL7gzcnyeyKUqQCJsvc5qbc9hqaopuRLtfuyLNsgn5oS",
			want:    tests.ReadJsonBytesForTest(t, "SushiXSwapMetadata"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := provider.GetMetadataByCID(tt.hash)
			if tt.wantErr {
				tAssert.Error(err)
			} else {
				tAssert.NoError(err)
				jsonResponse, err := response.ToJSON()
				tAssert.NoError(err)
				tAssert.Equal(tt.want, jsonResponse)
			}
		})
	}
}
