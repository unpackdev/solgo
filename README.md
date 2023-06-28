# Solidity Parser in Golang

SolGo contains a Solidity parser written in Golang, using Antlr for grammar parsing. The aim of this project is to provide a tool to parse Solidity source code into a structured format, enabling further analysis and manipulation within Go programs.

The parser is generated from a Solidity grammar file using Antlr, producing a lexer, parser, and listener in Golang. This allows for the syntactic analysis of Solidity code, transforming it into a parse tree that can be traversed and manipulated.

Please note that this project currently only supports syntactic analysis. Semantic analysis, such as type checking, is not yet implemented.

This project is ideal for developers working with Solidity smart contracts who wish to leverage the power and efficiency of Golang for their analysis tools.


## ANTLR Grammar

We are using grammar files that are maintained by the Solidity team.
Link to the grammar files can be found [here](https://github.com/ethereum/solidity/tree/develop/docs/grammar)