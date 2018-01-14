package pgsearch_api

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"

)

const (
	DB_HOST		= "postgres-server"
	DB_USER     = "postgres"
	DB_NAME     = "voucher_db"
	DB_PORT		= 5432
)

func GetConnection() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_NAME)
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