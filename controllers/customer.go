package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theerasakkeng/GO-CRUD/db"
	"github.com/theerasakkeng/GO-CRUD/models"
)

func GetCustomers(c *gin.Context) {
	// id := c.Query("id")
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
