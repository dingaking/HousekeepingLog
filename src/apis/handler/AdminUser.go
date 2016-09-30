package handler

import (
	"apis/checker"
	"apis/model"
	"apis/query"
	"fmt"
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

	err = query.AdminUserU(session, "hlog", "user", &req)
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

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	err, msg := query.AdminUserR(session, "hlog", "user", &req)
	if err != nil {
		WriteError(w, err)
		return
	}
	if err == nil && msg != "" {
		response := model.AdminUserRRes{
			Result:  "success",
			Message: msg}
		err = WriteSuccess(w, response)
		if err != nil {
			panic(err)
		}
		return
	}

	response := model.AdminUserRRes{
		Result:      "success",
		AccessToken: req.AccessToken}
	err = WriteSuccess(w, response)
	if err != nil {
		panic(err)
	}
}

func AdminUserC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AdminUserC")
}

func AdminUserL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AdminUserL")
}

func AdminUserD(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AdminUserD")
}

func AdminUserS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AdminUserS")
}
