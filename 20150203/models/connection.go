package models

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "os"
)

const (
    connection_format string = "user=%s dbname=%s sslmode=disable password=%s host=%s"
    db_name                  = "ciudad_gourmet"
)

func stablishConnection() (*sql.DB, error) {
    user := os.Getenv("POSTGRESQL_USER")
    pass := os.Getenv("POSTGRESQL_PASS")
    host := os.Getenv("PGHOST")
    connection_params := fmt.Sprintf(connection_format, user, db_name, pass, host)
    db, err := sql.Open("postgres", connection_params)

    if err != nil {
        println("error open")
        return nil, err
    }
    return db, nil
}
