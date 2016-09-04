package model

type CommonResponse struct {
	ErrorMessage string `json:"err_msg"`
	Result       string `json:"result"`
}
