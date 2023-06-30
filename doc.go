// Package solgo provides a suite of tools for parsing, analyzing, and interacting with Solidity contracts.
// It includes a contextual parser that maintains a stack of contexts as it parses a contract,
// allowing it to keep track of the current context (e.g., within a contract definition, function definition, etc.).
// It also includes a contract listener that extracts information about contracts as they are parsed,
// including the contract name, implemented interfaces, imported contracts, pragmas, and comments.
// Additionally, it includes a syntax error listener that listens for syntax errors in contracts and categorizes them by severity.
// The package also provides functionality for generating and working with Ethereum contract ABIs (Application Binary Interfaces).
// This includes parsing contract definitions to extract ABI information, normalizing type names, and handling complex types like mappings.
// These tools can be used together to provide a comprehensive interface for working with Solidity contracts,
// making it easier to understand their structure, identify potential issues, and interact with them on the Ethereum network.
package solgo
