package handlers

import (
	"context"
	"encoding/json"
	"go_server/storage"
	"log"
	"net/http"
)

type Handler struct {
	Path    string
	HFunc   func(w http.ResponseWriter, req *http.Request)
	Storage *storage.Storage
}

func UserGetHand() *Handler {
	return &Handler{
		Path:  "/user",
		HFunc: userGet,
	}
}

func userGet(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	data, err := storage.DB.GetById(context.TODO(), "1")
	if err != nil {
		log.Println(err)
	}

	body, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(body)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusAccepted)
}
