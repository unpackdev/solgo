package cfg

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/txpull/solgo/ir"
)

// Builder is responsible for constructing the control flow graph.
type Builder struct {
	ctx     context.Context    // Context for the builder operations
	builder *ir.Builder        // IR builder from solgo
	viz     *graphviz.Graphviz // Graphviz instance for graph operations
}

// NewBuilder initializes a new CFG builder.
func NewBuilder(ctx context.Context, builder *ir.Builder) *Builder {
	return &Builder{
		ctx:     ctx,
		builder: builder,
		viz:     graphviz.New(),
	}
}

// Close releases any resources used by the Graphviz instance.
func (b *Builder) Close() error {
	return b.viz.Close()
}

// GetGraphviz returns the underlying Graphviz instance.
func (b *Builder) GetGraphviz() *graphviz.Graphviz {
	return b.viz
}

// Build constructs the control flow graph for the given IR.
func (b *Builder) Build() (*cgraph.Graph, error) {
	if b.viz == nil {
		return nil, errors.New("graphviz instance is not set")
	}

	root := b.builder.GetRoot()
	if root == nil {
		return nil, errors.New("root node is not set")
	}

	graph, err := b.viz.Graph()
	if err != nil {
		return nil, err
	}

	if _, err := b.traverseIR(root, graph); err != nil {
		return nil, err
	}

	return graph, nil
}

// traverseIR recursively traverses the IR to build nodes and edges for the graph.
func (b *Builder) traverseIR(root *ir.RootSourceUnit, graph *cgraph.Graph) (*cgraph.Node, error) {
	rootNode, err := graph.CreateNode("You")
	if err != nil {
		return nil, err
	}

	nodeMap := make(map[string]*cgraph.Node)
	nodeMap["You"] = rootNode

	if len(root.Contracts) == 0 {
		return nil, nil
	}

	for _, contract := range root.Contracts {
		// Create a subgraph for the contract with the "cluster" prefix
		clusterName := fmt.Sprintf("cluster_%s", contract.GetName())
		contractSubGraph := graph.SubGraph(clusterName, 1)
		contractSubGraph.SetLabel(contract.GetName())

		// Create a node for the contract within the subgraph
		contractNode, err := contractSubGraph.CreateNode(contract.GetName())
		if err != nil {
			return nil, err
		}

		// Link the rootNode to the contractNode
		if _, err := graph.CreateEdge("", rootNode, contractNode); err != nil {
			return nil, err
		}

		// Traverse functions within the contract (assuming there's a method to get functions)
		for _, function := range contract.GetFunctions() {
			funcNode, err := contractSubGraph.CreateNode(function.GetName())
			if err != nil {
				return nil, err
			}
			nodeMap[function.GetName()] = funcNode

			// Link the contractNode to the funcNode
			if _, err := graph.CreateEdge("", contractNode, funcNode); err != nil {
				return nil, err
			}

			refFns := b.builder.LookupReferencedFunctionsByNode(function.GetAST())
			for _, refFn := range refFns {
				refFnNode, exists := nodeMap[refFn.GetName()]
				if !exists {
					refFnNode, err = graph.CreateNode(refFn.GetName())
					if err != nil {
						return nil, err
					}
					nodeMap[refFn.GetName()] = refFnNode
				}

				// Create an edge from the current function node to the referenced function node
				if _, err := graph.CreateEdge("", funcNode, refFnNode); err != nil {
					return nil, err
				}
			}
		}
	}

	return rootNode, nil
}

// GenerateDOT produces the DOT representation of the given graph.
func (b *Builder) GenerateDOT(graph *cgraph.Graph) (string, error) {
	if b.viz == nil {
		return "", errors.New("graphviz instance is not set")
	}

	if graph == nil {
		return "", errors.New("graph is not set")
	}

	var buf bytes.Buffer
	if err := b.viz.Render(graph, "dot", &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// SaveAs renders the graph to a file in the specified format.
func (b *Builder) SaveAs(graph *cgraph.Graph, format graphviz.Format, file string) error {
	if b.viz == nil {
		return errors.New("graphviz instance is not set")
	}

	if graph == nil {
		return errors.New("graph is not set")
	}

	if err := b.viz.RenderFilename(graph, format, file); err != nil {
		return err
	}
	return nil
}
