package controller

import (
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
	code, response := c.service.Register(employee)

	ctx.JSON(code, response)
	return
}