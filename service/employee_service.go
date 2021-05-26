package service

import (
	"fmt"
	"net/http"

	"github.com/ryanadiputraa/pawn-shop/config"
	"github.com/ryanadiputraa/pawn-shop/entity"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeService interface {
	GetAllEmployees() (int, interface{})
	GetEmployeeById(employee_id string) (int, interface{})
	Register(entity.Employee) (int, interface{})
	DeleteEmployee(employee_id string) (int, interface{})
}

type employeeService struct {}

func New() EmployeeService {
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

func (service *employeeService)  GetEmployeeById(employee_id string) (int, interface{}) {
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(employee.Password), bcrypt.DefaultCost)
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadRequest,
			Error: "can't hash user password",
		}
		return http.StatusInternalServerError, response
	}

	query := `INSERT INTO employees (firstname, lastname, gender, birthdate, address, password) VALUES ($1, $2, $3, $4, $5, $6)`
	
	_, err = db.Exec(query, employee.Firstname, employee.Lastname, employee.Gender, employee.Birthdate, employee.Address, hashedPassword)
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadRequest,
			Error: "can't insert data into db",
		}
		return http.StatusBadRequest, response
	}

	response := entity.HTTPCode {
		Code: http.StatusAccepted,
	}
	return http.StatusCreated, response
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