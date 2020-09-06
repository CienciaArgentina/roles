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

// GetAssignedRole Returns assigned roles for given authID
func (s *ServiceImpl) GetAssignedRole(id int64) (*AssignedRole, error) {
	return s.dao.GetAssignedRole(id)
}

// AssignRole Assigned new role to given auth ID
func (s *ServiceImpl) AssignRole(authID int64, roleID int) error {
	return s.dao.UpsertAssignedRole(authID, roleID)
}

// DeleteAssignedRole Deletes auth ID role assignment
func (s *ServiceImpl) DeleteAssignedRole(authID int64) error {
	return s.dao.DeleteAssignedRole(authID)
}
