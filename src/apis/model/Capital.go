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
