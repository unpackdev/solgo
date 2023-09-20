// Package clients provides tools for managing and interacting with Ethereum clients.
//
// The package offers a ClientPool structure that manages a pool of Ethereum clients
// for different networks and types. The ClientPool allows for retrieving clients based
// on various criteria, such as group and type, in a round-robin fashion. It also provides
// functionality to close all clients in the pool.
//
// Additionally, the package provides a Client structure that wraps the Ethereum client
// with additional context and options. This structure offers methods to retrieve various
// details about the client, such as its network ID, group, type, and endpoint.
//
// The package is designed to be flexible and efficient, ensuring that Ethereum clients
// can be easily managed and accessed based on the specific needs of the application.
package clients
