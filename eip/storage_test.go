package eip

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEIPFunctions(t *testing.T) {
	tests := []struct {
		name           string
		eip            EIP
		expectedExists bool
		expectedError  string
	}{
		{
			name:           "Test ERC20",
			eip:            NewErc20(),
			expectedExists: true,
			expectedError:  "",
		},
		{
			name:           "Test ERC721",
			eip:            NewErc721(),
			expectedExists: true,
			expectedError:  "",
		},
		{
			name:           "Test ERC1155",
			eip:            NewErc1155(),
			expectedExists: true,
			expectedError:  "",
		},
		{
			name:           "Test EIP1820",
			eip:            NewEip1820(),
			expectedExists: true,
			expectedError:  "",
		},
		{
			name:           "Test EIP1822",
			eip:            NewEip1822(),
			expectedExists: true,
			expectedError:  "",
		},
		{
			name:           "Test EIP1967",
			eip:            NewEip1967(),
			expectedExists: true,
			expectedError:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if all methods of the EIP interface are implemented
			eipType := reflect.TypeOf((*EIP)(nil)).Elem()
			assert.True(t, reflect.TypeOf(tt.eip).Implements(eipType), "EIP methods not fully implemented")

			// Test GetName
			assert.NotEmpty(t, tt.eip.GetName())

			// Test GetType
			assert.NotNil(t, tt.eip.GetType())

			// Test GetFunctions
			assert.NotEmpty(t, tt.eip.GetFunctions())

			// Test GetEvents
			assert.NotEmpty(t, tt.eip.GetEvents())

			// Test GetStandard
			assert.NotNil(t, tt.eip.GetStandard())

			// Test TokenCount
			assert.GreaterOrEqual(t, tt.eip.TokenCount(), 0)

			// Test ToProto
			assert.NotNil(t, tt.eip.ToProto())

			// Test String representation
			assert.NotEmpty(t, tt.eip.String())

			// Test RegisterStandard
			err := RegisterStandard(tt.eip.GetType(), tt.eip.GetStandard())
			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}

			// Test Exists
			exists := Exists(tt.eip.GetType())
			assert.Equal(t, tt.expectedExists, exists)

			// Test GetRegisteredStandards
			registeredStandards := GetRegisteredStandards()
			_, ok := registeredStandards[tt.eip.GetType()]
			assert.True(t, ok, "standard %v not found in registered standards", tt.eip.GetType())
		})
	}
}
