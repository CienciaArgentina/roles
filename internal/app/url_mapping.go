package app

// MapURLs Maps URLs to their corresponding handler
func (a *App) MapURLs() {
	a.Router.GET("/ping")

	roles := a.Router.Group("/roles")
	{
		// TODO: Implement controller
		roles.GET("")

		assign := roles.Group("/assign")
		{
			// TODO: Implement
			assign.GET("/:auth_id")
			// TODO: Implement
			assign.POST("")
			// TODO: Implement
			assign.PATCH("")
			// TODO: Implement
			assign.DELETE("/:auth_id")
			// TODO: Implement
		}
	}
}
