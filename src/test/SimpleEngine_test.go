package test

import (
	"encoding/json"
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"src/server/api"
	"testing"
)

func TestArithemeticEngine(t *testing.T) {

	//Stage-1
	//Add Expressions to Engine
	AddExpression(t, &api.Expression{Name: "First", Expr: "2+3"})
	AddExpression(t, &api.Expression{Name: "Second", Expr: "10%3"})

	//State-2
	//Engine is asked to evaluate all Expressions
	Evaluate(t)

	//Stage-3
	//Fetch Result From Engine
	resultMap := FetchResult(t)

	//Stage-4
	//Validate the Result
	assert.Equal(t, 2, len(resultMap))
	//Map Should have a key called First
	assert.Contains(t, resultMap, "First")
	//Map should have the correct value for the key called First
	assert.Equal(t, "5", resultMap["First"])
	//Map Should have a key called First
	assert.Contains(t, resultMap, "Second")
	//Map should have the correct value for the key called Second
	assert.Equal(t, "1", resultMap["Second"])

	//Stage-5
	//Delete the Result

	DeleteExpression(t, "First")

	//Stage-6
	//Fetch Result From Engine
	resultMap2 := FetchResult(t)
	assert.Equal(t, 1, len(resultMap2))

	defer ClearEngine(t)
}

func TestComparisonEngine(t *testing.T) {

	//Stage-1
	//Add Expressions to Engine
	AddExpression(t, &api.Expression{Name: "First", Expr: "3>2"})
	AddExpression(t, &api.Expression{Name: "Second", Expr: "4<2"})

	//State-2
	//Engine is asked to evaluate all Expressions
	Evaluate(t)

	//Stage-3
	//Fetch Result From Engine
	resultMap := FetchResult(t)

	//Stage-4
	//Validate the Result
	assert.Equal(t, 2, len(resultMap))
	//Map Should have a key called First
	assert.Contains(t, resultMap, "First")
	//Map should have the correct value for the key called First
	assert.Equal(t, "true", resultMap["First"])
	//Map Should have a key called First
	assert.Contains(t, resultMap, "Second")
	//Map should have the correct value for the key called Second
	assert.Equal(t, "false", resultMap["Second"])

	//Stage-5
	//Delete the Result

	DeleteExpression(t, "First")

	//Stage-6
	//Fetch Result From Engine
	resultMap2 := FetchResult(t)
	assert.Equal(t, 1, len(resultMap2))

	defer ClearEngine(t)
}

func AddExpression(t *testing.T, expr *api.Expression) {
	code, err := PostData("http://localhost:8000/engine/expr", MakeJsonString(expr))
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func DeleteExpression(t *testing.T, Name string) {
	url := fmt.Sprintf("http://localhost:8000/engine/expr?name=%s", Name)
	code, err := Delete(url)
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func Evaluate(t *testing.T) {
	code, err := PostData("http://localhost:8000/engine/evaluate", "")
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func ClearEngine(t *testing.T) {
	code, err := PostData("http://localhost:8000/engine/clear", "")
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func FetchResult(t *testing.T) map[string]string {
	body, code, err := Get("http://localhost:8000/engine/result")
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)

	var resultMap map[string]string
	jErr := json.Unmarshal(body, &resultMap)

	assert.Nil(t, jErr)

	return resultMap
}

func MakeJsonString(expr *api.Expression) string {
	exprString, _ := json.Marshal(expr)
	return string(exprString)
}

// PostData is a utility method to Post Rest Data
func PostData(url string, body string) (int, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)

	return resp.StatusCode(), err
}

// Get is a utility method to Get Rest Data
func Get(url string) ([]byte, int, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get(url)

	return resp.Body(), resp.StatusCode(), err
}

// Delete is a utility method to Delete Rest Data
func Delete(url string) (int, error) {
	fmt.Println(url)
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Delete(url)

	return resp.StatusCode(), err
}
