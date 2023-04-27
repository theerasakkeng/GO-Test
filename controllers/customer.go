package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/theerasakkeng/GO-CRUD/db"
	"github.com/theerasakkeng/GO-CRUD/models"
)

func GetCustomers(c *gin.Context) {
	db, err := db.UseDB()
	query := "SELECT * FROM sales.customers"
	customers := []models.Customer_Model{}
	err = db.Select(&customers, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": customers})
	}
	defer db.Close()
}

func GetCustomerDetails(c *gin.Context) {
	id := c.Query("id")
	fmt.Println(reflect.TypeOf(id))
	db, err := db.UseDB()
	if err != nil {
		panic(err)
	}
	query := "SELECT customer_id, first_name FROM sales.customers WHERE first_name=@id"
	customer := models.Customer_Model{}
	row := db.QueryRow(query, sql.Named("id", id))
	err = row.Scan(&customer.Customer_Id, &customer.First_Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": customer})
	}
	defer db.Close()
}
