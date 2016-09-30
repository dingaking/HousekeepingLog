package model

import "labix.org/v2/mgo/bson"

// 구분항목 모델
type Category struct {
	CategoryNo   bson.ObjectId `bson:"_id,omitempty"` // 구분항목키
	UserNo       string        `json:"userno"`        // 회원키
	IsInput      int           `json:"is_input"`      // type 1:수입, 0:지출
	CategoryName string        `json:"category_name"` // 구분항목 이름
}

type CategoryCReq struct {
	AccessToken string `json:"access_token"`
}

type CategoryCRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type CategoryRReq struct {
	AccessToken string `json:"access_token"`
}

type CategoryRRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type CategoryUReq struct {
	AccessToken string `json:"access_token"`
}

type CategoryURep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type CategoryDReq struct {
	AccessToken string `json:"access_token"`
}

type CategoryDRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type CategoryLReq struct {
	AccessToken string `json:"access_token"`
}

type CategoryLRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type CategorySReq struct {
	AccessToken string `json:"access_token"`
}

type CategorySRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}
