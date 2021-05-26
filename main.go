package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ryanadiputraa/pawn-shop/controller"
	"github.com/ryanadiputraa/pawn-shop/service"
)

var (
	employeeService service.EmployeeService = service.New()
	employeeController controller.EmployeeController = controller.New(employeeService)
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("no env found")
	} 
}

func main() {
	r := gin.Default()
	api := r.Group("/api")

	api.GET("/employees", func(c *gin.Context) {
		employeeController.GetAllEmployees(c)
	})
	api.GET("/employees/:employee_id", func(c *gin.Context) {
		employeeController.GetEmployeeById(c)
	})
	api.POST("/employees", func(c *gin.Context) {
		employeeController.Register(c)
	})
	api.DELETE("/employees/:employee_id", func(c *gin.Context) {
		employeeController.DeleteEmployee(c)
	})

	port := os.Getenv("PORT")
	if port == "" { port = "8000" }
	
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}