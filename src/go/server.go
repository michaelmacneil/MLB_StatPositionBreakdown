package main

import(
  "github.com/gorilla/mux"
)

type server struct {
  configuration Configuration
  router *mux.Router
}

func (s *server) Init() {
    s.configuration = getConfigurationFromFile("../data/conf/conf.json")
    s.router = mux.NewRouter()
}
