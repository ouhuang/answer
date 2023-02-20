package v1

import (
	"2333/answer/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(ctx *gin.Context) {
	key := ctx.Query("key")

	data := service.SearchService(key)

	// if data
	ctx.JSON(http.StatusOK, gin.H{"data": data})
	//  else
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"value": data})

}
