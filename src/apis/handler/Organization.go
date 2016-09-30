package handler

import (
	"apis/checker"
	"apis/model"
	"errors"
	"net/http"
)

func OrganizationR(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.OrganizationRReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.OrganizationR(req); err != nil {
		WriteError(w, err)
		return
	}
}

func OrganizationL(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.OrganizationLReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.OrganizationL(req); err != nil {
		WriteError(w, err)
		return
	}
}

func OrganizationS(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.OrganizationSReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.OrganizationS(req); err != nil {
		WriteError(w, err)
		return
	}
}
