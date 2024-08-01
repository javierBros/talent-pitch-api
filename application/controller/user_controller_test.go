package controller_test

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"project/application/controller"
	"project/application/core/domain"
	"project/application/core/entities"
	"strings"
	"testing"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) GetUserByID(id int) (*entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserService) ListUsers(limit, offset int) ([]entities.User, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]entities.User), args.Error(1)
}

func (m *MockUserService) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name":"John Doe","email":"john.doe@example.com","image_path":"http://example.com/image.jpg"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserService)
	controller := &controller.UserController{UserService: mockService}

	user := &entities.User{Name: "John Doe", Email: "john.doe@example.com", ImagePath: "http://example.com/image.jpg"}
	mockService.On("CreateUser", user).Return(nil)

	if assert.NoError(t, controller.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var response domain.UserResponse
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "John Doe", response.Name)
	}
}

func TestCreateUser_InvalidData(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name":"","email":"invalid_email","image_path":"invalid_url"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserService)
	controller := &controller.UserController{UserService: mockService}

	if assert.NoError(t, controller.CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
	}
}

func TestGetUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockService := new(MockUserService)
	controller := &controller.UserController{UserService: mockService}

	user := &entities.User{ID: 1, Name: "John Doe", Email: "john.doe@example.com", ImagePath: "http://example.com/image.jpg"}
	mockService.On("GetUserByID", 1).Return(user, nil)

	if assert.NoError(t, controller.GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response domain.UserResponse
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "John Doe", response.Name)
	}
}

func TestGetUser_InvalidID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/invalid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	mockService := new(MockUserService)
	controller := &controller.UserController{UserService: mockService}

	if assert.NoError(t, controller.GetUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Invalid user ID", response["message"])
	}
}

func TestGetUser_NotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockService := new(MockUserService)
	controller := &controller.UserController{UserService: mockService}

	mockService.On("GetUserByID", 1).Return(&entities.User{}, assert.AnError)
	if assert.NoError(t, controller.GetUser(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "User not found", response["message"])
	}
}

func TestListUsers(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users?limit=10&offset=0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserService)
	controller := &controller.UserController{UserService: mockService}

	users := []entities.User{
		{ID: 1, Name: "John Doe", Email: "john.doe@example.com", ImagePath: "http://example.com/image.jpg"},
		{ID: 2, Name: "Jane Doe", Email: "jane.doe@example.com", ImagePath: "http://example.com/image2.jpg"},
	}
	mockService.On("ListUsers", 10, 0).Return(users, nil)

	if assert.NoError(t, controller.ListUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response []domain.UserResponse
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Len(t, response, 2)
	}
}

func TestListUsers_Error(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users?limit=10&offset=0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserService)
	controller := &controller.UserController{UserService: mockService}

	users := make([]entities.User, 0)
	mockService.On("ListUsers", 10, 0).Return(users, assert.AnError)

	if assert.NoError(t, controller.ListUsers(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Failed to list users", response["message"])
	}
}

func TestDeleteUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockService := new(MockUserService)
	controller := &controller.UserController{UserService: mockService}

	mockService.On("DeleteUser", 1).Return(nil)

	if assert.NoError(t, controller.DeleteUser(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func TestDeleteUser_NotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockService := new(MockUserService)
	controller := &controller.UserController{UserService: mockService}

	mockService.On("DeleteUser", 1).Return(assert.AnError)

	if assert.NoError(t, controller.DeleteUser(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "User not found", response["message"])
	}
}
