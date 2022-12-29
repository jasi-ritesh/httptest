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
	engineMap map[string]*server.Engine
	router    *mux.Router
}

func NewServer() *RestServer {
	s := RestServer{engineMap: make(map[string]*server.Engine, 0)}
	s.router = mux.NewRouter().StrictSlash(true)

	s.router.HandleFunc("/engine", s.AddEngine).Methods("POST")
	s.router.HandleFunc("/engine", s.DeleteEngine).Methods("DELETE")

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

func (s *RestServer) AddEngine(w http.ResponseWriter, r *http.Request) {
	engineName := r.URL.Query().Get("engine")
	fmt.Println("Add Engine", engineName)
	engine := server.NewEngine()
	s.engineMap[engineName] = engine
}

func (s *RestServer) DeleteEngine(w http.ResponseWriter, r *http.Request) {
	engineName := r.URL.Query().Get("engine")
	fmt.Println("Delete Engine", engineName)
	engine := server.NewEngine()
	if engine != nil {
		delete(s.engineMap, engineName)
	}
}

func (s *RestServer) GetEngine(name string) *server.Engine {
	if engine, present := s.engineMap[name]; present {
		return engine
	}
	return nil
}

func (s *RestServer) AddExpression(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var expr api.Expression
	engineName := r.URL.Query().Get("engine")
	json.Unmarshal(reqBody, &expr)
	fmt.Println("AddExpression", expr)
	engine := s.GetEngine(engineName)
	engine.AddExpression(expr.Name, expr.Expr)

}

func (s *RestServer) DeleteExpression(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	engineName := r.URL.Query().Get("engine")
	fmt.Println("Delete Expression", name)
	engine := s.GetEngine(engineName)
	engine.DeleteExpression(name)

}

func (s *RestServer) GetResult(w http.ResponseWriter, r *http.Request) {
	engineName := r.URL.Query().Get("engine")
	engine := s.GetEngine(engineName)
	result := engine.GetResult()
	json.NewEncoder(w).Encode(result)

}

func (s *RestServer) Clear(w http.ResponseWriter, r *http.Request) {
	engineName := r.URL.Query().Get("engine")
	engine := s.GetEngine(engineName)
	engine.Clear()
}

func (s *RestServer) Evaluate(w http.ResponseWriter, r *http.Request) {
	engineName := r.URL.Query().Get("engine")
	engine := s.GetEngine(engineName)
	engine.Evaluate()

}

func main() {
	rest := NewServer()
	rest.start()
}
