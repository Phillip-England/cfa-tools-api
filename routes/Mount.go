package routes

func Mount(db DB) {
	UserRoutes(ctx)
	LocationRoutes(ctx)
	CaresRoutes(ctx)
}
