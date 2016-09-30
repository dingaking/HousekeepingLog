package model

type SystemRReq struct {
	AccessToken string `json:"access_token"`
}

type SystemRRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}
