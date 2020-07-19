package role

import (
	"github.com/gin-gonic/gin"
)

// DAO Describes role DAO interface
type DAO interface {
	GetAll() ([]Role, error)
	Get(int) (*Role, error)
	Create(*Role) error
	Update(*Role) error
}

// Service Describes role service interface
type Service interface {
	GetAll() ([]Role, error)
	GetSingle(int) (*Role, error)
	Create(string, []Claim) error
	Update(int, string, []Claim) error
}

// Controller Describes role controller interface
type Controller interface {
	GetAll(c *gin.Context) error
	Get(c *gin.Context) error
	Create(c *gin.Context) error
	Update(c *gin.Context) error
}
