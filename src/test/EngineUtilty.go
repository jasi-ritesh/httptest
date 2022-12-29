package test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"src/server/api"
	"testing"
)

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



