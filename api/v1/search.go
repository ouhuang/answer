package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"key": 1, "value": 2})
}
