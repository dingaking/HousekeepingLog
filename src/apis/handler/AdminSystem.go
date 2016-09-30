package handler

import (
	"apis/checker"
	"apis/model"
	"apis/query"
	"errors"
	"net/http"
)

func AdminSystemL(w http.ResponseWriter, r *http.Request) {

	var req model.AdminSystemLReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminSystemL(req); err != nil {
		WriteError(w, err)
		return
	}

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	err = query.CheckPermission(session, "hlog", "user", req.AccessToken)
	if err != nil {
		WriteError(w, err)
		return
	}

	rep := model.AdminSystemLRep{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.SystemL(session, "hlog", "admin", &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func AdminSystemR(w http.ResponseWriter, r *http.Request) {

	var req model.AdminSystemRReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminSystemR(req); err != nil {
		WriteError(w, err)
		return
	}

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	err = query.CheckPermission(session, "hlog", "user", req.AccessToken)
	if err != nil {
		WriteError(w, err)
		return
	}

	rep := model.AdminSystemRRep{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.SystemR(session, "hlog", "admin", req.AdminNo, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func AdminSystemS(w http.ResponseWriter, r *http.Request) {

	var req model.AdminSystemSReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminSystemS(req); err != nil {
		WriteError(w, err)
		return
	}

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	err = query.CheckPermission(session, "hlog", "user", req.AccessToken)
	if err != nil {
		WriteError(w, err)
		return
	}

	rep := model.AdminSystemSRep{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.SystemS(session, "hlog", "admin", req.Search, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func AdminSystemU(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.AdminSystemUReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminSystemU(req); err != nil {
		WriteError(w, err)
		return
	}
}
