package repository

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/pawn-shop/config"
	"github.com/ryanadiputraa/pawn-shop/entity"
)

type EmployeeRepository interface {
	GetAll(ctx *gin.Context) (employees []entity.Employee, code int, err error) 
	GetById(employeeId string) (employee entity.Employee, code int, err error) 
}

type employeeRepository struct {}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{}
}

func (r *employeeRepository) GetAll(ctx *gin.Context) (employees []entity.Employee, code int, err error) {
	db, err := config.OpenConnection()
	if err != nil {
		return employees, http.StatusBadGateway, err
	}
	defer db.Close()

	param := ctx.Request.URL.Query()
	var query string
	if len(param) != 0 {
		query = fmt.Sprintf("SELECT * FROM employees WHERE CAST(employee_id AS TEXT) LIKE '%v%%' OR LOWER(firstname) LIKE LOWER('%v%%') OR LOWER(lastname) LIKE LOWER('%v%%') OR LOWER(gender) LIKE LOWER('%v%%') OR CAST(birthdate AS TEXT) LIKE '%v%%' OR LOWER(address) LIKE LOWER('%v%%') OR LOWER(password) LIKE LOWER('%v%%')", param["query"][0], param["query"][0], param["query"][0], param["query"][0], param["query"][0], param["query"][0], param["query"][0])
	} else {
		query = `SELECT * FROM employees`
	}

	rows, err := db.Query(query)
	if err != nil {
		return employees, http.StatusBadRequest, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee entity.Employee
		rows.Scan(&employee.ID, &employee.Firstname, &employee.Lastname, &employee.Gender, &employee.Birthdate, &employee.Address, &employee.Password)
		employees = append(employees, employee)
	}

	return employees, http.StatusOK, nil
}

func (r *employeeRepository) GetById(employeeId string) (employee entity.Employee, code int, err error) {
	db, err := config.OpenConnection()
	if err != nil {
		return employee, http.StatusBadGateway, err
	}
	defer db.Close()

	row, err := db.Query(fmt.Sprintf("SELECT * FROM employees WHERE employee_id = %v", employeeId))
	if err != nil {
		return employee, http.StatusBadRequest, err
	}
	defer row.Close()

	isNotNull := row.Next()
	if !isNotNull {
		return employee, http.StatusNotFound, errors.New("no employee with given id")
	}
	row.Scan(&employee.ID, &employee.Firstname, &employee.Lastname, &employee.Gender, &employee.Birthdate, &employee.Address, &employee.Password)

	return employee, http.StatusOK, nil
}
