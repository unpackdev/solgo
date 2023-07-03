package ast

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

func TestAstBuilder(t *testing.T) {
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
			name:     "Empty Contract",
			contract: tests.ReadContractFileForTest(t, "Empty").Content,
			expected: `{"interfaces":[],"contracts":[]}`,
		},
		{
			name:     "Complex AST Contract",
			contract: tests.ReadContractFileForTest(t, "AstComplex").Content,
			expected: `{"interfaces":[{"name":"IBaseContract","functions":null}],"contracts":[{"name":"BaseContract","variables":[],"structs":null,"events":null,"errors":null,"constructor":null,"functions":[{"name":"baseFunction","parameters":[],"return_parameters":[{"name":"","type":"uint"}],"body":[{"raw":["return","1",";"],"text_raw":"return1;"}],"visibility":["public"],"mutability":["nonpayable"],"modifiers":[],"is_virtual":true,"is_receive":false,"is_fallback":false,"overrides":true}],"kind":"contract","inherits":["IBaseContract"],"using":null},{"name":"TestContract","variables":[{"name":"stateVariable","type":"uint","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"privateStateVariable","type":"uint","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"internalStateVariable","type":"uint","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"constantStateVariable","type":"uint","visibility":"internal","is_constant":true,"is_immutable":false,"initial_value":"1"},{"name":"immutableStateVariable","type":"uint","visibility":"internal","is_constant":false,"is_immutable":true,"initial_value":""},{"name":"owner","type":"address","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"arrayStateVariable","type":"uint[]","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"mappingStateVariable","type":"mapping(address=\u003euint)","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"structStateVariable","type":"StructType","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""}],"structs":[{"Name":"StructType","Members":[{"Name":"field1","Type":"uint"},{"Name":"field2","Type":"address"}]}],"events":[{"name":"StateVariableChanged","anonymous":false,"parameters":[{"name":"oldValue","type":"uint","indexed":false},{"name":"newValue","type":"uint","indexed":false}]},{"name":"OwnerChanged","anonymous":false,"parameters":[{"name":"oldOwner","type":"address","indexed":true},{"name":"newOwner","type":"address","indexed":true}]},{"name":"CustomEvent","anonymous":false,"parameters":[{"name":"message","type":"string","indexed":false}]}],"errors":[{"name":"CustomError","values":[{"name":"message","type":"string","code":0}]}],"constructor":{"parameters":[{"name":"_stateVariable","type":"uint"},{"name":"_privateStateVariable","type":"uint"}],"body":[{"raw":["stateVariable","=","_stateVariable",";"],"text_raw":"stateVariable=_stateVariable;"},{"raw":["privateStateVariable","=","_privateStateVariable",";"],"text_raw":"privateStateVariable=_privateStateVariable;"},{"raw":["owner","=","msg",".","sender",";"],"text_raw":"owner=msg.sender;"},{"raw":["immutableStateVariable","=","block",".","timestamp",";"],"text_raw":"immutableStateVariable=block.timestamp;"}]},"functions":[{"name":"publicFunction","parameters":[{"name":"a","type":"uint"},{"name":"b","type":"uint"}],"return_parameters":[{"name":"","type":"uint"}],"body":[{"raw":["return","a",".","add","(","b",")",";"],"text_raw":"returna.add(b);"}],"visibility":["public"],"mutability":["nonpayable"],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"privateFunction","parameters":[{"name":"a","type":"uint"},{"name":"b","type":"uint"}],"return_parameters":[{"name":"","type":"uint"}],"body":[{"raw":["uint","result",";"],"text_raw":"uintresult;"},{"raw":["assembly","{","result",":=","sub","(","a",",","b",")","}"],"text_raw":"assembly{result:=sub(a,b)}"},{"raw":["return","result",";"],"text_raw":"returnresult;"}],"visibility":["private"],"mutability":["nonpayable"],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"baseFunction","parameters":[],"return_parameters":[{"name":"","type":"uint"}],"body":[{"raw":["return","2",";"],"text_raw":"return2;"}],"visibility":["public"],"mutability":["nonpayable"],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":true},{"name":"functionWithModifier","parameters":[{"name":"a","type":"uint"},{"name":"b","type":"uint"}],"return_parameters":[{"name":"","type":"uint"}],"body":[{"raw":["return","a",".","add","(","b",")",";"],"text_raw":"returna.add(b);"}],"visibility":["public"],"mutability":["nonpayable"],"modifiers":["onlyOwner"],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"changeStateVariable","parameters":[{"name":"newValue","type":"uint"}],"return_parameters":[],"body":[{"raw":["emit","StateVariableChanged","(","stateVariable",",","newValue",")",";"],"text_raw":"emitStateVariableChanged(stateVariable,newValue);"},{"raw":["stateVariable","=","newValue",";"],"text_raw":"stateVariable=newValue;"}],"visibility":["public"],"mutability":["nonpayable"],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"changeOwner","parameters":[{"name":"newOwner","type":"address"}],"return_parameters":[],"body":[{"raw":["emit","OwnerChanged","(","owner",",","newOwner",")",";"],"text_raw":"emitOwnerChanged(owner,newOwner);"},{"raw":["owner","=","newOwner",";"],"text_raw":"owner=newOwner;"}],"visibility":["public"],"mutability":["nonpayable"],"modifiers":["onlyOwner"],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"fallback","parameters":[],"return_parameters":[],"body":[{"raw":["revert","(","\"Fallback function called\"",")",";"],"text_raw":"revert(\"Fallback function called\");"}],"visibility":["external"],"mutability":["nonpayable"],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":true,"overrides":false},{"name":"receive","parameters":[],"return_parameters":[],"body":[{"raw":["emit","CustomEvent","(","\"Receive function called\"",")",";"],"text_raw":"emitCustomEvent(\"Receive function called\");"}],"visibility":[],"mutability":[],"modifiers":[],"is_virtual":false,"is_receive":true,"is_fallback":false,"overrides":false},{"name":"throwError","parameters":[],"return_parameters":[],"body":[{"raw":["revert","(","CustomError","(","\"Error occurred\"",")",")",";"],"text_raw":"revert(CustomError(\"Error occurred\"));"}],"visibility":["public"],"mutability":["pure"],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"getEnumValue","parameters":[{"name":"option","type":"CustomEnum"}],"return_parameters":[{"name":"","type":"uint"}],"body":[{"raw":["if","(","option","==","CustomEnum",".","Option1",")","{","return","1",";","}","else","if","(","option","==","CustomEnum",".","Option2",")","{","return","2",";","}","else","if","(","option","==","CustomEnum",".","Option3",")","{","return","3",";","}"],"text_raw":"if(option==CustomEnum.Option1){return1;}elseif(option==CustomEnum.Option2){return2;}elseif(option==CustomEnum.Option3){return3;}"},{"raw":["return","0",";"],"text_raw":"return0;"}],"visibility":["public"],"mutability":["pure"],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false}],"kind":"contract","inherits":["BaseContract"],"using":[{"alias":"SafeMath","type":"uint","is_wildcard":false,"is_global":false,"is_user_defined":false}]}]}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parser, err := solgo.New(context.TODO(), strings.NewReader(testCase.contract))
			assert.NoError(t, err)
			assert.NotNil(t, parser)

			// Register the contract information listener
			astBuilder := NewAstBuilder()
			err = parser.RegisterListener(solgo.ListenerAst, astBuilder)
			assert.NoError(t, err)

			syntaxErrs := parser.Parse()
			assert.Empty(t, syntaxErrs)

			astJson, err := astBuilder.ToJSON()
			assert.NoError(t, err)
			assert.NotEmpty(t, astJson)

			assert.Equal(t, testCase.expected, astJson)
		})
	}
}
