package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	"github.com/yuuki-tsujimura/architecture-study/src/infra/datamodel"
)

type RdbUserRepositoryImpl struct {
	conn   *sqlx.DB
}

func NewRdbUserRepository(conn *sqlx.DB) *RdbUserRepositoryImpl {
	return &RdbUserRepositoryImpl{
		conn: conn,
	}
}

func (repo *RdbUserRepositoryImpl) Store(ctx context.Context, u *user.User) error {
	userData := user.ConvertUserToUserData(u)

	tx, err := repo.conn.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		}
	}()

	query := `
		INSERT INTO users (user_id, name, email, password, profile, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err = tx.ExecContext(ctx, query, userData.UserID, userData.Name, userData.Email, userData.Password, userData.Profile, userData.CreatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, career := range userData.Careers {
		query = `
			INSERT INTO careers (career_id, user_id, detail, start_year, end_year, created_at)
			VALUES (?, ?, ?, ?, ?, ?)
		`

		_, err = tx.ExecContext(ctx, query, career.CareerID, userData.UserID, career.Detail, career.StartYear, career.EndYear, career.CreatedAt)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, skill := range userData.Skills {
		query = `
			INSERT INTO skills (skill_id, user_id, tag_id, evaluation, years, created_at)
			VALUES (?, ?, ?, ?, ?, ?)
		`

		_, err = tx.ExecContext(ctx, query, skill.SkillID, userData.UserID, skill.TagID, skill.Evaluation, skill.Years, skill.CreatedAt)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}


func (repo *RdbUserRepositoryImpl) FindByName(ctx context.Context, name string) (*user.User, error) {
	var dbUser datamodel.User
	err := repo.conn.Get(&dbUser, "SELECT * FROM users WHERE name = $1 LIMIT 1", name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Fatal(err)
	}

	var dbCareers []*datamodel.Career
	err = repo.conn.Select(&dbCareers, "SELECT * FROM careers WHERE user_id = $1", dbUser.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Fatal(err)
	}

	var dbSkills []*datamodel.Skill
	err = repo.conn.Select(&dbSkills, "SELECT * FROM skills WHERE user_id = $1", dbUser.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Fatal(err)
	}

	careersData := make([]*user.CareerData, len(dbCareers))
	for i, dbCareer := range dbCareers {
		careersData[i] = &user.CareerData{
			CareerID:  dbCareer.ID,
			Detail:    dbCareer.Detail,
			StartYear: dbCareer.StartYear,
			EndYear:   dbCareer.EndYear,
			CreatedAt: dbCareer.CreatedAt,
		}
	}

	skillsData := make([]*user.SkillData, len(dbSkills))
	for i, dbSkill := range dbSkills {
		skillsData[i] = &user.SkillData{
			SkillID:    dbSkill.ID,
			TagID:      dbSkill.TagID,
			Evaluation: dbSkill.Evaluation,
			Years:      dbSkill.Years,
			CreatedAt:  dbSkill.CreatedAt,
		}
	}

	userData := user.UserData{
		UserID:    dbUser.ID,
		Name:      dbUser.Name,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		Profile:   dbUser.Profile,
		Careers:   careersData,
		Skills:    skillsData,
		CreatedAt: dbUser.CreatedAt,
	}

	user, err := user.ReconstructUserFromData(userData)
	if err != nil {
		return nil, err
	}

	return user, nil
}
