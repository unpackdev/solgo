package etherscan

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/metadata"
)

type Contract struct {
	SourceCode           interface{} `json:"SourceCode"`
	ABI                  string      `json:"ABI"`
	Name                 string      `json:"ContractName"`
	CompilerVersion      string      `json:"CompilerVersion"`
	OptimizationUsed     string      `json:"OptimizationUsed"`
	Runs                 string      `json:"Runs"`
	ConstructorArguments string      `json:"ConstructorArguments"`
	EVMVersion           string      `json:"EVMVersion"`
	Library              string      `json:"Library"`
	LicenseType          string      `json:"LicenseType"`
	Proxy                string      `json:"Proxy"`
	Implementation       string      `json:"Implementation"`
	SwarmSource          string      `json:"SwarmSource"`
}

func (c Contract) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Contract) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

type ContractResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Result  []Contract `json:"result"`
}

func (e *EtherScanProvider) ScanContract(addr common.Address) (*Contract, error) {
	cacheKey := e.CacheKey("getsourcecode_method_1", addr.Hex())

	if e.cache != nil {
		if result, err := e.cache.Exists(e.ctx, cacheKey).Result(); err == nil && result == 1 {
			if result, err := e.cache.Get(e.ctx, cacheKey).Result(); err == nil {
				var response *Contract
				if err := json.Unmarshal([]byte(result), &response); err != nil {
					return nil, fmt.Errorf("failed to unmarshal %s response: %s", e.ProviderName(), err)
				}
				return response, nil
			}
		}
	}

	// @FIXME: For now we're going to sleep a bit to avoid rate limiting.
	// There are multiple ways we can sort this by applying pool of the API keys and rotating them.
	// In addition, what's a must at first is basically purchasing the API key from Etherscan.
	time.Sleep(200 * time.Millisecond)

	url := fmt.Sprintf(
		"%s?module=contract&action=getsourcecode&address=%s&apikey=%s",
		e.opts.Endpoint, addr.Hex(), e.GetNextKey(),
	)

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
		var contractResponse ErrorResponse
		if err := json.Unmarshal(body, &contractResponse); err != nil {
			return nil, fmt.Errorf("failed to unmarshal error response: %s", err)
		}
		return nil, errors.New(contractResponse.Result)
	}

	var contractResponse ContractResponse
	if err := json.Unmarshal(body, &contractResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal etherscan response: %s", err)
	}

	toReturn := contractResponse.Result[0]

	if toReturn.ABI == "Contract source code not verified" {
		return nil, fmt.Errorf("contract source code not verified")
	}

	// Well we have standard metadata response probably from IPFS within the source code and we need to parse it...
	// Not to mention with excuse Etherscan bullshit we need to unquote it manually...
	if len(toReturn.SourceCode.(string)) > 2 && (toReturn.SourceCode.(string)[:2] == "{{" || toReturn.SourceCode.(string)[:1] == "{") {
		cm := metadata.ContractMetadata{}

		// This is just nasty but I don't really care at this moment...
		sourceCode, _ := unprettyJSON(manualUnquote(toReturn.SourceCode.(string)))

		if err := json.Unmarshal([]byte(sourceCode), &cm); err != nil {
			var onlySources map[string]metadata.ContractSource
			if err := json.Unmarshal([]byte(fmt.Sprint(toReturn.SourceCode.(string))), &onlySources); err != nil {
				return nil, fmt.Errorf("failed to unmarshal metadata response (string): %s", err)
			} else {
				cm.Sources = onlySources
			}

		}

		toReturn.SourceCode = cm
	} else if _, ok := toReturn.SourceCode.(map[string]interface{}); ok {
		cm := metadata.ContractMetadata{}
		encoded, _ := json.Marshal(toReturn.SourceCode)
		if err := json.Unmarshal(encoded, &cm); err != nil {
			return nil, fmt.Errorf("failed to unmarshal metadata response (map): %s", err)
		}
		toReturn.SourceCode = cm
	}

	if toReturn.ABI == "Contract source code not verified" {
		return nil, fmt.Errorf("contract not found")
	}

	if e.cache != nil {
		if err := e.cache.Set(e.ctx, cacheKey, toReturn, 1*time.Hour).Err(); err != nil {
			return nil, fmt.Errorf("failed to write to cache: %s", err)
		}
	}

	return &toReturn, nil
}

func manualUnquote(s string) string {
	// Check if it starts with `{{` and ends with `}}`
	if strings.HasPrefix(s, "{{") && strings.HasSuffix(s, "}}") {
		s = "{" + s[2:len(s)-2] + "}"
	}

	return s
}

// UnprettyJSON takes a pretty JSON string and returns a compact, unprettified version of it.
func unprettyJSON(prettyJSON string) (string, error) {
	var data interface{}

	// Unmarshal the pretty JSON into an interface{}
	err := json.Unmarshal([]byte(prettyJSON), &data)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	// Marshal the data back into JSON without indentation
	compactJSON, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %w", err)
	}

	return string(compactJSON), nil
}
