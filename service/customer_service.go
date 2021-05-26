package service

import (
	"net/http"

	"github.com/ryanadiputraa/pawn-shop/config"
	"github.com/ryanadiputraa/pawn-shop/entity"
)

type CustomerService interface {
	GetAllCustomer() (int, interface{})
}

type customerService struct {}

func NewCustomerService() CustomerService {
	return &customerService{}
}

func (service *customerService) GetAllCustomer() (int, interface{}) {
	db, err := config.OpenConnection()
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadGateway,
			Error: "can't open db connection",
		}
		return http.StatusBadGateway, response
	}
	defer db.Close()

	query := `SELECT customer_id, firstname, lastname, gender, contact nominal, interest, item_name, status FROM customers INNER JOIN loans ON loan = loan_id INNER JOIN insurance_items ON insurance_item = item_id`

	rows, err := db.Query(query)
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadRequest,
			Error: "can't get customers data",
		}
		return http.StatusBadRequest, response
	}
	defer rows.Close()

	var customers []entity.Customer
	for rows.Next() {
		var customer entity.Customer
		rows.Scan(&customer.CustomerId, &customer.Firstname, &customer.Lastname, &customer.Gender, &customer.Contact, &customer.Loan, &customer.Interest, &customer.InsuranceItem, &customer.ItemStatus)
		customers = append(customers, customer)
	}

	return http.StatusOK, customers
}