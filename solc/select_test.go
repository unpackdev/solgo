package solc

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUse(t *testing.T) {
	tests := []struct {
		name           string
		version        string
		expectedResult bool
		expectedError  error
	}{
		{
			name:           "Successful switch",
			version:        "0.8.20",
			expectedResult: true,
			expectedError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if os.Getenv("TEST_SOLGO_SOLC_SELECT_DISABLED") == "true" {
				t.Skip("Skipping test that requires solc-select")
			}

			s, err := NewSelect()
			assert.NoError(t, err, "NewSelect() should not return an error")

			result, _, _, err := s.Use(tt.version)
			assert.Equal(t, tt.expectedResult, result)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestInstall(t *testing.T) {
	tests := []struct {
		name           string
		version        string
		expectedResult bool
		expectedError  error
	}{
		{
			name:           "Successful install",
			version:        "0.8.20",
			expectedResult: true,
			expectedError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if os.Getenv("TEST_SOLGO_SOLC_SELECT_DISABLED") == "true" {
				t.Skip("Skipping test that requires solc-select")
			}

			s, err := NewSelect()
			assert.NoError(t, err, "NewSelect() should not return an error")
			result, _, err := s.Install(tt.version)
			assert.Equal(t, tt.expectedResult, result)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestVersions(t *testing.T) {
	tests := []struct {
		name          string
		expectedError error
	}{
		{
			name:          "Successful retrieval of versions",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if os.Getenv("TEST_SOLGO_SOLC_SELECT_DISABLED") == "true" {
				t.Skip("Skipping test that requires solc-select")
			}

			s, err := NewSelect()
			assert.NoError(t, err, "NewSelect() should not return an error")
			versions, err := s.Versions()
			assert.GreaterOrEqual(t, len(versions), 1)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestUpgrade(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if os.Getenv("TEST_SOLGO_SOLC_SELECT_DISABLED") == "true" {
				t.Skip("Skipping test that requires solc-select")
			}

			s, err := NewSelect()
			assert.NoError(t, err, "NewSelect() should not return an error")
			result, err := s.Upgrade()
			assert.Equal(t, tt.expectedResult, result)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestNewSelect(t *testing.T) {
	t.Skip("Skipping test that requires solc-select to be installed")
	tests := []struct {
		name             string
		expectedError    error
		expectCurrentSet bool
	}{
		{
			name:             "Successful initialization with solc-select installed",
			expectedError:    nil,
			expectCurrentSet: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if os.Getenv("TEST_SOLGO_SOLC_SELECT_DISABLED") == "true" {
				t.Skip("Skipping test that requires solc-select")
			}

			s, err := NewSelect()
			assert.NoError(t, err, "NewSelect() should not return an error")

			// If we expect the current version to be set, check it
			if tt.expectCurrentSet {
				assert.NotEmpty(t, s.Current())
			}
		})
	}
}
