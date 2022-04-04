package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"my_blog/domain/model"
)

type Dao interface {
	close()
	CheckLogin(c *gin.Context, where string, paras ...interface{}) (user *model.Users, err error)
}

type dao struct {
	DbMyBlog *gorm.DB
}

func NewDmDao(blgDb *gorm.DB) Dao {
	return &dao{
		DbMyBlog: blgDb,
	}
}

func (d *dao) close() {
	d.DbMyBlog.Close()
}
