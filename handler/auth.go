package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth(ctx *gin.Context) {
	user, pass, ok := ctx.Request.BasicAuth()
	if !ok || user != "admin" || pass != "admin" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Next()
}
