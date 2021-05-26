package entity

type Employee struct {
	ID 			int 	`json:"id"`
	Firstname 	string 	`json:"firstname"`
	Lastname 	string 	`json:"lastname"`
	Gender 		string 	`json:"gender"`
	Birthdate 	string 	`json:"birthdate"`
	Address 	string 	`json:"address"`
	Password 	[]byte 	`json:"password"`
}