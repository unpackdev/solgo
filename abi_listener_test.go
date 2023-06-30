package solgo

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbiListener(t *testing.T) {
	// Define multiple test cases
	testCases := []struct {
		name     string
		contract string
		expected string
	}{
		{
			name:     "Dummy Contract",
			contract: ReadContractFileForTest(t, "Dummy").Content,
			expected: `[{"inputs":null,"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"x","type":"function","stateMutability":"view"},{"inputs":null,"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"y","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"uint256","name":"_x","type":"uint256"},{"internalType":"uint256","name":"_y","type":"uint256"}],"type":"constructor"}]`,
		},
		{
			name:     "Complex ERC20",
			contract: ReadContractFileForTest(t, "Complex").Content,
			expected: `[{"inputs":[{"internalType":"uint256","name":"available","type":"uint256"},{"internalType":"uint256","name":"required","type":"uint256"}],"name":"InsufficientBalance","type":"error","stateMutability":"view"},{"inputs":null,"outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"MINTER_ROLE","type":"function","stateMutability":"view"},{"inputs":null,"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"subscriptionAmount","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"subscriptionBalance","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"outputs":[{"internalType":"bool","name":"","type":"bool"}],"name":"isSubscribed","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"rewards","type":"function","stateMutability":"view"},{"inputs":null,"outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"name":"rewardedUsers","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"address","name":"","type":"address"},{"internalType":"address","name":"","type":"address"}],"outputs":[{"internalType":"bool","name":"","type":"bool"}],"name":"nestedMappingTest","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"string","name":"name","type":"string"},{"internalType":"string","name":"symbol","type":"string"},{"internalType":"uint256","name":"initialSupply","type":"uint256"},{"internalType":"uint256","name":"_subscriptionAmount","type":"uint256"}],"outputs":[],"name":"initialize","type":"function","stateMutability":"nonpayable"},{"inputs":[],"outputs":[],"name":"pause","type":"function","stateMutability":"nonpayable"},{"inputs":[],"outputs":[],"name":"purchaseSubscription","type":"function","stateMutability":"nonpayable"},{"inputs":[],"outputs":[],"name":"cancelSubscription","type":"function","stateMutability":"nonpayable"},{"inputs":[{"internalType":"address","name":"user","type":"address"}],"outputs":[{"internalType":"bool","name":"","type":"bool"}],"name":"getSubscriptionStatus","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"address","name":"user","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"outputs":[],"name":"reward","type":"function","stateMutability":"nonpayable"},{"inputs":[{"internalType":"address","name":"user","type":"address"}],"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"getRewards","type":"function","stateMutability":"view"},{"inputs":[],"outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"name":"getRewardedUsers","type":"function","stateMutability":"view"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"outputs":[],"name":"_beforeTokenTransfer","type":"function","stateMutability":"nonpayable"},{"inputs":[{"internalType":"uint256","name":"newAmount","type":"uint256"}],"outputs":[],"name":"updateSubscriptionAmount","type":"function","stateMutability":"nonpayable"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"outputs":[],"name":"mint","type":"function","stateMutability":"nonpayable"},{"type":"fallback","stateMutability":"payable"},{"type":"receive","stateMutability":"payable"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"user","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"SubscriptionPurchased","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"user","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"SubscriptionCanceled","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"user","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"UserRewarded","type":"event"},{"anonymous":false,"inputs":[{"internalType":"uint256","name":"a","type":"uint256"},{"internalType":"bytes32","name":"b","type":"bytes32"}],"name":"Event2","type":"event"}]`,
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

			//
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
