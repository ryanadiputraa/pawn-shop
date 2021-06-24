package entity

type FinancialStatements struct {
	TotalLoans		int					`json:"totalLoans"`
	LoanStatus		[]LoanStatements	`json:"loanStatus"`
}

type LoanStatements struct {
	Status	string 	`json:"status"`
	Total	int		`json:"total"`
}