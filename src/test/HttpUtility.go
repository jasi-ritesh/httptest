package test

import (
	"encoding/json"
	"src/server/api"

	"github.com/go-resty/resty/v2"
)

func MakeJsonString(expr *api.Expression) string {
	exprString, _ := json.Marshal(expr) //expr will implement the marshal interface to produce a Json string
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
	//fmt.Println(url)
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Delete(url)

	return resp.StatusCode(), err
}
