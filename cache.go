package repository

import (
	"context"
	"time"

	"github.com/ginger-core/errors"
)

type Cache interface {
	GetDB() any

	Store(ctx context.Context, key string, value any,
		expiration ...time.Duration) errors.Error
	MarshalStore(ctx context.Context, key string, value any,
		expiration ...time.Duration) errors.Error
	Expire(ctx context.Context, key string, expiration time.Duration) errors.Error
	Lock(ctx context.Context, key string, expiration time.Duration) errors.Error

	InsertItem(ctx context.Context, sourceKey string, value Entity) errors.Error
	MarshalInsertItem(ctx context.Context, sourceKey string, value Entity) errors.Error

	ListKeys(ctx context.Context, pattern string) ([]string, errors.Error)

	Fetch(ctx context.Context, key string) ([]byte, errors.Error)
	Load(ctx context.Context, key string, resultRef any) errors.Error
	Count(ctx context.Context, pattern string) (int, errors.Error)
	CountItems(ctx context.Context, sourceKey string) (int, errors.Error)
	LoadItems(ctx context.Context, sourceKey string,
		page, pageSize int) ([]string, errors.Error)
	LoadItem(ctx context.Context, sourceKey string,
		id int64, resultRef any) errors.Error

	GetExpiration(ctx context.Context, key string) (time.Duration, errors.Error)

	Rename(ctx context.Context, key, new string) errors.Error

	Delete(ctx context.Context, key string) errors.Error
	DeleteItem(ctx context.Context, sourceKey string, id int64) errors.Error

	SetItem(ctx context.Context, key, id string, value any) errors.Error
	MarshalSetItem(ctx context.Context, key, id string, value any) errors.Error
	GetItem(ctx context.Context, key, id string, resultRef any) errors.Error
	UnsetItem(ctx context.Context, sourceKey, id string) errors.Error
}
