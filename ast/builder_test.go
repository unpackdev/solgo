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
			name:     "ERC20 Token",
			contract: tests.ReadContractFileForTest(t, "ERC20_Token").Content,
			expected: `{"interfaces":[],"contracts":[{"name":"MyToken","variables":[{"name":"MINTER_ROLE","type":"bytes32","visibility":"internal","is_constant":true,"is_immutable":false,"initial_value":"keccak256(\"MINTER_ROLE\")"},{"name":"subscriptionAmount","type":"uint256","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"subscriptionBalance","type":"mapping(address=\u003euint256)","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"isSubscribed","type":"mapping(address=\u003ebool)","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"rewards","type":"mapping(address=\u003euint256)","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"rewardedUsers","type":"address[]","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""}],"enums":null,"structs":null,"events":[{"name":"SubscriptionPurchased","anonymous":false,"parameters":[{"name":"user","type":"address","indexed":true},{"name":"amount","type":"uint256","indexed":false}]},{"name":"SubscriptionCanceled","anonymous":false,"parameters":[{"name":"user","type":"address","indexed":true},{"name":"amount","type":"uint256","indexed":false}]},{"name":"UserRewarded","anonymous":false,"parameters":[{"name":"user","type":"address","indexed":true},{"name":"amount","type":"uint256","indexed":false}]}],"errors":null,"constructor":null,"functions":[{"name":"initialize","parameters":[{"name":"name","type":"string"},{"name":"symbol","type":"string"},{"name":"initialSupply","type":"uint256"},{"name":"_subscriptionAmount","type":"uint256"}],"return_parameters":[],"body":[{"raw":["__ERC20_init","(","name",",","symbol",")",";"],"text_raw":"__ERC20_init(name,symbol);"},{"raw":["__AccessControl_init","(",")",";"],"text_raw":"__AccessControl_init();"},{"raw":["__Pausable_init","(",")",";"],"text_raw":"__Pausable_init();"},{"raw":["subscriptionAmount","=","_subscriptionAmount",";"],"text_raw":"subscriptionAmount=_subscriptionAmount;"},{"raw":["_mint","(","msg",".","sender",",","initialSupply",")",";"],"text_raw":"_mint(msg.sender,initialSupply);"},{"raw":["_setupRole","(","DEFAULT_ADMIN_ROLE",",","msg",".","sender",")",";"],"text_raw":"_setupRole(DEFAULT_ADMIN_ROLE,msg.sender);"},{"raw":["_setupRole","(","MINTER_ROLE",",","msg",".","sender",")",";"],"text_raw":"_setupRole(MINTER_ROLE,msg.sender);"}],"visibility":[{"value":"public"}],"mutability":[{"value":"nonpayable"}],"modifiers":[{"value":"initializer"}],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"pause","parameters":[],"return_parameters":[],"body":[{"raw":["_pause","(",")",";"],"text_raw":"_pause();"}],"visibility":[{"value":"external"}],"mutability":[{"value":"nonpayable"}],"modifiers":[{"value":"onlyRole(DEFAULT_ADMIN_ROLE)"}],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"unpause","parameters":[],"return_parameters":[],"body":[{"raw":["_unpause","(",")",";"],"text_raw":"_unpause();"}],"visibility":[{"value":"external"}],"mutability":[{"value":"nonpayable"}],"modifiers":[{"value":"onlyRole(DEFAULT_ADMIN_ROLE)"}],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"purchaseSubscription","parameters":[],"return_parameters":[],"body":[{"raw":["require","(","subscriptionBalance","[","msg",".","sender","]","\u003c","subscriptionAmount",",","\"Already subscribed\"",")",";"],"text_raw":"require(subscriptionBalance[msg.sender]\u003csubscriptionAmount,\"Already subscribed\");"},{"raw":["transfer","(","address","(","this",")",",","subscriptionAmount",")",";"],"text_raw":"transfer(address(this),subscriptionAmount);"},{"raw":["subscriptionBalance","[","msg",".","sender","]","+=","subscriptionAmount",";"],"text_raw":"subscriptionBalance[msg.sender]+=subscriptionAmount;"},{"raw":["isSubscribed","[","msg",".","sender","]","=","true",";"],"text_raw":"isSubscribed[msg.sender]=true;"},{"raw":["emit","SubscriptionPurchased","(","msg",".","sender",",","subscriptionAmount",")",";"],"text_raw":"emitSubscriptionPurchased(msg.sender,subscriptionAmount);"}],"visibility":[{"value":"external"}],"mutability":[{"value":"nonpayable"}],"modifiers":[{"value":"whenNotPaused"}],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"cancelSubscription","parameters":[],"return_parameters":[],"body":[{"raw":["require","(","isSubscribed","[","msg",".","sender","]",",","\"Not subscribed\"",")",";"],"text_raw":"require(isSubscribed[msg.sender],\"Not subscribed\");"},{"raw":["uint256","refundAmount","=","subscriptionBalance","[","msg",".","sender","]",";"],"text_raw":"uint256refundAmount=subscriptionBalance[msg.sender];"},{"raw":["require","(","refundAmount","\u003e","0",",","\"No subscription balance to refund\"",")",";"],"text_raw":"require(refundAmount\u003e0,\"No subscription balance to refund\");"},{"raw":["transfer","(","msg",".","sender",",","refundAmount",")",";"],"text_raw":"transfer(msg.sender,refundAmount);"},{"raw":["subscriptionBalance","[","msg",".","sender","]","=","0",";"],"text_raw":"subscriptionBalance[msg.sender]=0;"},{"raw":["isSubscribed","[","msg",".","sender","]","=","false",";"],"text_raw":"isSubscribed[msg.sender]=false;"},{"raw":["emit","SubscriptionCanceled","(","msg",".","sender",",","refundAmount",")",";"],"text_raw":"emitSubscriptionCanceled(msg.sender,refundAmount);"}],"visibility":[{"value":"external"}],"mutability":[{"value":"nonpayable"}],"modifiers":[{"value":"whenNotPaused"}],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"getSubscriptionStatus","parameters":[{"name":"user","type":"address"}],"return_parameters":[{"name":"","type":"bool"}],"body":[{"raw":["return","isSubscribed","[","user","]",";"],"text_raw":"returnisSubscribed[user];"}],"visibility":[{"value":"external"}],"mutability":[{"value":"view"}],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"reward","parameters":[{"name":"user","type":"address"},{"name":"amount","type":"uint256"}],"return_parameters":[],"body":[{"raw":["transfer","(","user",",","amount",")",";"],"text_raw":"transfer(user,amount);"},{"raw":["rewards","[","user","]","+=","amount",";"],"text_raw":"rewards[user]+=amount;"},{"raw":["if","(","rewards","[","user","]","==","amount",")","{","rewardedUsers",".","push","(","user",")",";","}"],"text_raw":"if(rewards[user]==amount){rewardedUsers.push(user);}"},{"raw":["emit","UserRewarded","(","user",",","amount",")",";"],"text_raw":"emitUserRewarded(user,amount);"}],"visibility":[{"value":"external"}],"mutability":[{"value":"nonpayable"}],"modifiers":[{"value":"onlyRole(DEFAULT_ADMIN_ROLE)"}],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"getRewards","parameters":[{"name":"user","type":"address"}],"return_parameters":[{"name":"","type":"uint256"}],"body":[{"raw":["return","rewards","[","user","]",";"],"text_raw":"returnrewards[user];"}],"visibility":[{"value":"external"}],"mutability":[{"value":"view"}],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"getRewardedUsers","parameters":[],"return_parameters":[{"name":"","type":"address[]"}],"body":[{"raw":["return","rewardedUsers",";"],"text_raw":"returnrewardedUsers;"}],"visibility":[{"value":"external"}],"mutability":[{"value":"view"}],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"_beforeTokenTransfer","parameters":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"amount","type":"uint256"}],"return_parameters":[],"body":[{"raw":["super",".","_beforeTokenTransfer","(","from",",","to",",","amount",")",";"],"text_raw":"super._beforeTokenTransfer(from,to,amount);"},{"raw":["if","(","isSubscribed","[","from","]","\u0026\u0026","subscriptionBalance","[","from","]","\u003e=","amount",")","{","subscriptionBalance","[","from","]","-=","amount",";","subscriptionBalance","[","to","]","+=","amount",";","}"],"text_raw":"if(isSubscribed[from]\u0026\u0026subscriptionBalance[from]\u003e=amount){subscriptionBalance[from]-=amount;subscriptionBalance[to]+=amount;}"}],"visibility":[{"value":"internal"}],"mutability":[{"value":"nonpayable"}],"modifiers":[{"value":"whenNotPaused"}],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":true},{"name":"updateSubscriptionAmount","parameters":[{"name":"newAmount","type":"uint256"}],"return_parameters":[],"body":[{"raw":["subscriptionAmount","=","newAmount",";"],"text_raw":"subscriptionAmount=newAmount;"}],"visibility":[{"value":"external"}],"mutability":[{"value":"nonpayable"}],"modifiers":[{"value":"onlyRole(DEFAULT_ADMIN_ROLE)"}],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"mint","parameters":[{"name":"to","type":"address"},{"name":"amount","type":"uint256"}],"return_parameters":[],"body":[{"raw":["_mint","(","to",",","amount",")",";"],"text_raw":"_mint(to,amount);"}],"visibility":[{"value":"external"}],"mutability":[{"value":"nonpayable"}],"modifiers":[{"value":"onlyRole(MINTER_ROLE)"}],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false}],"kind":"contract","inherits":["Initializable","ERC20Upgradeable","AccessControlUpgradeable","PausableUpgradeable"],"using":[{"alias":"SafeERC20Upgradeable","type":"IERC20Upgradeable","is_wildcard":false,"is_global":false,"is_user_defined":false}]}]}`,
		},
		{
			name:     "Complex AST Contract",
			contract: tests.ReadContractFileForTest(t, "AstComplex").Content,
			expected: `{"interfaces":[{"name":"IBaseContract","functions":null}],"contracts":[{"name":"BaseContract","variables":[],"enums":null,"structs":null,"events":null,"errors":null,"constructor":null,"functions":[{"name":"baseFunction","parameters":[],"return_parameters":[{"name":"","type":"uint"}],"body":[{"raw":["return","1",";"],"text_raw":"return1;"}],"visibility":[{"value":"public"}],"mutability":[{"value":"nonpayable"}],"modifiers":[],"is_virtual":true,"is_receive":false,"is_fallback":false,"overrides":true}],"kind":"contract","inherits":["IBaseContract"],"using":null},{"name":"TestContract","variables":[{"name":"stateVariable","type":"uint","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"privateStateVariable","type":"uint","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"internalStateVariable","type":"uint","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"constantStateVariable","type":"uint","visibility":"internal","is_constant":true,"is_immutable":false,"initial_value":"1"},{"name":"immutableStateVariable","type":"uint","visibility":"internal","is_constant":false,"is_immutable":true,"initial_value":""},{"name":"owner","type":"address","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"arrayStateVariable","type":"uint[]","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"mappingStateVariable","type":"mapping(address=\u003euint)","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""},{"name":"structStateVariable","type":"StructType","visibility":"internal","is_constant":false,"is_immutable":false,"initial_value":""}],"enums":[{"name":"CustomEnum","memberValues":[{"name":"Option1"},{"name":"Option2"},{"name":"Option3"}]}],"structs":[{"Name":"StructType","Members":[{"Name":"field1","Type":"uint"},{"Name":"field2","Type":"address"}]}],"events":[{"name":"StateVariableChanged","anonymous":false,"parameters":[{"name":"oldValue","type":"uint","indexed":false},{"name":"newValue","type":"uint","indexed":false}]},{"name":"OwnerChanged","anonymous":false,"parameters":[{"name":"oldOwner","type":"address","indexed":true},{"name":"newOwner","type":"address","indexed":true}]},{"name":"CustomEvent","anonymous":false,"parameters":[{"name":"message","type":"string","indexed":false}]}],"errors":[{"name":"CustomError","values":[{"name":"message","type":"string","code":0}]}],"constructor":{"parameters":[{"name":"_stateVariable","type":"uint"},{"name":"_privateStateVariable","type":"uint"}],"body":[{"raw":["stateVariable","=","_stateVariable",";"],"text_raw":"stateVariable=_stateVariable;"},{"raw":["privateStateVariable","=","_privateStateVariable",";"],"text_raw":"privateStateVariable=_privateStateVariable;"},{"raw":["owner","=","msg",".","sender",";"],"text_raw":"owner=msg.sender;"},{"raw":["immutableStateVariable","=","block",".","timestamp",";"],"text_raw":"immutableStateVariable=block.timestamp;"}]},"functions":[{"name":"publicFunction","parameters":[{"name":"a","type":"uint"},{"name":"b","type":"uint"}],"return_parameters":[{"name":"","type":"uint"}],"body":[{"raw":["return","a",".","add","(","b",")",";"],"text_raw":"returna.add(b);"}],"visibility":[{"value":"public"}],"mutability":[{"value":"nonpayable"}],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"privateFunction","parameters":[{"name":"a","type":"uint"},{"name":"b","type":"uint"}],"return_parameters":[{"name":"","type":"uint"}],"body":[{"raw":["uint","result",";"],"text_raw":"uintresult;"},{"raw":["assembly","{","result",":=","sub","(","a",",","b",")","}"],"text_raw":"assembly{result:=sub(a,b)}"},{"raw":["return","result",";"],"text_raw":"returnresult;"}],"visibility":[{"value":"private"}],"mutability":[{"value":"nonpayable"}],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"baseFunction","parameters":[],"return_parameters":[{"name":"","type":"uint"}],"body":[{"raw":["return","2",";"],"text_raw":"return2;"}],"visibility":[{"value":"public"}],"mutability":[{"value":"nonpayable"}],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":true},{"name":"functionWithModifier","parameters":[{"name":"a","type":"uint"},{"name":"b","type":"uint"}],"return_parameters":[{"name":"","type":"uint"}],"body":[{"raw":["return","a",".","add","(","b",")",";"],"text_raw":"returna.add(b);"}],"visibility":[{"value":"public"}],"mutability":[{"value":"nonpayable"}],"modifiers":[{"value":"onlyOwner"}],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"changeStateVariable","parameters":[{"name":"newValue","type":"uint"}],"return_parameters":[],"body":[{"raw":["emit","StateVariableChanged","(","stateVariable",",","newValue",")",";"],"text_raw":"emitStateVariableChanged(stateVariable,newValue);"},{"raw":["stateVariable","=","newValue",";"],"text_raw":"stateVariable=newValue;"}],"visibility":[{"value":"public"}],"mutability":[{"value":"nonpayable"}],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"changeOwner","parameters":[{"name":"newOwner","type":"address"}],"return_parameters":[],"body":[{"raw":["emit","OwnerChanged","(","owner",",","newOwner",")",";"],"text_raw":"emitOwnerChanged(owner,newOwner);"},{"raw":["owner","=","newOwner",";"],"text_raw":"owner=newOwner;"}],"visibility":[{"value":"public"}],"mutability":[{"value":"nonpayable"}],"modifiers":[{"value":"onlyOwner"}],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"fallback","parameters":[],"return_parameters":[],"body":[{"raw":["revert","(","\"Fallback function called\"",")",";"],"text_raw":"revert(\"Fallback function called\");"}],"visibility":[{"value":"external"}],"mutability":[{"value":"nonpayable"}],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":true,"overrides":false},{"name":"receive","parameters":[],"return_parameters":[],"body":[{"raw":["emit","CustomEvent","(","\"Receive function called\"",")",";"],"text_raw":"emitCustomEvent(\"Receive function called\");"}],"visibility":[],"mutability":[],"modifiers":[],"is_virtual":false,"is_receive":true,"is_fallback":false,"overrides":false},{"name":"throwError","parameters":[],"return_parameters":[],"body":[{"raw":["revert","(","CustomError","(","\"Error occurred\"",")",")",";"],"text_raw":"revert(CustomError(\"Error occurred\"));"}],"visibility":[{"value":"public"}],"mutability":[{"value":"pure"}],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false},{"name":"getEnumValue","parameters":[{"name":"option","type":"CustomEnum"}],"return_parameters":[{"name":"","type":"uint"}],"body":[{"raw":["if","(","option","==","CustomEnum",".","Option1",")","{","return","1",";","}","else","if","(","option","==","CustomEnum",".","Option2",")","{","return","2",";","}","else","if","(","option","==","CustomEnum",".","Option3",")","{","return","3",";","}"],"text_raw":"if(option==CustomEnum.Option1){return1;}elseif(option==CustomEnum.Option2){return2;}elseif(option==CustomEnum.Option3){return3;}"},{"raw":["return","0",";"],"text_raw":"return0;"}],"visibility":[{"value":"public"}],"mutability":[{"value":"pure"}],"modifiers":[],"is_virtual":false,"is_receive":false,"is_fallback":false,"overrides":false}],"kind":"contract","inherits":["BaseContract"],"using":[{"alias":"SafeMath","type":"uint","is_wildcard":false,"is_global":false,"is_user_defined":false}]}]}`,
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
