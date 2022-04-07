package server

import (
	"github.com/gin-gonic/gin"
	"my_blog/admin/configs"
	_ "my_blog/admin/docs" // 千万不要忘了导入把你上一步生成的docs
	"my_blog/admin/internal/service"
	"my_blog/utils"

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	r := engine.Group("/admin")
	{
		r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
		r.POST("/login", login)
	}

	g := engine.Group("/my-blog/admin")
	g.Use(utils.JwtToken())
	{
		g.POST("/user/add", addUser)
		g.POST("/user/delete", delUser)
		g.POST("/user/update", updateUser)
		g.POST("/user/updatePwd", updateUser)

		g.GET("/category/getCategory", getCategory)
		g.POST("/category/addCategory", addCategory)
		g.POST("/category/delCategory", delCategory)
		g.POST("/category/updateCategory", updateCategory)
	}
}
