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
