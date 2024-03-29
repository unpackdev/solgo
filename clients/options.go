package clients

// Options represents the configuration options for network nodes.
type Options struct {
	// Nodes is a slice of Node representing the network nodes.
	Nodes []Node `mapstructure:"nodes" yaml:"nodes" json:"nodes"`
}

// GetNodes returns the slice of network nodes from the Options.
func (o *Options) GetNodes() []Node {
	return o.Nodes
}

// Node represents the configuration and details of a network node.
type Node struct {
	// Group represents the group name of the node.
	Group string `mapstructure:"group" yaml:"group" json:"group"`

	// Type represents the type of the node.
	Type string `mapstructure:"type" yaml:"type" json:"type"`

	// FailoverGroup represents the failover group name of the node.
	FailoverGroup string `mapstructure:"failoverGroup" yaml:"failoverGroup" json:"failoverGroup"`

	// FailoverType represents the type of failover for the node.
	FailoverType string `mapstructure:"failoverType" yaml:"failoverType" json:"failoverType"`

	// NetworkId represents the network ID of the node.
	NetworkId int `mapstructure:"networkId" yaml:"networkId" json:"networkId"`

	// Endpoint represents the network endpoint of the node.
	Endpoint string `mapstructure:"endpoint" yaml:"endpoint" json:"endpoint"`

	// ConcurrentClients represents the number of concurrent clients for the node.
	ConcurrentClients int `mapstructure:"concurrentClients" yaml:"concurrentClients" json:"concurrentClients"`
}

// GetGroup returns the group name of the node.
func (n *Node) GetGroup() string {
	return n.Group
}

// GetType returns the type of the node.
func (n *Node) GetType() string {
	return n.Type
}

// GetNetworkID returns the network ID of the node.
func (n *Node) GetNetworkID() int64 {
	return int64(n.NetworkId)
}

// GetEndpoint returns the network endpoint of the node.
func (n *Node) GetEndpoint() string {
	return n.Endpoint
}

// GetFailoverGroup returns the failover group name of the node.
func (n *Node) GetFailoverGroup() string {
	return n.FailoverGroup
}

// GetFailoverType returns the type of failover for the node.
func (n *Node) GetFailoverType() string {
	return n.FailoverType
}

// GetConcurrentClientsNumber returns the number of concurrent clients for the node.
func (n *Node) GetConcurrentClientsNumber() int {
	return n.ConcurrentClients
}
