// Package abis provides functionality for parsing and manipulating Solidity contract ABIs (Application Binary Interfaces).
// It includes features for normalizing type names, parsing mapping types, handling struct types, detecting mapping and struct types,
// and converting ABIs to JSON and Go-ethereum ABI formats.
//
// The package consists of several key components:
//
//   - The AbiListener struct is a listener for the Solidity parser that constructs an ABI as it walks the parse tree.
//     It extends the BaseSolidityParserListener and uses an AbiParser to build the ABI.
//
//   - The AbiParser struct is responsible for parsing a Solidity contract ABI and converting it into an ABI object
//     that can be easily manipulated. It maintains an internal representation of the ABI and provides methods for injecting
//     various contract elements, such as constructors, functions, events, errors, state variables, and fallback/receive functions,
//     into the ABI. It also handles the resolution of struct components and supports the parsing of mapping and enum types.
//
//   - The package includes functions for normalizing type names in Solidity to their canonical forms. For example, it can
//     convert "uint" to "uint256" and "addresspayable" to "address".
//
//   - The package provides functions for detecting mapping types, struct types, and enum types in Solidity. It can determine
//     if a given type name represents a mapping, if a type name corresponds to a defined struct, or if a type is an enumerated type.
//
//   - The package offers a function for parsing mapping types in Solidity ABI. It can extract the key and value types from a
//     mapping type string of the form "mapping(keyType => valueType)".
//
//   - The AbiParser can convert the internal ABI representation into a JSON string or an Ethereum/go-ethereum ABI object.
//     It provides methods for converting the ABI to these formats.
//
// This package is part of a larger system for working with Solidity contracts and their ABIs.
package abis
