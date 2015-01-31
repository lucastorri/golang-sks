package server

import (
    "io/ioutil"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
    "github.com/lucastorri/sks/store"
)

type Server struct {
    server *http.Server
    Store store.Store
}

func New(port int, store store.Store) (s *Server) {

    r := mux.NewRouter()
    server := &http.Server { Addr: fmt.Sprintf(":%d", port), Handler: r }

    s = &Server { server, store }

    r.HandleFunc("/{key}", GetHandler(s)).Methods("GET")
    r.HandleFunc("/{key}", AddHandler(s)).Methods("POST")
    server.ListenAndServe()

    return
}

func GetHandler(s *Server) func (http.ResponseWriter, *http.Request) {
    return func(res http.ResponseWriter, req *http.Request) {
        key := mux.Vars(req)["key"]
        if content, ok := s.Store.Get(key); ok {
            res.Write([]byte(content))
        } else {
            http.NotFound(res, req)
        }
    }
}

func AddHandler(s *Server) func (http.ResponseWriter, *http.Request) {
    return func(res http.ResponseWriter, req *http.Request) {
        key := mux.Vars(req)["key"]
        if body, err := ioutil.ReadAll(req.Body); err != nil {
            http.Error(res, err.Error(), 500)
        } else if err := s.Store.Add(key, string(body)); err != nil {
            http.Error(res, err.Error(), 500)
        }
    }
}
