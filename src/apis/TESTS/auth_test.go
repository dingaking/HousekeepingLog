package main

import (
	"apis/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"gopkg.in/metakeule/fmtdate.v1"
)

// 회원가입 정상 테스트.
// 준비 : 존재하지 않는 아이디로 테스트한다. 아이디는 현재일시를 사용한다. id20160820_123398
// 결과 : result에 "success"가 있어야 한다.
func TestAuthC_0(t *testing.T) {

	date := fmtdate.Format("YYYYMMDD_hhmmss", time.Now())
	user := map[string]string{
		"userid":   "test_" + date,
		"password": "12345678",
	}

	pbytes, _ := json.Marshal(user)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://localhost:8082/api/authC", "application/json", buff)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		return
	}

	var res model.AuthCRep
	json.Unmarshal(respBody, &res)

	if res.Result != "success" {
		t.Error("Wrong result(not success)")
	}

}

// 회원가입 에러 테스트 : 요청 paramter에 아이디가 없으면 응답으로 fail을 받아야한다.
func TestAuthC_1(t *testing.T) {

	user := map[string]string{
		//"userid":   "test",
		"password": "12345678",
	}

	pbytes, _ := json.Marshal(user)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://localhost:8082/api/authC", "application/json", buff)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		return
	}

	var res model.AuthCRep
	json.Unmarshal(respBody, &res)

	if res.Result != "fail" {
		t.Error("Wrong result(not success)")
	}

}

// 회원가입
func TestAuthC_2(t *testing.T) {

	date := fmtdate.Format("YYYYMMDD_hhmmss", time.Now())
	user := map[string]string{
		"userid":   "test_" + date + "@gmail.com",
		"password": "12345678",
	}

	pbytes, _ := json.Marshal(user)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://localhost:8082/api/authC", "application/json", buff)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		return
	}

	var res model.AuthCRep
	json.Unmarshal(respBody, &res)

	fmt.Println(res)

	if res.Result != "success" {
		t.Error("Wrong result(not success)")
	}

}
