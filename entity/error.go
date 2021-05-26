package entity

type Error struct {
	Code	int 	`json:"code"`
	Error string	`json:"message"`
}