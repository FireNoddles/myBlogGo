package model

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

type DelUserReq struct {
	Id uint `json:"id"`
}

type UpdateUserReq struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}

type UpdateUserPwdReq struct {
	Id     uint   `json:"id"`
	OldPwd string `json:"old_pwd"`
	NewPwd string `json:"new_pwd"`
}

type LoginResp struct {
	Username string `json:"username"`
	Id       uint   `json:"id"`
	Token    string `json:"token"`
}
