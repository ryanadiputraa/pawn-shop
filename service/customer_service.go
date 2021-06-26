package service

import (
	"fmt"
	"os"
	"strconv"

	// "io"
	"net/http"
	// "os"

	// "strconv"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
	"github.com/ryanadiputraa/pawn-shop/entity"
	"github.com/ryanadiputraa/pawn-shop/repository"
)

type CustomerService interface {
	GetAllCustomer(ctx *gin.Context) (code int, response interface{})
	CreateLoan(ctx *gin.Context) (code int, response interface{})
	PayOffLoan(ctx *gin.Context) (code int, response interface{})
	GetFinancialStatements(ctx *gin.Context) (code int, response interface{})
}

type customerService struct {
	repository repository.CustomerRepository
}

func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	return &customerService {
		repository: repository,
	}
}

func (service *customerService) GetAllCustomer(ctx *gin.Context) (code int, response interface{}) {
	URLQuery := ctx.Request.URL.Query()
	var queryParam string
	if len(URLQuery) > 0 {
		queryParam = URLQuery["query"][0]
	}

	customers, code, err := service.repository.GetAll(queryParam)
	if err != nil {
		response = entity.Error {
			Code: code,
			Error: err.Error(),
		}
		return code, response
	}

	if customers == nil {
		response = entity.Error {
			Code: http.StatusNotFound,
			Error: "no customers with given name",
		}
		return http.StatusNotFound, response
	}

	return code, customers
}

func (service *customerService) CreateLoan(ctx *gin.Context) (code int, response interface{}) {
	file, err := ctx.FormFile("upload")
	if err != nil {
		response = entity.Error {
			Code: http.StatusBadRequest,
			Error: err.Error(),
		}
		return http.StatusBadRequest, response
	}

	err = ctx.SaveUploadedFile(file, "./uploads/"+file.Filename)
	if err != nil {
		response = entity.Error {
			Code: http.StatusBadRequest,
			Error: err.Error(),
		}
		return http.StatusBadRequest, response
	}

	var customer entity.Customer
	customer.Firstname = ctx.Request.FormValue("firstname")
	customer.Lastname = ctx.Request.FormValue("lastname")
	customer.Gender = ctx.Request.FormValue("gender")
	customer.Contact = ctx.Request.FormValue("contact")
	customer.Loan, err = strconv.Atoi(ctx.Request.FormValue("loan"))
	if err != nil {
		response = entity.Error {
			Code: http.StatusBadRequest,
			Error: err.Error(),
		}
		return http.StatusBadRequest, response
	}
	customer.Interest = customer.Loan / 10
	customer.InsuranceItem = ctx.Request.FormValue("insuranceItem")
	customer.ItemStatus = "jaminan"
	customer.Image = fmt.Sprintf("%v/%v", os.Getenv("BASE_URL"), file.Filename)

	loanId, err := uuid.NewUUID()
	if err != nil {
		response = entity.Error {
			Code: http.StatusInternalServerError,
			Error: "can't generate uuid",
		}
		return http.StatusInternalServerError, response
	}

	itemId, err := uuid.NewUUID()
	if err != nil {
		response = entity.Error {
			Code: http.StatusInternalServerError,
			Error: "can't generate uuid",
		}
		return http.StatusInternalServerError, response
	}

	customer.ID , err = uuid.NewUUID()
	if err != nil {
		response = entity.Error {
			Code: http.StatusInternalServerError,
			Error: "can't generate uuid",
		}
		return http.StatusInternalServerError, response
	}	

	code, err = service.repository.CreateLoan(loanId, itemId, customer)
	if err != nil {
		response = entity.Error {
			Code:  code,
			Error: err.Error(),
		}
		return code, response
	}

	response = entity.HTTPCode { Code: code }

	return code, response
}

func (service *customerService) PayOffLoan(ctx *gin.Context) (code int, response interface{}) {
	customerId := ctx.Param("customer_id")
	parsedCustomerId, err := uuid.FromBytes([]byte(customerId))
	code, err = service.repository.PayOffLoan(parsedCustomerId)
	if err != nil {
		response = entity.Error {
			Code: code,
			Error: err.Error(),
		}
		return code, response
	}
	
	response = entity.HTTPCode { Code: code }

	return code, response
}

func (service *customerService) GetFinancialStatements(ctx *gin.Context) (code int, response interface{}) {
	financialStatements, code, err := service.repository.GetFinancialStatements()
	if err != nil {
		response = entity.Error {
			Code: code,
			Error: err.Error(),
		}
		return code, response
	}

	return code, financialStatements
}