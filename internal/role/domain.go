package role

// AssignedRole Role assigned to auth ID
type AssignedRole struct {
	AuthID string `json:"auth_id"`
	Role   Role   `json:"role"`
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
