package pgsearch_api

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"

)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = ""
	DB_NAME     = "voucher_db"
)

func GetConncetion() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)
	//defer db.Close()

	return db
}

func closeConnection()  {

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}