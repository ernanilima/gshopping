package app

import (
	"github.com/ernanilima/gshopping/app/controller"
	"github.com/ernanilima/gshopping/app/repository"
	"github.com/ernanilima/gshopping/app/repository/database"
)

// Init inicializa os repositories e os controllers
func Init(connector database.DatabaseConnector) controller.Controller {
	repo := repository.NewRepository(connector)
	return controller.NewController(repo)
}
