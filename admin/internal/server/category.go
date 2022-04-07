package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"my_blog/admin/internal/model"
	. "my_blog/utils"
)

func updateCategory(c *gin.Context) {
	req := &model.UpdateCategoryReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("bind updateCategory req false")
		svc.Render(c, ERROR_REQ_PARAS, GetErrMsg(ERROR_REQ_PARAS), nil)

	}
	log.Info("start updateCategory, req [%v]", req)
	s, m := svc.UpdateCategory(c, req)
	svc.Render(c, s, m, nil)

}

func delCategory(c *gin.Context) {
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

func addCategory(c *gin.Context) {
	req := &model.AddCategoryReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("bind addCategory req false")
		svc.Render(c, ERROR_REQ_PARAS, GetErrMsg(ERROR_REQ_PARAS), nil)

	}
	log.Info("start addCategory, req [%v]", req)
	s, m := svc.AddCategory(c, req)
	svc.Render(c, s, m, nil)

}

func getCategory(c *gin.Context) {
	req := &model.GetCategoryReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("bind getCategory req false")
		svc.Render(c, ERROR_REQ_PARAS, GetErrMsg(ERROR_REQ_PARAS), nil)

	}
	log.Info("start getCategory, req [%v]", req)
	s, m, d := svc.GetCategory(c, req)
	svc.Render(c, s, m, d)

}
