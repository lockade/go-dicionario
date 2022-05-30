package banco

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dicionario")
	if err != nil {
		log.Print(err.Error())
	}
	return db
}

func CloseConnection(db *sql.DB) {
	db.Close()
}
