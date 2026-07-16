package users

import (
	"context"

	"github.com/alfredoprograma/mks/internal/queries"
)

type Service interface {
	CreateOne(ctx context.Context, args queries.CreateUserParams) (queries.CreateUserRow, error)
	GetByID(ctx context.Context, id int32) (queries.GetUserByIDRow, error)
}

type service struct {
	repo Repo
}

func (s *service) CreateOne(ctx context.Context, args queries.CreateUserParams) (queries.CreateUserRow, error) {
	return s.repo.CreateOne(ctx, args)
}

func (s *service) GetByID(ctx context.Context, id int32) (queries.GetUserByIDRow, error) {
	return s.repo.GetByID(ctx, id)
}

func NewService(repo Repo) Service {
	return &service{repo}
}
