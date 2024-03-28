package bitquery

// Options contains configuration settings for the bitquery client.
// These settings specify how the client connects to the blockchain data provider.
type Options struct {
	// Endpoint specifies the URL of the blockchain data service.
	// This is where the bitquery client will send its requests.
	Endpoint string `mapstructure:"endpoint" yaml:"endpoint" json:"endpoint"`

	// Key is the authentication key required by the blockchain data service.
	// This key is used to authorize the client's requests.
	Key string `mapstructure:"key" yaml:"key" json:"key"`
}
