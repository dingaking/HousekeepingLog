package model

import "labix.org/v2/mgo/bson"

type Capital struct {
	CapitalNo   bson.ObjectId `bson:"_id,omitempty"`
	UserNo      string        `json:"userno"`
	TokenId     string        `json:"token_id"`
	CapitalName string        `json:"capital_name"`
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
