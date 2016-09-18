package handler

import (
	"apis/checker"
	"apis/model"
	"apis/query"
	"net/http"
)

func AdminGroupC(w http.ResponseWriter, r *http.Request) {

	var req model.AdminGroupCReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminGroupC(req); err != nil {
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

	err = query.InsertGroup(session, "hlog", "group", &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	response := model.AdminGroupCRep{
		Result:       "success",
		ErrorMessage: "",
	}
	err = WriteSuccess(w, response)
	if err != nil {
		panic(err)
	}
}

func AdminGroupL(w http.ResponseWriter, r *http.Request) {

	var req model.AdminGroupLReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminGroupL(req); err != nil {
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

	var rep model.AdminGroupLRep
	err = query.GroupL(session, "hlog", "group", &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}
