package handler

import (
	"errors"
	"net/http"
)

func SystemR(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return
}
