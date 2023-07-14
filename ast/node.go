package ast

type Node struct {
	ID       int64    `json:"id"`
	Literals []string `json:"literals,omitempty"`
	NodeType NodeType `json:"node_type"`
	Src      Src      `json:"src"`
	Nodes    []Node   `json:"nodes,omitempty"`

	AbsolutePath             string                    `json:"absolute_path,omitempty"`
	File                     string                    `json:"file,omitempty"`
	Scope                    int                       `json:"scope,omitempty"`
	SourceUnit               int                       `json:"source_unit,omitempty"`
	SymbolAliases            []interface{}             `json:"symbol_aliases,omitempty"`
	UnitAlias                string                    `json:"unit_alias,omitempty"`
	Abstract                 bool                      `json:"abstract,omitempty"`
	BaseContracts            []interface{}             `json:"base_contracts,omitempty"`
	ContractDependencies     []interface{}             `json:"contract_dependencies,omitempty"`
	ContractKind             string                    `json:"contract_kind,omitempty"`
	FullyImplemented         bool                      `json:"fully_implemented,omitempty"`
	LinearizedBaseContracts  []int64                   `json:"linearized_base_contracts,omitempty"`
	Name                     string                    `json:"name,omitempty"`
	Body                     *Body                     `json:"body,omitempty"`
	FunctionSelector         string                    `json:"function_selector,omitempty"`
	Implemented              bool                      `json:"implemented,omitempty"`
	Kind                     string                    `json:"kind,omitempty"`
	Modifiers                []interface{}             `json:"modifiers,omitempty"`
	Parameters               *ParametersList           `json:"parameters,omitempty"`
	ReturnParameters         *ParametersList           `json:"return_parameters,omitempty"`
	StateMutability          string                    `json:"state_mutability,omitempty"`
	Visibility               string                    `json:"visibility,omitempty"`
	Virtual                  bool                      `json:"virtual,omitempty"`
	Override                 bool                      `json:"override,omitempty"`
	LibraryName              *LibraryName              `json:"library_name,omitempty"`
	TypeName                 *TypeName                 `json:"type_name,omitempty"`
	StateVariable            bool                      `json:"state_variable,omitempty"`
	StorageLocation          string                    `json:"storage_location,omitempty"`
	TypeDescriptions         *TypeDescriptions         `json:"type_descriptions,omitempty"`
	Constant                 bool                      `json:"constant,omitempty"`
	Mutability               string                    `json:"mutability,omitempty"`
	Expression               *Expression               `json:"expression,omitempty"`
	FunctionReturnParameters *FunctionReturnParameters `json:"function_feturn_parameters,omitempty"`
}
