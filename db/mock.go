package db

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type DBMockStore struct {
	mock.Mock
}

func (m *DBMockStore) RegisterFarmer(ctx context.Context, farmer Farmer) (err error) {
	args := m.Called(ctx, farmer)
	return args.Error(0)
}
