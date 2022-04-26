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

func (s *Service) UpdateArticle(c *gin.Context, req *model.UpdateArticleReq) (state int, message string) {
	art := &dmDao.Article{
		ID:        req.Id,
		Name:      req.Name,
		Cid:       req.Cid,
		Desc:      req.Desc,
		Content:   req.Content,
		Img:       req.Img,
		UpdatedAt: time.Now(),
	}
	err := s.dmDao.UpdateArticle(c, art)
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

func (s *Service) DelArticle(c *gin.Context, req *model.DelArticleReq) (state int, message string) {
	art := &dmDao.Article{
		ID:        req.Id,
		UpdatedAt: time.Now(),
		State:     dmDao.ArticleDelete,
	}
	err := s.dmDao.DelArticle(c, art)
	if err != nil {
		log.Error("s.dmDao.DelArticle err ->", err)
		state = ERROR
		message = GetErrMsg(state)
		return
	}
	state = SUCCSE
	message = GetErrMsg(state)
	return
}

func (s *Service) AddArticle(c *gin.Context, req *model.AddArticleReq) (state int, message string) {
	art := &dmDao.Article{
		State:     dmDao.ArticleInUse,
		Name:      req.Name,
		Cid:       req.Cid,
		Desc:      req.Desc,
		Content:   req.Content,
		Img:       req.Img,
		CreatedAt: time.Now(),
	}
	err := s.dmDao.AddArticle(c, art)
	if err != nil {
		log.Error("s.dmDao.AddArticle err ->", err)
		state = ERROR
		message = GetErrMsg(state)
	}
	state = SUCCSE
	message = GetErrMsg(state)
	return
}

func formatSearchArticleListParas(req *model.GetArticleListReq) (where string, paras []interface{}) {
	whereArr := []string{}
	if req.Id != 0 {
		whereArr = append(whereArr, "id = ?")
		paras = append(paras, req.Id)
	}
	if req.Name != "" {
		whereArr = append(whereArr, "name like ?")
		paras = append(paras, "%"+req.Name+"%")
	}

	if req.Cid != 0 {
		whereArr = append(whereArr, "cid = ?")
		paras = append(paras, req.Cid)
	}

	whereArr = append(whereArr, "state <> ?")
	paras = append(paras, dmDao.ArticleDelete)

	where = strings.Join(whereArr, " and ")
	return
}

func (s *Service) GetArticleList(c *gin.Context, req *model.GetArticleListReq) (state int, message string, data *model.GetArticleListResp) {
	where, paras := formatSearchArticleListParas(req)
	res, total, err := s.dmDao.GetArticleList(c, req.PageSize, req.PageNum, where, paras...)
	if err != nil {
		log.Error("s.dmDao.GetCategoryList err ->", err)
		state = ERROR
		message = GetErrMsg(state)
		return
	}
	data = new(model.GetArticleListResp)
	data.Total = total
	for _, val := range res {
		temp := &model.GetArticleListData{
			Id:           val.ID,
			Name:         val.Name,
			Desc:         val.Desc,
			Cid:          val.Cid,
			CommentCount: val.CommentCount,
			ReadCount:    val.ReadCount,
			UpdateTime:   val.UpdatedAt,
		}
		data.List = append(data.List, temp)
	}
	state = SUCCSE
	message = GetErrMsg(state)
	return
}

func (s *Service) GetArticleInfo(c *gin.Context, req *model.GetArticleInfoReq) (state int, message string, data *model.GetArticleInfoResp) {
	res, err := s.dmDao.GetArticleInfo(c, "id = ?", req.Id)
	data = new(model.GetArticleInfoResp)
	data.Content = res.Content
	data.Img = res.Img
	data.Name = res.Name
	data.Cid = res.Cid
	data.Id = res.ID
	data.Desc = res.Desc
	data.CommentCount = res.CommentCount
	data.ReadCount = res.ReadCount
	if err != nil {
		log.Error("s.dmDao.GetArticleInfo err ->", err)
		state = ERROR
		message = GetErrMsg(state)
		return
	}
	state = ERROR
	message = GetErrMsg(state)
	return
}
