package solc

import (
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
			args:    []string{"--optimize-runs", "-"},
			want:    []string{"--optimize-runs", "-"},
			wantErr: "",
		},
		{
			name:    "Invalid Argument",
			args:    []string{"--optimize", "--invalid"},
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
			args:    []string{"--overwrite", "--combined-json", "--optimize", "200", "-"},
			wantErr: "",
		},
		{
			name:    "Missing Required Argument",
			args:    []string{"--overwrite", "--combined-json"},
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

	// Test SetArguments
	newArgs := []string{"-"}
	config.SetArguments(newArgs)
	assert.Equal(t, config.GetArguments(), newArgs)

	// Test AppendArguments
	appendArgs := []string{"--json"}
	config.AppendArguments(appendArgs...)
	assert.Equal(t, config.GetArguments(), []string{"-", "--json"})
}

func TestNewDefaultConfig(t *testing.T) {
	_, err := NewDefaultConfig()
	assert.NoError(t, err)
}
