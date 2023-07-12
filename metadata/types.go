package metadata

import "encoding/json"

// ContractMetadata represents the metadata of a contract stored in IPFS
// See https://docs.soliditylang.org/en/v0.8.19/metadata.html
// @TODO: Devdoc and userdoc are not included in the struct as I see very small amount of
// reasons to use them in this moment. If you need them, please consider doing a PR to add them. :)
// Basically lots of the times devdoc and userdoc lack information so they are not very useful.
// You can use .Raw to extract dev and user doc if you need them.
type ContractMetadata struct {
	Raw      string `json:"raw"`
	Version  int    `json:"version"`
	Compiler struct {
		Version   string `json:"version"`
		Keccak256 string `json:"keccak256"`
	} `json:"compiler"`
	Language string `json:"language"`
	Settings struct {
		EvmVersion        string            `json:"evmVersion"`
		CompilationTarget map[string]string `json:"compilationTarget"`
		Libraries         map[string]string `json:"libraries"`
		Remappings        []string          `json:"remappings"`
		Metadata          struct {
			BytecodeHash      string `json:"bytecodeHash"`
			UseLiteralContent bool   `json:"useLiteralContent"`
			AppendCBOR        bool   `json:"appendCBOR"`
		} `json:"metadata"`
		Optimizer struct {
			Enabled bool `json:"enabled"`
			Runs    int  `json:"runs"`
			Details struct {
				Peephole          bool `json:"peephole"`
				Inliner           bool `json:"inliner"`
				JumpdestRemover   bool `json:"jumpdestRemover"`
				OrderLiterals     bool `json:"orderLiterals"`
				Deduplicate       bool `json:"deduplicate"`
				Cse               bool `json:"cse"`
				ConstantOptimizer bool `json:"constantOptimizer"`
				Yul               bool `json:"yul"`
				YulDetails        struct {
					StackAllocation bool `json:"stackAllocation"`
					OptimizerSteps  int  `json:"optimizerSteps"`
				} `json:"yulDetails"`
			} `json:"details"`
		} `json:"optimizer"`
	} `json:"settings"`
	Output struct {
		Abi []interface{} `json:"abi"`
	}
	Sources map[string]struct {
		Content string `json:"content"`
		Keccak  string `json:"keccak256"`
		License string `json:"license"`
	} `json:"sources"`
}

// AbiToJSON returns the ABI as a JSON string
func (c *ContractMetadata) AbiToJSON() (string, error) {
	abi, err := json.Marshal(c.Output.Abi)
	if err != nil {
		return "", err
	}

	return string(abi), nil
}

func (c *ContractMetadata) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}
