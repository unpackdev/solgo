# Application Binary Interface (ABI) Parser

The Application Binary Interface (ABI) Parser is a tool that allows developers to parse and interact with ethereum-based (EVM) smart contracts written in Solidity. The parser is built on the Go programming language and uses the solgo package to parse Solidity source code and generate an ABI (Application Binary Interface).

## How to Use the ABI Parser

Here is a step-by-step guide on how to use the ABI Parser:

### Step 1: Import Required Packages

The first step is to import the required packages. The main packages required are context, encoding/json, fmt, strings, github.com/txpull/solgo, and github.com/txpull/solgo/abis.

```
import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/txpull/solgo"
	"github.com/txpull/solgo/abis"
)
```

### Step 2: Define the Smart Contract

Define the smart contract that you want to parse. This is done by creating a string variable that contains the Solidity source code of the contract.

```go
var contract = `
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract MyContract {
	uint256 public myUint;
	address public myAddress;
	string public myString = "Hello, World!";
	bytes32 public myBytes32 = "Hello, World!";
	bool public myBool = true;
	uint256[] public myUintArr = [1,2,3];

    function example() public pure returns (string memory) {
        return "Hello, World!";
    }
}`
```

### Step 3: Create a New Parser

Create a new parser using the solgo.New() function. This function takes in a context and a reader that reads the contract source code.

```go
parser, err := solgo.New(context.Background(), strings.NewReader(contract))
if err != nil {
	panic(err)
}
```

### Step 4: Register the ABI Listener

Register the ABI listener using the parser.RegisterListener() function. This function takes in a listener type and a listener. In this case, the listener type is solgo.ListenerAbi and the listener is an instance of abis.NewAbiListener().

```go
abiListener := abis.NewAbiListener()
if err := parser.RegisterListener(solgo.ListenerAbi, abiListener); err != nil {
	panic(err)
}
```

### Step 5: Parse the Contract

Parse the contract using the parser.Parse() function. This function returns a list of errors if any occur during parsing.

```go
if errs := parser.Parse(); len(errs) > 0 {
	for _, err := range errs {
		fmt.Println(err)
	}
	return
}
```

### Step 6: Get the ABI Parser

Get the ABI parser from the listener using the abiListener.GetParser() function.

```go
abiParser := abiListener.GetParser()
```

### Step 7: Get the ABI

Get the JSON representation of the ABI using the abiParser.ToJSON() function. You can also get the go-ethereum ABI representation using the abiParser.ToABI() function.

```go
_, err = abiParser.ToJSON()
if err != nil {
	panic(err)
}

_, err = abiParser.ToABI()
if err != nil {
	panic(err)
}
```

Return values from the ToJSON() and ToABI() functions are ommited for brevity.

### Step 8: Print the ABI

Finally, you can print the ABI in a structured format using the json.MarshalIndent() function.

```go
jsonResponse, _ := json.MarshalIndent(abiParser.ToStruct(), "", "  ")
fmt.Println(string(jsonResponse))
```

This will print the ABI in a structured format, making it easier to understand and interact with.

That's it! You have successfully parsed a Solidity contract and generated its ABI using the ABI Parser. You can now use this ABI to interact with the contract on the ethereum based blockchains.


## Example

Here is a complete example of what one abi parsed json looks like:

```json
[
  {
    "inputs": [],
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "name": "myUint",
    "type": "function",
    "stateMutability": "view"
  },
  {
    "inputs": [],
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "name": "myAddress",
    "type": "function",
    "stateMutability": "view"
  },
  {
    "inputs": [],
    "outputs": [
      {
        "internalType": "uint256[]",
        "name": "",
        "type": "uint256[]"
      }
    ],
    "name": "myUintArr",
    "type": "function",
    "stateMutability": "view"
  }
]
```