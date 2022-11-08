package gorm_transaction

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type transaction struct {
	DB *gorm.DB
}

func NewTransaction(DB *gorm.DB) *transaction {
	return &transaction{DB: DB}
}

func (t transaction) WithTransaction(ctx context.Context, fn func(c context.Context) error) error {
	tx := t.DB.Begin()
	ctx = context.WithValue(ctx, "tx", tx)
	err := fn(ctx)
	if err != nil {
		tx.Rollback()
		ctx = context.WithValue(ctx, "tx", nil)
		return err
	} else {
		tx.Commit()
		ctx = context.WithValue(ctx, "tx", nil)
	}

	return nil
}

func GetTx(ctx context.Context) (*gorm.DB, error) {
	db, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		return nil, fmt.Errorf("without transaction")
	}

	return db, nil
}
