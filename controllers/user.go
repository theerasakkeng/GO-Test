package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/theerasakkeng/GO-CRUD/models"
)

var db *sqlx.DB

func GetCustomers(c *gin.Context) {
	var err error
	db, err = sqlx.Open("sqlserver", "sqlserver://sa:Keng1234@localhost?database=TestDB")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "select brand_id, brand_name from production.brands"
	brand := []models.Brands_Model{}
	err = db.Select(&brand, query)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": brand})
}
