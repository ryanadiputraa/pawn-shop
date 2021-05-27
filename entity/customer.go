package entity

type Customer struct {
	CustomerId		string	`json:"customerId"`
	Firstname		string	`json:"firstname"`
	Lastname		string 	`json:"lastname"`
	Gender			string	`json:"gender"`
	Contact			string 	`json:"contact"`
	Loan			int		`json:"loan"`
	Interest		int		`json:"interest"`
	InsuranceItem	string	`json:"insuranceItem"`
	ItemStatus		string	`json:"itemStatus"`
}	