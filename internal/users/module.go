package users

import (
	"github.com/alfredoprograma/mks/internal/queries"
)

type Module struct {
	Repo       Repo
	Service    Service
	Controller Controller
}

func NewModule(querier queries.Querier) *Module {
	usersRepo := NewRepo(querier)
	usersService := NewService(usersRepo)
	usersController := NewController(usersService)

	return &Module{Repo: usersRepo, Service: usersService, Controller: usersController}
}
