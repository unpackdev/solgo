package ast

type Statement struct {
	Expression *Expression `json:"expression"`
	ID         int         `json:"id"`
	NodeType   string      `json:"node_type"`
	Src        string      `json:"src"`
}

type Expression struct {
	ID                     int               `json:"id"`
	IsConstant             bool              `json:"is_constant"`
	IsLValue               bool              `json:"is_l_value"`
	IsPure                 bool              `json:"is_pure"`
	LValueRequested        bool              `json:"l_value_requested"`
	LeftHandSide           *LeftHandSide     `json:"left_hand_side,omitempty"`
	NodeType               string            `json:"node_type"`
	Operator               string            `json:"operator,omitempty"`
	RightHandSide          *RightHandSide    `json:"right_hand_side,omitempty"`
	Src                    string            `json:"src"`
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
	ID                     int               `json:"id"`
	Name                   string            `json:"name"`
	NodeType               string            `json:"node_type"`
	OverloadedDeclarations []interface{}     `json:"overloaded_declarations"`
	ReferencedDeclaration  int               `json:"referenced_declaration"`
	Src                    string            `json:"src"`
	TypeDescriptions       *TypeDescriptions `json:"type_descriptions"`
}

type RightHandSide struct {
	Arguments        []Argument        `json:"arguments"`
	Expression       *Expression       `json:"expression"`
	ID               int               `json:"id"`
	IsConstant       bool              `json:"is_constant"`
	IsLValue         bool              `json:"is_l_value"`
	IsPure           bool              `json:"is_pure"`
	Kind             string            `json:"kind"`
	LValueRequested  bool              `json:"l_value_requested"`
	Names            []interface{}     `json:"names"`
	NodeType         string            `json:"node_type"`
	Src              string            `json:"src"`
	TryCall          bool              `json:"try_call"`
	TypeDescriptions *TypeDescriptions `json:"type_descriptions"`
}

type Argument struct {
	ID                     int               `json:"id"`
	Name                   string            `json:"name"`
	NodeType               string            `json:"node_type"`
	OverloadedDeclarations []interface{}     `json:"overloaded_declarations"`
	ReferencedDeclaration  int               `json:"referenced_declaration"`
	Src                    string            `json:"src"`
	TypeDescriptions       *TypeDescriptions `json:"type_descriptions"`
}

type ParametersList struct {
	ID         int64       `json:"id"`
	NodeType   string      `json:"node_type"`
	Parameters []Parameter `json:"parameters"`
	Src        Src         `json:"src"`
}

type Parameter struct {
	Constant         bool              `json:"constant"`
	ID               int64             `json:"id"`
	Mutability       string            `json:"mutability"`
	Name             string            `json:"name"`
	NodeType         string            `json:"node_type"`
	Scope            int64             `json:"scope"`
	Src              Src               `json:"src"`
	StateVariable    bool              `json:"state_variable"`
	StorageLocation  string            `json:"storage_location"`
	TypeDescriptions *TypeDescriptions `json:"type_descriptions"`
	TypeName         *TypeName         `json:"type_name"`
	Visibility       string            `json:"visibility"`
}

type FunctionReturnParameters struct {
	ID         int           `json:"id"`
	NodeType   string        `json:"node_type"`
	Parameters []interface{} `json:"parameters"`
	Src        string        `json:"src"`
}
