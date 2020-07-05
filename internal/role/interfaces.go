package role

import (
	errors "github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
	"github.com/gin-gonic/gin"
)

// DAO Describes role DAO interface
type DAO interface {
	GetAll() ([]Role, errors.ApiError)
	Get(int) (*Role, errors.ApiError)
	Create(*Role) errors.ApiError
	Update(*Role) errors.ApiError
}

// Service Describes role service interface
type Service interface {
	GetAll() ([]Role, errors.ApiError)
	GetSingle(int) (*Role, errors.ApiError)
	Create(string, []Claim) (*Role, errors.ApiError)
	Update(int, string, []Claim) (*Role, errors.ApiError)
}

// Controller Describes role controller interface
type Controller interface {
	GetAll(c *gin.Context) error
	Get(c *gin.Context) error
	Create(c *gin.Context) error
	Update(c *gin.Context) error
}
