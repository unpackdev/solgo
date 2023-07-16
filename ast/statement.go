package ast

/* type Statement struct {
	Expression *Expression `json:"expression"`
	ID         int64       `json:"id"`
	NodeType   string      `json:"node_type"`
	Src        ast_pb.Src  `json:"src"`
}

type Expression struct {
	ID                     int64             `json:"id"`
	IsConstant             bool              `json:"is_constant"`
	IsLValue               bool              `json:"is_l_value"`
	IsPure                 bool              `json:"is_pure"`
	LValueRequested        bool              `json:"l_value_requested"`
	LeftHandSide           *LeftHandSide     `json:"left_hand_side,omitempty"`
	NodeType               string            `json:"node_type"`
	Operator               string            `json:"operator,omitempty"`
	RightHandSide          *RightHandSide    `json:"right_hand_side,omitempty"`
	Src                    ast_pb.Src        `json:"src"`
	TypeDescriptions       *TypeDescriptions `json:"type_descriptions"`
	Name                   string            `json:"name,omitempty"`
	ReferencedDeclarations []int             `json:"referenced_declaration,omitempty"`
	OverloadedDeclarations []interface{}     `json:"overloaded_declarations,omitempty"`
	Arguments              []Argument        `json:"arguments,omitempty"`
	Expression             *Expression       `json:"expression,omitempty"`
	MemberName             string            `json:"member_name,omitempty"`
	Kind                   string            `json:"kind,omitempty"`
	TryCall                bool              `json:"try_call,omitempty"`
}

type LeftHandSide struct {
	ID                     int64             `json:"id"`
	Name                   string            `json:"name"`
	NodeType               string            `json:"node_type"`
	OverloadedDeclarations []interface{}     `json:"overloaded_declarations"`
	ReferencedDeclaration  int               `json:"referenced_declaration"`
	Src                    ast_pb.Src        `json:"src"`
	TypeDescriptions       *TypeDescriptions `json:"type_descriptions"`
}

type RightHandSide struct {
	Arguments        []Argument        `json:"arguments"`
	Expression       *Expression       `json:"expression"`
	ID               int64             `json:"id"`
	IsConstant       bool              `json:"is_constant"`
	IsLValue         bool              `json:"is_l_value"`
	IsPure           bool              `json:"is_pure"`
	Kind             string            `json:"kind"`
	LValueRequested  bool              `json:"l_value_requested"`
	Names            []interface{}     `json:"names"`
	NodeType         string            `json:"node_type"`
	Src              ast_pb.Src        `json:"src"`
	TryCall          bool              `json:"try_call"`
	TypeDescriptions *TypeDescriptions `json:"type_descriptions"`
}

type Argument struct {
	ID                     int64             `json:"id"`
	Name                   string            `json:"name"`
	NodeType               string            `json:"node_type"`
	OverloadedDeclarations []interface{}     `json:"overloaded_declarations"`
	ReferencedDeclaration  int               `json:"referenced_declaration"`
	Src                    ast_pb.Src        `json:"src"`
	TypeDescriptions       *TypeDescriptions `json:"type_descriptions"`
}
*/
