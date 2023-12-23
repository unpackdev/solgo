package exchanges

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestExchangeManager(t *testing.T) {
	tAssert := assert.New(t)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientOptions := &clients.Options{
		Nodes: []clients.Node{
			{
				Group:                   string(utils.Ethereum),
				Type:                    "mainnet",
				Endpoint:                "https://ethereum.publicnode.com",
				NetworkId:               1,
				ConcurrentClientsNumber: 1,
			},
		},
	}

	pool, err := clients.NewClientPool(ctx, clientOptions)
	tAssert.NoError(err)
	tAssert.NotNil(pool)

	manager, err := NewManager(ctx, pool, DefaultOptions())
	tAssert.NoError(err)
	tAssert.NotNil(manager)
}

func TestUniswapV2Exchange(t *testing.T) {
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
				Endpoint:                "https://ethereum.publicnode.com",
				NetworkId:               1,
				ConcurrentClientsNumber: 1,
			},
		},
	}

	pool, err := clients.NewClientPool(ctx, clientOptions)
	tAssert.NoError(err)
	tAssert.NotNil(pool)

	manager, err := NewManager(ctx, pool, DefaultOptions())
	tAssert.NoError(err)
	tAssert.NotNil(manager)

	exchange, found := manager.GetExchange(utils.UniswapV2)
	tAssert.True(found)
	tAssert.NotNil(exchange)

	// Lets cast the exchange to a UniswapV2Exchange so we can properly use it...
	// I want to avoid using generics as they are not yet where I want them to be. Therefore, we are going to
	// type cast interface into a proper type and using from there.
	// What you can do is have a manager of your own, a small one that would have all of necessary castings done in the
	// beginning and then you can use it as you wish with minimal impact on the performance.
	uniswapv2 := ToUniswapV2(exchange)
	tAssert.NotNil(uniswapv2)
	tAssert.IsType(&UniswapV2Exchange{}, uniswapv2)

}
