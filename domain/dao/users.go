package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"my_blog/domain/model"
)

type UserDao interface {
	CheckLogin(c *gin.Context, where string, paras ...interface{}) (user *model.Users, err error)
	ExistUserPreCheck(c *gin.Context, where string, paras ...interface{}) (user *model.Users, err error)
	AddUser(c *gin.Context, user *model.Users) (err error)
	DelUser(c *gin.Context, where string, paras ...interface{}) (err error)
	UpdateUser(c *gin.Context, user *model.Users) (err error)
	UpdateUserPwd(c *gin.Context, user *model.Users) (err error)
	GetUsersList(c *gin.Context, where string, paras ...interface{}) (users []*model.Users, err error)
}

func (d *dao) GetUsersList(c *gin.Context, where string, paras ...interface{}) (users []*model.Users, err error) {
	return d.GetUsersListInTx(c, where, paras...)
}

func (d *dao) GetUsersListInTx(c *gin.Context, where string, paras ...interface{}) (users []*model.Users, err error) {
	err = d.DbMyBlog.Where(where, paras...).Find(&users).Error
	if err != nil {
		log.Error("d.GetUsersListInTx err -> ", err)
		return
	}
	return
}

func (d *dao) CheckLogin(c *gin.Context, where string, paras ...interface{}) (user *model.Users, err error) {
	return d.GetUserInTx(c, where, paras...)
}

func (d *dao) DelUser(c *gin.Context, where string, paras ...interface{}) (err error) {
	return d.DelUserInTx(c, where, paras...)
}

func (d *dao) UpdateUser(c *gin.Context, user *model.Users) (err error) {
	data := map[string]interface{}{
		"username": user.Username,
		"role":     user.Role,
	}
	return d.UpdateUserInTx(c, user, data)
}

func (d *dao) UpdateUserPwd(c *gin.Context, user *model.Users) (err error) {
	data := map[string]interface{}{
		"password": user.Password,
	}
	return d.UpdateUserInTx(c, user, data)
}

func (d *dao) DelUserInTx(c *gin.Context, where string, paras ...interface{}) (err error) {
	err = d.DbMyBlog.Model(&model.Users{}).Where(where, paras...).Update("state", model.UserDelete).Error
	if err != nil {
		log.Error("d.DelUser false err -> ", err)
		return
	}
	return
}

func (d *dao) UpdateUserInTx(c *gin.Context, user *model.Users, data map[string]interface{}) (err error) {
	err = d.DbMyBlog.Model(&user).Update(data).Error
	if err != nil {
		log.Error("d.UpdateUserInTx false err -> ", err)
		return
	}
	return
}

func (d *dao) ExistUserPreCheck(c *gin.Context, where string, paras ...interface{}) (user *model.Users, err error) {
	return d.GetUserInTx(c, where, paras...)
}

func (d *dao) AddUser(c *gin.Context, users *model.Users) (err error) {
	return d.AddUserInTx(c, users)
}

func (d *dao) GetUserInTx(c *gin.Context, where string, paras ...interface{}) (user *model.Users, err error) {
	user = new(model.Users)
	err = d.DbMyBlog.Where(where, paras...).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return
		}
		log.Error("d.GetUser false err[%v]", err)
		return
	}
	return
}

func (d *dao) AddUserInTx(c *gin.Context, user *model.Users) (err error) {
	err = d.DbMyBlog.Create(user).Error
	if err != nil {
		log.Error("d.AddUserInTx err[%v]", err)
		return
	}
	return
}
