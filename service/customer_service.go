package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ryanadiputraa/pawn-shop/config"
	"github.com/ryanadiputraa/pawn-shop/entity"
)

type CustomerService interface {
	GetAllCustomer(ctx *gin.Context) (int, interface{})
	CreateLoan(entity.Customer) (int, interface{})
	PayOffLoan(customerId string) (int, interface{})
}

type customerService struct {}

func NewCustomerService() CustomerService {
	return &customerService{}
}

func (service *customerService) GetAllCustomer(ctx *gin.Context) (int, interface{}) {
	// authenticate
	cookie, err := ctx.Cookie("jwt")
	if err != nil {
		response := entity.Error {
			Code: http.StatusUnauthorized,
			Error: "no cookie found",
		}
		return http.StatusUnauthorized, response
	}

	_, err = jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetSecretKey()), nil
	})
	if err != nil {
		response := entity.Error {
			Code: http.StatusUnauthorized,
			Error: "unauthorized",
		}
		return http.StatusUnauthorized, response
	}

	db, err := config.OpenConnection()
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadGateway,
			Error: "can't open db connection",
		}
		return http.StatusBadGateway, response
	}
	defer db.Close()

	var query string
	URLQueryParam := ctx.Request.URL.Query()
	if len(URLQueryParam) != 0 {
		query = fmt.Sprintf("SELECT customer_id, firstname, lastname, gender, contact, nominal, interest, item_name, status FROM customers INNER JOIN loans ON loan = loan_id INNER JOIN insurance_items ON insurance_item = item_id WHERE LOWER(firstname) LIKE LOWER('%v%%') OR LOWER(lastname) LIKE LOWER('%v%%')", URLQueryParam["name"][0], URLQueryParam["name"][0])
	} else {
		query = `SELECT customer_id, firstname, lastname, gender, contact, nominal, interest, item_name, status FROM customers INNER JOIN loans ON loan = loan_id INNER JOIN insurance_items ON insurance_item = item_id`
	}

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
	defer rows.Close()
	if customers == nil {
		response := entity.Error {
			Code: http.StatusNotFound,
			Error: "no customers with given name",
		}
		return http.StatusNotFound, response
	}

	return http.StatusOK, customers
}

func (service *customerService) CreateLoan(customer entity.Customer) (int, interface{}) {
	db, err := config.OpenConnection()
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadGateway,
			Error: "can't open db connection",
		}
		return http.StatusBadGateway, response
	}
	defer db.Close()

	loan_id, err := uuid.NewUUID()
	if err != nil {
		response := entity.Error {
			Code: http.StatusInternalServerError,
			Error: "can't generate uuid",
		}
		return http.StatusInternalServerError, response
	}
	item_id, err := uuid.NewUUID()
	if err != nil {
		response := entity.Error {
			Code: http.StatusInternalServerError,
			Error: "can't generate uuid",
		}
		return http.StatusInternalServerError, response
	}
	customer_id, err := uuid.NewUUID()
	if err != nil {
		response := entity.Error {
			Code: http.StatusInternalServerError,
			Error: "can't generate uuid",
		}
		return http.StatusInternalServerError, response
	}

	query := `INSERT INTO insurance_items (item_id, item_name, status) VALUES ($1, $2, $3)`
	_, err = db.Exec(query, item_id, customer.InsuranceItem, "jaminan")
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadRequest,
			// Error: "fail to insert insurance item",
			Error: err.Error(),
		}
		return http.StatusBadRequest, response	
	}

	query = `INSERT INTO loans (loan_id, nominal, interest) VALUES ($1, $2, $3)`
	_, err = db.Exec(query, loan_id, strconv.Itoa(customer.Loan), strconv.Itoa(customer.Interest))
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadRequest,
			Error: "fail to insert loan",
		}
		return http.StatusBadRequest, response	
	}

	query = `INSERT INTO customers (customer_id, firstname, lastname, gender, loan, insurance_item, contact) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.Exec(query, customer_id, customer.Firstname, customer.Lastname, customer.Gender, loan_id, item_id, customer.Contact)
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadRequest,
			Error: "fail to insert customer",
		}
		return http.StatusBadRequest, response	
	}	

	response := entity.HTTPCode { Code: http.StatusOK }

	return http.StatusOK, response
}

func (service *customerService) PayOffLoan(customerId string) (int, interface{}) {
	db, err := config.OpenConnection()
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadGateway,
			Error: "can't open db connection",
		}
		return http.StatusBadGateway, response
	}
	defer db.Close()

	row, err := db.Query(fmt.Sprintf("SELECT insurance_item FROM customers WHERE customer_id = '%v'", customerId))
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadRequest,
			Error: "can't get customer data",
		}
		return http.StatusBadRequest, response
	}
	defer row.Close()

	var customer entity.Customer
	isNotNull := row.Next()
	if !isNotNull {
		response := entity.Error {
			Code: http.StatusNotFound,
			Error: "no customer with given id",
		}
		return http.StatusNotFound, response	
	}

	row.Scan(&customer.InsuranceItem)

	query := fmt.Sprintf("UPDATE insurance_items SET status = 'ditebus' WHERE item_id = '%v'", customer.InsuranceItem)
	_, err = db.Exec(query)
	if err != nil {
		response := entity.Error {
			Code: http.StatusBadRequest,
			Error: "fail to secure paymant",
		}
		return http.StatusBadRequest, response
	}
	
	response := entity.HTTPCode { Code: http.StatusOK }

	return http.StatusOK, response
}