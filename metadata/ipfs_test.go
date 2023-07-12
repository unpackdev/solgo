package metadata

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"strings"
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

func TestNewIpfsProvider(t *testing.T) {
	_, err := NewIpfsProvider(context.Background(), nil)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrInvalidIpfsClient))

	_, err = NewIpfsProvider(context.Background(), &MockShell{})
	assert.NoError(t, err)
}

func TestGetMetadataByCID(t *testing.T) {
	tests := []struct {
		name           string
		cid            string
		mockCatContent string
		expectError    bool
		expectedRaw    string
		jsonByteSize   int
	}{
		{
			name:        "invalid CID",
			cid:         "invalidCID",
			expectError: true,
		},
		{
			name: "valid CID",
			cid:  "ipfs://QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG",
			mockCatContent: `{
				"raw": "raw",
				"version": 1,
				"compiler": {
					"version": "0.8.0",
					"keccak256": "keccak256"
				},
				"language": "Solidity",
				"settings": {
					"evmVersion": "istanbul",
					"compilationTarget": {
						"contract.sol": "Contract"
					},
					"libraries": {},
					"remappings": [],
					"metadata": {
						"bytecodeHash": "ipfs",
						"useLiteralContent": true,
						"appendCBOR": false
					},
					"optimizer": {
						"enabled": true,
						"runs": 200,
						"details": {
							"peephole": true,
							"inliner": true,
							"jumpdestRemover": true,
							"orderLiterals": true,
							"deduplicate": true,
							"cse": true,
							"constantOptimizer": true,
							"yul": true,
							"yulDetails": {
								"stackAllocation": true,
								"optimizerSteps": 1000
							}
						}
					}
				},
				"output": {
					"abi": []
				},
				"sources": {
					"contract.sol": {
						"content": "contract Contract {}",
						"keccak256": "keccak256",
						"license": "MIT"
					}
				}
			}`,
			expectedRaw:  "{\n\t\t\t\t\"raw\": \"raw\",\n\t\t\t\t\"version\": 1,\n\t\t\t\t\"compiler\": {\n\t\t\t\t\t\"version\": \"0.8.0\",\n\t\t\t\t\t\"keccak256\": \"keccak256\"\n\t\t\t\t},\n\t\t\t\t\"language\": \"Solidity\",\n\t\t\t\t\"settings\": {\n\t\t\t\t\t\"evmVersion\": \"istanbul\",\n\t\t\t\t\t\"compilationTarget\": {\n\t\t\t\t\t\t\"contract.sol\": \"Contract\"\n\t\t\t\t\t},\n\t\t\t\t\t\"libraries\": {},\n\t\t\t\t\t\"remappings\": [],\n\t\t\t\t\t\"metadata\": {\n\t\t\t\t\t\t\"bytecodeHash\": \"ipfs\",\n\t\t\t\t\t\t\"useLiteralContent\": true,\n\t\t\t\t\t\t\"appendCBOR\": false\n\t\t\t\t\t},\n\t\t\t\t\t\"optimizer\": {\n\t\t\t\t\t\t\"enabled\": true,\n\t\t\t\t\t\t\"runs\": 200,\n\t\t\t\t\t\t\"details\": {\n\t\t\t\t\t\t\t\"peephole\": true,\n\t\t\t\t\t\t\t\"inliner\": true,\n\t\t\t\t\t\t\t\"jumpdestRemover\": true,\n\t\t\t\t\t\t\t\"orderLiterals\": true,\n\t\t\t\t\t\t\t\"deduplicate\": true,\n\t\t\t\t\t\t\t\"cse\": true,\n\t\t\t\t\t\t\t\"constantOptimizer\": true,\n\t\t\t\t\t\t\t\"yul\": true,\n\t\t\t\t\t\t\t\"yulDetails\": {\n\t\t\t\t\t\t\t\t\"stackAllocation\": true,\n\t\t\t\t\t\t\t\t\"optimizerSteps\": 1000\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t},\n\t\t\t\t\"output\": {\n\t\t\t\t\t\"abi\": []\n\t\t\t\t},\n\t\t\t\t\"sources\": {\n\t\t\t\t\t\"contract.sol\": {\n\t\t\t\t\t\t\"content\": \"contract Contract {}\",\n\t\t\t\t\t\t\"keccak256\": \"keccak256\",\n\t\t\t\t\t\t\"license\": \"MIT\"\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t}",
			jsonByteSize: 2111,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider, err := NewIpfsProvider(context.Background(), &MockShell{
				CatFunc: func(path string) (io.ReadCloser, error) {
					// Mock the Cat function to return the metadata content
					return ioutil.NopCloser(strings.NewReader(tt.mockCatContent)), nil
				},
			})
			assert.NoError(t, err)

			metadata, err := provider.GetMetadataByCID(tt.cid)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedRaw, metadata.Raw)
				mJson, err := metadata.ToJSON()
				assert.NoError(t, err)
				assert.GreaterOrEqual(t, len(mJson), tt.jsonByteSize)

				aJson, err := metadata.AbiToJSON()
				assert.NoError(t, err)
				assert.GreaterOrEqual(t, len(aJson), 1)
			}
		})
	}
}
