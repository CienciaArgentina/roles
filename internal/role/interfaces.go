package role

import (
	errors "github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
)

// DAO Describes role DAO interface
type DAO interface {
	Get(int) (*Role, errors.ApiError)
	Create() // TODO: Define
	Update() // TODO: Define
}

// Service Describes role service interface
type Service interface {
	Get(int) (*Role, errors.ApiError)
}
