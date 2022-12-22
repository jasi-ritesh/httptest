package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maja42/goval"
	"log"
	"net/http"
)

func execute(writer http.ResponseWriter, request *http.Request) {

	data := request.URL.Query().Get("expr")
	eval := goval.NewEvaluator()
	fmt.Println("Input Recieved", data)

	val, err := eval.Evaluate(data, nil, nil)
	fmt.Println("Output", val)

	if err == nil {
		writer.Write([]byte(fmt.Sprint(val)))
	} else {
		writer.Write([]byte(err.Error()))
	}

}
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/execute", execute).Methods("GET")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}
}
