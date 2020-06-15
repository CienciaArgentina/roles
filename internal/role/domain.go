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
