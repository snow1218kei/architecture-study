package database

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/requirementusecase"
)

type RequirementQueryServiceImpl struct {
	conn *sqlx.DB
}

func NewRequirementQueryServiceImpl(conn *sqlx.DB) *RequirementQueryServiceImpl {
	return &RequirementQueryServiceImpl{
		conn,
	}
}

func (queryService *RequirementQueryServiceImpl) GetAll(ctx context.Context) ([]*requirementusecase.GetMentorRequirementDTO, error) {
	// SQLは省略
	return nil, nil
}
