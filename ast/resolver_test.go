package ast

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestResolver(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

	// Define multiple test cases
	testCases := []struct {
		name                 string
		sources              solgo.Sources
		expected             string
		unresolvedReferences int64
	}{
		{
			name: "Test Case 1 - Simple Test Contract",
			sources: solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "TestContract",
						Path:    "TestContract.sol",
						Content: "pragma solidity ^0.5.0; contract TestContract { uint256 public value; function setValue(uint256 _value) public { value = _value; } }",
					},
				},
				EntrySourceUnitName: "TestContract",
			},
			expected:             ``,
			unresolvedReferences: 0,
		},
		{
			name: "Test Case 2 - Contract with Inheritance",
			sources: solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "BaseContract",
						Path:    "BaseContract.sol",
						Content: "pragma solidity ^0.5.0; contract BaseContract { uint256 public baseValue; }",
					},
					{
						Name:    "DerivedContract",
						Path:    "DerivedContract.sol",
						Content: "pragma solidity ^0.5.0; import './BaseContract.sol'; contract DerivedContract is BaseContract { uint256 public derivedValue; }",
					},
				},
				EntrySourceUnitName: "DerivedContract",
			},
			expected:             ``,
			unresolvedReferences: 0,
		},
		{
			name: "Test Case 3 - Contract with Library",
			sources: solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "SafeMath",
						Path:    "SafeMath.sol",
						Content: "pragma solidity ^0.5.0; library SafeMath { function add(uint256 a, uint256 b) internal pure returns (uint256) { uint256 c = a + b; require(c >= a, 'SafeMath: addition overflow'); return c; } }",
					},
					{
						Name:    "TestContract",
						Path:    "TestContract.sol",
						Content: "pragma solidity ^0.5.0; import './SafeMath.sol'; contract TestContract { using SafeMath for uint256; uint256 public value; function increaseValue(uint256 _value) public { value = value.add(_value); } }",
					},
				},
				EntrySourceUnitName: "TestContract",
			},
			expected:             ``,
			unresolvedReferences: 0,
		},
		{
			name: "Test Case 4 - Contract with Interface",
			sources: solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "IERC20",
						Path:    "IERC20.sol",
						Content: "pragma solidity ^0.5.0; interface IERC20 { function totalSupply() external view returns (uint256); function balanceOf(address account) external view returns (uint256); function transfer(address recipient, uint256 amount) external returns (bool); }",
					},
					{
						Name:    "TestContract",
						Path:    "TestContract.sol",
						Content: "pragma solidity ^0.5.0; import './IERC20.sol'; contract TestContract is IERC20 { uint256 public totalSupply; mapping(address => uint256) private _balances; function balanceOf(address account) public view returns (uint256) { return _balances[account]; } function transfer(address recipient, uint256 amount) public returns (bool) { _balances[msg.sender] -= amount; _balances[recipient] += amount; return true; } }",
					},
				},
				EntrySourceUnitName: "TestContract",
			},
			expected:             ``,
			unresolvedReferences: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parser, err := solgo.NewParserFromSources(context.TODO(), testCase.sources)
			assert.NoError(t, err)
			assert.NotNil(t, parser)

			astBuilder := NewAstBuilder(
				parser.GetParser(),
				parser.GetSources(),
			)

			err = parser.RegisterListener(solgo.ListenerAst, astBuilder)
			assert.NoError(t, err)

			syntaxErrs := parser.Parse()
			assert.Empty(t, syntaxErrs)

			errs := astBuilder.ResolveReferences()
			assert.Equal(t, int(testCase.unresolvedReferences), len(errs))

			for _, sourceUnit := range astBuilder.GetRoot().GetSourceUnits() {
				prettyJson, err := astBuilder.ToPrettyJSON(sourceUnit)
				assert.NoError(t, err)
				assert.NotEmpty(t, prettyJson)
			}

			prettyJson, err := astBuilder.ToPrettyJSON(astBuilder.GetRoot())
			assert.NoError(t, err)
			assert.NotEmpty(t, prettyJson)

			astJson, err := astBuilder.ToJSON()
			assert.NoError(t, err)
			assert.NotEmpty(t, astJson)

			assert.Equal(t, testCase.expected, string(astJson))
		})
	}
}
