[![Go Report Card](https://goreportcard.com/badge/github.com/txpull/solgo)](https://goreportcard.com/report/github.com/txpull/solgo)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/txpull/solgo)](https://pkg.go.dev/github.com/txpull/solgo)

# Solidity Parser in Golang

SolGo contains a Solidity parser written in Golang, using Antlr and AntlrGo for grammar parsing. The aim of this project is to provide a tool to parse Solidity source code into a structured format, enabling further analysis and manipulation within Go programs.

The parser is generated from a Solidity grammar file using Antlr, producing a lexer, parser, and listener in Golang using AntlrGo. This allows for the syntactic analysis of Solidity code, transforming it into a parse tree that can be traversed and manipulated.

This project is ideal for developers working with Solidity smart contracts who wish to leverage the power and efficiency of Golang for their analysis tools.

**Note: This project is still in development and is not yet ready for use in production.**

## Why?

In my projects, I have a strong preference for using Golang or Rust over Javascript. I found myself needing to parse Solidity code and conduct some analysis on it. During my research, I discovered several projects that aimed to do this. However, they were either incomplete, not maintained, or not written in Go at all. This led me to the decision to create this project. The goal is to provide a Solidity parser in Golang that is as "complete as possible" and well-maintained. I'm excited to see how this project will develop and evolve.

This project will be integrated within [unpack](https://github.com/txpull/unpack). I've found that for many deployed contracts, the source code is available, but the ABIs are not. This project will help bridge that gap.

## ANTLR Grammar

We are using grammar files that are maintained by the Solidity team.
Link to the grammar files can be found [here](https://github.com/ethereum/solidity/tree/develop/docs/grammar)

## ANTLR Go

We are using the ANTLR4 Go runtime library to generate the parser. Repository can be found [here](https://github.com/antlr4-go/antlr).


## Features

- **Syntactic Analysis**: SolGo transforms Solidity code into a parse tree that can be traversed and manipulated, providing a detailed syntactic analysis of the code.
- **Listener Registration**: SolGo allows for the registration of custom listeners that can be invoked as the parser walks the parse tree, enabling custom handling of parse events.
- **Error Handling**: SolGo includes a SyntaxErrorListener which collects syntax errors encountered during parsing, providing detailed error information including line, column, message, severity, and context.
- **Contextual Parsing**: SolGo provides a ContextualSolidityParser that maintains a stack of contexts, allowing for context-specific parsing rules.

## Getting Started

To use SolGo, you will need to have Golang installed on your machine. You can then import SolGo into your Go programs like this:

```go
import "github.com/txpull/solgo"
```

## Usage

### Parse Solidity Code and Retrieve Contract Information

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/txpull/solgo"
)

var contract = `
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// Some additional comments that can be extracted

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol";

contract MyToken is Initializable, ERC20Upgradeable, AccessControlUpgradeable, PausableUpgradeable {
    using SafeERC20Upgradeable for IERC20Upgradeable;
}
`

func main() {
	parser, err := solgo.New(context.Background(), strings.NewReader(contract))
	if err != nil {
		panic(err)
	}

	// Register the contract information listener
	contractListener := solgo.NewContractListener(parser.GetParser())
	if err := parser.RegisterListener(solgo.ListenerContractInfo, contractListener); err != nil {
		panic(err)
	}

	if errs := parser.Parse(); len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		return
	}

	// Get the contract information from listener that is built for testing purposes
	jsonResponse, _ := json.MarshalIndent(contractListener.ToStruct(), "", "  ")
	fmt.Println(string(jsonResponse))
}
```

Response from the example above:

```json
{
  "comments": [
    "// Some additional comments that can be extracted"
  ],
  "license": "MIT",
  "pragmas": [
    "solidity ^0.8.0"
  ],
  "imports": [
    "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol",
    "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol",
    "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol",
    "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol",
    "@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol"
  ],
  "name": "MyToken",
  "implements": [
    "Initializable",
    "ERC20Upgradeable",
    "AccessControlUpgradeable",
    "PausableUpgradeable",
    "SafeERC20Upgradeable"
  ]
}
```

You can try to play with the code here: [Go Play Tools](https://goplay.tools/snippet/G6jmkP3L1mr)
Sometimes it works, sometimes it does not, returns 502. If it does not work, you can try to run it locally.


## Contributing

Contributions to SolGo are always welcome! If you have a feature request, bug report, or proposal for improvement, please open an issue. If you wish to contribute code, please fork the repository, make your changes, and submit a pull request.


## License

SolGo is licensed under the Apache 2.0. See [LICENSE](LICENSE) for the full license text.


## Acknowledgements

We would like to express our gratitude to the Solidity team for maintaining the Solidity grammar files, and to the Antlr and AntlrGo team for providing the powerful Antlr tool that makes this project possible.