package role

import (
	"database/sql"
	"fmt"

	"github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
	"github.com/prometheus/common/log"
)

const (
	daoLogKey = "[ROLES][DAO]"
)

// DAOImpl Productive role DAO implementation
type DAOImpl struct {
	db *sql.DB
}

// NewDAO Returns new productive role DAO implementation
func NewDAO(db *sql.DB) DAO {
	return &DAOImpl{
		db: db,
	}
}

func (d *DAOImpl) getRolesFromRows(rows *sql.Rows) ([]Role, error) {
	roleMap := map[int]*Role{}
	for rows.Next() {
		var roleID int
		var claimID int
		var claimDescription string
		var roleDescription string

		err := rows.Scan(&roleID, &roleDescription, &claimID, &claimDescription)
		if err != nil {
			msg := "Error unmarshalling role"
			log.Errorf("%s %s - %+v", daoLogKey, msg, err)
			return nil, apierror.NewInternalServerApiError(msg, err)
		}

		role, exist := roleMap[roleID]
		if !exist {
			roleMap[roleID] = &Role{
				ID:          roleID,
				Description: roleDescription,
				Claims: []Claim{
					{
						ID:          claimID,
						Description: claimDescription,
					},
				},
			}
			continue
		}

		role.Claims = append(role.Claims, Claim{
			ID:          claimID,
			Description: claimDescription,
		})
	}

	roles := []Role{}
	for _, role := range roleMap {
		roles = append(roles, *role)
	}

	return roles, nil
}

// GetAll Returns all active roles
func (d *DAOImpl) GetAll() ([]Role, error) {
	query := `
	SELECT 
		r.id AS role_id,
		r.description AS role_description,
		c.id AS claim_id,
		c.description AS claim_description
		
	FROM roles_x_claims rxc
		INNER JOIN roles r ON r.id = rxc.role_id
		INNER JOIN claims c ON c.id = rxc.claim_id
	`

	rows, err := d.db.Query(query)
	if err != nil {
		msg := "Error retrieving roles from DB"
		log.Errorf("%s %s - %+v", daoLogKey, msg, err)
		return nil, apierror.NewInternalServerApiError(msg, err)
	}
	defer rows.Close()

	return d.getRolesFromRows(rows)
}

// Get Returns role with given id
func (d *DAOImpl) Get(id int) (*Role, error) {
	query := fmt.Sprintf(`
	SELECT 
		r.id AS role_id,
		r.description AS role_description,
		c.id AS claim_id,
		c.description AS claim_description

	FROM roles_x_claims rxc
		INNER JOIN roles r ON r.id = rxc.role_id
		INNER JOIN claims c ON c.id = rxc.claim_id

	WHERE r.id = %d
	`, id)

	rows, err := d.db.Query(query)
	if err != nil {
		msg := "Error retrieving role from DB"
		log.Errorf("%s %s - %+v", daoLogKey, msg, err)
		return nil, apierror.NewInternalServerApiError(msg, err)
	}
	defer rows.Close()

	roles, err := d.getRolesFromRows(rows)
	if err != nil {
		return nil, err
	}

	if len(roles) == 0 {
		msg := fmt.Sprintf("Role (%d) not found", id)
		return nil, apierror.NewNotFoundApiError(msg)
	}

	return &roles[0], nil
}

// GetAssignedRole Returns assigned role for given auth ID
func (d *DAOImpl) GetAssignedRole(id string) (*AssignedRole, error) {
	query := fmt.Sprintf(`
	SELECT 
		assigned.auth_id AS auth_id,
		r.id AS role_id,
		r.description AS role_description,
		c.id AS claim_id,
		c.description AS claim_description

	FROM assigned_roles assigned
		INNER JOIN roles_x_claims rxc ON rxc.role_id = assigned.role_id
		INNER JOIN roles r ON r.id = rxc.role_id
		INNER JOIN claims c ON c.id = rxc.claim_id

	WHERE auth_id = '%s';
	`, id)

	rows, err := d.db.Query(query)
	if err != nil {
		msg := fmt.Sprintf("Error retrieving assigned role from DB for auth_id (%s)", id)
		log.Errorf("%s %s - %+v", daoLogKey, msg, err)
		return nil, apierror.NewInternalServerApiError(msg, err)
	}
	defer rows.Close()

	roleMap := map[string]*Role{}
	for rows.Next() {
		var authID string
		var roleID int
		var claimID int
		var claimDescription string
		var roleDescription string

		err := rows.Scan(&authID, &roleID, &roleDescription, &claimID, &claimDescription)
		if err != nil {
			msg := "Error unmarshalling role"
			log.Errorf("%s %s - %+v", daoLogKey, msg, err)
			return nil, apierror.NewInternalServerApiError(msg, err)
		}

		role, exist := roleMap[authID]
		if !exist {
			roleMap[authID] = &Role{
				ID:          roleID,
				Description: roleDescription,
				Claims: []Claim{
					{
						ID:          claimID,
						Description: claimDescription,
					},
				},
			}
			continue
		}

		role.Claims = append(role.Claims, Claim{
			ID:          claimID,
			Description: claimDescription,
		})
	}

	assignedRoles := []AssignedRole{}
	for authID, role := range roleMap {
		assignedRoles = append(assignedRoles, AssignedRole{
			AuthID: authID,
			Role:   *role,
		})
	}
	if len(assignedRoles) == 0 {
		msg := fmt.Sprintf("Assigned Role for ID (%s) not found", id)
		return nil, apierror.NewNotFoundApiError(msg)
	}

	return &assignedRoles[0], nil
}

// UpsertAssignedRole Inserts new element if record doesn't exist, updates otherwise
func (d *DAOImpl) UpsertAssignedRole(authID string, roleID int) error {
	statement := fmt.Sprintf(`
	INSERT INTO assigned_roles (auth_id, role_id) VALUES ('%s', %d)
	`, authID, roleID)

	_, err := d.db.Exec(statement)
	if err != nil {
		msg := fmt.Sprintf("Error assigning role (%d) to auth ID (%s)", roleID, authID)
		log.Errorf("%s %s - %+v", daoLogKey, msg, err)
		return apierror.NewInternalServerApiError(msg, err)
	}

	return nil
}

// DeleteAssignedRole Deletes assigned role from auth ID
func (d *DAOImpl) DeleteAssignedRole(authID string) error {
	statement := fmt.Sprintf(`
		DELETE FROM assigned_roles 
		WHERE auth_id = '%s'
	`, authID)

	_, err := d.db.Exec(statement)
	if err != nil {
		msg := fmt.Sprintf("Error deleting auth ID (%s)", authID)
		log.Errorf("%s %s - %+v", daoLogKey, msg, err)
		return apierror.NewInternalServerApiError(msg, err)
	}

	return nil
}
