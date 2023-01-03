package test

import (
	"github.com/stretchr/testify/assert"
	"src/server/api"
	"testing"
)

func TestArithmeticEngine(t *testing.T) {

	engineName := "Simple_Arithmetic"
	AddEngine(t, engineName)
	defer DeleteEngine(t, engineName)

	//Stage-1 Add Expressions to Engine
	AddExpression(t, engineName, &api.Expression{Name: "First", Expr: "2+3"})
	AddExpression(t, engineName, &api.Expression{Name: "Second", Expr: "10%3"})

	//Stage-2 Engine is asked to evaluate all Expressions
	Evaluate(t, engineName)

	//Stage-3 Fetch Result From Engine
	resultMap := FetchResult(t, engineName)

	//Stage-4 Validate the Result
	assert.Equal(t, 2, len(resultMap))
	//Map Should have a key called First
	assert.Contains(t, resultMap, "First")
	//Map should have the correct value for the key called First
	assert.Equal(t, "5", resultMap["First"])
	//Map Should have a key called Second
	assert.Contains(t, resultMap, "Second")
	//Map should have the correct value for the key called Second
	assert.Equal(t, "1", resultMap["Second"])

	//Stage-5  Delete the Result
	DeleteExpression(t, engineName, "First")

	//Stage-6 Fetch Result From Engine
	resultMap2 := FetchResult(t, engineName)
	assert.Equal(t, 1, len(resultMap2))

	defer ClearEngine(t, engineName)
}

func TestComparisonEngine(t *testing.T) {

	engineName := "Simple_Arithemetic"
	AddEngine(t, engineName)
	defer DeleteEngine(t, engineName)

	//Stage-1 	//Add Expressions to Engine
	AddExpression(t, engineName, &api.Expression{Name: "First", Expr: "3>2"})
	AddExpression(t, engineName, &api.Expression{Name: "Second", Expr: "4<2"})
	AddExpression(t, engineName, &api.Expression{Name: "Third", Expr: "100<=(50+60)"})

	//State-2 	//Engine is asked to evaluate all Expressions
	Evaluate(t, engineName)

	//Stage-3	//Fetch Result From Engine
	resultMap := FetchResult(t, engineName)

	//Stage-4 	//Validate the Result
	assert.Equal(t, 3, len(resultMap))
	//Map Should have a key called First
	assert.Contains(t, resultMap, "First")
	//Map should have the correct value for the key called First
	assert.Equal(t, "true", resultMap["First"])
	//Map Should have a key called Second
	assert.Contains(t, resultMap, "Second")
	//Map should have the correct value for the key called Second
	assert.Equal(t, "false", resultMap["Second"])
	//Map Should have a key called Third
	assert.Contains(t, resultMap, "Third")
	//Map should have the correct value for the key called Second
	assert.Equal(t, "true", resultMap["Third"])

	//Stage-5 	//Delete the Result

	DeleteExpression(t, engineName, "First")

	//Stage-6 	//Fetch Result From Engine
	resultMap2 := FetchResult(t, engineName)
	assert.Equal(t, 2, len(resultMap2))

	defer ClearEngine(t, engineName)
}
