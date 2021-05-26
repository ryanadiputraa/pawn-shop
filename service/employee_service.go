package service

import (
	"net/http"

	"github.com/ryanadiputraa/pawn-shop/config"
	"github.com/ryanadiputraa/pawn-shop/entity"
)

type EmployeeService interface {
	GetAllEmployees() entity.EmployeesResponse
	Register(entity.Employee) int
}

type employeeService struct {
	employees entity.EmployeesResponse
}

func New() EmployeeService {
	return &employeeService{}
}


func (service *employeeService) GetAllEmployees() entity.EmployeesResponse {
	return service.employees
}

func (service *employeeService) Register(employee entity.Employee) int {
	db, err := config.OpenConnection()
	if err != nil {
		return http.StatusBadGateway
	}
	defer db.Close()

	query := `INSERT INTO employees (firstname, lastname, gender, birthdate, address, password) VALUES ($1, $2, $3, $4, $5, $6)`
	
	_, err = db.Exec(query, employee.Firstname, employee.Lastname, employee.Gender, employee.Birthdate, employee.Address, employee.Password)
	if err != nil {
		return http.StatusBadRequest
	}

	return http.StatusAccepted
}