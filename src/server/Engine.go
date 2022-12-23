package server

import (
	"fmt"
	"github.com/maja42/goval"
)

type Engine struct {
	exprMap   map[string]string
	resultMap map[string]string
}

func NewEngine() *Engine {
	e := Engine{}
	e.exprMap = make(map[string]string, 0)
	e.resultMap = make(map[string]string, 0)
	return &e
}
func (e *Engine) AddExpression(Name string, expr string) {
	e.exprMap[Name] = expr
}

func (e *Engine) Evaluate() {
	eval := goval.NewEvaluator()
	for name, expr := range e.exprMap {
		val, err := eval.Evaluate(expr, nil, nil)
		if err == nil {
			e.resultMap[name] = fmt.Sprint(val)
		} else {
			e.resultMap[name] = err.Error()

		}
		fmt.Println(name, val, expr)
	}
}

func (e *Engine) GetResult() map[string]string {
	return e.resultMap
}

func (e *Engine) DeleteExpression(Name string) {

	if _, ok := e.exprMap[Name]; ok {
		delete(e.exprMap, Name)
	}
	if _, ok := e.resultMap[Name]; ok {
		delete(e.resultMap, Name)
	}
}
func (e *Engine) Clear() {
	e.exprMap = make(map[string]string, 0)
	e.resultMap = make(map[string]string, 0)
}
