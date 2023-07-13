package handler

import (
	"context"

	"github.com/doffy007/golang-nicepay/config"
	"github.com/doffy007/golang-nicepay/internal/repository"
	"github.com/doffy007/golang-nicepay/internal/usecase"
)

type handler struct {
	ctx        context.Context
	config     *config.Config
	usecase    usecase.Usecase
	repository repository.Repository
}

func NewHandler(ctx context.Context, conf *config.Config) Handler {
	return handler{
		ctx:     ctx,
		config:  conf,
		usecase: usecase.NewUsecase(ctx, conf),
	}
}
