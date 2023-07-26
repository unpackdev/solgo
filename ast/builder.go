package ast

import (
	"encoding/json"
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
	comments              []*CommentNode
	commentsParsed        bool
	astRoot               *RootNode
	entrySourceUnit       *ast_pb.Node
	sourceUnits           []*SourceUnit[Node[ast_pb.SourceUnit]]
	currentSourceUnit     *ast_pb.SourceUnit
	currentStateVariables []*StateVariableDeclaration
	currentEvents         []*ast_pb.Node
	currentEnums          []*ast_pb.Node
	currentStructs        []*ast_pb.Node
	currentErrors         []*ast_pb.Node
}

func NewAstBuilder(parser *parser.SolidityParser, sources solgo.Sources) *ASTBuilder {
	return &ASTBuilder{
		parser:      parser,
		sources:     sources,
		comments:    make([]*CommentNode, 0),
		sourceUnits: make([]*SourceUnit[Node[ast_pb.SourceUnit]], 0),
		nextID:      1,
	}
}

func (b *ASTBuilder) GetRoot() *RootNode {
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
	/* for _, unit := range b.sourceUnits {

				for _, root := range unit.Nodes.Nodes {
			fmt.Println(root.NodeType)
		}
	} */
	return nodes, false
}

func (b *ASTBuilder) FindNodeById(nodeId int64) (*ast_pb.Node, bool) {
	/* 	for _, unit := range b.astRoot.GetSourceUnits() {
	   		for _, node := range unit.GetRoot().GetNodes() {
	   			if node.GetId() == nodeId {
	   				return node, true
	   			}
	   		}
	   	}
	*/
	return nil, false
}

func (b *ASTBuilder) ResolveReferences() error {
	/* 	for _, unit := range b.astRoot.GetSourceUnits() {
		for _, node := range unit.GetRoot().GetNodes() {
			fmt.Println(node.NodeType)
		}
	} */
	return nil
}
