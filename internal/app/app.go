package app

import (
	"github.com/CienciaArgentina/go-backend-commons/pkg/middleware"
	"github.com/CienciaArgentina/roles/internal/role"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// App App definition
type App struct {
	Router *gin.Engine

	// App specific handlers
	RoleController role.Controller
}

// New Returns new app
func New(
	roleController role.Controller,
) *App {
	return &App{
		RoleController: roleController,
	}
}

// Start Initializes app
func (a *App) Start() {
	a.Router = gin.Default()
	a.Router.Use(gzip.Gzip(gzip.DefaultCompression))
	a.Router.Use(middleware.ResponseMiddleware)
	a.Router.Use(middleware.ErrorMiddleware)

	a.MapURLs()
	a.Router.Run()
}
