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
func (s *ServiceImpl) GetAll() ([]Role, error) {
	return s.dao.GetAll()
}

// GetSingle Returns single existing role
func (s *ServiceImpl) GetSingle(id int) (*Role, error) {
	return s.dao.Get(id)
}

// Create Creates new role
func (s *ServiceImpl) Create(description string, claims []Claim) error {
	role := NewRole(2, description, claims)
	return s.dao.Create(role)
}

// Update Updates existing role
func (s *ServiceImpl) Update(id int, description string, claims []Claim) error {
	// TODO: Use validate package instead!
	if id <= 0 {
		msg := fmt.Sprintf("Invalid role ID %d", id)
		return errors.NewBadRequestApiError(msg)
	}

	for _, claim := range claims {
		if claim.ID <= 0 {
			msg := fmt.Sprintf("Invalid claim ID %d", claim.ID)
			return errors.NewBadRequestApiError(msg)
		}
	}

	role := NewRole(id, description, claims)
	return s.dao.Update(role)
}
