package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"payment_service/config"
)

func Connection() (*sql.DB, error) {
	cnf := config.Load()
	conDb := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", cnf.PostgresHost, cnf.PostgresPort, cnf.PostgresUser, cnf.PostgresDatabase, cnf.PostgresPassword)
	return sql.Open("postgres", conDb)
}
