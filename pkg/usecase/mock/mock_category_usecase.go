// Code generated by MockGen. DO NOT EDIT.
// Source: /home/jasim/CityVibe-Ecommerce-CleanCode-Project/pkg/usecase/interface/category.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	domain "github.com/14jasimmtp/CityVibe-Project-Clean-Architecture/pkg/domain"
	models "github.com/14jasimmtp/CityVibe-Project-Clean-Architecture/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

// MockCategoryUsecase is a mock of CategoryUsecase interface.
type MockCategoryUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryUsecaseMockRecorder
}

// MockCategoryUsecaseMockRecorder is the mock recorder for MockCategoryUsecase.
type MockCategoryUsecaseMockRecorder struct {
	mock *MockCategoryUsecase
}

// NewMockCategoryUsecase creates a new mock instance.
func NewMockCategoryUsecase(ctrl *gomock.Controller) *MockCategoryUsecase {
	mock := &MockCategoryUsecase{ctrl: ctrl}
	mock.recorder = &MockCategoryUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryUsecase) EXPECT() *MockCategoryUsecaseMockRecorder {
	return m.recorder
}

// AddCategory mocks base method.
func (m *MockCategoryUsecase) AddCategory(category models.Category) (domain.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCategory", category)
	ret0, _ := ret[0].(domain.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCategory indicates an expected call of AddCategory.
func (mr *MockCategoryUsecaseMockRecorder) AddCategory(category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCategory", reflect.TypeOf((*MockCategoryUsecase)(nil).AddCategory), category)
}

// DeleteCategory mocks base method.
func (m *MockCategoryUsecase) DeleteCategory(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockCategoryUsecaseMockRecorder) DeleteCategory(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockCategoryUsecase)(nil).DeleteCategory), id)
}

// GetCategory mocks base method.
func (m *MockCategoryUsecase) GetCategory() ([]domain.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory")
	ret0, _ := ret[0].([]domain.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockCategoryUsecaseMockRecorder) GetCategory() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockCategoryUsecase)(nil).GetCategory))
}

// UpdateCategory mocks base method.
func (m *MockCategoryUsecase) UpdateCategory(current, new string) (domain.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", current, new)
	ret0, _ := ret[0].(domain.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategory indicates an expected call of UpdateCategory.
func (mr *MockCategoryUsecaseMockRecorder) UpdateCategory(current, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockCategoryUsecase)(nil).UpdateCategory), current, new)
}
