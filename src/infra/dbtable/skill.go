package dbtable

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func createSkillTable() {
    dsn := "user=postgres password=postgres dbname=architecture_study host=localhost port=5432 sslmode=disable"
    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatalln(err)
    }
    defer db.Close()

    createTableQuery := `CREATE TABLE IF NOT EXISTS skills (
			skill_id     SERIAL PRIMARY KEY,
			user_id      INTEGER NOT NULL REFERENCES users(user_id),
			tag_id       INTEGER NOT NULL REFERENCES tags(tag_id),
			evaluation   SMALLINT NOT NULL,
			years        SMALLINT NOT NULL,
			created_at   TIMESTAMP NOT NULL DEFAULT NOW()
			updated_at	 TIMESTAMP NOT NULL DEFAULT NOW()
		);`

    _, err = db.Exec(createTableQuery)
    if err != nil {
        log.Fatalln(err)
    }
}
