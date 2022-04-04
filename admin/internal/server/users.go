package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"my_blog/admin/internal/model"
)

func login(c *gin.Context) {
	req := &model.LoginReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("bind login req false")
		c.JSON(400, err)

	}
	log.Info("start login, req [%v]", req)
	s, m, d := svc.Login(c, req)
	svc.Render(c, s, m, d)

}
