package role

import (
	errors "github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
)

// Service Describes role service interface
type Service interface {
	GetRole(int) (*Role, errors.ApiError)
}
