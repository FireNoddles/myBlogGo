package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"my_blog/admin/internal/model"
	. "my_blog/utils"
)

func (s *Service) Login(c *gin.Context, req *model.LoginReq) (status int, message string, data *model.LoginResp) {

	user, err := s.dmDao.CheckLogin(c, "username = ?", req.Username)
	if err != nil {
		log.Error("s.dmDao.CheckLogin err, req[%v], err[%v]", req, err)
		status = 500
		message = GetErrMsg(status)
		return
	}
	if user.Username == "" || user.Role != 1 {
		status = ERROR_PASSWORD_WRONG
		message = GetErrMsg(status)
	} else {
		PasswordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if PasswordErr != nil {
			status = ERROR_PASSWORD_WRONG
			message = GetErrMsg(status)
		} else {
			token, err := setToken(c, req)
			if err != nil {
				log.Error("login setToken err [%v]", err)
				status = 500
				message = GetErrMsg(status)
				return
			}
			data = new(model.LoginResp)
			data.Id = user.ID
			data.Username = user.Username
			data.Token = token

			status = SUCCSE
			message = GetErrMsg(status)

			return
		}
	}
	return

}

func setToken(c *gin.Context, user *model.LoginReq) (string, error) {
	j := NewJWT()

	token, err := j.SetToken(user.Username, user.Password)

	if err != nil {
		log.Error("j.SetToken err [%v]", err)
		return "", err
	}

	return token, err
}
