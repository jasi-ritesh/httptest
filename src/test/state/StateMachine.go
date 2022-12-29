package state

import (
	"github.com/stretchr/testify/assert"
	"src/server/api"
	"src/test"
	"testing"
)

type ExpressionData struct {
	InputMap          map[string]string
	ExpectedOutputMap map[string]string
	DeletedNames      []string
}

type StateMachineInterface interface {
	addEngine()
	deleteEngine()
	addExpression()
	evaluate()
	validate()
	delete()
	clear()
}

type StateMachineEngine struct {
	engineName string
	t          *testing.T ``
	data       *ExpressionData
}

func (s *StateMachineEngine) addEngine() {
	test.AddEngine(s.t, s.engineName)
}

func (s *StateMachineEngine) deleteEngine() {
	test.DeleteEngine(s.t, s.engineName)
}
func (s *StateMachineEngine) addExpression() {
	for name, expr := range s.data.InputMap {
		test.AddExpression(s.t, s.engineName, &api.Expression{Name: name, Expr: expr})
	}
}

func (s *StateMachineEngine) evaluate() {
	test.Evaluate(s.t, s.engineName)
}

func (s *StateMachineEngine) validate() {
	resultMap := test.FetchResult(s.t, s.engineName)

	//Stage-4
	//Validate the Result
	assert.Equal(s.t, len(s.data.ExpectedOutputMap), len(resultMap))

	for name, val := range s.data.ExpectedOutputMap {
		//Map Should have a key called First
		assert.Contains(s.t, resultMap, name)

		//Map should have the correct value for the key called First
		assert.Equal(s.t, val, resultMap[name])
	}
}

func (s *StateMachineEngine) delete() {
	for _, name := range s.data.DeletedNames {
		test.DeleteExpression(s.t, s.engineName, name)
	}

	resultMap := test.FetchResult(s.t, s.engineName)
	//Validate the Result
	assert.Equal(s.t, len(s.data.ExpectedOutputMap)-len(s.data.DeletedNames), len(resultMap))

}

func (s *StateMachineEngine) clear() {
	test.ClearEngine(s.t, s.engineName)
}

func (s *StateMachineEngine) RunStrategy() {

	s.addEngine()
	defer s.deleteEngine()
	
	s.addExpression()
	s.evaluate()
	s.validate()
	s.delete()
	s.clear()
}
