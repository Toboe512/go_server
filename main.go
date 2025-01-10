package main

import (
	"log"
	"net/http"
)

const (
	localhost = "127.0.0.1:80"
)

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_, err := w.Write([]byte("page"))
	if err != nil {
		log.Println(err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("header", "value")
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	http.HandleFunc("/hello", hello)

	http.Handle("/page", &Handler{})
	err := http.ListenAndServe(localhost, nil)

	println(err)

}
