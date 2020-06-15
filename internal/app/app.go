package app

import (
	"github.com/gin-gonic/gin"
)

// App App definition
type App struct {
	Router *gin.Engine
}

// New Returns new app
func New() *App {
	return &App{}
}

// Start Initializes app
func (a *App) Start() {
	a.Router = gin.Default()
	a.MapURLs()
	a.Router.Run()
}
