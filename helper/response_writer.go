package helper

import "github.com/gin-gonic/gin"

func WriteResponse(ctx *gin.Context, code int, response interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, response)
	return
}