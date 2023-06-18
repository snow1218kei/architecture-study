package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/yuuki-tsujimura/architecture-study/src/config"
	"golang.org/x/xerrors"
)

var (
	once   sync.Once
	driver = "postgres"
)

func NewConnection() (*sqlx.DB, error) {
	var conn *sqlx.DB
	var err error
	once.Do(func() {
		connInfo := fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			config.Env.DBUser, config.Env.DBPassword, config.Env.DBHost, config.Env.DBPort, config.Env.DBName)

		conn, err = sqlx.Open(driver, connInfo)
		if err != nil {
			return
		}

		if e := conn.Ping(); e != nil {
			err = e
			return
		}

		// DBへの全体のコネクション総数
		conn.SetMaxOpenConns(config.Env.RDBMaxConn)
		// DB接続を待機させておくコネクション総数
		conn.SetMaxIdleConns(config.Env.RDBMaxIdle)
	})

	return conn, err
}

func ExecFromCtx(ctx context.Context) (RDBHandler, error) {
	val := ctx.Value(config.DBKey)
	if val == nil {
		return nil, xerrors.New("fail to get connection from context")
	}

	conn, ok := val.(RDBHandler)
	if !ok {
		return nil, xerrors.New("can't get context executor")
	}
	return conn, nil
}
