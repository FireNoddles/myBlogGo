package model

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Username string `json:"username"`
	Id       uint   `json:"id"`
	Token    string `json:"token"`
}
