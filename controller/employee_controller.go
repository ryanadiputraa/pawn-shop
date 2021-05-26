package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/pawn-shop/entity"
	"github.com/ryanadiputraa/pawn-shop/helper"
	"github.com/ryanadiputraa/pawn-shop/service"
)

type EmployeeController interface {
	GetAllEmployees(ctx *gin.Context)
	Register(ctx *gin.Context)
	GetEmployeeById(ctx *gin.Context)
	DeleteEmployee(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type controller struct {
	service service.EmployeeService
}

func New(service service.EmployeeService) EmployeeController {
	return &controller {
		service: service,
	}
}

func (c *controller) GetAllEmployees(ctx *gin.Context) {
	code, response := c.service.GetAllEmployees()

	helper.WriteResponse(ctx, code, response)
}

func (c *controller) GetEmployeeById(ctx *gin.Context) {
	employeeId := ctx.Param("employee_id")
	code, response := c.service.GetEmployeeById(employeeId)

	helper.WriteResponse(ctx, code, response)
}

func (c *controller) Register(ctx *gin.Context) {
	var employee entity.Employee
	ctx.BindJSON(&employee)
	code, response := c.service.Register(employee)

	helper.WriteResponse(ctx, code, response)
}

func (c *controller) Login(ctx *gin.Context) {
	var loginData entity.LoginEmployee
	ctx.BindJSON(&loginData)
	code, response := c.service.Login(loginData, ctx)

	helper.WriteResponse(ctx, code, response)
}

func (c *controller) DeleteEmployee(ctx *gin.Context) {
	employeeId := ctx.Param("employee_id")
	code, response := c.service.DeleteEmployee(employeeId)

	helper.WriteResponse(ctx, code, response)
}