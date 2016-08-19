package main

import (
	"apis/model"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSum(t *testing.T) {

	user := model.User{UserId: "jungtek_whang@hanmail.net"}
	pbytes, _ := json.Marshal(user)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://localhost:8082/api/authC", "application/json", buff)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}

}
