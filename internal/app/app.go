package app

import (
	"github.com/CienciaArgentina/roles/internal/role"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
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
	a.MapURLs()

	// a.Router.Use(middleware.ResponseMiddleware)
	// a.Router.Use(middleware.ErrorMiddleware)
	a.Router.Use(middle)

	a.Router.Run()
}

func middle(c *gin.Context) {
	log.Info("asjdasodujoiasdi")
	panic("ashusausaduasduiuiasduiasduhisd")
}
