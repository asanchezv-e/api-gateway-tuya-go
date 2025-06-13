package models

type TuyaTokenResponse struct {
	Result struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpireTime   int64  `json:"expire_time"`
	} `json:"result"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
}