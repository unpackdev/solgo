[![Build Status](https://app.travis-ci.com/txpull/solgo.svg?branch=main)](https://app.travis-ci.com/txpull/solgo)
[![Coverage Status](https://coveralls.io/repos/github/txpull/solgo/badge.svg?branch=main)](https://coveralls.io/github/txpull/solgo?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/txpull/solgo)](https://goreportcard.com/report/github.com/txpull/solgo)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/txpull/solgo)](https://pkg.go.dev/github.com/txpull/solgo)

# Solidity Parser in Go

SolGo contains a Solidity parser written in Go, using Antlr and AntlrGo for grammar parsing. The aim of this project is to provide a tool to parse Solidity source code into a structured format, enabling further analysis.

The parser is generated from a Solidity grammar file using Antlr, producing a lexer, parser, and listener using AntlrGo. This allows for the syntactic analysis of Solidity code, transforming it into a parse tree that can be traversed and manipulated.

This project is ideal for developers working with Solidity smart contracts who wish to leverage the power and efficiency of Go for their analysis tools.

## Disclaimer

Please be aware that this project is still under active development. While it is approaching a state suitable for production use, there may still be undiscovered issues or limitations. Over the next few weeks, extensive testing will be conducted on 1-2 million contracts to evaluate its performance and stability. Additional tests and comprehensive documentation will also be added during this phase.

Once we are confident that the project is fully ready for production, this disclaimer will be removed. Until then, please use the software with caution and report any potential issues or feedback to help us improve its quality.

## Documentation

The SolGo basic documentation is hosted on GitHub, ensuring it's always up-to-date with the latest changes and features. You can access the full documentation [here](https://github.com/txpull/solgo/wiki).

## Getting Started

Detailed examples of how to install and use this package can be found in the [Usage](https://github.com/txpull/solgo/wiki/Getting-Started) section.

## Solidity Language Grammar

Latest Solidity language grammar higher overview and detailed description can be found [here](https://docs.soliditylang.org/en/v0.8.19/grammar.html).

## ANTLR Grammar

We are using grammar files that are maintained by the Solidity team.
Link to the grammar files can be found [here](https://github.com/ethereum/solidity/tree/develop/docs/grammar).

## ANTLR Go

We are using the ANTLR4 Go runtime library to generate the parser. Repository can be found [here](https://github.com/antlr4-go/antlr).

## Crytic Slither

We are using Slither to detect vulnerabilities in smart contracts. Repository can be found [here](https://github.com/crytic/slither).

Makes no sense to rewrite all of that hard work just to be written in Go. Therefore, a bit of python will not hurt. In the future we may change direction.


## Features

- **Protocol Buffers**: SolGo uses [Protocol Buffers](https://github.com/txpull/protos) to provide a structured format for the data, enabling more effective analysis and a way to build a common interface for other tools. Supported languages: Go and Javascript. In the future, will be adding Rust and Python.
- **Abstract Syntax Tree (AST) Generation:** SolGo includes an builder that constructs an Abstract Syntax Tree for Solidity code.
- **Intermediate Representation (IR) Generation**: SolGo can generate an Intermediate Representation (IR) from the AST. The IR provides a language-agnostic representation of the contract, capturing key elements such as functions, state variables, events, and more. This enables more advanced analysis and manipulation of the contract.
- **Application Binary Interface (ABI) Generation:** SolGo includes an builder that can parse contract definitions to generate ABI for group of contracts or each contract individually. 
- **Opcode Decompilation and Execution Trees:** The opcode package facilitates the decompilation of bytecode into opcodes, and offers tools for constructing and visualizing opcode execution trees. This provides a comprehensive view of opcode sequences in smart contracts, aiding in deeper analysis and understanding.
- **Syntax Error Handling**: SolGo includes listener which collects syntax errors encountered during parsing, providing detailed error information including line, column, message, severity, and context.
- **Automatic Source Detection**: SolGo automatically loads and integrates Solidity contracts from well-known libraries such as [OpenZeppelin](https://github.com/OpenZeppelin/openzeppelin-contracts).
- **Ethereum Improvement Proposals (EIP) Registry**: A package designed to provide a structured representation of Ethereum Improvement Proposals (EIPs) and Ethereum Request for Comments (ERCs). It simplifies the interaction with various contract standards by including functions, events, and a registry mechanism crafted for efficient management.


## Contributing

Contributions to SolGo are always welcome! Please visit [Contributing](https://github.com/txpull/solgo/wiki/Contributing) for more information on how to get started.


## License

SolGo is licensed under the Apache 2.0. See [LICENSE](LICENSE) for the full license text.


## Acknowledgements

We would like to express our gratitude to the Solidity team for maintaining the Solidity grammar files, and to the Antlr and AntlrGo team for providing the powerful Antlr tool that makes this project possible. Not to forget the Slither team for their hard work on the Slither tool as without it we would have hard time detecting vulnerabilities in smart contracts.