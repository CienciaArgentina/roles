package role

import (
	"github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
)

// ServiceImpl Productive role service implemenatation
type ServiceImpl struct {
	dao DAO
}

// NewService Returns new productive service
func NewService() Service {
	return &ServiceImpl{}
}

// Get Returns existing role
func (s *ServiceImpl) Get(id int) (*Role, apierror.ApiError) {
	return s.dao.Get(id)
}
