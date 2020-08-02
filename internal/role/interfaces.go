package role

import (
	"github.com/gin-gonic/gin"
)

// DAO Describes role DAO interface
type DAO interface {
	GetAll(int) ([]Role, error)
	Get(int) (*Role, error)
}

// Service Describes role service interface
type Service interface {
	GetAll() ([]Role, error)
	GetSingle(int) (*Role, error)
}

// Controller Describes role controller interface
type Controller interface {
	GetAll(c *gin.Context) error
	Get(c *gin.Context) error
}
