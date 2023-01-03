package api

import "fmt"

type Expression struct {
	Name string `json: "name"`
	Expr string `json: "expr"`
}

func (e *Expression) String() string {
	return fmt.Sprintf("%s=%s", e.Name, e.Expr)
}
