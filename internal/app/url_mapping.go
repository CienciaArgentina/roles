package app

import "github.com/CienciaArgentina/go-backend-commons/pkg/middleware"

// MapURLs Maps URLs to their corresponding handler
func (a *App) MapURLs() {
	a.Router.GET("/ping")

	roles := a.Router.Group("/roles")
	{
		roles.GET("", middleware.AdaptController(a.RoleController.GetAll))
		roles.GET("/:id", middleware.AdaptController(a.RoleController.Get))

		// assign := roles.Group("/assign")
		// {
		// 	// TODO: Implement
		// 	assign.GET("/:auth_id")
		// 	// TODO: Implement
		// 	assign.POST("")
		// 	// TODO: Implement
		// 	assign.PATCH("")
		// 	// TODO: Implement
		// 	assign.DELETE("/:auth_id")
		// 	// TODO: Implement
		// }
	}
}
