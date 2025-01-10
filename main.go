package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	localhost = ":80"
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

	body, err := json.Marshal(User{ID: "1234"})
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(body)
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
