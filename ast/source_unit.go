package ast

type SourceUnit struct {
	Comments        []*CommentNode    `json:"comments,omitempty"`
	AbsolutePath    string            `json:"absolute_path,omitempty"`
	ExportedSymbols []ExportedSymbols `json:"exported_symbols,omitempty"`
	ID              int64             `json:"id"`
	License         string            `json:"license"`
	NodeType        string            `json:"node_type"`
	Nodes           []Node            `json:"nodes"`
	Src             Src               `json:"src"`
}

type ExportedSymbols struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type RootSourceUnit struct {
	SourceUnits []SourceUnit `json:"source_units"`
}
