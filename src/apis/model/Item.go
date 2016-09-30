package model

import "labix.org/v2/mgo/bson"

// 구분항목 모델
type Item struct {
	ItemNo     bson.ObjectId `bson:"_id,omitempty"` // 항목키
	UserNo     string        `json:"userno"`        // 회원키
	CategoryNo string        `json:"categoryno"`    // 구분항목키
	CapitalNo  string        `json:"capitalno"`     // 재산항목키
	ItemName   string        `json:"item_name"`     // 항목 이름
	ItemValue  string        `json:"item_value"`    // 항목금액
}

type ItemCReq struct {
	AccessToken string `json:"access_token"`
}

type ItemCRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type ItemRReq struct {
	AccessToken string `json:"access_token"`
}

type ItemRRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type ItemUReq struct {
	AccessToken string `json:"access_token"`
}

type ItemURep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type ItemDReq struct {
	AccessToken string `json:"access_token"`
}

type ItemDRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type ItemLReq struct {
	AccessToken string `json:"access_token"`
}

type ItemLRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type ItemSReq struct {
	AccessToken string `json:"access_token"`
}

type ItemSRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}
