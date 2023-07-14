package ast

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
	ID         int64         `json:"id"`
	NodeType   string        `json:"node_type"`
	Parameters []interface{} `json:"parameters"`
	Src        Src           `json:"src"`
}

type TypeName struct {
	ID               int64             `json:"id"`
	Name             string            `json:"name"`
	NodeType         string            `json:"node_type"`
	Src              Src               `json:"src"`
	TypeDescriptions *TypeDescriptions `json:"type_descriptions"`
}

type TypeDescriptions struct {
	TypeIdentifier string `json:"type_identifier"`
	TypeString     string `json:"type_string"`
}
