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
	return s.dao.GetAll(-1)
}

// GetSingle Returns single existing role
func (s *ServiceImpl) GetSingle(id int) (*Role, error) {
	return s.dao.Get(id)
}
