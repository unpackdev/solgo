package contracts

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo"
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
		expected ContractInfo
	}{
		{
			name:     "Empty Contract",
			contract: tests.ReadContractFileForTest(t, "Empty").Content,
			expected: ContractInfo{},
		},
		{
			name:     "Contract Without Interfaces",
			contract: tests.ReadContractFileForTest(t, "NoInterfaces").Content,
			expected: ContractInfo{
				Name: "NoInterfaces",
				Pragmas: []string{
					"solidity ^0.8.5",
				},
				License: "MIT",
			},
		},
		{
			name:     "Contract Without Imports",
			contract: tests.ReadContractFileForTest(t, "NoImports").Content,
			expected: ContractInfo{
				Name: "NoImports",
				Pragmas: []string{
					"solidity ^0.8.5",
				},
				License: "MIT",
			},
		},
		{
			name:     "Contract Without Pragmas",
			contract: tests.ReadContractFileForTest(t, "NoPragmas").Content,
			expected: ContractInfo{
				Name:    "NoPragmas",
				License: "MIT",
			},
		},
		{
			name:     "Contract With Single-Line Comment",
			contract: tests.ReadContractFileForTest(t, "SingleLineComment").Content,
			expected: ContractInfo{
				Comments: []string{
					"// This is a single-line comment",
				},
				Name: "SingleLineComment",
				Pragmas: []string{
					"solidity ^0.8.5",
				},
				License: "MIT",
			},
		},
		{
			name:     "Contract With Multi-Line Comment",
			contract: tests.ReadContractFileForTest(t, "MultiLineComment").Content,
			expected: ContractInfo{
				Comments: []string{
					"/* This is a\n multi-line comment */",
				},
				Name: "MultiLineComment",
				Pragmas: []string{
					"solidity ^0.8.5",
				},
				License: "MIT",
			},
		},
		{
			name:     "Contract With Different SPDX License Identifier",
			contract: tests.ReadContractFileForTest(t, "DifferentLicense").Content,
			expected: ContractInfo{
				License: "GPL-3.0",
				Name:    "DifferentLicense",
				Pragmas: []string{
					"solidity ^0.8.5",
				},
			},
		},
		{
			name:     "Dummy Contract",
			contract: tests.ReadContractFileForTest(t, "Dummy").Content,
			expected: ContractInfo{
				Comments: []string{
					"// Some additional comments that can be extracted",
					"/** \n * Multi line comments\n * are supported as well\n*/",
				},
				License: "MIT",
				Name:    "Dummy",
				Pragmas: []string{
					"solidity ^0.8.5",
				},
			},
		},
		{
			name:     "OpenZeppelin ERC20",
			contract: tests.ReadContractFileForTest(t, "ERC20_Token").Content,
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
				IsProxy:         true,
				ProxyConfidence: 100,
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
