package role

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
func NewRole(id int, description string, claims []Claim) *Role {
	return &Role{
		ID:          id,
		Description: description,
		Claims:      claims,
	}
}
