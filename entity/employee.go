package entity

type Employee struct {
	ID 			int 	`json:"id"`
	Firstname 	string 	`json:"firstname"`
	Lastname 	string 	`json:"lastname"`
	Gender 		string 	`json:"gender"`
	Birthdate 	string 	`json:"birthdate"`
	Address 	string 	`json:"address"`
	Password 	string 	`json:"password"`
}

type EmployeesResponse struct {
	Code	int			`json:"code"`
	Data	[]Employee	`json:"data"`
}