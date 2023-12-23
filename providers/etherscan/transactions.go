package etherscan

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type ContractCreation struct {
	Address         string `json:"contractAddress"`
	CreatorAddress  string `json:"contractCreator"`
	TransactionHash string `json:"txHash"`
}

func (c *ContractCreation) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *ContractCreation) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *ContractCreation) GetTransactionHash() common.Hash {
	return common.HexToHash(c.TransactionHash)
}

type ContractCreationResponse struct {
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Result  []*ContractCreation `json:"result"`
}

func (p *EtherScanProvider) QueryContractCreationTx(ctx context.Context, addr common.Address) (*ContractCreation, error) {
	if p.cache != nil {
		cacheKey := p.CacheKey("getcontractcreation", addr.Hex())
		if result, err := p.cache.Exists(p.ctx, cacheKey).Result(); err == nil && result == 1 {
			if result, err := p.cache.Get(p.ctx, cacheKey).Result(); err == nil {
				var response *ContractCreation
				if err := json.Unmarshal([]byte(result), &response); err != nil {
					return nil, fmt.Errorf("failed to unmarshal %s response: %s", p.ProviderName(), err)
				}
				return response, nil
			}
		}
	}

	// @FIXME: For now we're going to sleep a bit to avoid rate limiting.
	// There are multiple ways we can sort this by applying pool of the API keys and rotating them.
	// In addition, what's a must at first is basically purchasing the API key from Etherscan.
	time.Sleep(200 * time.Millisecond)

	url := fmt.Sprintf("%s?module=contract&action=getcontractcreation&contractaddresses=%s&apikey=%s", p.opts.Endpoint, addr.Hex(), p.GetNextKey())

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	if strings.Contains(string(body), "NOTOK") {
		var creationResponse ErrorResponse
		if err := json.Unmarshal(body, &creationResponse); err != nil {
			return nil, fmt.Errorf("failed to unmarshal error response: %s", err)
		}

		return nil, errors.New(creationResponse.Result)
	}

	var creationResponse ContractCreationResponse
	if err := json.Unmarshal(body, &creationResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal contract creation response: %s", err)
	}

	if p.cache != nil {
		cacheKey := p.CacheKey("getcontractcreation", addr.Hex())
		if err := p.cache.Set(p.ctx, cacheKey, creationResponse.Result[0], 1*time.Hour).Err(); err != nil {
			return nil, fmt.Errorf("failed to write to cache: %s", err)
		}
	}

	return creationResponse.Result[0], nil
}
