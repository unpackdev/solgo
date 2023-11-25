package inspector

import (
	"github.com/ethereum/go-ethereum/common"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

type Mintable struct {
	Enabled            bool              `json:"enabled"`
	Visibility         ast_pb.Visibility `json:"visibility"`
	ExternallyCallable bool              `json:"externally_callable"`
	Function           *ast.Function     `json:"function"`
}

func (m Mintable) IsEnabled() bool {
	return m.Enabled
}

func (m Mintable) IsVisible() bool {
	return m.Visibility == ast_pb.Visibility_PUBLIC || m.Visibility == ast_pb.Visibility_EXTERNAL
}

type Burnable struct {
	Enabled bool `json:"enabled"`
}

type StateVariable struct {
	Variable *ast.VariableDeclaration `json:"variable"`
}

type Transfer struct {
	Function   *ast.Function
	Expression *ast.MemberAccessExpression
}

type Report struct {
	Addresses      []common.Address `json:"addresses"`
	StateVariables []StateVariable  `json:"state_variables"`
	UsesTransfers  bool             `json:"uses_transfers"`
	Transfers      []Transfer       `json:"transfers"`
	Mintable       Mintable         `json:"mintable"`
	Burnable       Burnable         `json:"burnable"`
}
