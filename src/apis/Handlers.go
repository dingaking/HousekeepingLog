package main

import (
	"apis/model"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {

	var user model.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	if chkResult := CheckAuthC(user); chkResult != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		apiResponse := model.APIResponse{ErrorMessage: chkResult, Result: "fail"}
		if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
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

	c := session.DB("hlog").C("user")

	err = c.Insert(&model.User{UserId: user.UserId, Password: user.Password})
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	apiResponse := model.APIResponse{Result: "success"}

	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		panic(err)
	}
}

func UserInfo(w http.ResponseWriter, r *http.Request) {

	var user model.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	if chkResult := CheckAuthR(user); chkResult != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		apiResTokenId := model.APIResTokenId{ErrorMessage: chkResult, Result: "fail"}
		if err := json.NewEncoder(w).Encode(apiResTokenId); err != nil {
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
	c.Find(bson.M{"userid": user.UserId, "password": user.Password}).Select(bson.M{"password": user.Password}).One(&qryResult)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	apiResTokenId := model.APIResTokenId{Result: "success"}

	if err := json.NewEncoder(w).Encode(apiResTokenId); err != nil {
		panic(err)
	}
}
