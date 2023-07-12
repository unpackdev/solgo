package metadata

// Provider is the interface that wraps the basic interaction with contract metadata from different
// sources such as IPFS, SWARM, etc....
type Provider interface {

	// GetMetadataByCID returns the metadata of a contract by the CID (Content Identifier) of the contract
	GetMetadataByCID(cid string) (*ContractMetadata, error)
}
