package main

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

// var db *sql.DB
var db *sqlx.DB

type Brands_Model struct {
	Brand_Id   int    `db:"brand_id"`
	Brand_Name string `db:"brand_name"`
}

func main() {
	var err error
	db, err = sqlx.Open("sqlserver", "sqlserver://sa:Keng1234@localhost?database=TestDB")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	brands, err := GetBrandsX()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(brands)
}

func GetBrandsX() ([]Brands_Model, error) {
	query := "select brand_id from production.brands"
	brand := []Brands_Model{}
	err := db.Select(&brand, query)
	if err != nil {
		return nil, err
	}
	return brand, nil
}
