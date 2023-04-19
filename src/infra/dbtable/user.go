package dbtable

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func createUserTable() {
    dsn := "user=postgres password=postgres dbname=architecture_study host=localhost port=5432 sslmode=disable"
    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatalln(err)
    }
    defer db.Close()

    createTableQuery := `CREATE TABLE IF NOT EXISTS users (
			user_id       SERIAL PRIMARY KEY,
			name          VARCHAR(255) NOT NULL UNIQUE,
			email         VARCHAR(255) NOT NULL UNIQUE,
			password      VARCHAR(255) NOT NULL,
			profile       TEXT,
			created_at    TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at    TIMESTAMP NOT NULL DEFAULT NOW()
    );`

    _, err = db.Exec(createTableQuery)
    if err != nil {
        log.Fatalln(err)
    }
}
