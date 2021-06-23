package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/pawn-shop/helper"
	"github.com/ryanadiputraa/pawn-shop/service"
)

type EmployeeController interface {
	GetAllEmployees(ctx *gin.Context)
	Register(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetEmployeeById(ctx *gin.Context)
	DeleteEmployee(ctx *gin.Context)
	Login(ctx *gin.Context)
	LoginAdmin(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type employeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) EmployeeController {
	return &employeeController {
		service: service,
	}
}

func (c *employeeController) GetAllEmployees(ctx *gin.Context) {
	code, response := c.service.GetAllEmployees(ctx)
	helper.WriteResponse(ctx, code, response)
}

func (c *employeeController) GetEmployeeById(ctx *gin.Context) {
	code, response := c.service.GetEmployeeById(ctx)
	helper.WriteResponse(ctx, code, response)
}

func (c *employeeController) Register(ctx *gin.Context) {
	code, response := c.service.Register(ctx)
	helper.WriteResponse(ctx, code, response)
}

func (c *employeeController) Update(ctx *gin.Context) {
	code, response := c.service.Update(ctx)
	helper.WriteResponse(ctx, code, response)
}

func (c *employeeController) Login(ctx *gin.Context) {
	code, response := c.service.Login(ctx)
	helper.WriteResponse(ctx, code, response)
}

func (c *employeeController) LoginAdmin(ctx *gin.Context) {
	code, response := c.service.LoginAdmin(ctx)
	helper.WriteResponse(ctx, code, response)
}

func (c *employeeController) Logout(ctx *gin.Context) {
	code, response := c.service.Logout(ctx)
	helper.WriteResponse(ctx, code, response)
}

func (c *employeeController) DeleteEmployee(ctx *gin.Context) {
	code, response := c.service.DeleteEmployee(ctx)
	helper.WriteResponse(ctx, code, response)
}