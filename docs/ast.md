## Parsing a Solidity Contract and Building an AST

This example demonstrates how to parse a Solidity contract and build an Abstract Syntax Tree (AST) using the `solgo` package.

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/txpull/solgo"
	"github.com/txpull/solgo/ast"
)

var contract = `
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract MyContract {
	uint256 public myUint;
	address public myAddress;
	uint256[] public myUintArr = [1,2,3];
}`

func main() {
	parser, err := solgo.New(context.Background(), strings.NewReader(contract))
	if err != nil {
		panic(err)
	}

	// Register the abi listener
	astBuilder := ast.NewAstBuilder()
	if err := parser.RegisterListener(solgo.ListenerAst, astBuilder); err != nil {
		panic(err)
	}

	if errs := parser.Parse(); len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		return
	}

	jsonResponse, _ := json.MarshalIndent(astBuilder.GetTree(), "", "  ")
	fmt.Println(string(jsonResponse))
}
```

The output of this program is a JSON representation of the AST for the given contract. Here's an example of what the output will look like:

```
{
    "interfaces": [],
    "contracts": [
        {
        "name": "MyContract",
        "variables": [
            {
            "name": "myUint",
            "type": "uint256",
            "visibility": "internal",
            "is_constant": false,
            "is_immutable": false,
            "initial_value": ""
            },
            {
            "name": "myAddress",
            "type": "address",
            "visibility": "internal",
            "is_constant": false,
            "is_immutable": false,
            "initial_value": ""
            },
            {
            "name": "myUintArr",
            "type": "uint256[]",
            "visibility": "internal",
            "is_constant": false,
            "is_immutable": false,
            "initial_value": "[1,2,3]"
            }
        ],
        "structs": null,
        "events": null,
        "errors": null,
        "constructor": null,
        "functions": null,
        "kind": "contract",
        "inherits": null,
        "using": null
        }
    ]
}
```
