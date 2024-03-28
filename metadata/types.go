package metadata

import (
	"encoding/json"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	metadata_pb "github.com/unpackdev/protos/dist/go/metadata"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type ContractSource struct {
	Content   string   `json:"content"`
	Keccak256 string   `json:"keccak256"`
	License   string   `json:"license"`
	Urls      []string `json:"urls"`
}

// ContractMetadata represents the metadata of a contract stored in IPFS.
// The metadata includes information about the compiler, language, settings, output, and sources.
// The Raw field contains the raw metadata as a string.
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
		Libraries         interface{}       `json:"libraries"`
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
	Sources map[string]ContractSource `json:"sources"`
}

func (c *ContractMetadata) ToProto() *metadata_pb.Metadata {
	abiValue := &any.Any{}
	bytes, _ := json.Marshal(c.Output.Abi)
	abiBytesValue := &wrappers.BytesValue{
		Value: bytes,
	}
	if err := anypb.MarshalFrom(abiValue, abiBytesValue, proto.MarshalOptions{}); err != nil {
		panic(err)
	}

	sources := map[string]*metadata_pb.Metadata_Source{}
	for sourceName, source := range c.Sources {
		sources[sourceName] = &metadata_pb.Metadata_Source{
			Content:   source.Content,
			Keccak256: source.Keccak256,
			License:   source.License,
			Urls:      source.Urls,
		}
	}

	return &metadata_pb.Metadata{
		Raw:     c.Raw,
		Version: int32(c.Version),
		Compiler: &metadata_pb.Metadata_Compiler{
			Version:   c.Compiler.Version,
			Keccak256: c.Compiler.Keccak256,
		},
		Language: c.Language,
		Settings: &metadata_pb.Metadata_Settings{
			EvmVersion:        c.Settings.EvmVersion,
			CompilationTarget: c.Settings.CompilationTarget,
			// TODO: Libraries can be different types, we'll need to figure this shit out...
			//Libraries:         c.Settings.Libraries.(map[string]string),
			Remappings: c.Settings.Remappings,
			Metadata: &metadata_pb.Metadata_Settings_MetadataSettings{
				BytecodeHash:      c.Settings.Metadata.BytecodeHash,
				UseLiteralContent: c.Settings.Metadata.UseLiteralContent,
				AppendCbor:        c.Settings.Metadata.AppendCBOR,
			},
			Optimizer: &metadata_pb.Metadata_Settings_Optimizer{
				Enabled: c.Settings.Optimizer.Enabled,
				Runs:    int32(c.Settings.Optimizer.Runs),
				Details: &metadata_pb.Metadata_Settings_Optimizer_Details{
					Peephole:          c.Settings.Optimizer.Details.Peephole,
					Inliner:           c.Settings.Optimizer.Details.Inliner,
					JumpdestRemover:   c.Settings.Optimizer.Details.JumpdestRemover,
					OrderLiterals:     c.Settings.Optimizer.Details.OrderLiterals,
					Deduplicate:       c.Settings.Optimizer.Details.Deduplicate,
					Cse:               c.Settings.Optimizer.Details.Cse,
					ConstantOptimizer: c.Settings.Optimizer.Details.ConstantOptimizer,
					Yul:               c.Settings.Optimizer.Details.Yul,
					YulDetails: &metadata_pb.Metadata_Settings_Optimizer_Details_YulDetails{
						StackAllocation: c.Settings.Optimizer.Details.YulDetails.StackAllocation,
						OptimizerSteps:  int32(c.Settings.Optimizer.Details.YulDetails.OptimizerSteps),
					},
				},
			},
		},
		Output: &metadata_pb.Metadata_Output{
			Abi: abiValue,
		},
		Sources: sources,
	}
}

// AbiToJSON converts the ABI of the contract to a JSON string.
// It returns the JSON string or an error if the conversion fails.
func (c *ContractMetadata) AbiToJSON() (string, error) {
	abi, err := json.Marshal(c.Output.Abi)
	if err != nil {
		return "", err
	}

	return string(abi), nil
}

// ToJSON converts the ContractMetadata object to a JSON byte array.
// It returns the byte array or an error if the conversion fails.
func (c *ContractMetadata) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}
