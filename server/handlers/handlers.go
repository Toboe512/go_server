package handlers

import (
	"context"
	"encoding/json"
	"go_server/storage"
	"go_server/storage/dao"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Handler struct {
	Path    string
	HFunc   func(w http.ResponseWriter, req *http.Request)
	Storage *storage.Storage
}

func UserGetHand() *Handler {
	return &Handler{
		Path:  "/user",
		HFunc: user,
	}
}

func user(w http.ResponseWriter, req *http.Request) {

	var data *dao.Data
	var err error = nil
	var respBody []byte

	switch req.Method {
	case http.MethodGet:
		queryVars, err := url.ParseQuery(req.URL.RawQuery)
		if err != nil {
			log.Println(err)
		}
		data, err = storage.DB.GetById(context.TODO(), queryVars.Get("id"))
		if err != nil {
			log.Println(err)
		}
	case http.MethodPost:
		rq, err := io.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
		}

		log.Println(rq)

		err = json.Unmarshal(rq, &data)
		if err != nil {
			log.Println(err)
		}

		err = storage.DB.Save(context.TODO(), data)
		if err != nil {
			log.Println(err)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	respBody, err = json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(respBody)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusAccepted)

}
