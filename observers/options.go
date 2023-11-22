package observers

import (
	"math/big"

	"github.com/unpackdev/solgo/utils"
)

type WorkerOptions struct {
	WorkerType        string `json:"worker_type" yaml:"worker_type" mapstructure:"worker_type"`
	WorkerCount       int    `json:"worker_count" yaml:"worker_count" mapstructure:"worker_count"`
	WorkerChannelSize int    `json:"worker_channel_size" yaml:"worker_channel_size" mapstructure:"worker_channel_size"`
}

type Options struct {
	NetworkID  utils.NetworkID  `json:"network_id" yaml:"network_id" mapstructure:"network_id"`
	Network    utils.Network    `json:"network" yaml:"network" mapstructure:"network"`
	Strategies []utils.Strategy `json:"strategies" yaml:"strategies" mapstructure:"strategies"`
	StartBlock *big.Int         `json:"start_block" yaml:"start_block" mapstructure:"start_block"`
	EndBlock   *big.Int         `json:"end_block" yaml:"end_block" mapstructure:"end_block"`
	Workers    []WorkerOptions  `json:"workers" yaml:"workers" mapstructure:"workers"`
}
