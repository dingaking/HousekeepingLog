package handler

import (
	"apis/checker"
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

	if err = checker.ItemC(req); err != nil {
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

	if err = checker.ItemR(req); err != nil {
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

	if err = checker.ItemU(req); err != nil {
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

	if err = checker.ItemD(req); err != nil {
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

	if err = checker.ItemL(req); err != nil {
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

	if err = checker.ItemS(req); err != nil {
		WriteError(w, err)
		return
	}
}
