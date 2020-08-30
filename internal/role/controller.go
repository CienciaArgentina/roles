package role

import (
	"fmt"
	"strconv"

	"github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
	"github.com/CienciaArgentina/go-backend-commons/pkg/middleware"
	"github.com/CienciaArgentina/roles/internal/adapter"
	"github.com/gin-gonic/gin"
)

const (
	bodyKey = middleware.ResponseBodyKey
	codeKey = middleware.ResponseCodeKey
)

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
	roles, err := ctr.service.GetAll()
	if err != nil {
		return err
	}

	c.Set(bodyKey, adapter.Adapt(roles))
	return nil
}

// Get Returns single role
func (ctr *ControllerImpl) Get(c *gin.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		msg := fmt.Sprintf("Couldn't convert ID %s to string", idStr)
		return apierror.NewBadRequestApiError(msg)
	}

	role, err := ctr.service.GetSingle(id)
	if err != nil {
		return err
	}

	c.Set(bodyKey, adapter.Adapt(role))
	return nil
}

// GetAssignedRole Returns assigned role for auth_id in path
func (ctr *ControllerImpl) GetAssignedRole(c *gin.Context) error {
	id := c.Param("auth_id")
	if id == "" {
		return apierror.NewBadRequestApiError("Empty auth ID")
	}

	role, err := ctr.service.GetAssignedRole(id)
	if err != nil {
		return err
	}

	c.Set(bodyKey, adapter.Adapt(role))
	return nil
}

// AssignRole Assigns role to given auth ID
func (ctr *ControllerImpl) AssignRole(c *gin.Context) error {
	var body AssignRoleRequest
	if err := c.BindJSON(&body); err != nil {
		return apierror.NewBadRequestApiError("Error reading body")
	}
	if body.AuthID == "" || body.RoleID <= 0 {
		return apierror.NewBadRequestApiError("Invalid auth ID or role ID")
	}

	err := ctr.service.AssignRole(body.AuthID, body.RoleID)
	if err != nil {
		return err
	}

	c.Set(bodyKey, map[string]string{
		"status": "CREATED",
	})
	return nil
}

// DeleteAssignedRole Deletes all assigned roles for given auth ID
func (ctr *ControllerImpl) DeleteAssignedRole(c *gin.Context) error {
	id := c.Param("auth_id")
	if id == "" {
		return apierror.NewBadRequestApiError("Empty auth ID")
	}

	err := ctr.service.DeleteAssignedRole(id)
	if err != nil {
		return err
	}

	c.Set(bodyKey, map[string]string{
		"status": "DELETED",
	})
	return nil
}
