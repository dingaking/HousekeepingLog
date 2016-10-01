package handler

import (
	"apis/model"
	"apis/checker"
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

	if err = checker.SystemR(req); err != nil {
		WriteError(w, err)
		return
	}
}
