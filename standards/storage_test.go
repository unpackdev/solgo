package standards

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	eip_pb "github.com/unpackdev/protos/dist/go/eip"
)

func TestEIPStorage(t *testing.T) {
	// Register all EIPs.
	assert.NoError(t, LoadStandards())

	// Test non existing yet standard...
	s, err := GetContractByStandard(Standard("CORRUPTED"))
	assert.Error(t, err)
	assert.Nil(t, s)

	unknown := Standard("NOT_DEFINED_YET")
	assert.Equal(t, unknown.ToProto(), eip_pb.Standard_UNKNOWN)

	tests := []struct {
		name           string
		standard       EIP
		expectedExists bool
		isStagnant     bool
		expectedError  string
	}{
		{
			name: "Test ERC20",
			standard: func() EIP {
				standard, err := GetContractByStandard(ERC20)
				assert.NoError(t, err)
				assert.NotNil(t, standard)
				return standard
			}(),
			expectedExists: true,
			expectedError:  "standard ERC20 already exists",
		},
		{
			name: "Test ERC721",
			standard: func() EIP {
				standard, err := GetContractByStandard(ERC721)
				assert.NoError(t, err)
				assert.NotNil(t, standard)
				return standard
			}(),
			expectedExists: true,
			expectedError:  "standard ERC721 already exists",
		},
		{
			name: "Test ERC1155",
			standard: func() EIP {
				standard, err := GetContractByStandard(ERC1155)
				assert.NoError(t, err)
				assert.NotNil(t, standard)
				return standard
			}(),
			expectedExists: true,
			expectedError:  "standard ERC1155 already exists",
		},
		{
			name: "Test ERC1820",
			standard: func() EIP {
				standard, err := GetContractByStandard(ERC1820)
				assert.NoError(t, err)
				assert.NotNil(t, standard)
				return standard
			}(),
			expectedExists: true,
			expectedError:  "standard ERC1820 already exists",
		},
		{
			name: "Test ERC1822",
			standard: func() EIP {
				standard, err := GetContractByStandard(ERC1822)
				assert.NoError(t, err)
				assert.NotNil(t, standard)
				return standard
			}(),
			expectedExists: true,
			isStagnant:     true,
			expectedError:  "standard ERC1822 already exists",
		},
		{
			name: "Test ERC1967",
			standard: func() EIP {
				standard, err := GetContractByStandard(ERC1967)
				assert.NoError(t, err)
				assert.NotNil(t, standard)
				return standard
			}(),
			expectedExists: true,
			expectedError:  "standard ERC1967 already exists",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if all methods of the EIP interface are implemented
			eipType := reflect.TypeOf((*EIP)(nil)).Elem()
			assert.True(t, reflect.TypeOf(tt.standard).Implements(eipType), "EIP methods not fully implemented")

			// Test GetName
			assert.NotEmpty(t, tt.standard.GetName())

			// Test GetType
			assert.NotNil(t, tt.standard.GetType())

			// Test GetUrl
			assert.NotEmpty(t, tt.standard.GetUrl())

			// Test GetFunctions
			assert.NotEmpty(t, tt.standard.GetFunctions())

			// Test GetEvents
			assert.NotEmpty(t, tt.standard.GetEvents())

			// Test GetStandard
			assert.NotNil(t, tt.standard.GetStandard())

			// Test TokenCount
			assert.GreaterOrEqual(t, tt.standard.TokenCount(), 0)

			// Test IsStagnant
			assert.Equal(t, tt.isStagnant, tt.standard.IsStagnant())

			// Test GetABI
			assert.NotEmpty(t, tt.standard.GetABI(), "ABI is empty")

			// Test ToProto
			assert.NotNil(t, tt.standard.ToProto())

			// Test String representation
			assert.NotEmpty(t, tt.standard.String())

			// Test RegisterStandard
			err := RegisterStandard(tt.standard.GetType(), tt.standard)
			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}

			// Test Exists
			exists := Exists(tt.standard.GetType())
			assert.Equal(t, tt.expectedExists, exists)

			// Test GetRegisteredStandards
			registeredStandards := GetRegisteredStandards()
			_, ok := registeredStandards[tt.standard.GetType()]
			assert.True(t, ok, "standard %v not found in registered standards", tt.standard.GetType())

			standard, bool := GetStandard(tt.standard.GetType())
			assert.True(t, bool, "standard %v not found in registered standards", tt.standard.GetType())
			assert.NotNil(t, standard)
			assert.Equal(t, tt.standard.GetStandard(), standard.GetStandard())

			assert.NotNil(t, GetSortedRegisteredStandards())
			assert.True(t, StandardsLoaded())
		})
	}
}
