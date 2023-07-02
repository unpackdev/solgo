// Package contracts provides a set of utilities and listeners for working with Solidity contracts.
// It includes a ContractListener, which is a listener for the Solidity parser that extracts information
// about contracts, such as the contract name, implemented interfaces, imported contracts, pragmas, and comments.
// The ContractListener is designed to be used in conjunction with the Solidity parser to provide a convenient
// interface for working with Solidity contracts.
//
// The ContractListener extracts information from Solidity contracts by traversing the Solidity parse tree
// and capturing relevant information during the parsing process. It identifies pragma directives, import directives,
// contract definitions, inheritance specifiers, using directives, and comments within the contract source code.
// Additionally, it can detect SPDX license identifiers if present.
//
// The extracted contract information can be accessed using the provided getter methods, which return various aspects
// of the contract, such as the contract name, implemented interfaces, imported contracts, pragmas, comments, and SPDX license.
// The ContractListener also provides a convenience method, ToStruct, that returns a struct containing all the extracted information.
package contracts
