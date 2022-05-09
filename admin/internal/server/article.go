package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"my_blog/admin/internal/model"
	. "my_blog/utils"
)

func updateArticle(c *gin.Context) {
	req := &model.UpdateArticleReq{}
	err := c.ShouldBindJSON(req)
	if err != nil || req.Id == 0 {
		log.Error("bind updateArticle req false")
		svc.Render(c, ERROR_REQ_PARAS, GetErrMsg(ERROR_REQ_PARAS), nil)
		return

	}
	log.Info("start updateArticle, req [%v]", req)
	s, m := svc.UpdateArticle(c, req)
	svc.Render(c, s, m, nil)

}

func delArticle(c *gin.Context) {
	req := &model.DeleteCategoryReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("bind delCategory req false")
		svc.Render(c, ERROR_REQ_PARAS, GetErrMsg(ERROR_REQ_PARAS), nil)
	}
	log.Info("start updatePwd, req [%v]", req)
	s, m := svc.DelCategory(c, req)
	svc.Render(c, s, m, nil)

}

func addArticle(c *gin.Context) {
	req := &model.AddArticleReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("bind addArticle req false")
		svc.Render(c, ERROR_REQ_PARAS, GetErrMsg(ERROR_REQ_PARAS), nil)
		return
	}
	log.Info("start addArticle, req [%v]", req)
	s, m := svc.AddArticle(c, req)
	svc.Render(c, s, m, nil)

}

func getArticleInfo(c *gin.Context) {
	req := &model.GetArticleInfoReq{}
	err := c.ShouldBind(req)
	if err != nil {
		log.Error("bind getArticleInfo req false")
		svc.Render(c, ERROR_REQ_PARAS, GetErrMsg(ERROR_REQ_PARAS), nil)
		return
	}
	log.Info("start getArticleInfo, req [%v]", req)
	s, m, d := svc.GetArticleInfo(c, req)
	svc.Render(c, s, m, d)

}

func getArticleList(c *gin.Context) {
	req := &model.GetArticleListReq{}
	err := c.ShouldBind(req)
	if err != nil {
		log.Error("bind getArticleList req false")
		svc.Render(c, ERROR_REQ_PARAS, GetErrMsg(ERROR_REQ_PARAS), nil)
		return
	}
	log.Info("start getCategory, req [%v]", req)
	s, m, d := svc.GetArticleList(c, req)
	svc.Render(c, s, m, d)

}
