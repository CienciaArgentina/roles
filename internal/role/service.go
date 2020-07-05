package role

import (
	"fmt"

	errors "github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
)

// ServiceImpl Productive role service implemenatation
type ServiceImpl struct {
	dao DAO
}

// NewService Returns new productive service
func NewService(d DAO) Service {
	return &ServiceImpl{
		dao: d,
	}
}

// GetAll Returns all existing roles
func (s *ServiceImpl) GetAll() ([]Role, errors.ApiError) {
	return s.dao.GetAll()
}

// GetSingle Returns single existing role
func (s *ServiceImpl) GetSingle(id int) (*Role, errors.ApiError) {
	return s.dao.Get(id)
}

// Create Creates new role
func (s *ServiceImpl) Create(description string, claims []Claim) (*Role, errors.ApiError) {
	role := NewRole(2, description, claims)
	return role, s.dao.Create(role)
}

// Update Updates existing role
func (s *ServiceImpl) Update(id int, description string, claims []Claim) (*Role, errors.ApiError) {
	if id <= 0 {
		msg := fmt.Sprintf("Invalid role ID %d", id)
		return nil, errors.NewBadRequestApiError(msg)
	}
	for _, claim := range claims {
		if claim.ID <= 0 {
			msg := fmt.Sprintf("Invalid claim ID %d", claim.ID)
			return nil, errors.NewBadRequestApiError(msg)
		}
	}

	role := NewRole(id, description, claims)
	return role, s.dao.Update(role)
}
