package handler

import (
	"apis/checker"
	"apis/model"
	"apis/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func AuthC(w http.ResponseWriter, r *http.Request) {

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

	if chkResult := checker.AuthC(user); chkResult != "" {
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

	var qryResult model.User
	c.Find(bson.M{"userid": user.UserId}).One(&qryResult)
	if qryResult.UserNo != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		apiResponse := model.APIResponse{ErrorMessage: "duplicated id error."}
		if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
			panic(err)
		}
		return
	}

	var insertUser = model.User{UserId: user.UserId, Password: user.Password}
	insertUser.TokenList = append(insertUser.TokenList, model.TokenInfo{TokenId: util.SHA1()})

	err = c.Insert(&insertUser)
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

func AuthR(w http.ResponseWriter, r *http.Request) {

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

	if chkResult := checker.AuthR(user); chkResult != "" {
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
	c.Find(bson.M{"userid": user.UserId, "password": user.Password}).One(&qryResult)

	if qryResult.UserNo == "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		apiResTokenId := model.APIResTokenId{ErrorMessage: "id/pw not match", Result: "fail"}
		if err := json.NewEncoder(w).Encode(apiResTokenId); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	apiResTokenId := model.APIResTokenId{Result: "success"}

	if len(qryResult.TokenList) > 0 {
		apiResTokenId.TokenId = qryResult.TokenList[0].TokenId
	}

	if err := json.NewEncoder(w).Encode(apiResTokenId); err != nil {
		panic(err)
	}
}

func AuthU(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AuthU")
}

func AuthD(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AuthD")
}
