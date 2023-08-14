package audit

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeArguments(t *testing.T) {
	config := &Config{}

	tests := []struct {
		name    string
		args    []string
		want    []string
		wantErr string
	}{
		{
			name:    "Valid Arguments",
			args:    []string{"--json", "-"},
			want:    []string{"--json", "-"},
			wantErr: "",
		},
		{
			name:    "Invalid Argument",
			args:    []string{"--json", "--invalid"},
			want:    nil,
			wantErr: "invalid argument: --invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := config.SanitizeArguments(tt.args)
			assert.Equal(t, tt.want, got)
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr string
	}{
		{
			name:    "Valid Arguments",
			args:    []string{"--json", "-"},
			wantErr: "",
		},
		{
			name:    "Missing Required Argument",
			args:    []string{"--json"},
			wantErr: "missing required argument: -",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &Config{Arguments: tt.args}
			err := config.Validate()
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}

func TestConfigFunctions(t *testing.T) {
	config := &Config{Arguments: []string{"--json"}}

	// Test GetTempDir
	assert.Equal(t, config.GetTempDir(), "")

	// Test SetArguments
	newArgs := []string{"-"}
	config.SetArguments(newArgs)
	assert.Equal(t, config.GetArguments(), newArgs)

	// Test AppendArguments
	appendArgs := []string{"--json"}
	config.AppendArguments(appendArgs)
	assert.Equal(t, config.GetArguments(), []string{"-", "--json"})
}

func TestNewDefaultConfig(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := os.TempDir()

	_, err := NewDefaultConfig(tempDir)
	assert.NoError(t, err)
}
