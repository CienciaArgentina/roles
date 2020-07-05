package role

import (
	errors "github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
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
