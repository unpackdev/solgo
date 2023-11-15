package standards

import (
	"strings"

	eip_pb "github.com/unpackdev/protos/dist/go/eip"
)

// Standard represents the type for Ethereum standards and EIPs.
type Standard string

// String returns the string representation of the Standard.
func (s Standard) String() string {
	return string(s)
}

// ToProto() converts a string representation of an Ethereum standard
// to its corresponding protobuf enum value. If the standard is not recognized,
// it returns unknown.
func (s Standard) ToProto() eip_pb.Standard {
	if standardValue, ok := eip_pb.Standard_value[strings.ToUpper(string(s))]; ok {
		return eip_pb.Standard(standardValue)
	}
	return eip_pb.Standard_UNKNOWN
}

// Constants representing various Ethereum standards and EIPs.
const (
	ERC20                       Standard = "ERC20"                    // ERC-20 Token Standard.
	ERC721                      Standard = "ERC721"                   // ERC-721 Non-Fungible Token Standard.
	ERC1822                     Standard = "ERC1822"                  // ERC-1822 Universal Proxy Standard (UPS).
	ERC1820                     Standard = "ERC1820"                  // ERC-1820 Pseudo-introspection Registry Contract.
	ERC777                      Standard = "ERC777"                   // ERC-777 Token Standard.
	ERC1155                     Standard = "ERC1155"                  // ERC-1155 Multi Token Standard.
	ERC1337                     Standard = "ERC1337"                  // ERC-1337 Subscription Standard.
	ERC1400                     Standard = "ERC1400"                  // ERC-1400 Security Token Standard.
	ERC1410                     Standard = "ERC1410"                  // ERC-1410 Partially Fungible Token Standard.
	ERC165                      Standard = "ERC165"                   // ERC-165 Standard Interface Detection.
	ERC820                      Standard = "ERC820"                   // ERC-820 Registry Standard.
	ERC1014                     Standard = "ERC1014"                  // ERC-1014 Create2 Standard.
	ERC1948                     Standard = "ERC1948"                  // ERC-1948 Non-Fungible Data Token Standard.
	ERC1967                     Standard = "ERC1967"                  // ERC-1967 Proxy Storage Slots Standard.
	ERC2309                     Standard = "ERC2309"                  // ERC-2309 Consecutive Transfer Standard.
	ERC2535                     Standard = "ERC2535"                  // ERC-2535 Diamond Standard.
	ERC2771                     Standard = "ERC2771"                  // ERC-2771 Meta Transactions Standard.
	ERC2917                     Standard = "ERC2917"                  // ERC-2917 Interest-Bearing Tokens Standard.
	ERC3156                     Standard = "ERC3156"                  // ERC-3156 Flash Loans Standard.
	ERC3664                     Standard = "ERC3664"                  // ERC-3664 BitWords Standard.
	CAMELOT_POOL                Standard = "CamelotPool"              // Camelot Pool.
	SUSHISWAP_V2_POOL           Standard = "SushiSwapV2Pool"          // Sushiswap V2 Pool.
	UNISWAP_UNIVERSAL_ROUTER    Standard = "UniswapUniversalRouter"   // Uniswap Universal Router.
	UNISWAP_UNIVERSAL_ROUTER_V2 Standard = "UniswapUniversalRouterV2" // Uniswap Universal Router V2.
	UNISWAP_V2_FACTORY          Standard = "UniswapV2Factory"         // Uniswap V2 Factory.
	UNISWAP_V2_POOL             Standard = "UniswapV2Pool"            // Uniswap V2 Pool.
	UNISWAP_V2_ROUTER           Standard = "UniswapV2Router"          // Uniswap V2 Router.
	UNISWAP_V3_FACTORY          Standard = "UniswapV3Factory"         // Uniswap V3 Factory.
	UNISWAP_V3_POOL             Standard = "UniswapV3Pool"            // Uniswap V3 Pool.
	UNISWAP_V3_ROUTER           Standard = "UniswapV3Router"          // Uniswap V3 Router.
	UNISWAP_V3_ROUTER2          Standard = "UniswapV3Router2"         // Uniswap V3 Router2.
	UNISWAP_V3_TICKLENS         Standard = "UniswapV3Ticklens"        // Uniswap V3 TickLens.
)

// LoadStandards loads list of supported Ethereum EIPs into the registry.
func LoadStandards() error {
	for name, standard := range standards {
		if err := RegisterStandard(name, NewContract(standard)); err != nil {
			return err
		}
	}

	return nil
}
