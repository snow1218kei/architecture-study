package user

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	"github.com/yuuki-tsujimura/architecture-study/src/infra/datamodel"
)

type RdbUserRepository struct {
	users []*user.User
}

func NewRdbUserRepository() *RdbUserRepository {
	return &RdbUserRepository{
		users: make([]*user.User, 0),
	}
}

func (repo *RdbUserRepository) Store(u *user.User) error {
	repo.users = append(repo.users, u)
	return nil
}

func (repo *RdbUserRepository) FindByName(name string) (*user.User, error) {
	db, err := sqlx.Open("postgresql", "root/sample")
	rows, err := db.Queryx("SELECT * FROM users WHERE name = $1 LIMIT 1", name)
	if err != nil {
			log.Fatal(err)
	}

	defer rows.Close()

	results := make([]datamodel.User, 0)
	for rows.Next() {

			var user datamodel.User

			err := rows.StructScan(&user)

			if err != nil {
					log.Fatal(err)
			}

			results = append(results, user)
	}

	fmt.Println(results)
	return nil, fmt.Errorf("user not found")
}
