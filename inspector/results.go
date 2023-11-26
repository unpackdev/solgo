package inspector

import (
	"github.com/ethereum/go-ethereum/common"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/standards"
)

type Proxy struct {
	Enabled  bool               `json:"enabled"`
	Standard standards.Standard `json:"standard"`
}

type Burnable struct {
	Enabled            bool              `json:"enabled"`
	Visibility         ast_pb.Visibility `json:"visibility"`
	ExternallyCallable bool              `json:"externally_callable"`
	Function           *ast.Function     `json:"function"`
}

func (b Burnable) IsEnabled() bool {
	return b.Enabled
}

func (b Burnable) IsVisible() bool {
	return b.Visibility == ast_pb.Visibility_PUBLIC || b.Visibility == ast_pb.Visibility_EXTERNAL
}

type StateVariable struct {
	Variable *ast.VariableDeclaration `json:"variable"`
}

type Transfer struct {
	Function   *ast.Function
	Expression *ast.MemberAccessExpression
}

type Report struct {
	Addresses     []common.Address     `json:"addresses"`
	UsesTransfers bool                 `json:"uses_transfers"`
	Detectors     map[DetectorType]any `json:"detectors"`
}
