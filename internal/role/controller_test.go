package role
//
//import (
//	"errors"
//	"net/http"
//	"net/http/httptest"
//	"reflect"
//	"strings"
//	"testing"
//
//	"github.com/gin-gonic/gin"
//)
//
//// ServiceMock MOCK
//type ServiceMock struct {
//	FailAll bool
//}
//
//// NewServiceMock MOCK
//func NewServiceMock(failAll bool) Service {
//	return &ServiceMock{
//		FailAll: failAll,
//	}
//}
//
//// GetAll MOCK
//func (s *ServiceMock) GetAll() ([]Role, error) {
//	if s.FailAll {
//		return nil, errors.New("Error")
//	}
//
//	return []Role{
//		{
//			ID: 1,
//		},
//	}, nil
//}
//
//// GetSingle MOCK
//func (s *ServiceMock) GetSingle(id int) (*Role, error) {
//	if s.FailAll {
//		return nil, errors.New("Error")
//	}
//
//	return &Role{
//		ID: id,
//	}, nil
//}
//
//// GetAssignedRole MOCK
//func (s *ServiceMock) GetAssignedRole(id string) (*AssignedRole, error) {
//	if s.FailAll {
//		return nil, errors.New("Error")
//	}
//
//	return &AssignedRole{AuthID: id}, nil
//}
//
//// AssignRole MOCK
//func (s *ServiceMock) AssignRole(authID string, roleID int) error {
//	if s.FailAll {
//		return errors.New("Error")
//	}
//	return nil
//}
//
//// DeleteAssignedRole MOCK
//func (s *ServiceMock) DeleteAssignedRole(authID string) error {
//	if s.FailAll {
//		return errors.New("Error")
//	}
//	return nil
//}
//
//func TestControllerImpl_GetAll(t *testing.T) {
//	type fields struct {
//		service Service
//	}
//	tests := []struct {
//		name         string
//		fields       fields
//		expectedBody interface{}
//		wantErr      bool
//	}{
//		{
//			name: "ok",
//			fields: fields{
//				service: NewServiceMock(false),
//			},
//			expectedBody: map[string]interface{}{
//				"total": 1,
//				"results": []Role{
//					{
//						ID: 1,
//					},
//				},
//			},
//			wantErr: false,
//		},
//		{
//			name: "error",
//			fields: fields{
//				service: NewServiceMock(true),
//			},
//			expectedBody: nil,
//			wantErr:      true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			w := httptest.NewRecorder()
//			c, _ := gin.CreateTestContext(w)
//
//			ctr := NewController(tt.fields.service)
//
//			err := ctr.GetAll(c)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("[GetAll] error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//
//			body, exists := c.Get(bodyKey)
//			if exists && !tt.wantErr {
//				if !reflect.DeepEqual(body, tt.expectedBody) {
//					t.Errorf("[GetAll] Expected body = %v, got %v", tt.expectedBody, body)
//					return
//				}
//			}
//
//			if !exists && !tt.wantErr {
//				t.Error("[GetAll] Expected body")
//			}
//		})
//	}
//}
//
//func TestControllerImpl_Get(t *testing.T) {
//	type fields struct {
//		service Service
//	}
//
//	tests := []struct {
//		name         string
//		fields       fields
//		expectedBody interface{}
//		params       []gin.Param
//		wantErr      bool
//	}{
//		{
//			name: "ok",
//			fields: fields{
//				service: NewServiceMock(false),
//			},
//			expectedBody: map[string]interface{}{
//				"total": 1,
//				"results": []interface{}{
//					&Role{
//						ID: 2,
//					},
//				},
//			},
//			params: []gin.Param{
//				{
//					Key:   "id",
//					Value: "2",
//				},
//			},
//			wantErr: false,
//		},
//		{
//			name: "invalid_param",
//			fields: fields{
//				service: NewServiceMock(false),
//			},
//			expectedBody: nil,
//			params: []gin.Param{
//				{
//					Key:   "id",
//					Value: "holi",
//				},
//			},
//			wantErr: true,
//		},
//		{
//			name: "service_error",
//			fields: fields{
//				service: NewServiceMock(true),
//			},
//			expectedBody: nil,
//			params: []gin.Param{
//				{
//					Key:   "id",
//					Value: "2",
//				},
//			},
//			wantErr: true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			w := httptest.NewRecorder()
//			c, _ := gin.CreateTestContext(w)
//
//			for _, param := range tt.params {
//				c.Params = append(c.Params, param)
//			}
//
//			ctr := NewController(tt.fields.service)
//
//			err := ctr.Get(c)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("[GetAll] error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//
//			body, exists := c.Get(bodyKey)
//			if exists && !tt.wantErr {
//				if !reflect.DeepEqual(body, tt.expectedBody) {
//					t.Errorf("[GetAll] Expected body = %v, got %v", tt.expectedBody, body)
//					return
//				}
//			}
//
//			if !exists && !tt.wantErr {
//				t.Error("[GetAll] Expected body")
//			}
//		})
//	}
//}
//
//func TestControllerImpl_GetAssignedRole(t *testing.T) {
//	type fields struct {
//		service Service
//	}
//
//	tests := []struct {
//		name         string
//		fields       fields
//		expectedBody interface{}
//		params       []gin.Param
//		wantErr      bool
//	}{
//		{
//			name: "ok",
//			fields: fields{
//				service: NewServiceMock(false),
//			},
//			expectedBody: map[string]interface{}{
//				"total": 1,
//				"results": []interface{}{
//					&AssignedRole{AuthID: "2"},
//				},
//			},
//			params: []gin.Param{
//				{
//					Key:   "auth_id",
//					Value: "2",
//				},
//			},
//			wantErr: false,
//		},
//		{
//			name: "no_auth_id",
//			fields: fields{
//				service: NewServiceMock(false),
//			},
//			expectedBody: nil,
//			params:       []gin.Param{},
//			wantErr:      true,
//		},
//		{
//			name: "error",
//			fields: fields{
//				service: NewServiceMock(true),
//			},
//			expectedBody: nil,
//			params: []gin.Param{
//				{
//					Key:   "auth_id",
//					Value: "2",
//				},
//			},
//			wantErr: true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			w := httptest.NewRecorder()
//			c, _ := gin.CreateTestContext(w)
//
//			for _, param := range tt.params {
//				c.Params = append(c.Params, param)
//			}
//
//			ctr := NewController(tt.fields.service)
//
//			err := ctr.GetAssignedRole(c)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("[GetAll] error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//
//			body, exists := c.Get(bodyKey)
//			if exists && !tt.wantErr {
//				if !reflect.DeepEqual(body, tt.expectedBody) {
//					t.Errorf("[GetAll] Expected body = %v, got %v", tt.expectedBody, body)
//					return
//				}
//			}
//
//			if !exists && !tt.wantErr {
//				t.Error("[GetAll] Expected body")
//			}
//		})
//	}
//}
//
//func TestControllerImpl_AssignRole(t *testing.T) {
//	type fields struct {
//		service Service
//	}
//
//	tests := []struct {
//		name         string
//		fields       fields
//		expectedBody interface{}
//		params       []gin.Param
//		requestBody  interface{}
//		wantErr      bool
//	}{
//		{
//			name: "no_body",
//			fields: fields{
//				service: NewServiceMock(false),
//			},
//			requestBody:  "{2",
//			expectedBody: nil,
//			params:       []gin.Param{},
//			wantErr:      true,
//		},
//		{
//			name: "wrong_body",
//			fields: fields{
//				service: NewServiceMock(false),
//			},
//			requestBody:  `{"auth_id": ""}`,
//			expectedBody: nil,
//			params:       []gin.Param{},
//			wantErr:      true,
//		},
//		{
//			name: "ok",
//			fields: fields{
//				service: NewServiceMock(true),
//			},
//			requestBody:  `{"auth_id": "jeje", "role_id": 1}`,
//			expectedBody: nil,
//			params:       []gin.Param{},
//			wantErr:      true,
//		},
//		{
//			name: "ok",
//			fields: fields{
//				service: NewServiceMock(false),
//			},
//			requestBody: `{"auth_id": "jeje", "role_id": 1}`,
//			expectedBody: map[string]string{
//				"status": "CREATED",
//			},
//			params:  []gin.Param{},
//			wantErr: false,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			w := httptest.NewRecorder()
//			c, _ := gin.CreateTestContext(w)
//
//			if tt.requestBody != nil {
//				req, err := http.NewRequest(http.MethodPost, "", strings.NewReader(tt.requestBody.(string)))
//				if err != nil {
//					panic(err)
//				}
//				c.Request = req
//			}
//
//			for _, param := range tt.params {
//				c.Params = append(c.Params, param)
//			}
//
//			ctr := NewController(tt.fields.service)
//			err := ctr.AssignRole(c)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("[GetAll] error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//
//			body, exists := c.Get(bodyKey)
//			if exists && !tt.wantErr {
//				if !reflect.DeepEqual(body, tt.expectedBody) {
//					t.Errorf("[GetAll] Expected body = %v, got %v", tt.expectedBody, body)
//					return
//				}
//			}
//
//			if !exists && !tt.wantErr {
//				t.Error("[GetAll] Expected body")
//			}
//		})
//	}
//}
//
//func TestControllerImpl_DeleteAssignedRole(t *testing.T) {
//	type fields struct {
//		service Service
//	}
//
//	tests := []struct {
//		name         string
//		fields       fields
//		expectedBody interface{}
//		params       []gin.Param
//		wantErr      bool
//	}{
//		{
//			name: "ok",
//			fields: fields{
//				service: NewServiceMock(false),
//			},
//			expectedBody: map[string]string{
//				"status": "DELETED",
//			},
//			params: []gin.Param{
//				{
//					Key:   "auth_id",
//					Value: "2",
//				},
//			},
//			wantErr: false,
//		},
//		{
//			name: "no_auth_id",
//			fields: fields{
//				service: NewServiceMock(false),
//			},
//			expectedBody: nil,
//			params:       []gin.Param{},
//			wantErr:      true,
//		},
//		{
//			name: "error",
//			fields: fields{
//				service: NewServiceMock(true),
//			},
//			expectedBody: nil,
//			params: []gin.Param{
//				{
//					Key:   "auth_id",
//					Value: "2",
//				},
//			},
//			wantErr: true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			w := httptest.NewRecorder()
//			c, _ := gin.CreateTestContext(w)
//
//			for _, param := range tt.params {
//				c.Params = append(c.Params, param)
//			}
//
//			ctr := NewController(tt.fields.service)
//
//			err := ctr.DeleteAssignedRole(c)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("[GetAll] error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//
//			body, exists := c.Get(bodyKey)
//			if exists && !tt.wantErr {
//				if !reflect.DeepEqual(body, tt.expectedBody) {
//					t.Errorf("[GetAll] Expected body = %v, got %v", tt.expectedBody, body)
//					return
//				}
//			}
//
//			if !exists && !tt.wantErr {
//				t.Error("[GetAll] Expected body")
//			}
//		})
//	}
//}
