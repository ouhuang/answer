package router

import (
	v1 "2333/answer/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(Router *gin.RouterGroup) {
	Router.Any("search", v1.Search)
}
