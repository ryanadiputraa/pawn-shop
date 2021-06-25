package repository

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/ryanadiputraa/pawn-shop/config"
	"github.com/ryanadiputraa/pawn-shop/entity"
)

type CustomerRepository interface {
	GetAll(queryParam string) (customers []entity.Customer, code int, err error)
	CreateLoan(loanId uuid.UUID, itemId uuid.UUID, customer entity.Customer) (code int, err error)
	PayOffLoan(customerId uuid.UUID) (code int, err error)
	GetFinancialStatements() (financialStatements entity.FinancialStatements, code int, err error)
}

type customerRepository struct {}

func NewCustomerRepository() CustomerRepository {
	return &customerRepository{}
}

func (r *customerRepository) GetAll(queryParam string) (customers []entity.Customer, code int, err error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, http.StatusBadGateway, err
	}
	defer db.Close()


	var query string
	if len(queryParam) > 0 {
		query = fmt.Sprintf("SELECT customer_id, firstname, lastname, gender, contact, nominal, interest, item_name, status, image FROM customers INNER JOIN loans ON loan = loan_id INNER JOIN insurance_items ON insurance_item = item_id WHERE LOWER(firstname) LIKE LOWER('%v%%') OR LOWER(lastname) LIKE LOWER('%v%%') OR LOWER(gender) LIKE LOWER('%v%%') OR contact LIKE '%v%%' OR CAST(nominal as TEXT) LIKE '%v%%' OR CAST((nominal + interest) AS TEXT) LIKE '%v%%' OR LOWER(item_name) LIKE LOWER('%v%%') OR LOWER(status) LIKE LOWER('%v%%')", queryParam, queryParam, queryParam, queryParam, queryParam, queryParam, queryParam, queryParam)
	} else {
		query = `SELECT customer_id, firstname, lastname, gender, contact, nominal, interest, item_name, status, image FROM customers INNER JOIN loans ON loan = loan_id INNER JOIN insurance_items ON insurance_item = item_id`
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer entity.Customer
		rows.Scan(&customer.ID, &customer.Firstname, &customer.Lastname, &customer.Gender, &customer.Contact, &customer.Loan, &customer.Interest, &customer.InsuranceItem, &customer.ItemStatus, &customer.Image)
		customers = append(customers, customer)
	}

	return customers, http.StatusOK, nil
}

func (r *customerRepository) CreateLoan(loanId uuid.UUID, itemId uuid.UUID, customer entity.Customer) (code int, err error) {
	db, err := config.OpenConnection()
	if err != nil {
		return http.StatusBadGateway, err
	}
	defer db.Close()

	query := `INSERT INTO loans (loan_id, nominal, interest) VALUES ($1, $2, $3)`
	_, err = db.Exec(query, loanId, strconv.Itoa(customer.Loan), strconv.Itoa(customer.Interest))
	if err != nil {
		return http.StatusBadRequest, err	
	}

	query = `INSERT INTO insurance_items (item_id, item_name, image, status) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(query, itemId, customer.InsuranceItem, customer.Image, "jaminan")
	if err != nil {
		return http.StatusBadRequest, err
	}

	query = `INSERT INTO customers (customer_id, firstname, lastname, gender, loan, insurance_item, contact) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.Exec(query, customer.ID, customer.Firstname, customer.Lastname, customer.Gender, loanId, itemId, customer.Contact)
	if err != nil {
		return http.StatusBadRequest, err	
	}

	return http.StatusCreated, nil
}

func (r *customerRepository) PayOffLoan(customerId uuid.UUID) (code int, err error) {
	db, err := config.OpenConnection()
	if err != nil {
		return http.StatusBadGateway, err
	}
	defer db.Close()

	row, err := db.Query(fmt.Sprintf("SELECT insurance_item FROM customers WHERE customer_id = '%v'", customerId))
	if err != nil {
		return http.StatusBadRequest, err
	}
	defer row.Close()

	var customer entity.Customer
	isNotNull := row.Next()
	if !isNotNull {
		return http.StatusNotFound, errors.New("no customer with given id")	
	}

	row.Scan(&customer.InsuranceItem)

	query := fmt.Sprintf("UPDATE insurance_items SET status = 'ditebus' WHERE item_id = '%v'", customer.InsuranceItem)
	_, err = db.Exec(query)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

func (r *customerRepository) GetFinancialStatements() (financialStatements entity.FinancialStatements, code int, err error) {
	db, err := config.OpenConnection()
	if err != nil {
		return financialStatements, http.StatusBadGateway, err
	}
	defer db.Close()

	query := "SELECT SUM (nominal) FROM loans"
	rows, err := db.Query(query)
	if err != nil {
		return financialStatements, http.StatusBadRequest, err
	}

	for rows.Next() {
		rows.Scan(&financialStatements.TotalLoans)
	}

	query = "SELECT status, SUM (nominal+interest) AS total FROM customers INNER JOIN loans ON loan = loan_id INNER JOIN insurance_items ON insurance_item = item_id GROUP BY status"

	rows, err = db.Query(query)
	if err != nil {
		return financialStatements, http.StatusBadRequest, err
	}
	defer rows.Close()

	for rows.Next() {
		var loanStatus entity.LoanStatements
		rows.Scan(&loanStatus.Status, &loanStatus.Total)
		financialStatements.LoanStatus = append(financialStatements.LoanStatus, loanStatus)
	}

	return financialStatements, http.StatusOK, nil
}
