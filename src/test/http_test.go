package test

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestArtithmeticOperations(t *testing.T) {

	assert.Equal(t, "5", evaluate("2+3"), "")
	assert.Equal(t, "906", evaluate("2*453"), "")
	assert.Equal(t, "-32", evaluate("2-34"), "")
}

func evaluate(input string) string {
	baseUrl, _ := url.Parse("http://localhost:8000/execute")

	params := url.Values{}
	params.Add("expr", input)

	baseUrl.RawQuery = params.Encode()

	res, err := http.Get(baseUrl.String())
	if err != nil {
		log.Fatal(err)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return string(body)
}
