package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"my_blog/admin/internal/model"
)

func uploadFile(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	req := &model.UploadReq{File: file, Size: fileHeader.Size}
	log.Info("start updateArticle, req [%v]", req)
	s, m, d := svc.UploadFile(c, req)
	svc.Render(c, s, m, d)

}
