package storage

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/providers/etherscan"
	"github.com/unpackdev/solgo/simulator"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestStorage(t *testing.T) {
	tAssert := assert.New(t)

	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	tAssert.NoError(err)
	zap.ReplaceGlobals(logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientOptions := &clients.Options{
		Nodes: []clients.Node{
			{
				Group:                   string(utils.Ethereum),
				Type:                    "mainnet",
				Endpoint:                "ws://localhost:8545",
				NetworkId:               1,
				ConcurrentClientsNumber: 1,
			},
		},
	}

	pool, err := clients.NewClientPool(ctx, clientOptions)
	tAssert.NoError(err)
	tAssert.NotNil(pool)

	simulator, err := simulator.NewSimulator(ctx, pool, nil)
	tAssert.NoError(err)
	tAssert.NotNil(simulator)

	etherscanProvider := etherscan.NewEtherScanProvider(ctx, nil, &etherscan.Options{
		Provider: etherscan.EtherScan,
		Endpoint: "https://api.etherscan.io/api",
		Keys: []string{
			"Q54IYAWC4NE6XC9Q9YNTUM54U9YGI8QGUY", //unpackdev
			"DW4CPPGR8TWYYNAZQDEKT332UJI9R68C8S", //nevio
			"NYH2MKD67GGQAMG8UXDVYBGX2K6J9HMS7E", //inorbit
		},
	})

	bindManager, err := bindings.NewManager(ctx, pool)
	tAssert.NoError(err)
	tAssert.NotNil(bindManager)

	storage, err := NewStorage(ctx, utils.Ethereum, pool, simulator, etherscanProvider, nil, bindManager, NewDefaultOptions())
	tAssert.NoError(err)
	tAssert.NotNil(storage)

	// Now we have base prepared and we can start with actual testing scenarios...
	uniswapV3Factory := common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")

	// Look up for storage descriptor at latest block...
	descriptor, err := storage.GetStorageDescriptor(ctx, uniswapV3Factory, nil)
	tAssert.NoError(err)
	tAssert.NotNil(descriptor)

}
