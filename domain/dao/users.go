package dao

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"my_blog/domain/model"
)

func (d *dao) CheckLogin(c *gin.Context, where string, paras ...interface{}) (user *model.Users, err error) {
	return d.GetUser(c, where, paras...)
}

func (d *dao) GetUser(c *gin.Context, where string, paras ...interface{}) (user *model.Users, err error) {
	user = new(model.Users)
	err = d.DbMyBlog.Where(where, paras...).First(user).Error
	if err != nil {
		log.Error("d.GetUser false err[%v]", err)
		return
	}
	return
}
