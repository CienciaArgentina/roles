package role

import "github.com/gin-gonic/gin"

// ControllerImpl Productive role controller implementation
type ControllerImpl struct {
	service Service
}

// NewController Returns new productive controller
func NewController(s Service) Controller {
	return &ControllerImpl{
		service: s,
	}
}

// GetAll Returns all existing roles
func (ctr *ControllerImpl) GetAll(c *gin.Context) error {
	// TODO: Implement
	return nil
}

// Get Returns single role
func (ctr *ControllerImpl) Get(c *gin.Context) error {
	// TODO: Implement
	return nil
}

// Create Creates new role
func (ctr *ControllerImpl) Create(c *gin.Context) error {
	// TODO: Implement
	return nil
}

// Update Updates existing role
func (ctr *ControllerImpl) Update(c *gin.Context) error {
	// TODO: Implement
	return nil
}
