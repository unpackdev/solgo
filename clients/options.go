package clients

type Options struct {
	Nodes []Node `mapstructure:"network_nodes" yaml:"network_nodes" json:"network_nodes"`
}

func (o *Options) GetNodes() []Node {
	return o.Nodes
}

type Node struct {
	Group                   string `mapstructure:"group" yaml:"group" json:"group"`
	Type                    string `mapstructure:"type" yaml:"type" json:"type"`
	NetworkId               int64  `mapstructure:"network_id" yaml:"network_id" json:"network_id"`
	Endpoint                string `mapstructure:"endpoint" yaml:"endpoint" json:"endpoint"`
	ConcurrentClientsNumber int    `mapstructure:"concurrent_clients_number" yaml:"concurrent_clients_number" json:"concurrent_clients_number"`
}

func (n *Node) GetGroup() string {
	return n.Group
}

func (n *Node) GetType() string {
	return n.Type
}

func (n *Node) GetNetworkID() int64 {
	return n.NetworkId
}

func (n *Node) GetEndpoint() string {
	return n.Endpoint
}

func (n *Node) GetConcurrentClientsNumber() int {
	return n.ConcurrentClientsNumber
}
