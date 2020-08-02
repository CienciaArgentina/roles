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

// GetAll Returns all active roles
func (d *DAOImpl) GetAll(id int) ([]Role, error) {
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
	if id != -1 {
		query += fmt.Sprintf("WHERE r.id = %d", id)
	}

	rows, err := d.db.Query(query)
	if err != nil {
		msg := "Error retrieving all roles from DB"
		log.Errorf("%s %s", err, daoLogKey, msg)
		return nil, apierror.NewInternalServerApiError(msg, err)
	}
	defer rows.Close()

	roleMap := map[int]*Role{}
	for rows.Next() {
		var roleID int
		var claimID int
		var claimDescription string
		var roleDescription string

		err := rows.Scan(&roleID, &roleDescription, &claimID, &claimDescription)
		if err != nil {
			msg := "Error unmarshalling role"
			log.Errorf("%s %s", err, daoLogKey, msg)
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

// Get Returns role with given id
func (d *DAOImpl) Get(id int) (*Role, error) {
	roles, err := d.GetAll(id)
	if err != nil {
		return nil, err
	}
	if len(roles) == 0 {
		msg := fmt.Sprintf("Role (%d) not found", id)
		return nil, apierror.NewNotFoundApiError(msg)
	}

	return &roles[0], nil
}
