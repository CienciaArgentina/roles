package role

import (
	"fmt"
	"strconv"

	"github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
	"github.com/CienciaArgentina/go-backend-commons/pkg/middleware"
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

	c.Set(bodyKey, roles)
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

	c.Set(bodyKey, role)
	return nil
}

// Create Creates new role
func (ctr *ControllerImpl) Create(c *gin.Context) error {
	var body Role
	if err := c.BindJSON(&body); err != nil {
		return apierror.NewBadRequestApiError("Couldn't parse body")
	}

	err := ctr.service.Create(body.Description, body.Claims)
	if err != nil {
		return err
	}

	c.Set(bodyKey, map[string]string{"STATUS": "CREATED"})
	return nil
}

// Update Updates existing role
func (ctr *ControllerImpl) Update(c *gin.Context) error {
	// TODO: Implement
	return nil
}
