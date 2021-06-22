package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/pawn-shop/service"
)

type ImageController interface {
	ServeImage(ctx *gin.Context)
}

type imageController struct {
	service service.ImageService
}

func NewImageController(service service.ImageService) ImageController {
	return &imageController{
		service: service,
	}
}

func (c *imageController) ServeImage(ctx *gin.Context) {
	c.service.ServeImage(ctx)	
	return
}