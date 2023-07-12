package bytecode

import (
	"fmt"
	"reflect"
	"strings"

	abi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type Argument struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Value   string `json:"value"`
	Indexed bool   `json:"indexed"`
}

type Constructor struct {
	Abi          string `json:"abi"`
	SignatureRaw string `json:"signature_raw"`
	Arguments    []Argument
}

func DecodeConstructorFromAbi(bytecode []byte, constructorAbi string) (*Constructor, error) {
	if len(bytecode) == 0 || bytecode[0] != '[' {
		constructorAbi = "[" + constructorAbi + "]"
	}

	parsed, err := abi.JSON(strings.NewReader(constructorAbi))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	unpacked, err := parsed.Constructor.Inputs.UnpackValues(bytecode)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack values: %w", err)
	}

	if len(unpacked) != len(parsed.Constructor.Inputs) {
		return nil, fmt.Errorf("number of unpacked values does not match number of inputs")
	}

	arguments := make([]Argument, len(parsed.Constructor.Inputs))
	for i, input := range parsed.Constructor.Inputs {
		arguments[i] = Argument{
			Name: input.Name,
			Type: input.Type.String(),
			Value: func() string {
				if input.Type.T == abi.BytesTy {
					return fmt.Sprintf("0x%v", common.Bytes2Hex(unpacked[i].([]byte)))
				}

				if input.Type.T == abi.FixedBytesTy {
					v := reflect.ValueOf(unpacked[i])
					if v.Kind() != reflect.Array {
						return ""
					}
					b := make([]byte, v.Len())
					for i := 0; i < v.Len(); i++ {
						b[i] = byte(v.Index(i).Uint())
					}
					return fmt.Sprintf("0x%v", common.Bytes2Hex(b))
				}

				if input.Type.T == abi.AddressTy {
					return unpacked[i].(common.Address).Hex()
				}

				return fmt.Sprintf("%v", unpacked[i])
			}(),
			Indexed: input.Indexed,
		}
	}

	return &Constructor{
		Abi:          constructorAbi,
		SignatureRaw: parsed.Constructor.String(),
		Arguments:    arguments,
	}, nil
}
