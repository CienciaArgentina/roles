package role

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func sameRoleSliceUnordered(x, y []Role) bool {
	if len(x) != len(y) {
		return false
	}

	// create a map of Role -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x.Description]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y.Description]; !ok {
			return false
		}
		diff[_y.Description]--
		if diff[_y.Description] == 0 {
			delete(diff, _y.Description)
		}
	}

	return len(diff) == 0
}

func TestGetAllError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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
	mock.ExpectQuery(query).WillReturnError(errors.New("Internal error"))

	dao := NewDAO(db)
	_, err = dao.GetAll()
	if err == nil {
		t.Error("Expected error")
	}
}

func TestGetAllWrongTableSchema(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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

	table := sqlmock.NewRows([]string{"role_id", "role_description", "claim_id", "claim_description", "claim_description_2"})
	table.AddRow(123, "test_role", 321, "test_claim", "test")
	table.AddRow(123, "test_role", 323, "test_claim_2", "test")
	table.AddRow(321, "test_role_2", 321, "test_claim", "test")
	table.AddRow(321, "test_role_2", 323, "test_claim_2", "test")
	mock.ExpectQuery(query).WillReturnRows(table)

	dao := NewDAO(db)
	_, err = dao.GetAll()
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestGetAllOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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

	table := sqlmock.NewRows([]string{"role_id", "role_description", "claim_id", "claim_description"})
	table.AddRow(123, "test_role", 321, "test_claim")
	table.AddRow(123, "test_role", 323, "test_claim_2")
	table.AddRow(321, "test_role_2", 321, "test_claim")
	table.AddRow(321, "test_role_2", 323, "test_claim_2")
	mock.ExpectQuery(query).WillReturnRows(table)

	dao := NewDAO(db)
	roles, err := dao.GetAll()
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
		return
	}

	expectedRoles := []Role{
		{
			ID:          321,
			Description: "test_role_2",
			Claims: []Claim{
				{
					ID:          321,
					Description: "test_claim",
				},
				{
					ID:          323,
					Description: "test_claim_2",
				},
			},
		},
		{
			ID:          123,
			Description: "test_role",
			Claims: []Claim{
				{
					ID:          321,
					Description: "test_claim",
				},
				{
					ID:          323,
					Description: "test_claim_2",
				},
			},
		},
	}

	if !sameRoleSliceUnordered(roles, expectedRoles) {
		t.Errorf("Expected %+v got %+v", expectedRoles, roles)
	}
}

func TestGetError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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
	`, 123)

	mock.ExpectQuery(query).WillReturnError(errors.New("Internal error"))

	dao := NewDAO(db)
	_, err = dao.Get(123)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestGetOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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
	`, 123)

	table := sqlmock.NewRows([]string{"role_id", "role_description", "claim_id", "claim_description"})
	table.AddRow(123, "test_role", 321, "test_claim")
	table.AddRow(123, "test_role", 323, "test_claim_2")
	mock.ExpectQuery(query).WillReturnRows(table)

	dao := NewDAO(db)
	role, err := dao.Get(123)
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
		return
	}

	expected := &Role{
		ID:          123,
		Description: "test_role",
		Claims: []Claim{
			{
				ID:          321,
				Description: "test_claim",
			},
			{
				ID:          323,
				Description: "test_claim_2",
			},
		},
	}

	if !reflect.DeepEqual(role, expected) {
		t.Errorf("Expected %+v got %+v", expected, role)
	}
}

func TestGetNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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
	`, 123)

	table := sqlmock.NewRows([]string{"role_id", "role_description", "claim_id", "claim_description"})
	mock.ExpectQuery(query).WillReturnRows(table)

	dao := NewDAO(db)
	_, err = dao.Get(123)
	if err == nil {
		t.Error("Expected not found error")
		return
	}
}

func TestGetWrongSchema(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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
	`, 123)

	table := sqlmock.NewRows([]string{"role_id", "role_description", "claim_id", "claim_description", "test"})
	table.AddRow(123, "test_role", 321, "test_claim", "test")
	mock.ExpectQuery(query).WillReturnRows(table)

	dao := NewDAO(db)
	_, err = dao.Get(123)
	if err == nil {
		t.Error("Expected internal error")
		return
	}
}

func TestGetAssignedRoleError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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

	WHERE auth_id = '%d';
	`, 123)

	mock.ExpectQuery(query).WillReturnError(errors.New("Internal error"))

	dao := NewDAO(db)
	_, err = dao.GetAssignedRole(123)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestGetAssignedRoleNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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

	WHERE auth_id = '%d';
	`, 123)

	table := sqlmock.NewRows([]string{"auth_id", "role_id", "role_description", "claim_id", "claim_description"})
	mock.ExpectQuery(query).WillReturnRows(table)

	dao := NewDAO(db)
	_, err = dao.GetAssignedRole(123)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestGetAssignedRoleUnmarshal(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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

	WHERE auth_id = '%d';
	`, 123)

	table := sqlmock.NewRows([]string{"auth_id", "role_id", "role_description", "claim_id", "claim_description", "test"})
	table.AddRow(123, 123, "test_role", 321, "test_claim", "test")
	mock.ExpectQuery(query).WillReturnRows(table)

	dao := NewDAO(db)
	_, err = dao.GetAssignedRole(123)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestGetAssignedRoleOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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

	WHERE auth_id = '%d';
	`, 123)

	table := sqlmock.NewRows([]string{"auth_id", "role_id", "role_description", "claim_id", "claim_description"})
	table.AddRow(123, 123, "test_role", 321, "test_claim")
	table.AddRow(123, 123, "test_role", 323, "test_claim_2")
	table.AddRow(123, 321, "test_role_2", 321, "test_claim")
	table.AddRow(123, 321, "test_role_2", 323, "test_claim_2")
	mock.ExpectQuery(query).WillReturnRows(table)

	dao := NewDAO(db)
	got, err := dao.GetAssignedRole(123)
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
		return
	}

	expected := AssignedRole{
		AuthID: 123,
		Roles: []Role{
			{
				ID:          123,
				Description: "test_role",
				Claims: []Claim{
					{
						ID:          321,
						Description: "test_claim",
					},
					{
						ID:          323,
						Description: "test_claim_2",
					},
				},
			},
			{
				ID:          321,
				Description: "test_role_2",
				Claims: []Claim{
					{
						ID:          321,
						Description: "test_claim",
					},
					{
						ID:          323,
						Description: "test_claim_2",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(got, &expected) {
		t.Errorf("Expected %+v got %+v", expected, got)
	}
}

func TestUpsertAssignedRoleError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	statement := fmt.Sprintf(`
	INSERT INTO assigned_roles \(auth_id, role_id\) VALUES \(%d, %d\)
	`, 123, 321)

	mock.ExpectExec(statement).WillReturnError(errors.New("Internal error"))

	dao := NewDAO(db)
	err = dao.UpsertAssignedRole(123, 321)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestUpsertAssignedRoleOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	statement := fmt.Sprintf(`
	INSERT INTO assigned_roles \(auth_id, role_id\) VALUES \(%d, %d\)
	`, 123, 321)

	mock.ExpectExec(statement).WillReturnResult(sqlmock.NewResult(1, 1))

	dao := NewDAO(db)
	err = dao.UpsertAssignedRole(123, 321)
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}
}

func TestDeleteAssignedRoleError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	statement := fmt.Sprintf(`
		DELETE FROM assigned_roles 
		WHERE auth_id = '%d'
	`, 123)

	mock.ExpectExec(statement).WillReturnError(errors.New("Internal error"))

	dao := NewDAO(db)
	err = dao.DeleteAssignedRole(123)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestDeleteAssignedRoleOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	statement := fmt.Sprintf(`
		DELETE FROM assigned_roles 
		WHERE auth_id = '%d'
	`, 123)

	mock.ExpectExec(statement).WillReturnResult(sqlmock.NewResult(1, 1))

	dao := NewDAO(db)
	err = dao.DeleteAssignedRole(123)
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}
}
