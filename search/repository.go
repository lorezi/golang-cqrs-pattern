package search

import (
	"context"

	"github.com/lorezi/golang-cqrs-pattern/schema"
)

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	SearchMeows(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Meow, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func InsertMeow(ctx context.Context, meow schema.Meow) error {
	return impl.InsertMeow(ctx, meow)
}

func SearchMeows(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Meow, error) {
	return impl.SearchMeows(ctx, query, skip, take)
}
