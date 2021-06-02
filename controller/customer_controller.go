package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/pawn-shop/helper"
	"github.com/ryanadiputraa/pawn-shop/service"
)

type CustomerController interface {
	GetAllCustomer(ctx *gin.Context)
	CreateLoan(ctx *gin.Context)
	PayOffLoan(ctx *gin.Context)
}

type customerContoller struct {
	service service.CustomerService
}


func NewCustomerController(service service.CustomerService) CustomerController {
	return &customerContoller{
		service: service,
	}
}

func (c *customerContoller) GetAllCustomer(ctx *gin.Context) {
	code, response := c.service.GetAllCustomer(ctx)

	helper.WriteResponse(ctx, code, response)
}

func (c *customerContoller) CreateLoan(ctx *gin.Context) {
	code, response := c.service.CreateLoan(ctx)

	helper.WriteResponse(ctx, code, response)
}

func (c *customerContoller) PayOffLoan(ctx *gin.Context) {
	customerId := ctx.Param("customer_id")
	code, response := c.service.PayOffLoan(customerId)

	helper.WriteResponse(ctx, code, response)
}