package db_transaction

import (
	"context"
)

type DBTransaction interface {
	WithTransaction(ctx context.Context, fn func(c context.Context) error) error
}
