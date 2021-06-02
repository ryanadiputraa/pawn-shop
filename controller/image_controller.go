package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/pawn-shop/entity"
)

type ImageController interface {
	ServeImage(ctx *gin.Context)
}

type imageController struct {}

func NewImageController() ImageController {
	return &imageController{}
}

func (c *imageController) ServeImage(ctx *gin.Context) {
	filename := ctx.Param("image_path")

	img, err := os.Open(fmt.Sprintf("./uploads/%v", filename))
	if err != nil {
		response := entity.Error {
			Code: http.StatusInternalServerError,
			Error: err.Error(),
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	defer img.Close()

	ctx.Header("Content-Type", "image/jpeg")
	io.Copy(ctx.Writer, img)
	return	
}