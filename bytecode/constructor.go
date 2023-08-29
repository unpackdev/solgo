package bytecode

import (
	"fmt"
	"reflect"
	"strings"

	abi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Argument represents a single argument in a contract constructor.
// It includes the argument's name, type, value, and whether it is indexed.
type Argument struct {
	Name    string `json:"name"`    // Name of the argument
	Type    string `json:"type"`    // Type of the argument
	Value   string `json:"value"`   // Value of the argument
	Indexed bool   `json:"indexed"` // Indicates if the argument is indexed
}

// Constructor represents a contract constructor.
// It includes the ABI of the constructor, the raw signature, and the arguments.
type Constructor struct {
	Abi               string        `json:"abi"`           // ABI of the constructor
	Parsed            abi.ABI       `json:"-"`             // Parsed ABI of the constructor
	SignatureRaw      string        `json:"signature_raw"` // Raw signature of the constructor
	Arguments         []Argument    // List of arguments in the constructor
	UnpackedArguments []interface{} `json:"unpacked_arguments"` // List of unpacked arguments in the constructor
}

func (c *Constructor) Pack() ([]byte, error) {
	return c.Parsed.Constructor.Inputs.PackValues(c.UnpackedArguments)
}

// DecodeConstructorFromAbi decodes the constructor from the provided ABI and bytecode.
// It returns a Constructor object and an error if any occurred during the decoding process.
//
// The function first checks if the bytecode is empty or does not start with '['. If so, it prepends '[' to the constructorAbi.
// Then it attempts to parse the ABI using the abi.JSON function from the go-ethereum library.
// If the ABI parsing is successful, it unpacks the values from the bytecode using the UnpackValues function.
// It then checks if the number of unpacked values matches the number of inputs in the constructor.
// If they match, it creates an Argument object for each input and adds it to the arguments slice.
// Finally, it returns a Constructor object containing the ABI, raw signature, and arguments.
func DecodeConstructorFromAbi(bytecode []byte, constructorAbi string) (*Constructor, error) {
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
		Abi:               constructorAbi,
		Parsed:            parsed,
		SignatureRaw:      parsed.Constructor.String(),
		Arguments:         arguments,
		UnpackedArguments: unpacked,
	}, nil
}
