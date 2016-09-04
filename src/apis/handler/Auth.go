package handler

import (
	"apis/checker"
	"apis/model"
	"apis/query"
	"encoding/json"
	"fmt"
	"net/http"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
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

	var user model.User
	err := Parse(w, r, user)
	if err != nil {
		return
	}

	if err := checker.AuthR(user); err != nil {
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

	var qryResult model.User
	c := session.DB("hlog").C("user")
	c.Find(bson.M{"userid": user.UserId, "password": user.Password}).One(&qryResult)

	if qryResult.UserNo == "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		response := model.AuthR{ErrorMessage: "id/pw not match", Result: "fail"}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			WriteError(w, err)
			return
		}
		return
	}

	/*
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		response := model.AuthR{Result: "success"}

			if len(qryResult.TokenList) > 0 {
				response.TokenId = qryResult.TokenList[0].TokenId
			}
			response.UserNo = qryResult.UserNo.Hex()

			if err := json.NewEncoder(w).Encode(response); err != nil {
				panic(err)
			}
	*/
}

func AuthU(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AuthU")
}

func AuthD(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AuthD")
}
