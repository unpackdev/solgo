package eip

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEIPStorage(t *testing.T) {

	// Register all EIPs.
	assert.NoError(t, LoadStandards())

	tests := []struct {
		name           string
		eip            EIP
		expectedExists bool
		isStagnant     bool
		expectedError  string
	}{
		{
			name:           "Test EIP20",
			eip:            NewEip20(),
			expectedExists: true,
			expectedError:  "standard EIP20 already exists",
		},
		{
			name:           "Test EIP721",
			eip:            NewEip721(),
			expectedExists: true,
			expectedError:  "standard EIP721 already exists",
		},
		{
			name:           "Test EIP1155",
			eip:            NewEip1155(),
			expectedExists: true,
			expectedError:  "standard EIP1155 already exists",
		},
		{
			name:           "Test EIP1820",
			eip:            NewEip1820(),
			expectedExists: true,
			expectedError:  "standard EIP1820 already exists",
		},
		{
			name:           "Test EIP1822",
			eip:            NewEip1822(),
			expectedExists: true,
			isStagnant:     true,
			expectedError:  "standard EIP1822 already exists",
		},
		{
			name:           "Test EIP1967",
			eip:            NewEip1967(),
			expectedExists: true,
			expectedError:  "standard EIP1967 already exists",
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

			// Test GetUrl
			assert.NotEmpty(t, tt.eip.GetUrl())

			// Test GetFunctions
			assert.NotEmpty(t, tt.eip.GetFunctions())

			// Test GetEvents
			assert.NotEmpty(t, tt.eip.GetEvents())

			// Test GetStandard
			assert.NotNil(t, tt.eip.GetStandard())

			// Test TokenCount
			assert.GreaterOrEqual(t, tt.eip.TokenCount(), 0)

			// Test IsStagnant
			assert.Equal(t, tt.isStagnant, tt.eip.IsStagnant())

			// Test ToProto
			assert.NotNil(t, tt.eip.ToProto())

			// Test String representation
			assert.NotEmpty(t, tt.eip.String())

			// Test RegisterStandard
			err := RegisterStandard(tt.eip.GetType(), tt.eip)
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

			standard, bool := GetStandard(tt.eip.GetType())
			assert.True(t, bool, "standard %v not found in registered standards", tt.eip.GetType())
			assert.NotNil(t, standard)
			assert.Equal(t, tt.eip.GetStandard(), standard.GetStandard())

			assert.NotNil(t, GetSortedRegisteredStandards())
			assert.True(t, StandardsLoaded())
		})
	}
}
