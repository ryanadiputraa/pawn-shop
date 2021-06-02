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
)
const SecretKey = "secret"
type EmployeeService interface {
	GetAllEmployees() (int, interface{})
	GetEmployeeById(employee_id string) (int, interface{})
	Register(entity.Employee) (int, interface{})
	Login(loginData entity.LoginEmployee, ctx *gin.Context) (int, interface{})
	Logout(ctx *gin.Context) (int, interface{})
	DeleteEmployee(employee_id string) (int, interface{})
}

type employeeService struct {}

func NewEmployeeService() EmployeeService {
	return &employeeService{}
}

func (service *employeeService) GetAllEmployees() (int, interface{}) {
	db, err := config.OpenConnection()
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadGateway,
			Error: "can't open db connection",
		}
		return http.StatusBadGateway, response
	}
	defer db.Close()

	query := `SELECT * FROM employees`

	rows, err := db.Query(query)
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadRequest,
			Error: "can't get employees data",
		}
		return http.StatusBadRequest, response
	}
	defer rows.Close()

	var employees []entity.Employee
	for rows.Next() {
		var employee entity.Employee
		rows.Scan(&employee.ID, &employee.Firstname, &employee.Lastname, &employee.Gender, &employee.Birthdate, &employee.Address, &employee.Password)
		employees = append(employees, employee)
	}

	return http.StatusOK, employees
}

func (service *employeeService) GetEmployeeById(employee_id string) (int, interface{}) {
	db, err := config.OpenConnection()
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadGateway,
			Error: "can't open db connection",
		}
		return http.StatusBadGateway, response
	}
	defer db.Close()

	row, err := db.Query(fmt.Sprintf("SELECT * FROM employees WHERE employee_id = %v", employee_id))
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadRequest,
			Error: "can't get employee data",
		}
		return http.StatusBadRequest, response
	}
	defer row.Close()

	var employee entity.Employee
	isNotNull := row.Next()
	if !isNotNull {
		response := entity.Error {
			Code: http.StatusNotFound,
			Error: "no employee with given id",
		}
		return http.StatusNotFound, response	
	}
	row.Scan(&employee.ID, &employee.Firstname, &employee.Lastname, &employee.Gender, &employee.Birthdate, &employee.Address, &employee.Password)
	
	return http.StatusOK, employee
}

func (service *employeeService) Register(employee entity.Employee) (int, interface{}) {
	db, err := config.OpenConnection()
	if err != nil {
		response := entity.Error {
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
		response := entity.Error {
			Code: http.StatusBadRequest,
			Error: "can't insert data into db",
		}
		return http.StatusBadRequest, response
	}

	response := entity.HTTPCode {
		Code: http.StatusCreated,
	}
	return http.StatusCreated, response
}

func (service *employeeService) Login(loginData entity.LoginEmployee, ctx *gin.Context) (int, interface{}) {
	db, err := config.OpenConnection()
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadGateway,
			Error: "can't open db connection",
		}
		return http.StatusBadGateway, response
	}
	defer db.Close()
	
	// check employee id
	row, err := db.Query(fmt.Sprintf("SELECT employee_id, password FROM employees WHERE employee_id = %v", strconv.Itoa(loginData.ID)))
	if err != nil {
		response := entity.Error {
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
			response := entity.Error {
				Code: http.StatusNotFound,
				Error: "no employee with given id",
			}
			return http.StatusNotFound, response	
		}
	}
	
	// check employee password
	// if err = bcrypt.CompareHashAndPassword(employee.Password, []byte(loginData.Password)); err != nil {
	if (employee.Password != loginData.Password) {
		response := entity.Error {
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
		response := entity.Error {
			Code: http.StatusInternalServerError,
			Error: "can't generate token",
		}
		return http.StatusInternalServerError, response
	}

	ctx.SetCookie("jwt", token, 60*60*24, "", "", true, true)
	response := entity.HTTPCode { Code: http.StatusAccepted }

	return http.StatusAccepted, response
}

func (service *employeeService) Logout(ctx *gin.Context) (int, interface{}) {
	ctx.SetCookie("jwt", "", 0, "", "", true, true)
	response := entity.HTTPCode { Code: http.StatusOK }
	return http.StatusOK, response
}

func (service *employeeService) DeleteEmployee(employee_id string) (int, interface{}) {
	db, err := config.OpenConnection()
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadGateway,
			Error: "can't open db connection",
		}
		return http.StatusBadGateway, response
	}
	defer db.Close()

	_, err = db.Query(fmt.Sprintf("DELETE FROM employees WHERE employee_id = %v", employee_id))
	if err != nil {
		response := entity.Error {
			Code: http.StatusNotFound,
			Error: "no employee with given id",
		}
		return http.StatusNotFound, response
	}

	response := entity.HTTPCode { Code: http.StatusOK }

	return http.StatusOK, response
}