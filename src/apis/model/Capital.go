package model

import "labix.org/v2/mgo/bson"

// 재산항목 모델
type Capital struct {
	CapitalNo   bson.ObjectId `bson:"_id,omitempty"` // 재산항목키
	UserNo      string        `json:"userno"`        // 회원키
	CapitalName string        `json:"capital_name"`  // 재산항목 이름
}

type CapitalI struct {
	CapitalNo   bson.ObjectId `bson:"_id,omitempty"`
	UserNo      string        `json:"userno"`
	CapitalName string        `json:"capital_name"`
}

type CapitalC struct {
	ErrorMessage string `json:"err_msg"`
	Result       string `json:"result"`
}

type CapitalCReq struct {
	AccessToken string `json:"access_token"`
}

type CapitalCRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type CapitalRReq struct {
	AccessToken string `json:"access_token"`
}

type CapitalRRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type CapitalUReq struct {
	AccessToken string `json:"access_token"`
}

type CapitalURes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type CapitalDReq struct {
	AccessToken string `json:"access_token"`
}

type CapitalDRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type CapitalLReq struct {
	AccessToken string `json:"access_token"`
}

type CapitalLRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type CapitalSReq struct {
	AccessToken string `json:"access_token"`
}

type CapitalSRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}
