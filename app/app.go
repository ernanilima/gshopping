package app

import (
	"github.com/ernanilima/gshopping/app/config"
	"github.com/ernanilima/gshopping/app/controller"
	"github.com/ernanilima/gshopping/app/repository"
)

func Init(cfg *config.Config) controller.Controller {
	repo := repository.NewRepository(cfg)
	return controller.NewController(repo)
}
