package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"my_blog/admin/configs"
)

type Dao struct {
	Db *DB
}

func NewDao(db *DB) *Dao {
	return &Dao{
		Db: db,
	}
}

type DB struct {
	MyBlog *gorm.DB
}

func NewMysql(conf *configs.DbConfig) (*DB, func(), error) {
	db, err := gorm.Open("mysql", conf.MyBlog.DSN)
	if err != nil {
		log.Error("NewMysql err [%v]", err)
		return nil, nil, err
	}
	db.SingularTable(true) //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	db.LogMode(true)       //打印sql语句
	//开启连接池
	db.DB().SetMaxIdleConns(100)   //最大空闲连接
	db.DB().SetMaxOpenConns(10000) //最大连接数
	db.DB().SetConnMaxLifetime(30)
	cf := func() { db.Close() }
	return &DB{
		MyBlog: db,
	}, cf, err
}
