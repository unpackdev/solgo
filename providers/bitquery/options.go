package bitquery

type Options struct {
	Endpoint string `mapstructure:"endpoint" yaml:"endpoint" json:"endpoint"`
	Key      string `mapstructure:"key" yaml:"key" json:"key"`
}
