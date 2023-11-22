package etherscan

type Options struct {
	Provider ProviderType `json:"provider" yaml:"provider" mapstructure:"provider"`
	Endpoint string       `json:"endpoint" yaml:"endpoint" mapstructure:"endpoint"`
	Keys     []string     `json:"keys" yaml:"keys" mapstructure:"keys"`
}
