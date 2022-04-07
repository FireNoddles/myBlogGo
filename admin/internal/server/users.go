package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"my_blog/admin/internal/model"
	. "my_blog/utils"
)

// login 登录接口
// @Summary login 登录接口
// @Description 根据用户名和密码登录
// @Tags users
// @Accept application/json
// @Produce application/json
// @Param object query model.LoginReq false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} model.LoginResp
// @Router /login [post]
func login(c *gin.Context) {
	req := &model.LoginReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("bind login req false")
		svc.Render(c, ERROR_REQ_PARAS, GetErrMsg(ERROR_REQ_PARAS), nil)

	}
	log.Info("start login, req [%v]", req)
	s, m, d := svc.Login(c, req)
	svc.Render(c, s, m, d)

}

// addUser 增加用户
// @Summary 增加用户信息
// @Description 根据用户名密码角色增加用户
// @Tags users
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query model.AddUserReq false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} nil
// @Router /posts2 [get]
func addUser(c *gin.Context) {
	req := &model.AddUserReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("bind addUser req false")
		svc.Render(c, ERROR_REQ_PARAS, GetErrMsg(ERROR_REQ_PARAS), nil)

	}
	log.Info("start adduser, req [%v]", req)
	s, m := svc.AddUser(c, req)
	svc.Render(c, s, m, nil)

}

// delUser 删除用户
// @Summary 删除用户
// @Description 根据用户id删除用户
// @Tags users
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query model.DelUserReq false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} nil
// @Router /posts2 [get]
func delUser(c *gin.Context) {
	req := &model.DelUserReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("bind delUser req false")
		svc.Render(c, ERROR_REQ_PARAS, GetErrMsg(ERROR_REQ_PARAS), nil)

	}
	log.Info("start adduser, req [%v]", req)
	s, m := svc.DelUser(c, req)
	svc.Render(c, s, m, nil)

}

// updateUser
// @Summary updateUser
// @Description 根据用户id更新用户
// @Tags users
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query model.UpdateUserReq false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} nil
// @Router /posts2 [get]
func updateUser(c *gin.Context) {
	req := &model.UpdateUserReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("bind UpdateUser req false")
		svc.Render(c, ERROR_REQ_PARAS, GetErrMsg(ERROR_REQ_PARAS), nil)

	}
	log.Info("start UpdateUser, req [%v]", req)
	s, m := svc.UpdateUser(c, req)
	svc.Render(c, s, m, nil)

}

func updatePwd(c *gin.Context) {
	req := &model.UpdateUserPwdReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("bind updatePwd req false")
		svc.Render(c, ERROR_REQ_PARAS, GetErrMsg(ERROR_REQ_PARAS), nil)

	}
	log.Info("start updatePwd, req [%v]", req)
	s, m := svc.UpdateUserPwd(c, req)
	svc.Render(c, s, m, nil)

}
