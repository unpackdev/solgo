package ast

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/parser"
)

type ASTBuilder struct {
	*parser.BaseSolidityParserListener
	sources               solgo.Sources          // sources is the source code of the Solidity files.
	parser                *parser.SolidityParser // parser is the Solidity parser instance.
	nextID                int64                  // nextID is the next ID to assign to a node.
	comments              []*ast_pb.Comment
	commentsParsed        bool
	sourceUnits           []*ast_pb.SourceUnit
	currentSourceUnit     *ast_pb.SourceUnit
	currentStateVariables []*ast_pb.Node
	currentEvents         []*ast_pb.Node
	currentEnums          []*ast_pb.Node
	astRoot               *ast_pb.RootSourceUnit
	entrySourceUnit       *ast_pb.Node
}

func NewAstBuilder(parser *parser.SolidityParser, sources solgo.Sources) *ASTBuilder {
	return &ASTBuilder{
		parser:   parser,
		sources:  sources,
		comments: make([]*ast_pb.Comment, 0),
		nextID:   1,
	}
}

func (b *ASTBuilder) GetRoot() *ast_pb.RootSourceUnit {
	return b.astRoot
}

func (b *ASTBuilder) ToJSON() ([]byte, error) {
	return json.Marshal(b.astRoot)
}

func (b *ASTBuilder) ToJSONString() (string, error) {
	bts, err := b.ToJSON()
	if err != nil {
		return "", err
	}
	return string(bts), nil
}

func (b *ASTBuilder) ToPrettyJSON(data interface{}) ([]byte, error) {
	return json.MarshalIndent(data, "", "  ")
}

func (b *ASTBuilder) WriteJSONToFile(path string) error {
	bts, err := b.ToJSON()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bts, 0644)
}

func (b *ASTBuilder) WriteToFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}

func (b *ASTBuilder) NodeToJson(node interface{}) ([]byte, error) {
	return json.Marshal(node)
}

func (b *ASTBuilder) FindNodesByType(nodeType ast_pb.NodeType) ([]*ast_pb.Node, bool) {
	var nodes []*ast_pb.Node
	for _, unit := range b.sourceUnits {
		fmt.Println(unit)

		for _, root := range unit.Root.Nodes {
			fmt.Println(root.NodeType)
		}
	}
	return nodes, false
}

func (b *ASTBuilder) FindNodeById(nodeId int64) (*ast_pb.Node, bool) {
	for _, unit := range b.astRoot.GetSourceUnits() {
		for _, node := range unit.GetRoot().GetNodes() {
			if node.GetId() == nodeId {
				return node, true
			}
		}
	}

	return nil, false
}
