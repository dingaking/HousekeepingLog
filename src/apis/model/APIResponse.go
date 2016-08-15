package model

type APIResponse struct {
	ErrorMessage string `json:"err_msg"`
	Result       string `json:"result"`
}
