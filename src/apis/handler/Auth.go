package handler

import (
	"apis/checker"
	"apis/model"
	"apis/query"
	"errors"
	"net/http"

	"labix.org/v2/mgo"
)

func AuthC(w http.ResponseWriter, r *http.Request) {

	var authc model.AuthCReq
	err := Parse(w, r, &authc)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AuthC(authc); err != nil {
		WriteError(w, err)
		return
	}

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	err = query.AuthC(session, "hlog", "user", &authc)
	if err != nil {
		WriteError(w, err)
		return
	}

	response := model.AuthCRep{
		Result:      "success",
		AccessToken: authc.AccessToken}
	err = WriteSuccess(w, response)
	if err != nil {
		panic(err)
	}
}

func AuthR(w http.ResponseWriter, r *http.Request) {

	var req model.AuthRReq
	err := Parse(w, r, &req)
	if err != nil {
		return
	}

	if err := checker.AuthR(req); err != nil {
		WriteError(w, err)
		return
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	var result model.User
	if req.Action == "1" {
		result = query.GetUserInfoById(session, "hlog", "user", &req)
		if result.UserNo == "" {
			WriteError(w, errors.New("id/pw not match"))
			return
		}
	} else if req.Action == "2" {
		result = query.GetUserInfoByAccessToken(session, "hlog", "user", &req)
		if result.UserNo == "" {
			WriteError(w, errors.New("id/pw not match"))
			return
		}
	}

	response := model.AuthRRep{
		Result:      "success",
		AccessToken: result.AccessToken}
	err = WriteSuccess(w, response)
	if err != nil {
		panic(err)
	}
}

func AuthU(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return
}

func AuthD(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return
}
