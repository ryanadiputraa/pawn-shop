package service

import (
	"fmt"
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
	Register(entity.Employee) (code int, response interface{})
	Update(employee entity.Employee, employee_id string) (code int, response interface{})
	Login(loginData entity.LoginEmployee, ctx *gin.Context) (code int, response interface{})
	LoginAdmin(loginData entity.LoginEmployee, ctx *gin.Context) (code int, response interface{})
	Logout(ctx *gin.Context) (code int, response interface{})
	DeleteEmployee(employee_id string) (code int, response interface{})
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

	employees, code, err := service.repository.GetAll(ctx)
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

func (service *employeeService) Register(employee entity.Employee) (code int, response interface{}) {
	db, err := config.OpenConnection()
	if err != nil {
		response = entity.Error {
			Code: http.StatusBadGateway,
			Error: "can't open db connection",
		}
		return http.StatusBadGateway, response
	}
	defer db.Close()

	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(employee.Password), 14)
	// if err != nil {
	// 	response := entity.Error {
	// 		Code: http.StatusBadRequest,
	// 		Error: "can't hash user password",
	// 	}
	// 	return http.StatusInternalServerError, response
	// }

	query := `INSERT INTO employees (firstname, lastname, gender, birthdate, address, password) VALUES ($1, $2, $3, $4, $5, $6)`
	
	_, err = db.Exec(query, employee.Firstname, employee.Lastname, employee.Gender, employee.Birthdate, employee.Address, employee.Password)
	if err != nil {
		response = entity.Error {
			Code: http.StatusBadRequest,
			Error: "can't insert data into db",
		}
		return http.StatusBadRequest, response
	}

	response = entity.HTTPCode {
		Code: http.StatusCreated,
	}
	return http.StatusCreated, response
}

func (service *employeeService) Update(employee entity.Employee, employee_id string) (code int, response interface{}) {
	db, err := config.OpenConnection()
	if err != nil {
		response = entity.Error {
			Code: http.StatusBadGateway,
			Error: "can't open db connection",
		}
		return http.StatusBadGateway, response
	}
	defer db.Close()

	query := fmt.Sprintf("UPDATE employees SET firstname = '%v', lastname = '%v', gender = '%v', birthdate = '%v', address = '%v', password = '%v' WHERE employee_id = %v", employee.Firstname, employee.Lastname, employee.Gender, employee.Birthdate, employee.Address, employee.Password, employee_id)

	_, err = db.Exec(query)
	if err != nil {
		response = entity.Error {
			Code: http.StatusBadRequest,
			Error: "can't update employee data",
		}
		return http.StatusBadRequest, response
	}

	response = entity.HTTPCode {
		Code: http.StatusOK,
	}
	return http.StatusOK, response
}

func (service *employeeService) Login(loginData entity.LoginEmployee, ctx *gin.Context) (code int, response interface{}) {
	db, err := config.OpenConnection()
	if err != nil {
		response = entity.Error {
			Code: http.StatusBadGateway,
			Error: "can't open db connection",
		}
		return http.StatusBadGateway, response
	}
	defer db.Close()
	
	// check employee id
	row, err := db.Query(fmt.Sprintf("SELECT employee_id, password FROM employees WHERE employee_id = %v", strconv.Itoa(loginData.ID)))
	if err != nil {
		response = entity.Error {
			Code: http.StatusBadRequest,
			Error: "can't get employee data",
		}
		fmt.Println(loginData.ID)
		return http.StatusBadRequest, response
	}
	defer row.Close()

	var employee entity.LoginEmployee
	for row.Next() {
		row.Scan(&employee.ID, &employee.Password)
		if employee.ID == 0 {
			response = entity.Error {
				Code: http.StatusNotFound,
				Error: "no employee with given id",
			}
			return http.StatusNotFound, response	
		}
	}
	
	// check employee password
	// if err = bcrypt.CompareHashAndPassword(employee.Password, []byte(loginData.Password)); err != nil {
	if (employee.Password != loginData.Password) {
		response = entity.Error {
			Code: http.StatusUnauthorized,
			Error: "wrong password",
		}
		return http.StatusUnauthorized, response	
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
	response = entity.HTTPCode { Code: http.StatusAccepted }

	return http.StatusAccepted, response
}

func (service *employeeService) LoginAdmin(loginData entity.LoginEmployee, ctx *gin.Context) (code int, response interface{}) {
	if loginData.ID != 123 && loginData.Password != "admin" {
		response = entity.Error {
			Code: http.StatusUnauthorized,
			Error: "unauthorized",
		}
		return http.StatusUnauthorized, response
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(loginData.ID),
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

func (service *employeeService) DeleteEmployee(employee_id string) (code int, response interface{}) {
	db, err := config.OpenConnection()
	if err != nil {
		response = entity.Error {
			Code: http.StatusBadGateway,
			Error: "can't open db connection",
		}
		return http.StatusBadGateway, response
	}
	defer db.Close()

	_, err = db.Query(fmt.Sprintf("DELETE FROM employees WHERE employee_id = %v", employee_id))
	if err != nil {
		response = entity.Error {
			Code: http.StatusNotFound,
			Error: "no employee with given id",
		}
		return http.StatusNotFound, response
	}

	response = entity.HTTPCode { Code: http.StatusOK }

	return http.StatusOK, response
}