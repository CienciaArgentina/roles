package app

import (
	"github.com/CienciaArgentina/go-backend-commons/pkg/middleware"
	"github.com/gin-gonic/gin"
)

const (
	bodyKey = middleware.ResponseBodyKey
)

// MapURLs Maps URLs to their corresponding handler
func (a *App) MapURLs() {
	a.Router.GET("/ping", middleware.AdaptController(func(c *gin.Context) error { c.Set(bodyKey, "pong"); return nil }))

	roles := a.Router.Group("/roles")
	{
		roles.GET("", middleware.AdaptController(a.RoleController.GetAll))
		roles.GET("/:id", middleware.AdaptController(a.RoleController.Get))

		assign := a.Router.Group("/assign")
		{
			assign.GET("/:auth_id", middleware.AdaptController(a.RoleController.GetAssignedRole))
			assign.POST("", middleware.AdaptController(a.RoleController.AssignRole))
			assign.DELETE("/:auth_id", middleware.AdaptController(a.RoleController.DeleteAssignedRole))
		}
	}
}
