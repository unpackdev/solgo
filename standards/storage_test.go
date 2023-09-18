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
			name: "Test EIP20",
			standard: func() EIP {
				standard, err := GetContractByStandard(EIP20)
				assert.NoError(t, err)
				assert.NotNil(t, standard)
				return standard
			}(),
			expectedExists: true,
			expectedError:  "standard EIP20 already exists",
		},
		{
			name: "Test EIP721",
			standard: func() EIP {
				standard, err := GetContractByStandard(EIP721)
				assert.NoError(t, err)
				assert.NotNil(t, standard)
				return standard
			}(),
			expectedExists: true,
			expectedError:  "standard EIP721 already exists",
		},
		{
			name: "Test EIP1155",
			standard: func() EIP {
				standard, err := GetContractByStandard(EIP1155)
				assert.NoError(t, err)
				assert.NotNil(t, standard)
				return standard
			}(),
			expectedExists: true,
			expectedError:  "standard EIP1155 already exists",
		},
		{
			name: "Test EIP1820",
			standard: func() EIP {
				standard, err := GetContractByStandard(EIP1820)
				assert.NoError(t, err)
				assert.NotNil(t, standard)
				return standard
			}(),
			expectedExists: true,
			expectedError:  "standard EIP1820 already exists",
		},
		{
			name: "Test EIP1822",
			standard: func() EIP {
				standard, err := GetContractByStandard(EIP1822)
				assert.NoError(t, err)
				assert.NotNil(t, standard)
				return standard
			}(),
			expectedExists: true,
			isStagnant:     true,
			expectedError:  "standard EIP1822 already exists",
		},
		{
			name: "Test EIP1967",
			standard: func() EIP {
				standard, err := GetContractByStandard(EIP1967)
				assert.NoError(t, err)
				assert.NotNil(t, standard)
				return standard
			}(),
			expectedExists: true,
			expectedError:  "standard EIP1967 already exists",
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
