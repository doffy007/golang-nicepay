package repository

import (
	"context"

	"github.com/doffy007/golang-nicepay/config"
	"github.com/doffy007/golang-nicepay/database"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	ctx    context.Context
	config *config.Config
	db     *sqlx.DB
}

func NewRepository(ctx context.Context, conf *config.Config) Repository {
	return repository{
		ctx:    ctx,
		config: conf,
		db:     database.Mysql(),
	}
}
