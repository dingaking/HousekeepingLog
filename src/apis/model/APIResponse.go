package model

type APIResponse struct {
	ErrorMessage string `json:"err_msg"`
	Result       string `json:"result"`
}

type APIResTokenId struct {
	ErrorMessage string `json:"err_msg"`
	Result       string `json:"result"`
	TokenId      string `json:"token_id"`
}
