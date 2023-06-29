package solgo

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContractListener(t *testing.T) {
	// Define multiple test cases
	testCases := []struct {
		name     string
		contract string
		expected ContractInfo
	}{
		{
			name:     "Empty Contract",
			contract: ReadContractFileForTest(t, "Empty").Content,
			expected: ContractInfo{},
		},
		{
			name:     "Dummy Contract",
			contract: ReadContractFileForTest(t, "Dummy").Content,
			expected: ContractInfo{
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
			contract: ReadContractFileForTest(t, "MyToken").Content,
			expected: ContractInfo{
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
			parser, err := New(context.TODO(), strings.NewReader(testCase.contract))
			assert.NoError(t, err)
			assert.NotNil(t, parser)

			// Register the contract information listener
			contractListener := NewContractListener(parser.GetParser())
			err = parser.RegisterListener(ListenerContractInfo, contractListener)
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
