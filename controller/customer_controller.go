package controller

import "github.com/ryanadiputraa/pawn-shop/service"

type CustomerController interface {

}

type customerContoller struct {
	service service.CustomerService
}


func NewCustomerController(service service.CustomerService) CustomerController {
	return &customerContoller{
		service: service,
	}
}