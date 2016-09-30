package handler

import (
	"apis/checker"
	"apis/model"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func CapitalC(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var capital model.Capital
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &capital); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	if chkResult := checker.CapitalC(capital); chkResult != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		response := model.CapitalC{ErrorMessage: chkResult, Result: "fail"}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
		return
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	var qryResult model.User
	c := session.DB("hlog").C("user")
	c.Find(bson.M{"_id": bson.ObjectIdHex(capital.UserNo)}).One(&qryResult)

}

func CapitalR(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.CapitalRReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func CapitalU(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.CapitalUReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func CapitalD(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.CapitalDReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func CapitalL(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.CapitalLReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func CapitalS(w http.ResponseWriter, r *http.Request) {

	WriteError(w, errors.New("Not Implemented"))
	return

	var req model.CapitalSReq
	err := Parse(w, r, &req)
	if err != nil {
		WriteError(w, err)
		return
	}
}
