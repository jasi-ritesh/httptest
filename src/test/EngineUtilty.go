package test

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"src/server/api"
	"testing"
)

func AddEngine(t *testing.T, engineName string) {
	log.Println("Starting Engine ", engineName)
	url := fmt.Sprintf("http://localhost:8000/engine?engine=%s", engineName)
	code, err := PostData(url, "")
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}
func DeleteEngine(t *testing.T, engineName string) {
	log.Println("Deleting Engine ", engineName)
	url := fmt.Sprintf("http://localhost:8000/engine?engine=%s", engineName)
	code, err := Delete(url)
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func AddExpression(t *testing.T, engineName string, expr *api.Expression) {
	log.Printf("Add Expression %s to Engine %s\n ", expr, engineName)
	url := fmt.Sprintf("http://localhost:8000/engine/expr?engine=%s", engineName)
	code, err := PostData(url, MakeJsonString(expr))
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func DeleteExpression(t *testing.T, engineName string, Name string) {
	log.Printf("Delete Expression %s to Engine %s\n ", Name, engineName)
	url := fmt.Sprintf("http://localhost:8000/engine/expr?name=%s&engine=%s", Name, engineName)
	code, err := Delete(url)
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func Evaluate(t *testing.T, engineName string) {
	log.Println("Evaluate Engine ", engineName)
	url := fmt.Sprintf("http://localhost:8000/engine/evaluate?engine=%s", engineName)
	code, err := PostData(url, "")
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func ClearEngine(t *testing.T, engineName string) {
	log.Println("Clear Engine ", engineName)
	url := fmt.Sprintf("http://localhost:8000/engine/clear?engine=%s", engineName)
	code, err := PostData(url, "")
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func FetchResult(t *testing.T, engineName string) map[string]string {
	log.Println("Fetch Result ", engineName)
	url := fmt.Sprintf("http://localhost:8000/engine/result?engine=%s", engineName)
	body, code, err := Get(url)
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)

	var resultMap map[string]string
	jErr := json.Unmarshal(body, &resultMap)

	assert.Nil(t, jErr)

	return resultMap
}
