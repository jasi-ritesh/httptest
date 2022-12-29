package test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"src/server/api"
	"testing"
)

func AddEngine(t *testing.T, engineName string) {
	url := fmt.Sprintf("http://localhost:8000/engine?engine=%s", engineName)
	code, err := PostData(url, "")
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}
func DeleteEngine(t *testing.T, engineName string) {
	url := fmt.Sprintf("http://localhost:8000/engine?engine=%s", engineName)
	code, err := Delete(url)
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func AddExpression(t *testing.T, engineName string, expr *api.Expression) {
	url := fmt.Sprintf("http://localhost:8000/engine/expr?engine=%s", engineName)
	code, err := PostData(url, MakeJsonString(expr))
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func DeleteExpression(t *testing.T, engineName string, Name string) {
	url := fmt.Sprintf("http://localhost:8000/engine/expr?name=%s&engine=%s", Name, engineName)
	code, err := Delete(url)
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func Evaluate(t *testing.T, engineName string) {
	url := fmt.Sprintf("http://localhost:8000/engine/evaluate?engine=%s", engineName)
	code, err := PostData(url, "")
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func ClearEngine(t *testing.T, engineName string) {
	url := fmt.Sprintf("http://localhost:8000/engine/clear?engine=%s", engineName)
	code, err := PostData(url, "")
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func FetchResult(t *testing.T, engineName string) map[string]string {
	url := fmt.Sprintf("http://localhost:8000/engine/result?engine=%s", engineName)
	body, code, err := Get(url)
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)

	var resultMap map[string]string
	jErr := json.Unmarshal(body, &resultMap)

	assert.Nil(t, jErr)

	return resultMap
}
