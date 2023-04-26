package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"github.com/theerasakkeng/GO-CRUD/controllers"
)

type Brands_Model struct {
	Brand_Id   int    `json:"brand_id"`
	Brand_Name string `json:"brand_name"`
}

func main() {

	r := gin.Default()

	customer := r.Group("/api/customer")
	{
		customer.GET("/getcustomers", controllers.GetCustomers)
	}

	r.Run()
}
