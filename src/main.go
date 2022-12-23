package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"src/server"
	"src/server/api"
)

type RestServer struct {
	engine *server.Engine
	router *mux.Router
}

func NewServer() *RestServer {
	s := RestServer{engine: server.NewEngine()}
	s.router = mux.NewRouter().StrictSlash(true)

	s.router.HandleFunc("/engine/expr", s.AddExpression).Methods("POST")
	s.router.HandleFunc("/engine/evaluate", s.Evaluate).Methods("POST")
	s.router.HandleFunc("/engine/clear", s.Clear).Methods("POST")

	s.router.HandleFunc("/engine/result", s.GetResult).Methods("GET")

	s.router.HandleFunc("/engine/expr", s.DeleteExpression).Methods("DELETE")

	return &s
}
func (s *RestServer) start() {
	err := http.ListenAndServe(":8000", s.router)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}
}

func (s *RestServer) AddExpression(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var expr api.Expression
	json.Unmarshal(reqBody, &expr)
	fmt.Println("AddExpression", expr)
	s.engine.AddExpression(expr.Name, expr.Expr)

}

func (s *RestServer) DeleteExpression(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Println("Delete Expression", name)
	s.engine.DeleteExpression(name)

}

func (s *RestServer) GetResult(w http.ResponseWriter, r *http.Request) {
	result := s.engine.GetResult()
	json.NewEncoder(w).Encode(result)

}

func (s *RestServer) Clear(w http.ResponseWriter, r *http.Request) {
	s.engine.Clear()
}

func (s *RestServer) Evaluate(w http.ResponseWriter, r *http.Request) {

	s.engine.Evaluate()

}

func main() {
	rest := NewServer()
	rest.start()
}
