package model

import "labix.org/v2/mgo/bson"

type AdminItem struct {
	AdminNo   bson.ObjectId `bson:"_id,omitempty"` // 관리항목 no
	ItemKey   string        `bson:"item_key"`      // 관리항목 키
	ItemValue string        `bson:"item_value"`    // 관리항목 값
	ItemDesc  string        `bson:"item_desc"`     // 관리항목 설명
}

type AdminItemJ struct {
	AdminNo   string `json:"adminno"`    // 관리항목 no
	ItemKey   string `json:"item_key"`   // 관리항목 키
	ItemValue string `json:"item_value"` // 관리항목 값
	ItemDesc  string `json:"item_desc"`  // 관리항목 설명
}

type AdminSystemLReq struct {
	AccessToken string `json:"access_token"`
	Page        string `json:"page"`
}

type AdminSystemLRes struct {
	ErrorMessage string       `json:"err_msg"`
	Result       string       `json:"result"`
	Data         []AdminItemJ `json:"data"`
}

type AdminSystemRReq struct {
	AccessToken string `json:"access_token"`
	AdminNo     string `json:"adminno"` // 관리항목 no
}

type AdminSystemRRes struct {
	ErrorMessage string     `json:"err_msg"`
	Result       string     `json:"result"`
	Data         AdminItemJ `json:"data"`
}

type AdminSystemSReq struct {
	AccessToken string `json:"access_token"`
	Search      string `bson:"search"` // 관리항목 설명
}

type AdminSystemSRes struct {
	ErrorMessage string       `json:"err_msg"`
	Result       string       `json:"result"`
	Data         []AdminItemJ `json:"data"`
}

type AdminSystemUReq struct {
	AccessToken string `json:"access_token"`
	AdminNo     string `json:"adminno"`
	ItemKey     string `json:"item_key"`   // 관리항목 키
	ItemValue   string `json:"item_value"` // 관리항목 값
	ItemDesc    string `json:"item_desc"`  // 관리항목 설명
}

type AdminSystemURes struct {
	ErrorMessage string `json:"err_msg"`
	Result       string `json:"result"`
}
