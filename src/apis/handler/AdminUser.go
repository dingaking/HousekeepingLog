package handler

import (
	"apis/checker"
	"apis/model"
	"apis/query"
	"net/http"
)

func AdminUserU(w http.ResponseWriter, r *http.Request) {

	var req model.AdminUserUReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminUserU(req); err != nil {
		WriteError(w, err)
		return
	}

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	err = query.AdminUserU(session, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	response := model.AdminUserURes{
		Result:      "success",
		AccessToken: req.AccessToken}
	err = WriteSuccess(w, response)
	if err != nil {
		panic(err)
	}
}

func AdminUserR(w http.ResponseWriter, r *http.Request) {

	var req model.AdminUserRReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminUserR(req); err != nil {
		WriteError(w, err)
		return
	}

	if req.UserId == "admin" && req.Password == "admin" {
		rep := model.AdminUserRRes{
			Result:       "success",
			ErrorMessage: "You must change your initial Password.",
		}
		err = WriteSuccess(w, rep)
		if err != nil {
			panic(err)
		}
		return
	}

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	if req.Action == "2" {
		err = query.CheckPermission(session, req.AccessToken)
		if err != nil {
			WriteError(w, err)
			return
		}
	}

	rep := model.AdminUserRRes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.AdminUserR(session, req, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	if req.Action == "1" {
		rep2 := model.AdminUserRRes1{
			Result:       "success",
			ErrorMessage: "",
			AccessToken:  rep.AccessToken,
		}
		err = WriteSuccess(w, rep2)
		if err != nil {
			panic(err)
		}
	} else if req.Action == "2" {
		rep2 := model.AdminUserRRes2{
			Result:       "success",
			ErrorMessage: "",
			Data:         rep.Data,
		}
		err = WriteSuccess(w, rep2)
		if err != nil {
			panic(err)
		}
	}
}

func AdminUserC(w http.ResponseWriter, r *http.Request) {

	var req model.AdminUserCReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminUserC(req); err != nil {
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

	rep := model.AdminUserCRes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.AdminUserC(session, req, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func AdminUserL(w http.ResponseWriter, r *http.Request) {

	var req model.AdminUserLReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminUserL(req); err != nil {
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

	rep := model.AdminUserLRes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.AdminUserL(session, req, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func AdminUserD(w http.ResponseWriter, r *http.Request) {

	var req model.AdminUserDReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminUserD(req); err != nil {
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

	rep := model.AdminUserDRes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.AdminUserD(session, req, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func AdminUserS(w http.ResponseWriter, r *http.Request) {

	var req model.AdminUserSReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminUserS(req); err != nil {
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

	rep := model.AdminUserSRes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.AdminUserS(session, req, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}
