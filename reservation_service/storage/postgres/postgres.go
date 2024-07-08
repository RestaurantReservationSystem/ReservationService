package postgres

import (
	"database/sql"
	"fmt"
	"reservation_service/config"
)

func ConnectionDb() (*sql.DB, error) {
	cnf := config.Load()
	conDb := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", cnf.PostgresHost, cnf.PostgresPort, cnf.PostgresUser, cnf.PostgresDatabase, cnf.PostgresPassword)
	return sql.Open("postgres", conDb)

}
