package handler

import (
	"apis/model"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Parse는 request를 3번째 인자 타입으로 Parse하여 에러가 발생하면 client에 error를 직접 write함.
func Parse(w http.ResponseWriter, r *http.Request, v interface{}) error {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, &v); err != nil {
		return err
	}

	return nil
}

func BodyClose(r *http.Request) error {
	err := r.Body.Close()
	return err
}

func WriteError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response := model.CommonResponse{ErrorMessage: err.Error(), Result: "fail"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func WriteSuccess(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
	}
	return nil
}
