package role

import (
	"github.com/gin-gonic/gin"
)

// DAO Describes role DAO interface
type DAO interface {
	GetAll() ([]Role, error)
	Get(id int) (*Role, error)
	GetAssignedRole(id int64) (*AssignedRole, error)
	UpsertAssignedRole(authID int64, roleID int) error
	DeleteAssignedRole(authID int64) error
}

// Service Describes role service interface
type Service interface {
	GetAll() ([]Role, error)
	GetSingle(id int) (*Role, error)
	GetAssignedRole(id int64) (*AssignedRole, error)
	AssignRole(authID int64, roleID int) error
	DeleteAssignedRole(authID int64) error
}

// Controller Describes role controller interface
type Controller interface {
	GetAll(c *gin.Context) error
	Get(c *gin.Context) error
	GetAssignedRole(*gin.Context) error
	AssignRole(*gin.Context) error
	DeleteAssignedRole(*gin.Context) error
}
