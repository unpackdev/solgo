package solgo

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestAbiListener(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

	// Define multiple test cases
	testCases := []struct {
		name     string
		contract string
		expected string
	}{
		{
			name:     "Dummy Contract",
			contract: ReadContractFileForTest(t, "Dummy").Content,
			expected: `[{"inputs":[],"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"x","type":"function","stateMutability":"view"},{"inputs":[],"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"y","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"uint256","name":"_x","type":"uint256"},{"internalType":"uint256","name":"_y","type":"uint256"}],"type":"constructor"}]`,
		},
		{
			name:     "ERC20 Token",
			contract: ReadContractFileForTest(t, "ERC20_Token").Content,
			expected: `[{"inputs":[],"outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"MINTER_ROLE","type":"function","stateMutability":"view"},{"inputs":[],"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"subscriptionAmount","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"subscriptionBalance","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"outputs":[{"internalType":"bool","name":"","type":"bool"}],"name":"isSubscribed","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"rewards","type":"function","stateMutability":"view"},{"inputs":[],"outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"name":"rewardedUsers","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"string","name":"name","type":"string"},{"internalType":"string","name":"symbol","type":"string"},{"internalType":"uint256","name":"initialSupply","type":"uint256"},{"internalType":"uint256","name":"_subscriptionAmount","type":"uint256"}],"outputs":[],"name":"initialize","type":"function","stateMutability":"nonpayable"},{"inputs":[],"outputs":[],"name":"pause","type":"function","stateMutability":"nonpayable"},{"inputs":[],"outputs":[],"name":"unpause","type":"function","stateMutability":"nonpayable"},{"inputs":[],"outputs":[],"name":"purchaseSubscription","type":"function","stateMutability":"nonpayable"},{"inputs":[],"outputs":[],"name":"cancelSubscription","type":"function","stateMutability":"nonpayable"},{"inputs":[{"internalType":"address","name":"user","type":"address"}],"outputs":[{"internalType":"bool","name":"","type":"bool"}],"name":"getSubscriptionStatus","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"address","name":"user","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"outputs":[],"name":"reward","type":"function","stateMutability":"nonpayable"},{"inputs":[{"internalType":"address","name":"user","type":"address"}],"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"getRewards","type":"function","stateMutability":"view"},{"inputs":[],"outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"name":"getRewardedUsers","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"outputs":[],"name":"_beforeTokenTransfer","type":"function","stateMutability":"nonpayable"},{"inputs":[{"internalType":"uint256","name":"newAmount","type":"uint256"}],"outputs":[],"name":"updateSubscriptionAmount","type":"function","stateMutability":"nonpayable"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"outputs":[],"name":"mint","type":"function","stateMutability":"nonpayable"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"user","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"SubscriptionPurchased","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"user","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"SubscriptionCanceled","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"user","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"UserRewarded","type":"event"}]`,
		},
		{
			name:     "Mappings",
			contract: ReadContractFileForTest(t, "Mappings").Content,
			expected: `[{"inputs":[{"internalType":"address","name":"","type":"address"}],"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"simpleMapping","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"address","name":"","type":"address"},{"internalType":"address","name":"","type":"address"}],"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"doubleMapping","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"address","name":"","type":"address"},{"internalType":"address","name":"","type":"address"},{"internalType":"address","name":"","type":"address"}],"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"tripleMapping","type":"function","stateMutability":"view"}]`,
		},
		{
			name:     "Structs",
			contract: ReadContractFileForTest(t, "Structs").Content,
			expected: `[{"inputs":[{"internalType":"struct MyStructs.ClasicStruct","name":"ClasicStruct","type":"tuple","components":[{"internalType":"uint256","name":"one","type":"uint256"},{"internalType":"uint256","name":"two","type":"uint256"}]},{"internalType":"struct MyStructs.NestedStruct","name":"NestedStruct","type":"tuple","components":[{"internalType":"uint256","name":"one","type":"uint256"},{"internalType":"uint256","name":"two","type":"uint256"},{"internalType":"struct MyStructs.ClasicStruct","name":"myStruct","type":"tuple","components":[{"internalType":"uint256","name":"one","type":"uint256"},{"internalType":"uint256","name":"two","type":"uint256"}]}]},{"internalType":"uint256","name":"Integer","type":"uint256"}],"outputs":[{"internalType":"struct MyStructs.NestedStruct","name":"NestedStruct","type":"tuple","components":[{"internalType":"uint256","name":"one","type":"uint256"},{"internalType":"uint256","name":"two","type":"uint256"},{"internalType":"struct MyStructs.ClasicStruct","name":"myStruct","type":"tuple","components":[{"internalType":"uint256","name":"one","type":"uint256"},{"internalType":"uint256","name":"two","type":"uint256"}]}]}],"name":"nestedStructExample","type":"function","stateMutability":"pure"},{"inputs":[{"internalType":"struct MyStructs.StructWithArray","name":"StructWithArray","type":"tuple","components":[{"internalType":"uint256","name":"one","type":"uint256"},{"internalType":"uint256[]","name":"array","type":"uint256[]"}]}],"outputs":[],"name":"structWithArrayExample","type":"function","stateMutability":"pure"},{"inputs":[{"internalType":"struct MyStructs.StructWithNestedArray","name":"StructWithNestedArray","type":"tuple","components":[{"internalType":"uint256","name":"one","type":"uint256"},{"internalType":"struct MyStructs.ClasicStruct[]","name":"structArray","type":"tuple[]","components":[{"internalType":"uint256","name":"one","type":"uint256"},{"internalType":"uint256","name":"two","type":"uint256"}]}]}],"outputs":[],"name":"structWithNestedArrayExample","type":"function","stateMutability":"pure"}]`,
		},
		{
			name:     "Enums",
			contract: ReadContractFileForTest(t, "Enums").Content,
			expected: `[{"inputs":[],"outputs":[{"internalType":"enum EnumContract.State","name":"","type":"uint8"}],"name":"state","type":"function","stateMutability":"view"},{"inputs":[],"type":"constructor"},{"inputs":[],"outputs":[{"internalType":"bool","name":"","type":"bool"}],"name":"isWaiting","type":"function","stateMutability":"view"},{"inputs":[],"outputs":[{"internalType":"bool","name":"","type":"bool"}],"name":"isReady","type":"function","stateMutability":"view"},{"inputs":[],"outputs":[{"internalType":"bool","name":"","type":"bool"}],"name":"isActive","type":"function","stateMutability":"view"},{"inputs":[],"outputs":[],"name":"makeReady","type":"function","stateMutability":"nonpayable"},{"inputs":[],"outputs":[],"name":"makeActive","type":"function","stateMutability":"nonpayable"},{"inputs":[],"outputs":[],"name":"reset","type":"function","stateMutability":"nonpayable"}]`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parser, err := New(context.TODO(), strings.NewReader(testCase.contract))
			assert.NoError(t, err)
			assert.NotNil(t, parser)

			// Register the contract information listener
			abiListener := NewAbiListener()
			err = parser.RegisterListener(ListenerAbi, abiListener)
			assert.NoError(t, err)

			syntaxErrs := parser.Parse()
			assert.Empty(t, syntaxErrs)

			abiParser := abiListener.GetParser()

			abiJson, err := abiParser.ToJSON()
			assert.NoError(t, err)
			assert.NotEmpty(t, abiJson)

			abi, err := abiParser.ToABI()
			assert.NoError(t, err)
			assert.NotEmpty(t, abi)

			// Assert the parsed contract matches the expected result
			assert.Equal(t, testCase.expected, abiJson)
		})
	}
}
