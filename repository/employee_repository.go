package repository

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ryanadiputraa/pawn-shop/config"
	"github.com/ryanadiputraa/pawn-shop/entity"
)

type EmployeeRepository interface {
	GetAll(queryParam string) (employees []entity.Employee, code int, err error) 
	GetById(employeeId string) (employee entity.Employee, code int, err error) 
	Create(employee entity.Employee) (code int, err error)
	Update(employee entity.Employee) (code int, err error)
	CheckPassword(loginPayload entity.LoginEmployee) (employee entity.Employee, code int, err error)
	Delete(employeeId string) (code int, err error)
}

type employeeRepository struct {}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{}
}

func (r *employeeRepository) GetAll(queryParam string) (employees []entity.Employee, code int, err error) {
	db, err := config.OpenConnection()
	if err != nil {
		return employees, http.StatusBadGateway, err
	}
	defer db.Close()

	var query string
	if len(queryParam) > 0 {
		query = fmt.Sprintf("SELECT * FROM employees WHERE CAST(employee_id AS TEXT) LIKE '%v%%' OR LOWER(firstname) LIKE LOWER('%v%%') OR LOWER(lastname) LIKE LOWER('%v%%') OR LOWER(gender) LIKE LOWER('%v%%') OR CAST(birthdate AS TEXT) LIKE '%v%%' OR LOWER(address) LIKE LOWER('%v%%') OR LOWER(password) LIKE LOWER('%v%%')", queryParam, queryParam, queryParam, queryParam, queryParam, queryParam, queryParam)
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

func (r *employeeRepository) Create(employee entity.Employee) (code int, err error) {
	db, err := config.OpenConnection()
	if err != nil {
		return http.StatusBadGateway, err
	}
	defer db.Close()

	query := `INSERT INTO employees (firstname, lastname, gender, birthdate, address, password) VALUES ($1, $2, $3, $4, $5, $6)`
	
	_, err = db.Exec(query, employee.Firstname, employee.Lastname, employee.Gender, employee.Birthdate, employee.Address, employee.Password)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusCreated, nil
}

func (r *employeeRepository) Update(employee entity.Employee) (code int, err error) {
	db, err := config.OpenConnection()
	if err != nil {
		return http.StatusBadGateway, err
	}
	defer db.Close()

	query := fmt.Sprintf("UPDATE employees SET firstname = '%v', lastname = '%v', gender = '%v', birthdate = '%v', address = '%v', password = '%v' WHERE employee_id = %v", employee.Firstname, employee.Lastname, employee.Gender, employee.Birthdate, employee.Address, employee.Password, employee.ID)

	_, err = db.Exec(query)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

func (r *employeeRepository) CheckPassword(loginPayload entity.LoginEmployee) (employee entity.Employee, code int, err error) {
	db, err := config.OpenConnection()
	if err != nil {
		return employee, http.StatusBadGateway, err
	}
	defer db.Close()

	// check if employee with given id exist
	row, err := db.Query(fmt.Sprintf("SELECT employee_id, password FROM employees WHERE employee_id = %v", strconv.Itoa(loginPayload.ID)))
	if err != nil {
		return employee, http.StatusBadRequest, err
	}
	defer row.Close()

	for row.Next() {
		row.Scan(&employee.ID, &employee.Password)
		if employee.ID == 0 {
			return employee, http.StatusNotFound, err	
		}
	}

	// validate password
	if (employee.Password != loginPayload.Password) {
		return employee, http.StatusUnauthorized, errors.New("password didn't match")	
	}

	return employee, http.StatusAccepted, nil
}

func (r *employeeRepository) Delete(employeeId string) (code int, err error) {
	db, err := config.OpenConnection()
	if err != nil {
		return http.StatusBadGateway, err
	}
	defer db.Close()	

	_, err = db.Query(fmt.Sprintf("DELETE FROM employees WHERE employee_id = %v", employeeId))
	if err != nil {
		return http.StatusNotFound, errors.New("employee with given id didn't exist")
	}

	return http.StatusOK, nil
}
