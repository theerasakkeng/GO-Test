package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"github.com/theerasakkeng/GO-CRUD/controllers"
)

func main() {

	r := gin.Default()

	customer := r.Group("/api/customer")
	{
		customer.GET("/getcustomers", controllers.GetCustomers)
		customer.GET("/getcustomerdetail", controllers.GetCustomerDetails)
	}

	r.Run()
}
