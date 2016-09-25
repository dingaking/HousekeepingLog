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

func AdminGroupD(w http.ResponseWriter, r *http.Request) {

	var req model.AdminGroupDReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminGroupD(req); err != nil {
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

	err = query.GroupD(session, "hlog", "group", req.GroupNo)
	if err != nil {
		WriteError(w, err)
		return
	}

	rep := model.AdminGroupDRep{
		Result:       "success",
		ErrorMessage: "",
	}
	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func AdminGroupR(w http.ResponseWriter, r *http.Request) {

	var req model.AdminGroupRReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminGroupR(req); err != nil {
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

	err, rep_data := query.GroupR(session, "hlog", "group", req.GroupNo)
	if err != nil {
		WriteError(w, err)
		return
	}

	rep := model.AdminGroupRRep{
		Result:       "success",
		ErrorMessage: "",
		Data:         rep_data,
	}
	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func AdminGroupS(w http.ResponseWriter, r *http.Request) {

	var req model.AdminGroupSReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminGroupS(req); err != nil {
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

	err, rep_data := query.GroupS(session, "hlog", "group", req.GroupName)
	if err != nil {
		WriteError(w, err)
		return
	}

	rep := model.AdminGroupSRep{
		Result:       "success",
		ErrorMessage: "",
		Data:         rep_data,
	}
	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}

func AdminGroupU(w http.ResponseWriter, r *http.Request) {

	var req model.AdminGroupUReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}

	if err = checker.AdminGroupU(req); err != nil {
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

	err = query.GroupU(session, "hlog", "group", req)
	if err != nil {
		WriteError(w, err)
		return
	}

	rep := model.AdminGroupURep{
		Result:       "success",
		ErrorMessage: "",
	}
	err = WriteSuccess(w, rep)
	if err != nil {
		panic(err)
	}
}
