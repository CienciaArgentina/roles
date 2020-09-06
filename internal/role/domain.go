package role

// AssignedRole Role assigned to auth ID
type AssignedRole struct {
	AuthID int64 `json:"auth_id"`
	Roles  []Role `json:"roles"`
}

// Role Structure of a role to be assumed by a user
type Role struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Claims      []Claim `json:"claims"`
}

// Claim Defines role permissions
type Claim struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// NewRole Returns new role
func NewRole(description string, claims []Claim) *Role {
	return &Role{
		Description: description,
		Claims:      claims,
	}
}

// AssignRoleRequest Request of role assignment
type AssignRoleRequest struct {
	AuthID int64 `json:"auth_id"`
	RoleID int    `json:"role_id"`
}
