package repository

import (
	"context"
	"database/sql"

	"fmt"
	"strings"

	"github.com/doffy007/golang-nicepay/config"
	"github.com/doffy007/golang-nicepay/database/query"
	"github.com/doffy007/golang-nicepay/internal/entity"
	"github.com/jmoiron/sqlx"
)

type NicepayRepository interface {
	Register(entity.Nicepay) error
	Status(entity.Status, []string) (*entity.Nicepay, error)
}

type nicepayRepository struct {
	ctx    context.Context
	config *config.Config
	db     *sqlx.DB
}

func (r repository) NicepayRepository() NicepayRepository {
	return &nicepayRepository{
		ctx:    r.ctx,
		config: r.config,
		db:     r.db,
	}
}

// Register NicepayRepository.
func (r nicepayRepository) Register(payload entity.Nicepay) error {
	_, err := r.db.NamedExecContext(r.ctx, query.Register, payload)

	if err != nil {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Repository Register: %v\n", err.Error())
		return err
	}

	return nil
}

// Status implements NicepayRepository.
func (r nicepayRepository) Status(payload entity.Status, fields []string) (*entity.Nicepay, error) {
	var columns string
	if len(fields) == 0 {
		columns = "*"
	} else {
		columns = strings.Join(fields, ", ")
	}

	dest := entity.Nicepay{}

	var filterValues []interface{}
	var filterKeys []string

	if payload.Timestamp != "" {
		filterKeys = append(filterKeys, "timeStamp = ?")
		filterValues = append(filterValues, payload.Timestamp)
	}

	if payload.TransactionId != "" {
		filterKeys = append(filterKeys, "tXid = ?")
		filterValues = append(filterValues, payload.TransactionId)
	}

	if payload.MerchantId != "" {
		filterKeys = append(filterKeys, "iMid = ?")
		filterValues = append(filterValues, payload.MerchantId)
	}

	if payload.Amount != "" {
		filterKeys = append(filterKeys, "amt = ?")
		filterValues = append(filterValues, payload.Amount)
	}

	query := fmt.Sprintf(query.FindStatus, columns, strings.Join(filterKeys, " AND "))
	err := r.db.GetContext(r.ctx, &dest, query, filterValues...)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Printf("\033[1;31m [ERROR] \033[0m Repository FindOneUser: %v\n", err.Error())
		return nil, err
	}

	return &dest, nil
}
