package server

import (
	"github.com/gin-gonic/gin"
	"my_blog/admin/configs"
	"my_blog/admin/internal/service"
)

var (
	svc *service.Service
)

func NewServer(conf *configs.Configs, s *service.Service) (engine *gin.Engine, err error) {
	svc = s
	engine = gin.Default()
	initRouter(engine)
	err = engine.Run(conf.Server.Port)
	return engine, err

}

func initRouter(engine *gin.Engine) {
	g := engine.Group("/my-blog/admin")
	{
		g.POST("/login", login)
	}
}
