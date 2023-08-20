![Build Status](https://github.com/txpull/solgo/actions/workflows/test.yml/badge.svg)
![Security Status](https://github.com/txpull/solgo/actions/workflows/gosec.yml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/txpull/solgo/badge.svg?branch=main)](https://coveralls.io/github/txpull/solgo?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/txpull/solgo)](https://goreportcard.com/report/github.com/txpull/solgo)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/txpull/solgo)](https://pkg.go.dev/github.com/txpull/solgo)
[![OpenSSF Best Practices](https://bestpractices.coreinfrastructure.org/projects/7719/badge)](https://bestpractices.coreinfrastructure.org/projects/7719)

# Solidity Parser and Analyzer in Go

**SolGo** - a robust tool crafted in Go, designed to dissect and analyze Solidity's source code.

The parser is generated from a Solidity grammar file using **Antlr**, producing a lexer, parser, and listener using **AntlrGo**. This allows for the syntactic analysis of Solidity code, transforming it into a parse tree that offers offers a detailed syntactic representation of the code, allowing for intricate navigation and manipulation.

This project is ideal for those diving into data analysis, construction of robust APIs, developing advanced analysis tools, enhancing smart contract security, and anyone keen on harnessing Go for their Solidity endeavors.

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

- **Protocol Buffers**: Utilizing [Protocol Buffers](https://github.com/txpull/protos), SolGo offers a structured data format, paving the way for enhanced analysis and facilitating a unified interface for diverse tools. Currently, it supports Go and Javascript, with plans to incorporate Rust and Python in upcoming versions.
- **Abstract Syntax Tree (AST) Generation:** Package `ast` is equipped with a dedicated builder that crafts an Abstract Syntax Tree (AST) tailored for Solidity code.
- **Intermediate Representation (IR) Generation**: From the AST, SolGo is adept at generating an Intermediate Representation (IR). `ir` package serves as a language-neutral depiction of the contract, encapsulating pivotal components like functions, state variables, and events, thus broadening the scope for intricate analysis and contract manipulation.
- **Application Binary Interface (ABI) Generation:** SolGo's in-built `abi` package can interpret contract definitions, enabling the generation of ABI for a collective group of contracts or individual ones.
- **Opcode Tools**: The `opcode` package in SolGo demystifies bytecode by decompiling it into opcodes. Additionally, it provides tools for the creation and visualization of opcode execution trees, granting a holistic perspective of opcode sequences in smart contracts.
- **Library Integration**: SolGo is programmed to autonomously source and assimilate Solidity contracts from renowned libraries, notably [OpenZeppelin](https://github.com/OpenZeppelin/openzeppelin-contracts). This feature enables users to seamlessly import and utilize contracts from these libraries without the need for manual integration.
- **EIP & ERC Registry**: SolGo introduces a package `eip` exclusively for Ethereum Improvement Proposals (EIPs) and Ethereum Request for Comments (ERCs). This package streamlines interactions with diverse contract standards by encompassing functions, events, and a registry system optimized for proficient management.
- **Solidity Compiler Detection & Compilation:** SolGo intelligently identifies the Solidity version employed for contract compilation. This not only streamlines the process of determining the compiler version but also equips users with the capability to seamlessly compile contracts.
- **Security Audit Package**: Prioritizing security, SolGo has incorporated an `audit` package. This specialized package leverages [Slither](https://github.com/crytic/slither)'s sophisticated algorithms to scrutinize and pinpoint potential vulnerabilities in Solidity smart contracts, ensuring robust protection against adversarial threats.
- **Contract Bytecode Validation:** Enhanced `validation` package ensures the integrity and authenticity of contract bytecode. By comparing the bytecode of a deployed contract with the expected bytecode generated from its source code, SolGo can detect any discrepancies or potential tampering. This feature is crucial for verifying that a deployed contract's bytecode corresponds accurately to its source code, providing an added layer of security and trust for developers and users alike.

## Contributing

Contributions to SolGo are always welcome! Please visit [Contributing](https://github.com/txpull/solgo/wiki/Contributing) for more information on how to get started.


## License

SolGo is licensed under the Apache 2.0. See [LICENSE](LICENSE) for the full license text.


## Acknowledgements

We would like to express our gratitude to the Solidity team for maintaining the Solidity grammar files, and to the Antlr and AntlrGo team for providing the powerful Antlr tool that makes this project possible. Not to forget the Crytic and Slither contributors for their hard work on the Slither tool as without it we would have hard time detecting vulnerabilities in smart contracts.