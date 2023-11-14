package accounts

// Options defines the configuration parameters for account management.
type Options struct {
	// KeystorePath specifies the file system path to the directory where the keystore files are stored.
	// The keystore is used to securely store the private keys of Ethereum accounts.
	KeystorePath string `json:"keystore_path" yaml:"keystore_path"`

	// SupportedNetworks lists the Ethereum based networks that the account manager will interact with.
	// Each network has a corresponding keystore and set of account configurations.
	SupportedNetworks []Network `json:"supported_networks" yaml:"supported_networks"`
}
