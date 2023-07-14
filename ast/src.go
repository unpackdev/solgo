package ast

import "fmt"

type Src struct {
	Line   int   `json:"line"`
	Start  int   `json:"start"`
	End    int   `json:"end"`
	Length int   `json:"length"`
	Index  int64 `json:"index"`
}

func (s Src) String() string {
	return fmt.Sprintf("%d:%d:%d", s.Start, s.Length, s.Index)
}
