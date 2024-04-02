package etherscan

import "errors"

// Options holds the configuration settings for an etherscan client.
// These settings define how the client interacts with the blockchain explorer APIs.
type Options struct {
	// Provider specifies the blockchain explorer service (e.g., Etherscan, BscScan) to be used.
	Provider ProviderType `json:"provider" yaml:"provider" mapstructure:"provider"`

	// Endpoint is the base URL of the blockchain explorer API.
	// It determines where the client sends its requests.
	Endpoint string `json:"endpoint" yaml:"endpoint" mapstructure:"endpoint"`

	// RateLimit specifies the maximum number of requests that the client is allowed to make to the
	// blockchain explorer API within a fixed time window. Consult Etherscan documentation.
	RateLimit int `json:"rateLimit" yaml:"rateLimit" mapstructure:"rateLimit"`

	// Keys contains a list of API keys used for authenticating requests to the blockchain explorer API.
	// The client can rotate through these keys to manage rate limits.
	Keys []string `json:"keys" yaml:"keys" mapstructure:"keys"`
}

// Validate checks the integrity and completeness of the Options settings.
// It ensures that all necessary configurations are properly set and valid,
// including non-empty Endpoint, at least one API key, and a supported Provider.
func (o *Options) Validate() error {
	if o.Endpoint == "" {
		return errors.New("endpoint is required but not set")
	}

	if len(o.Keys) == 0 {
		return errors.New("at least one API key is required")
	}

	if o.RateLimit == 0 {
		return errors.New("rate limit needs to be specified")
	}

	return nil
}
