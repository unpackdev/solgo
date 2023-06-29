[![Go Report Card](https://goreportcard.com/badge/github.com/txpull/solgo)](https://goreportcard.com/report/github.com/txpull/solgo)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/txpull/solgo)](https://pkg.go.dev/github.com/txpull/solgo)

# Solidity Parser in Golang

SolGo contains a Solidity parser written in Golang, using Antlr and AntlrGo for grammar parsing. The aim of this project is to provide a tool to parse Solidity source code into a structured format, enabling further analysis and manipulation within Go programs.

The parser is generated from a Solidity grammar file using Antlr, producing a lexer, parser, and listener in Golang using AntlrGo. This allows for the syntactic analysis of Solidity code, transforming it into a parse tree that can be traversed and manipulated.

This project is ideal for developers working with Solidity smart contracts who wish to leverage the power and efficiency of Golang for their analysis tools.

**Note:** This project is still in development and is not yet ready for use in production.

## Why?

I prefer to use Golang/Rust for my projects instead of Javascript. I wanted to be able to parse Solidity code and do some analysis on it. I found a few projects that were doing this, but they were either incomplete or not maintained. I decided to create this project to provide a Solidity parser in Golang that is complete and maintained.

Basically I wanted to be able provide a solidity source code and extract as much information about it possible, including ABIs that I can trust.

This project will be used within the [unpack](https://github.com/txpull/unpack) as I've discovered that for many deployed contracts I have source code, but not ABIs.

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


## Contributing

Contributions to SolGo are always welcome! If you have a feature request, bug report, or proposal for improvement, please open an issue. If you wish to contribute code, please fork the repository, make your changes, and submit a pull request.


## License

SolGo is licensed under the Apache 2.0. See [LICENSE](LICENSE) for the full license text.


## Acknowledgements

We would like to express our gratitude to the Solidity team for maintaining the Solidity grammar files, and to the Antlr and AntlrGo team for providing the powerful Antlr tool that makes this project possible.