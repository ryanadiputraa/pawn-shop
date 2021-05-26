package helper

import "github.com/gin-gonic/gin"

func WriteResponse(ctx *gin.Context, code int, response interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(code, response)
}