package role

import (
	"errors"
	"reflect"
	"testing"
)

// MockDAO MOCK
type MockDAO struct {
}

// NewMockDAO MOCK
func NewMockDAO() DAO {
	return &MockDAO{}
}

// GetAll MOCK
func (d *MockDAO) GetAll() ([]Role, error) {
	return []Role{
		{
			ID:          -1,
			Description: "This is a test role!",
		},
	}, nil
}

// Get MOCK
func (d *MockDAO) Get(id int) (*Role, error) {
	if id < 0 {
		return nil, errors.New("Not found")
	}

	return &Role{
		ID:     id,
		Claims: []Claim{},
	}, nil
}

// GetAssignedRole MOCK
func (d *MockDAO) GetAssignedRole(id int64) (*AssignedRole, error) {
	if id == 0 {
		return nil, errors.New("Not found")
	}

	return &AssignedRole{
		AuthID: id,
	}, nil
}

// UpsertAssignedRole MOCK
func (d *MockDAO) UpsertAssignedRole(authID int64, roleID int) error {
	if authID == 0 {
		return errors.New("Not found")
	}

	return nil
}

// DeleteAssignedRole MOCK
func (d *MockDAO) DeleteAssignedRole(authID int64) error {
	if authID == 0 {
		return errors.New("Not found")
	}

	return nil
}

func TestServiceImpl_GetAll(t *testing.T) {
	type fields struct {
		dao DAO
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Role
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				dao: NewMockDAO(),
			},
			want: []Role{
				{
					ID:          -1,
					Description: "This is a test role!",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServiceImpl{
				dao: tt.fields.dao,
			}
			got, err := s.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceImpl.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceImpl_GetSingle(t *testing.T) {
	type fields struct {
		dao DAO
	}

	type args struct {
		id int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Role
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				dao: NewMockDAO(),
			},
			args: args{
				id: 2,
			},
			want: &Role{
				ID:     2,
				Claims: []Claim{},
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				dao: NewMockDAO(),
			},
			args: args{
				id: -1,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.fields.dao)
			got, err := s.GetSingle(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.GetSingle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceImpl.GetSingle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceImpl_GetAssignedRole(t *testing.T) {
	type fields struct {
		dao DAO
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AssignedRole
		wantErr bool
	}{
		{
			name: "get_assigned_internal_error",
			fields: fields{
				dao: NewMockDAO(),
			},
			args: args{
				id: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ok",
			fields: fields{
				dao: NewMockDAO(),
			},
			args: args{
				id: 123,
			},
			want: &AssignedRole{
				AuthID: 123,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServiceImpl{
				dao: tt.fields.dao,
			}
			got, err := s.GetAssignedRole(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.GetAssignedRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceImpl.GetAssignedRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceImpl_AssignRole(t *testing.T) {
	type fields struct {
		dao DAO
	}
	type args struct {
		authID int64
		roleID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				dao: NewMockDAO(),
			},
			args: args{
				authID: 123,
			},
			wantErr: false,
		},
		{
			name: "assign_internal_error",
			fields: fields{
				dao: NewMockDAO(),
			},
			args: args{
				authID: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServiceImpl{
				dao: tt.fields.dao,
			}
			if err := s.AssignRole(tt.args.authID, tt.args.roleID); (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.AssignRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServiceImpl_DeleteAssignedRole(t *testing.T) {
	type fields struct {
		dao DAO
	}
	type args struct {
		authID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				dao: NewMockDAO(),
			},
			args:    args{authID: 123},
			wantErr: false,
		},
		{
			name: "delete_internal_error",
			fields: fields{
				dao: NewMockDAO(),
			},
			args:    args{authID: 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServiceImpl{
				dao: tt.fields.dao,
			}
			if err := s.DeleteAssignedRole(tt.args.authID); (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.DeleteAssignedRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
