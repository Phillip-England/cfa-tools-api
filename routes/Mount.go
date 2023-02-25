package routes

import (
	"github.com/phillip-england/go-http/model"
)

func Mount(db model.Db) {
	UserRoutes(db)
	LocationRoutes(db)
	CaresRoutes(db)
}
