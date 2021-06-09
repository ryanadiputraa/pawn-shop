package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ryanadiputraa/pawn-shop/controller"
	"github.com/ryanadiputraa/pawn-shop/middlewares"
	"github.com/ryanadiputraa/pawn-shop/service"
)

var (
	employeeService service.EmployeeService = service.NewEmployeeService()
	employeeController controller.EmployeeController = controller.NewEmployeeController(employeeService)
	customerService service.CustomerService = service.NewCustomerService()
	customerContoller controller.CustomerController = controller.NewCustomerController(customerService)
	imageController controller.ImageController = controller.NewImageController()
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("no env found")
	} 
}

func main() {
	r := gin.New()

	r.Use(gin.Recovery(), gin.Logger(), middlewares.CORSMiddleware())

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
	api.PUT("/employees/:employee_id", func(c *gin.Context) {
		employeeController.Update(c)
	})
	api.DELETE("/employees/:employee_id", func(c *gin.Context) {
		employeeController.DeleteEmployee(c)
	})
	api.POST("/login", func(c *gin.Context) {
		employeeController.Login(c)
	})
	api.POST("/loginadmin", func(c *gin.Context) {
		employeeController.LoginAdmin(c)
	})
	api.POST("/logout", func(c *gin.Context) {
		employeeController.Logout(c)
	})

	api.GET("/customers" , func(c *gin.Context) {
		customerContoller.GetAllCustomer(c)
	})
	api.POST("/customers" , func(c *gin.Context) {
		customerContoller.CreateLoan(c)
	})
	api.PUT("/customers/:customer_id", func(c *gin.Context) {
		customerContoller.PayOffLoan(c)
	})
	api.GET("/customers/financial", func(c *gin.Context) {
		customerContoller.GetFinancialStatements(c)
	})

	r.GET("/:image_path", func(c *gin.Context) {
		imageController.ServeImage(c)
	})

	port := os.Getenv("PORT")
	if port == "" { port = "8000" }
	
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}