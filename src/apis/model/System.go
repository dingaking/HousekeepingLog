package model

type SystemRReq struct {
	AccessToken string `json:"access_token"`
}

type SystemRRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}
