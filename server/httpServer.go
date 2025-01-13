package server

import (
	"go_server/server/handlers"
	"net/http"
)

type HttpServer struct {
	host    string
	server  http.Server
	handler http.Handler
}

func New(host string) *HttpServer {
	return &HttpServer{
		host:   host,
		server: http.Server{Addr: host},
	}
}

func (s *HttpServer) Run() error {
	err := http.ListenAndServe(s.host, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *HttpServer) Handler(h *handlers.Handler) {
	http.HandleFunc(h.Path, h.HFunc)
}

func (s *HttpServer) HandlersAll(hds []*handlers.Handler) {
	for _, h := range hds {
		http.HandleFunc(h.Path, h.HFunc)
	}
}

//func (s *HttpServer) GET(path string, data interface{}) {
//
//}
//
//func (s *HttpServer) POST(f func(w http.ResponseWriter, r *http.Request)) {
//
//}
//
//func (s *HttpServer) PATH(f func(w http.ResponseWriter, r *http.Request)) {
//
//}
//
//func (s *HttpServer) DELETE(f func(w http.ResponseWriter, r *http.Request)) {
//
//}
