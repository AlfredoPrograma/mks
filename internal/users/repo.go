package users

import (
	"context"

	"github.com/alfredoprograma/mks/internal/queries"
)

type Repo interface {
	CreateOne(ctx context.Context, args queries.CreateUserParams) (queries.CreateUserRow, error)
	GetByID(ctx context.Context, id int32) (queries.GetUserByIDRow, error)
}

type repo struct {
	querier queries.Querier
}

func (r *repo) CreateOne(ctx context.Context, args queries.CreateUserParams) (queries.CreateUserRow, error) {
	return r.querier.CreateUser(ctx, args)
}

func (r *repo) GetByID(ctx context.Context, id int32) (queries.GetUserByIDRow, error) {
	return r.querier.GetUserByID(ctx, id)
}

func NewRepo(querier queries.Querier) Repo {
	return &repo{querier}
}
