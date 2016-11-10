package handler

import (
	"apis/checker"
	"apis/model"
	"apis/query"
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

	err = query.CheckPermission(session, req.AccessToken)
	if err != nil {
		WriteError(w, err)
		return
	}

	rep := model.AdminSystemLRes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.AdminSystemL(session, &rep)
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

	err = query.CheckPermission(session, req.AccessToken)
	if err != nil {
		WriteError(w, err)
		return
	}

	rep := model.AdminSystemRRes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.AdminSystemR(session, req.SystemNo, &rep)
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

	err = query.CheckPermission(session, req.AccessToken)
	if err != nil {
		WriteError(w, err)
		return
	}

	rep := model.AdminSystemSRes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.AdminSystemS(session, req.Search, &rep)
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

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	err = query.CheckPermission(session, req.AccessToken)
	if err != nil {
		WriteError(w, err)
		return
	}

	rep := model.AdminSystemURes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.AdminSystemU(session, req, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}
