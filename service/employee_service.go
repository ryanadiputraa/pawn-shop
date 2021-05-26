package service

import (
	"net/http"

	"github.com/ryanadiputraa/pawn-shop/config"
	"github.com/ryanadiputraa/pawn-shop/entity"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeService interface {
	GetAllEmployees() (int, interface{})
	Register(entity.Employee) (int, interface{})
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

	var employees []entity.Employee
	for rows.Next() {
		var employee entity.Employee
		rows.Scan(&employee.ID, &employee.Firstname, &employee.Lastname, &employee.Gender, &employee.Birthdate, &employee.Address, employee.Password)
		employees = append(employees, employee)
	}

	return 200, employees
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