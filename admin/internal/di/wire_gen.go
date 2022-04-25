package di

import (
	log "github.com/sirupsen/logrus"
	"my_blog/admin/configs"
	"my_blog/admin/internal/dao"
	"my_blog/admin/internal/server"
	"my_blog/admin/internal/service"
)

func Init() (func(), error) {
	dbConf := configs.NewDbConfig()
	sConf := configs.NewServerConfig()
	upConf := configs.NewUploadConfig()
	ac := configs.NewConfigs(dbConf, sConf, upConf)
	db, cf, err := dao.NewMysql(dbConf)

	if err != nil {
		log.Error("init dao.NewMysql err [%v]", err)
		return nil, err
	}
	daoDao := dao.NewDao(db)
	s := service.NewService(daoDao, ac)
	_, err = server.NewServer(ac, s)
	if err != nil {
		cf()
		return nil, err
	}
	return nil, nil

}
