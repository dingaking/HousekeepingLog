package handler

import (
	"apis/checker"
	"apis/model"
	"apis/query"
	"errors"
	"net/http"
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

	err = query.AuthC(session, &authc)
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

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	var result model.User
	if req.Action == "1" {
		result = query.GetUserInfoById(session, &req)
		if result.UserNo == "" {
			WriteError(w, errors.New("id/pw not match"))
			return
		}
	} else if req.Action == "2" {
		result = query.GetUserInfoByAccessToken(session, &req)
		if result.UserNo == "" {
			WriteError(w, errors.New("id/pw not match"))
			return
		}
	} else if req.Action == "3" {
		err = query.IsExistUserByUserNo(session, &req)
		if err != nil {
			WriteError(w, err)
			return
		}
		err = query.IsExistProfileImageByUserNo(session, &req)
		if err != nil {
			WriteError(w, err)
			return
		}

		gridFile := query.GetUserProfileImage(session, &req)
		if gridFile.Name() == "" {
			WriteError(w, errors.New("no image error."))
			return
		}
		var b []byte
		gridFile.Read(b)
		if b == nil {
			WriteError(w, errors.New("no profile raw image error."))
			return
		}

		w.Write(b)
		return
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

	var req model.AuthUReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AuthU(req); err != nil {
		WriteError(w, err)
		return
	}

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	rep := model.AuthURes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.AuthU(session, req, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func AuthD(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.AuthDReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AuthD(req); err != nil {
		WriteError(w, err)
		return
	}

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	rep := model.AuthDRes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.AuthD(session, req, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}
