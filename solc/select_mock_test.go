package solc

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockUse(t *testing.T) {
	tests := []struct {
		name           string
		version        string
		expectedResult bool
		expectedError  error
	}{
		{
			name:           "Successful switch",
			version:        "0.8.19",
			expectedResult: true,
			expectedError:  nil,
		},
		{
			name:           "Unsuccessful switch",
			version:        "0.8.20",
			expectedResult: false,
			expectedError:  errors.New("switch error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Select{
				commander: &MockCommand{},
			}

			result, _, _, err := s.Use(tt.version)
			assert.Equal(t, tt.expectedResult, result)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestMockInstall(t *testing.T) {
	tests := []struct {
		name           string
		version        string
		expectedResult bool
		expectedError  error
	}{
		{
			name:           "Successful install",
			version:        "0.8.19",
			expectedResult: true,
			expectedError:  nil,
		},
		{
			name:           "Unsuccessful install",
			version:        "0.8.20",
			expectedResult: false,
			expectedError:  errors.New("installation error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Select{
				commander: &MockCommand{},
			}

			result, _, err := s.Install(tt.version)
			assert.Equal(t, tt.expectedResult, result)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestMockVersions(t *testing.T) {
	tests := []struct {
		name             string
		expectedVersions []Version
		expectedError    error
	}{
		{
			name: "Successful retrieval of versions",
			expectedVersions: []Version{
				{Release: "0.8.19", Current: true},
				{Release: "0.8.18", Current: false},
				// ... add other versions as needed ...
			},
			expectedError: nil,
		},
		// Note: The MockCommand always returns the versions 0.8.19 (current) and 0.8.18.
		// If there are scenarios where the retrieval can fail or return different versions, you should add them here.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Select{
				current: "0.8.19",
				commander: &MockCommand{
					current: "0.8.19",
				},
			}

			versions, err := s.Versions()
			assert.Equal(t, tt.expectedVersions, versions)
			assert.Equal(t, tt.expectedError, err)
			assert.NotEmpty(t, s.Current())
			assert.NotEmpty(t, s.commander.Current())
		})
	}
}

func TestMockUpgrade(t *testing.T) {
	tests := []struct {
		name           string
		expectedResult bool
		expectedError  error
	}{
		{
			name:           "Successful upgrade",
			expectedResult: true,
			expectedError:  nil,
		},
		// Note: The MockCommand always returns a successful upgrade.
		// If there are scenarios where the upgrade can fail, you should add them here.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Select{
				commander: &MockCommand{},
			}

			result, err := s.Upgrade()
			assert.Equal(t, tt.expectedResult, result)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
