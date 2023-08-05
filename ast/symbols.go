package ast

type Symbol struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	AbsolutePath string `json:"absolute_path"`
}

func NewSymbol(id int64, name string, absolutePath string) Symbol {
	return Symbol{
		Id:           id,
		Name:         name,
		AbsolutePath: absolutePath,
	}
}

func (s Symbol) GetId() int64 {
	return s.Id
}

func (s Symbol) GetName() string {
	return s.Name
}

func (s Symbol) GetAbsolutePath() string {
	return s.AbsolutePath
}
