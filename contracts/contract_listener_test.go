package contracts

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/common"
	"github.com/txpull/solgo/tests"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestContractListener(t *testing.T) {
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
		expected common.ContractInfo
	}{
		{
			name:     "Empty Contract",
			contract: tests.ReadContractFileForTest(t, "Empty").Content,
			expected: common.ContractInfo{},
		},
		{
			name:     "Dummy Contract",
			contract: tests.ReadContractFileForTest(t, "Dummy").Content,
			expected: common.ContractInfo{
				Comments: nil,
				License:  "MIT",
				Name:     "Dummy",
				Pragmas: []string{
					"solidity ^0.8.5",
				},
			},
		},
		{
			name:     "OpenZeppelin ERC20",
			contract: tests.ReadContractFileForTest(t, "ERC20_Token").Content,
			expected: common.ContractInfo{
				Comments: []string{
					"// Rewards",
					"// Allows users to purchase subscription using the token",
					"// Allows users to cancel subscription and receive a refund",
					"// Gets the subscription status of a user",
					"// Allows the contract owner to reward users with tokens",
					"// If the user hasn't been rewarded before, add them to the list of rewarded users",
					"// Gets the total amount of rewards for a user",
					"// Gets the list of all rewarded users",
					"// Allows the contract owner to update the subscription amount",
					"// Allows minters to mint new tokens",
					"// Events",
				},
				License: "MIT",
				Name:    "MyToken",
				Pragmas: []string{
					"solidity ^0.8.0",
				},
				Imports: []string{
					"@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol",
					"@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol",
					"@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol",
					"@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol",
					"@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol",
				},
				Implements: []string{
					"Initializable",
					"ERC20Upgradeable",
					"AccessControlUpgradeable",
					"PausableUpgradeable",
					"SafeERC20Upgradeable",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parser, err := solgo.New(context.TODO(), strings.NewReader(testCase.contract))
			assert.NoError(t, err)
			assert.NotNil(t, parser)

			// Register the contract information listener
			contractListener := NewContractListener(parser.GetParser())
			err = parser.RegisterListener(solgo.ListenerContractInfo, contractListener)
			assert.NoError(t, err)

			syntaxErrs := parser.Parse()
			assert.Empty(t, syntaxErrs)

			// Get the contract information from listener that is built for testing purposes
			contractInfo := contractListener.ToStruct()

			// Assert the parsed contract matches the expected result
			assert.Equal(t, testCase.expected, contractInfo)
		})
	}
}
