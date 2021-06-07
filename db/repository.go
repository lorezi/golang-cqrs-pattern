package db

import (
	"context"

	"github.com/lorezi/golang-cqrs-pattern/schema"
)

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error)
}

var repoImpl Repository

func SetRepository(r Repository) {
	repoImpl = r
}

func Close() {
	repoImpl.Close()
}

func InsertMeow(ctx context.Context, meow schema.Meow) error {
	return repoImpl.InsertMeow(ctx, meow)
}

func ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error) {
	return repoImpl.ListMeows(ctx, skip, take)
}
