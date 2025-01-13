package handlers

import (
	"context"
	"encoding/json"
	"go_server/dto"
	"go_server/mappers"
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
		Path:  "/data",
		HFunc: user,
	}
}

func user(w http.ResponseWriter, req *http.Request) {

	var reqData *dao.Data
	var resData *dto.Data
	var err error = nil
	var respBody []byte

	switch req.Method {
	case http.MethodGet:
		queryVars, err := url.ParseQuery(req.URL.RawQuery)
		if err != nil {
			log.Println(err)
		}

		reqData, err = storage.DB.GetById(context.TODO(), queryVars.Get("id"))
		if err != nil {
			log.Println(err)
		}
		resData = mapers.DataToDto(reqData)
	case http.MethodPost:
		rq, err := io.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
		}

		log.Println(rq)

		err = json.Unmarshal(rq, &reqData)
		if err != nil {
			log.Println(err)
		}

		err = storage.DB.Save(context.TODO(), reqData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)

			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				log.Println(err)
			}
			return
		}
		resData = mapers.DataToDao(reqData)

	case http.MethodDelete:
		queryVars, err := url.ParseQuery(req.URL.RawQuery)
		if err != nil {
			log.Println(err)
		}

		reqData, err = storage.DB.GetById(context.TODO(), queryVars.Get("id"))
		if err != nil {
			log.Println(err)
		}

		err = storage.DB.Delete(context.TODO(), queryVars.Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)

			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				log.Println(err)
			}
			return
		}
		resData = mapers.DataToDto(reqData)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	respBody, err = json.Marshal(resData)
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(respBody)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusAccepted)

}
