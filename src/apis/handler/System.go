package handler

import (
	"apis/model"
	"errors"
	"net/http"
)

func SystemR(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.SystemRReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}
