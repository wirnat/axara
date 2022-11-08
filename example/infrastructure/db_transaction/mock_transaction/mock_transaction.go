package mock_transaction

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MockTransaction struct {
	mock.Mock
}

func (m MockTransaction) WithTransaction(ctx context.Context, fn func(c context.Context) error) (err error) {
	args := m.Called(ctx, fn)
	err, _ = args.Get(0).(error)
	return
}
