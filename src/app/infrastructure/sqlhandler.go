package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type sqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *sqlHandler {
	conn, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/sample")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(sqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
