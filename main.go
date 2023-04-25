package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/theerasakkeng/GO-CRUD/controllers"
)

var db *sqlx.DB

type Brands_Model struct {
	Brand_Id   int    `json:"brand_id"`
	Brand_Name string `json:"brand_name"`
}

func main() {
	r := gin.Default()

	customer := r.Group("/api/customers")
	{
		customer.GET("getcustomer", controllers.GetCustomers)
	}

	r.Run()

	// brands, err := GetBrandsX()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// // fmt.Printf("%#v\n%#v", brands)

	// r := gin.New()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"status": "success", "data": brands})
	// })
	// r.Run()
}

// func GetBrandsX() ([]Brands_Model, error) {
// 	query := "select brand_id, brand_name from production.brands"
// 	brand := []Brands_Model{}
// 	err := db.Select(&brand, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return brand, nil
// }
