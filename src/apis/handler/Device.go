package handler

import (
	"apis/model"
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
}
