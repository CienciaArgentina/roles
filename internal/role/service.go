package role

// ServiceImpl Productive role service implemenatation
type ServiceImpl struct {
	dao DAO
}

// NewService Returns new productive service
func NewService(d DAO) Service {
	return &ServiceImpl{
		dao: d,
	}
}

// GetAll Returns all existing roles
func (s *ServiceImpl) GetAll() ([]Role, error) {
	return s.dao.GetAll()
}

// GetSingle Returns single existing role
func (s *ServiceImpl) GetSingle(id int) (*Role, error) {
	return s.dao.Get(id)
}

// GetAssignedRole ...
func (s *ServiceImpl) GetAssignedRole(id string) (*AssignedRole, error) {
	return s.dao.GetAssignedRole(id)
}

// AssignRole ...
func (s *ServiceImpl) AssignRole(authID string, roleID int) error {
	return nil
}

// UpdateAssignedRole ...
func (s *ServiceImpl) UpdateAssignedRole(authID string, roleID int) error {
	return nil
}

// DeleteAssignedRole ...
func (s *ServiceImpl) DeleteAssignedRole(authID string, roleID int) error {
	return nil
}
