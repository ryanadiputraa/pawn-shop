package service

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/pawn-shop/config"
	"github.com/ryanadiputraa/pawn-shop/entity"
	"github.com/ryanadiputraa/pawn-shop/repository"
)
const SecretKey = "secret"
type EmployeeService interface {
	GetAllEmployees(ctx *gin.Context) (code int, response interface{})
	GetEmployeeById(ctx *gin.Context) (code int, response interface{})
	Register(ctx *gin.Context) (code int, response interface{})
	Update(ctx *gin.Context) (code int, response interface{})
	Login(ctx *gin.Context) (code int, response interface{})
	LoginAdmin(ctx *gin.Context) (code int, response interface{})
	Logout(ctx *gin.Context) (code int, response interface{})
	DeleteEmployee(ctx *gin.Context) (code int, response interface{})
}

type employeeService struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(repository repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		repository: repository,
	}
}

func (service *employeeService) GetAllEmployees(ctx *gin.Context) (code int, response interface{}) {
	cookie, err := ctx.Cookie("jwt")
	if err != nil {
		response = entity.Error {
			Code: http.StatusUnauthorized,
			Error: "no cookie found",
		}
		return http.StatusUnauthorized, response
	}

	_, err = jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetAdminKey()), nil
	})
	if err != nil {
		response = entity.Error {
			Code: http.StatusUnauthorized,
			Error: "unauthorized",
		}
		return http.StatusUnauthorized, response
	}

	URLQuery := ctx.Request.URL.Query()
	var queryParam string
	if len(URLQuery) > 0 {
		queryParam = URLQuery["query"][0]
	}

	employees, code, err := service.repository.GetAll(queryParam)
	if err != nil {
		response = entity.Error {
			Code: code,
			Error: err.Error(),
		}
		return code, response	
	}

	if employees == nil {
		response = entity.Error {
			Code: http.StatusNotFound,
			Error: "no employee found",
		}
		return http.StatusNotFound, response
	}

	return code, employees
}

func (service *employeeService) GetEmployeeById(ctx *gin.Context) (code int, response interface{}) {
	cookie, err := ctx.Cookie("jwt")
	if err != nil {
		response = entity.Error {
			Code: http.StatusUnauthorized,
			Error: "no cookie found",
		}
		return http.StatusUnauthorized, response
	}

	_, err = jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetAdminKey()), nil
	})
	if err != nil {
		response = entity.Error {
			Code: http.StatusUnauthorized,
			Error: "unauthorized",
		}
		return http.StatusUnauthorized, response
	}
	
	employeeId := ctx.Param("employee_id")
	employee, code, err := service.repository.GetById(employeeId)
	if err != nil {
		response = entity.Error {
			Code: code,
			Error: err.Error(),
		}
		return code, response
	}

	return code, employee
}

func (service *employeeService) Register(ctx *gin.Context) (code int, response interface{}) {
	cookie, err := ctx.Cookie("jwt")
	if err != nil {
		response = entity.Error {
			Code: http.StatusUnauthorized,
			Error: "no cookie found",
		}
		return http.StatusUnauthorized, response
	}

	_, err = jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetAdminKey()), nil
	})
	if err != nil {
		response = entity.Error {
			Code: http.StatusUnauthorized,
			Error: "unauthorized",
		}
		return http.StatusUnauthorized, response
	}

	var employee entity.Employee
	ctx.BindJSON(&employee)

	code, err = service.repository.Create(employee)
	if err != nil {
		response = entity.Error {
			Code: code,
			Error: err.Error(),
		}
		return code, response
	}

	response = entity.HTTPCode {
		Code: code,
	}

	return code, response
}

func (service *employeeService) Update(ctx *gin.Context) (code int, response interface{}) {
	cookie, err := ctx.Cookie("jwt")
	if err != nil {
		response = entity.Error {
			Code: http.StatusUnauthorized,
			Error: "no cookie found",
		}
		return http.StatusUnauthorized, response
	}

	_, err = jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetAdminKey()), nil
	})
	if err != nil {
		response = entity.Error {
			Code: http.StatusUnauthorized,
			Error: "unauthorized",
		}
		return http.StatusUnauthorized, response
	}

	employeeId := ctx.Param("employee_id")
	var employee entity.Employee
	ctx.BindJSON(&employee)

	employee.ID, err = strconv.Atoi(employeeId)
	if err != nil {
		response = entity.Error {
			Code: http.StatusInternalServerError,
			Error: "cant convert employee id",
		}
		return http.StatusInternalServerError, response
	}

	code, err = service.repository.Update(employee)
	if err != nil {
		response = entity.Error {
			Code: code,
			Error: err.Error(),
		}
		return code, response
	}

	response = entity.HTTPCode {
		Code: code,
	}
	return code, response
}

func (service *employeeService) Login(ctx *gin.Context) (code int, response interface{}) {
	var loginPayload entity.LoginEmployee
	ctx.BindJSON(&loginPayload)

	employee, code, err := service.repository.CheckPassword(loginPayload)
	if err != nil {
		response = entity.Error {
			Code: code,
			Error: err.Error(),
		}
		return code, response
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(employee.ID),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(config.GetSecretKey()))
	if err != nil {
		response = entity.Error {
			Code: http.StatusInternalServerError,
			Error: "can't generate token",
		}
		return http.StatusInternalServerError, response
	}

	ctx.SetCookie("jwt", token, 60*60*24, "", "", true, true)
	response = entity.HTTPCode { Code: code }

	return code, response
}

func (service *employeeService) LoginAdmin(ctx *gin.Context) (code int, response interface{}) {
	var loginPayload entity.LoginEmployee
	ctx.BindJSON(&loginPayload)
	
	if loginPayload.ID != 123 && loginPayload.Password != "admin" {
		response = entity.Error {
			Code: http.StatusUnauthorized,
			Error: "unauthorized",
		}
		return http.StatusUnauthorized, response
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(loginPayload.ID),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(config.GetAdminKey()))
	if err != nil {
		response = entity.Error {
			Code: http.StatusInternalServerError,
			Error: "can't generate token",
		}
		return http.StatusInternalServerError, response
	}

	ctx.SetCookie("jwt", token, 60*60*24, "", "", true, true)
	response = entity.HTTPCode { Code: http.StatusAccepted }

	return http.StatusAccepted, response
}

func (service *employeeService) Logout(ctx *gin.Context) (code int, response interface{}) {
	ctx.SetCookie("jwt", "", 60, "", "", true, true)
	response = entity.HTTPCode { Code: http.StatusOK }
	return http.StatusOK, response
}

func (service *employeeService) DeleteEmployee(ctx *gin.Context) (code int, response interface{}) {
	employeeId := ctx.Param("employee_id")

	code, err := service.repository.Delete(employeeId)
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