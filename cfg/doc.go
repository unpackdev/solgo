// Package cfg offers a comprehensive toolkit for constructing and visualizing
// control flow graphs (CFGs) of Solidity smart contracts.
//
// Key Features:
// - Initialization of CFG builders with context and solgo IR builders.
// - Seamless integration with the go-graphviz library for graph operations.
// - Recursive traversal of the IR to build nodes and edges for the CFG.
// - Capability to render the CFG into various formats including DOT and PNG.
// - Error handling and resource management for efficient graph operations.
package cfg
