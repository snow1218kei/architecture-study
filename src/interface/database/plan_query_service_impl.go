package database

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/planusecase"
)

type PlanQueryServiceImpl struct {
	conn *sqlx.DB
}

func NewPlanQueryServiceImpl(conn *sqlx.DB) *PlanQueryServiceImpl {
	return &PlanQueryServiceImpl{
		conn,
	}
}

func (queryService *PlanQueryServiceImpl) GetAll(ctx context.Context) ([]*planusecase.GetPlanDTO, error) {
	// SQLは省略
	return nil, nil
}
