package dao

import (
	"github.com/jinzhu/gorm"
)

type Dao interface {
	close()
	UserDao
	CategoryDao
	ArticleDao
	UploadDao
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
