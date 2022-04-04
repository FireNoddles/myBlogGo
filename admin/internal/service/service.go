package service

import (
	"my_blog/admin/configs"
	"my_blog/admin/internal/dao"
	dmDao "my_blog/domain/dao"
)

type Service struct {
	conf  *configs.Configs
	dao   *dao.Dao
	dmDao dmDao.Dao
}

func NewService(db *dao.Dao, conf *configs.Configs) *Service {
	s := &Service{
		conf: conf,
		dao:  db,
	}
	s.dmDao = dmDao.NewDmDao(db.Db.MyBlog)
	return s
}
