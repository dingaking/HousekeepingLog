package handler

import (
	"apis/checker"
	"apis/model"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func CapitalC(w http.ResponseWriter, r *http.Request) {

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

	/*
		found := false
			for _, token := range qryResult.TokenList {
				if token.TokenId == capital.TokenId {
					found = true
					break
				}
			}

			if found == false {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)
				response := model.CapitalC{ErrorMessage: "user & token_id not match.", Result: "fail"}
				if err := json.NewEncoder(w).Encode(response); err != nil {
					panic(err)
				}
				return
			}

			d := session.DB("hlog").C("capital")
			var insertCapital = model.CapitalI{UserNo: capital.UserNo, CapitalName: capital.CapitalName}
			err = d.Insert(&insertCapital)
			if err != nil {
				panic(err)
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			response := model.CapitalC{Result: "success"}

			if err := json.NewEncoder(w).Encode(response); err != nil {
				panic(err)
			}
	*/
}

func CapitalR(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CapitalR")
}

func CapitalU(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CapitalU")
}

func CapitalD(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CapitalD")
}

func CapitalL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CapitalL")
}

func CapitalS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CapitalS")
}
