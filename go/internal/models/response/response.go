package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Success bool   `json:"success,omitempty"`
	Msg     string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
}

func (r Response) WriteResp(w http.ResponseWriter, code int) error {
	res, err := json.Marshal(r)
	if err != nil {
		log.Println("Internal error while encoding response: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(res)
	if err != nil {
		log.Println("Internal error while writing response: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	return nil
}
