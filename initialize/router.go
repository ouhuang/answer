package initialize

import (
	"2333/answer/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	ApiGroup := Router.Group("api")

	router.InitRouter(ApiGroup)

	return Router
}
