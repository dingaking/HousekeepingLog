package model

import "labix.org/v2/mgo/bson"

// 구분항목 모델
type Category struct {
	CategoryNo   bson.ObjectId `bson:"_id,omitempty"` // 구분항목키
	UserNo       string        `json:"userno"`        // 회원키
	IsInput      int           `json:"is_input"`      // type 1:수입, 0:지출
	CategoryName string        `json:"category_name"` // 구분항목 이름
}
