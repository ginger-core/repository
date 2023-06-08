package repository

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
)

type Repository interface {
	GetDB(query query.Query) any

	Ping(ctx context.Context) errors.Error

	Create(query query.Query, entity any) errors.Error

	Count(query query.Query) (uint64, errors.Error)
	List(query query.Query) (any, errors.Error)
	Get(query query.Query) (any, errors.Error)

	Update(query query.Query, update any) errors.Error
	Upsert(query query.Query, entity any) errors.Error

	Delete(query query.Query) errors.Error
}
