package handler

import (
	"apis/model"
	"errors"
	"net/http"
)

func ItemC(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.ItemCReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func ItemR(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.ItemRReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func ItemU(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.ItemUReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func ItemD(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.ItemDReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func ItemL(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.ItemLReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func ItemS(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.ItemSReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}
