package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/txpull/solgo"
	"go.uber.org/zap"
)

func main() {
	// Initialize a logger
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	zap.ReplaceGlobals(logger)

	// Example Solidity code
	code := `
	pragma solidity ^0.8.0;

	contract ExampleContract {
		uint256 public counter; // State variable
		
		bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

		string private name; // Private state variable
		address payable public owner; // Public state variable
		mapping(address => uint256) public subscriptionBalance;
		
		constructor(address owner) {
			counter = 0;
			name = "Example";
			owner = payable(msg.sender);
		}


		// Events
    	event SubscriptionPurchased(address indexed user, uint256 amount);
    	event SubscriptionCanceled(address indexed user, uint256 amount);
    	event UserRewarded(address indexed user, uint256 amount);

		function add(uint256 a, uint256 b) public pure returns (uint256) {
			return a + b;
		}
	
		function multiply(uint256 a, uint256 b) public pure returns (uint256) {
			return a * b;
		}
	
		function deposit() external payable {
			// Perform deposit logic here
		}
	
		function withdraw(uint256 amount) external {
			// Perform withdrawal logic here 
		}
	}
	`

	// Create a new SolGo instance
	solGo, err := solgo.New(context.Background(), strings.NewReader(code))
	if err != nil {
		zap.L().Error("Error creating SolGo", zap.Error(err))
		return
	}

	abiListener := solgo.NewAbiTreeShapeListener()

	// Register the abi tree shape listener
	err = solGo.RegisterListener(solgo.ListenerAbiTreeShape, abiListener)
	if err != nil {
		zap.L().Error("failed to register abi listener", zap.Error(err))
		return
	}

	// Parse the input
	err = solGo.Parse()
	if err != nil {
		zap.L().Error("Error parsing input", zap.Error(err))
		return
	}

	// Log the result
	spew.Dump(abiListener.GetRawABI())

	jsonABI, err := abiListener.ToJSON()
	if err != nil {
		zap.L().Error("failed to convert abi to json", zap.Error(err))
		return
	}

	fmt.Println(jsonABI)

	parsedAbi, err := abiListener.ToABI()
	if err != nil {
		zap.L().Error("failed to convert json abi to abi", zap.Error(err))
		return
	}

	// Dump the abi.ABI
	spew.Dump(parsedAbi)
}
