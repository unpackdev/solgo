package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/txpull/solgo/opcode"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/* func main() {
	currentTick := time.Now()
	defer func() {
		zap.S().Infof("Total time taken: %v", time.Since(currentTick))
	}()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(logger)

	sources := solgo.Sources{
		SourceUnits: []*solgo.SourceUnit{
			{
				Name: "MyToken",     // Make sure name matches the contract name. Important!!!
				Path: "MyToken.sol", // Make sure path matches the contract name. Important!!!
				Content: `// SPDX-License-Identifier: MIT
	pragma solidity ^0.8.0;

	import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

	contract MyToken is ERC20 {
		constructor(uint256 initialSupply) ERC20("MyToken", "MTK") {
			_mint(msg.sender, initialSupply);
		}
	}`,
			},
		},
		EntrySourceUnitName:  "MyToken",
		MaskLocalSourcesPath: true,
		LocalSourcesPath:     "../sources/",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	detector, err := detector.NewDetectorFromSources(ctx, sources)
	if err != nil {
		zap.L().Error("failure to construct detector", zap.Error(err))
		return
	}

	// 1. Parse all of the sources and see if there are any syntax errors
	// discovered durring the process.
	if errs := detector.Parse(); errs != nil {
		for _, err := range errs {
			zap.L().Error("failure to parse sources", zap.Error(err))
		}
		return
	}

	// 2. Build all of the components and see if there are any errors.
	// For example, this will build ABIs.
	if err := detector.Build(); err != nil {
		zap.L().Error("failure to compile sources", zap.Error(err))
		return
	}

	// 1. Lets create some common variable for IR instance.
	irBuilder := detector.GetIR()

	// 2. Lets create common variable for root IR node.
	irRoot := irBuilder.GetRoot()

	// 3. Lets print out how many contracts were discovered.
	zap.L().Info(
		"Number of discovered contracts",
		zap.Int("count", int(irRoot.GetContractsCount())),
	)

	// 4. Lets print out the entry name.
	zap.L().Info(
		"Entry Contract Name and Internal Id",
		zap.String("name", irRoot.GetEntryName()),
		zap.Int64("internal_id", irRoot.GetEntryId()),
	)

	// 5. Lets print out the entry contract license.
	zap.L().Info(
		"Entry Contract License",
		zap.String("license", irRoot.GetEntryContract().License),
	)

	// 6. Lets print out the entry contract discovered types.
	zap.L().Info(
		"Discovered Contract Types",
		zap.Any("types", irRoot.GetContractTypes()),
	)

	// 7. Lets print out the entry contract discovered EIPs.
	// Following result might bring totally unrelated EIPs but you will see in confidence level
	// very low confidence. It goes from 0 to 1. 0 is the lowest confidence and 1 is the highest.
	// Always look for those that are highest.
	for _, eip := range irRoot.GetEips() {
		zap.L().Info(
			"Discovered Potential EIPs",
			zap.String("eip", eip.GetStandard().Name),
			zap.Any("confidence_level", eip.GetConfidence().Confidence.String()),
			zap.Any("confidence_points", eip.GetConfidence().ConfidencePoints),
		)
	}

	// 8. Now we can for example print all of the functions discovered in the entry contract.
	var discoveredFunctions []string
	for _, function := range irRoot.GetEntryContract().GetFunctions() {
		discoveredFunctions = append(discoveredFunctions, function.Name)
	}
	zap.L().Info(
		"Discovered Functions",
		zap.Any("functions", discoveredFunctions),
	)

	// 9. Because there are no functions, we should in this example display constructor...
	zap.L().Info(
		"Discovered Constructor Information",
		zap.Any("functions", irRoot.GetEntryContract().GetConstructor()),
	)

	// 8. Lets print out the entry contract discovered ABI.
	// Note that this won't return full ABI for all discovered contracts.
	// It will only return ABI for the entry contract.
	// If you want to get full ABI for all discovered contracts then you need to pass:
	// detector.GetABI().GetRoot().GetContracts()
	contractAbi, err := detector.GetABI().ToJSON(
		detector.GetABI().GetRoot().GetEntryContract(),
	)
	if err != nil {
		zap.L().Error("failure to convert ABI to ABIv2", zap.Error(err))
		return
	}

	zap.L().Info(
		"Discovered Contract ABI",
		zap.String("abi", string(contractAbi)),
	)

	// This concludes example...
}
*/

func main() {
	currentTick := time.Now()
	defer func() {
		zap.S().Infof("Total time took: %v", time.Since(currentTick))
	}()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bytecode, err := hex.DecodeString("608060405234801561001057600080fd5b506040516109b33803806109b38339818101604052606081101561003357600080fd5b8151602083015160408085018051915193959294830192918464010000000082111561005e57600080fd5b90830190602082018581111561007357600080fd5b825164010000000081118282018810171561008d57600080fd5b82525081516020918201929091019080838360005b838110156100ba5781810151838201526020016100a2565b50505050905090810190601f1680156100e75780820380516001836020036101000a031916815260200191505b5060408181527f656970313936372e70726f78792e696d706c656d656e746174696f6e0000000082525190819003601c0190208693508592508491508390829060008051602061095d8339815191526000199091011461014357fe5b610155826001600160e01b0361027a16565b80511561020d576000826001600160a01b0316826040518082805190602001908083835b602083106101985780518252601f199092019160209182019101610179565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146101f8576040519150601f19603f3d011682016040523d82523d6000602084013e6101fd565b606091505b505090508061020b57600080fd5b505b5050604080517f656970313936372e70726f78792e61646d696e000000000000000000000000008152905190819003601301902060008051602061093d8339815191526000199091011461025d57fe5b61026f826001600160e01b036102da16565b5050505050506102f2565b61028d816102ec60201b61054e1760201c565b6102c85760405162461bcd60e51b815260040180806020018281038252603681526020018061097d6036913960400191505060405180910390fd5b60008051602061095d83398151915255565b60008051602061093d83398151915255565b3b151590565b61063c806103016000396000f3fe60806040526004361061004e5760003560e01c80633659cfe6146100655780634f1ef286146100985780635c60da1b146101185780638f28397014610149578063f851a4401461017c5761005d565b3661005d5761005b610191565b005b61005b610191565b34801561007157600080fd5b5061005b6004803603602081101561008857600080fd5b50356001600160a01b03166101ab565b61005b600480360360408110156100ae57600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156100d957600080fd5b8201836020820111156100eb57600080fd5b8035906020019184600183028401116401000000008311171561010d57600080fd5b5090925090506101e5565b34801561012457600080fd5b5061012d610292565b604080516001600160a01b039092168252519081900360200190f35b34801561015557600080fd5b5061005b6004803603602081101561016c57600080fd5b50356001600160a01b03166102cf565b34801561018857600080fd5b5061012d610389565b6101996103b4565b6101a96101a4610414565b610439565b565b6101b361045d565b6001600160a01b0316336001600160a01b031614156101da576101d581610482565b6101e2565b6101e2610191565b50565b6101ed61045d565b6001600160a01b0316336001600160a01b031614156102855761020f83610482565b6000836001600160a01b031683836040518083838082843760405192019450600093509091505080830381855af49150503d806000811461026c576040519150601f19603f3d011682016040523d82523d6000602084013e610271565b606091505b505090508061027f57600080fd5b5061028d565b61028d610191565b505050565b600061029c61045d565b6001600160a01b0316336001600160a01b031614156102c4576102bd610414565b90506102cc565b6102cc610191565b90565b6102d761045d565b6001600160a01b0316336001600160a01b031614156101da576001600160a01b0381166103355760405162461bcd60e51b815260040180806020018281038252603a815260200180610555603a913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61035e61045d565b604080516001600160a01b03928316815291841660208301528051918290030190a16101d5816104c2565b600061039361045d565b6001600160a01b0316336001600160a01b031614156102c4576102bd61045d565b6103bc61045d565b6001600160a01b0316336001600160a01b0316141561040c5760405162461bcd60e51b81526004018080602001828103825260428152602001806105c56042913960600191505060405180910390fd5b6101a96101a9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e808015610458573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b61048b816104e6565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b6104ef8161054e565b61052a5760405162461bcd60e51b815260040180806020018281038252603681526020018061058f6036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b3b15159056fe5472616e73706172656e745570677261646561626c6550726f78793a206e65772061646d696e20697320746865207a65726f20616464726573735570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e74726163745472616e73706172656e745570677261646561626c6550726f78793a2061646d696e2063616e6e6f742066616c6c6261636b20746f2070726f787920746172676574a26469706673582212205c518be5ecdac9ebba758e8ce0b8e0dcacae92de07203f44e322e833b133a57564736f6c63430006040033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374000000000000000000000000ba5fe23f8a3a24bed3236f05f2fcf35fd0bf0b5c000000000000000000000000d2f93484f2d319194cba95c5171b18c1d8cfd6c400000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000")
	if err != nil {
		zap.S().Errorf("Error during decoding bytecode from hex: %v", err)
		return
	}

	decompiler, err := opcode.NewDecompiler(ctx, []byte(bytecode))
	if err != nil {
		zap.S().Errorf("Error during construction of new decompiler: %v", err)
		return
	}

	if err := decompiler.Decompile(); err != nil {
		zap.S().Errorf("Error during decompilation: %v", err)
		return
	}

	// Print the decompiled instructions
	fmt.Println(decompiler.String())
}
