package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"my_blog/admin/internal/model"
	dmDao "my_blog/domain/model"
	. "my_blog/utils"
	"strings"
	"time"
)

func (s *Service) UpdateCategory(c *gin.Context, req *model.UpdateCategoryReq) (state int, message string) {
	category := &dmDao.Category{
		ID:          req.Id,
		Name:        req.Name,
		UpdatedTime: time.Now(),
	}
	err := s.dmDao.UpdateCategory(c, category)
	if err != nil {
		log.Error("s.dmDao.UpdateCategory err ->", err)
		state = ERROR
		message = GetErrMsg(state)
		return
	}
	state = SUCCSE
	message = GetErrMsg(state)
	return
}

func (s *Service) DelCategory(c *gin.Context, req *model.DeleteCategoryReq) (state int, message string) {
	category := &dmDao.Category{
		ID:          req.Id,
		State:       dmDao.CategoryDelete,
		UpdatedTime: time.Now(),
	}
	err := s.dmDao.DelCategory(c, category)
	if err != nil {
		log.Error("s.dmDao.DelCategory err ->", err)
		state = ERROR
		message = GetErrMsg(state)
		return
	}
	state = SUCCSE
	message = GetErrMsg(state)
	return
}

func (s *Service) AddCategory(c *gin.Context, req *model.AddCategoryReq) (state int, message string) {
	category := &dmDao.Category{
		State:       dmDao.CategoryInUse,
		Name:        req.Name,
		CreatedTime: time.Now(),
	}
	err := s.dmDao.AddCategory(c, category)
	if err != nil {
		log.Error("s.dmDao.AddCategory err ->", err)
		state = ERROR
		message = GetErrMsg(state)
	}
	state = SUCCSE
	message = GetErrMsg(state)
	return
}

func formatParas(req *model.GetCategoryReq) (where string, paras []interface{}) {

	whereArr := []string{}
	if req.Id != 0 {
		whereArr = append(whereArr, "id = ?")
		paras = append(paras, req.Id)
	}
	if req.Name != "" {
		whereArr = append(whereArr, "name like ?")
		paras = append(paras, "%"+req.Name+"%")
	}

	whereArr = append(whereArr, "state <> ?")
	paras = append(paras, dmDao.CategoryDelete)

	where = strings.Join(whereArr, " and ")
	return
}

func (s *Service) GetCategory(c *gin.Context, req *model.GetCategoryReq) (state int, message string, data *model.GetCategoryListResp) {
	data = new(model.GetCategoryListResp)
	where, paras := formatParas(req)
	res, err := s.dmDao.GetCategoryList(c, where, paras...)
	if err != nil {
		log.Error("s.dmDao.GetCategoryList err ->", err)
		state = ERROR
		message = GetErrMsg(state)
		return
	}
	for _, val := range res {
		t := &model.CategoryList{
			Id:   val.ID,
			Name: val.Name,
		}
		data.List = append(data.List, t)
	}
	state = SUCCSE
	message = GetErrMsg(state)
	return
}
