package role

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
	"github.com/azer/crud"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/common/log"
)

const (
	daoLogKey = "[ROLES][DAO]"
)

// DAOImpl Productive role DAO implementation
type DAOImpl struct {
	db *crud.DB
}

// NewDAO Returns new productive role DAO implementation
func NewDAO(db *crud.DB) DAO {
	return &DAOImpl{
		db: db,
	}
}

// GetAll Returns all active roles
func (d *DAOImpl) GetAll() ([]Role, error) {
	roles := []Role{}

	err := d.db.Read(roles, "SELECT * FROM roles")
	if err != nil {
		msg := fmt.Sprintf("%s Error retrieving all roles from DB - %+v", daoLogKey, err)
		log.Errorf(msg)
		return nil, apierror.NewInternalServerApiError(msg, err)
	}

	return roles, nil
}

// Get Returns role with given id
func (d *DAOImpl) Get(id int) (*Role, error) {
	role := &Role{}

	err := d.db.Read(role, "SELECT * FROM roles WHERE id = ?", id)
	if errors.Is(err, sql.ErrNoRows) {
		msg := fmt.Sprintf("%s Couldn't find role with ID (%d)", daoLogKey, id)
		return nil, apierror.NewNotFoundApiError(msg)
	}
	if err != nil {
		msg := fmt.Sprintf("%s Error retrieving role with ID (%d) from DB - %+v", daoLogKey, id, err)
		log.Errorf(msg)
		return nil, apierror.NewInternalServerApiError(msg, err)
	}

	return role, nil
}

// Create Creates a new role
func (d *DAOImpl) Create(role *Role) error {
	err := d.db.Create(role)
	if err != nil {
		msg := fmt.Sprintf("%s Error creating role - %+v", daoLogKey, err)
		log.Errorf(msg)
		return apierror.NewInternalServerApiError(msg, err)
	}

	return nil
}

// Update Updates existing role
func (d *DAOImpl) Update(role *Role) error {
	err := d.db.Update(role)
	if errors.Is(err, sql.ErrNoRows) {
		msg := fmt.Sprintf("%s Couldn't find role with ID (%d)", daoLogKey, role.ID)
		return apierror.NewNotFoundApiError(msg)
	}
	if err != nil {
		msg := fmt.Sprintf("%s Error updating role - %+v", daoLogKey, err)
		log.Errorf(msg)
		return apierror.NewInternalServerApiError(msg, err)
	}

	return nil
}
