package audit

type Config struct {
	tempDir string // temp directory to store temporary contract files
}

// GetTempDir returns the temp directory to store temporary contract files
func (c *Config) GetTempDir() string {
	return c.tempDir
}
