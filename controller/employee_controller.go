package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/pawn-shop/entity"
	"github.com/ryanadiputraa/pawn-shop/service"
)

type EmployeeController interface {
	GetAllEmployees(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type controller struct {
	service service.EmployeeService
}

func New(service service.EmployeeService) EmployeeController {
	return &controller {
		service: service,
	}
}

func (c *controller) GetAllEmployees(ctx *gin.Context){
	c.service.GetAllEmployees()
	return
}

func (c *controller) Register(ctx *gin.Context){
	var employee entity.Employee
	ctx.BindJSON(&employee)
	code := c.service.Register(employee)

	response := entity.HTTPCode { Code: http.StatusAccepted }
	
	if code == http.StatusBadGateway {
		response := entity.Error{
			Code: code,
			Error: "fail to open db connection",
		}
		ctx.JSON(code, response)
		return
	} else if code == http.StatusBadRequest {
		response := entity.Error{
			Code: code,
			Error: "cant insert data into db",
		}
		ctx.JSON(code, response)
		return
	}

	ctx.JSON(code, response)
	return
}