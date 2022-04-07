package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"my_blog/admin/internal/model"
	dmDao "my_blog/domain/model"
	. "my_blog/utils"
)

func (s *Service) Login(c *gin.Context, req *model.LoginReq) (status int, message string, data *model.LoginResp) {

	user, err := s.dmDao.CheckLogin(c, "username = ? and state <> ?", req.Username, dmDao.UserDelete)
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

func (s *Service) AddUser(c *gin.Context, req *model.AddUserReq) (status int, message string) {
	existUser, err := s.dmDao.ExistUserPreCheck(c, "username = ? and state <> ?", req.Username, dmDao.UserDelete)
	if err != nil {
		log.Error("s.dmDao.AddUserPreCheck err [%v]", err)
		status = ERROR
		message = GetErrMsg(status)
		return
	}

	if existUser.Username != "" {
		log.Info("s.dmDao.AddUserPreCheck username repeat [%v]", existUser.Username)
		status = ERROR_USERNAME_USED
		message = GetErrMsg(ERROR_USERNAME_USED)
		return
	}

	addUser := &dmDao.Users{
		Username: req.Username,
		Password: ScryptPw(req.Password),
		State:    1,
		Role:     req.Role,
	}
	err = s.dmDao.AddUser(c, addUser)
	if err != nil {
		log.Error("s.dmDao.AddUser err [%v]", err)
		status = ERROR
		message = GetErrMsg(status)
		return
	}

	status = SUCCSE
	message = GetErrMsg(SUCCSE)
	return
}

func (s *Service) DelUser(c *gin.Context, req *model.DelUserReq) (status int, message string) {
	err := s.dmDao.DelUser(c, "id = ?", req.Id)
	if err != nil {
		log.Error("s.DelUser err -> ", err)
		status = ERROR
		message = GetErrMsg(status)
		return
	}
	status = SUCCSE
	message = GetErrMsg(status)
	return
}

func (s *Service) UpdateUser(c *gin.Context, req *model.UpdateUserReq) (status int, message string) {
	exUser, err := s.dmDao.ExistUserPreCheck(c, "id = ?", req.Id)
	if err != nil {
		log.Error("s.ExistUserPreCheck err -> ", err)
		status = ERROR
		message = GetErrMsg(status)
		return
	}

	if exUser.Username == "" {
		log.Info("s.UpdateUser user is not exist, userid -> ", req.Id)
		status = ERROR_USER_NOT_EXIST
		message = GetErrMsg(status)
		return
	}

	data := &dmDao.Users{
		Model: gorm.Model{
			ID: req.Id,
		},
		Username: req.Username,
		Role:     req.Role,
	}

	err = s.dmDao.UpdateUser(c, data)
	if err != nil {
		log.Error("s.DelUser err -> ", err)
		status = ERROR
		message = GetErrMsg(status)
		return
	}
	status = SUCCSE
	message = GetErrMsg(status)
	return
}

func (s *Service) UpdateUserPwd(c *gin.Context, req *model.UpdateUserPwdReq) (status int, message string) {
	exUser, err := s.dmDao.ExistUserPreCheck(c, "id = ?", req.Id)
	if err != nil {
		log.Error("s.ExistUserPreCheck err -> ", err)
		status = ERROR
		message = GetErrMsg(status)
		return
	}

	if exUser.Username == "" {
		log.Info("s.UpdateUser user is not exist, userid -> ", req.Id)
		status = ERROR_USER_NOT_EXIST
		message = GetErrMsg(status)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(exUser.Password), []byte(req.OldPwd))
	if err != nil {
		log.Info("s.UpdateUserPwd user old pwd is not right, userid -> ", req.Id)
		status = ERROR_PASSWORD_WRONG
		message = GetErrMsg(status)
		return
	}
	data := &dmDao.Users{
		Model: gorm.Model{
			ID: req.Id,
		},
		Password: ScryptPw(req.NewPwd),
	}

	err = s.dmDao.UpdateUserPwd(c, data)
	if err != nil {
		log.Error("s.UpdateUserPwd err -> ", err)
		status = ERROR
		message = GetErrMsg(status)
		return
	}
	status = SUCCSE
	message = GetErrMsg(status)
	return
}

func ScryptPw(password string) string {
	const cost = 10

	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(HashPw)
}
