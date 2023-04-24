package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

type Brands_Model struct {
	brand_id   int
	brand_name string
}

func main() {
	var err error
	db, err = sql.Open("sqlserver", "sqlserver://sa:Keng1234@localhost?database=TestDB")
	if err != nil {
		panic(err)
	}
	covers, err := GetBrands()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, cover := range covers {
		fmt.Println(cover)
	}
}

func GetBrands() ([]Brands_Model, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select * from production.brands"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	brands := []Brands_Model{}
	for rows.Next() {
		brand := Brands_Model{}
		err = rows.Scan(&brand.brand_id, &brand.brand_name)
		if err != nil {
			return nil, err
		}
		brands = append(brands, brand)
	}

	defer rows.Close()
	return brands, nil
}
