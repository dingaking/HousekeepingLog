package handler

import (
	"apis/checker"
	"apis/model"
	"apis/query"
	"errors"
	"net/http"
)

func DeviceC(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.DeviceCReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.DeviceC(req); err != nil {
		WriteError(w, err)
		return
	}

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	rep := model.DeviceCRes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.DeviceC(session, req, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func DeviceR(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.DeviceRReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.DeviceR(req); err != nil {
		WriteError(w, err)
		return
	}

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	rep := model.DeviceRRes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.DeviceR(session, req, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func DeviceU(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.DeviceUReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.DeviceU(req); err != nil {
		WriteError(w, err)
		return
	}

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	rep := model.DeviceURes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.DeviceU(session, req, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func DeviceD(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.DeviceDReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.DeviceD(req); err != nil {
		WriteError(w, err)
		return
	}

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	rep := model.DeviceDRes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.DeviceD(session, req, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func DeviceL(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.DeviceLReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.DeviceL(req); err != nil {
		WriteError(w, err)
		return
	}

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	rep := model.DeviceLRes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.DeviceL(session, req, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func DeviceS(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.DeviceSReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.DeviceS(req); err != nil {
		WriteError(w, err)
		return
	}

	session, err := query.GetConnect()
	if err != nil {
		WriteError(w, err)
		return
	}
	defer session.Close()

	rep := model.DeviceSRes{
		Result:       "success",
		ErrorMessage: "",
	}
	err = query.DeviceS(session, req, &rep)
	if err != nil {
		WriteError(w, err)
		return
	}

	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}
