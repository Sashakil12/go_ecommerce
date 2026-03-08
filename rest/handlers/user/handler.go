package user

import (
	"ecommerce/config"
	"ecommerce/repo"
)

type Handler struct {
	userRepo      repo.UserRepo
	configuration *config.Config
}

func NewHandler(r repo.UserRepo, c *config.Config) *Handler {
	return &Handler{
		userRepo:      r,
		configuration: c,
	}
}
