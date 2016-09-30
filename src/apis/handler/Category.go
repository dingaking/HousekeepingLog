package handler

import (
	"apis/model"
	"errors"
	"net/http"
)

func CategoryC(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.CategoryCReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func CategoryR(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.CategoryRReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func CategoryU(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.CategoryUReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func CategoryD(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.CategoryDReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func CategoryL(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.CategoryLReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func CategoryS(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.CategorySReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}
