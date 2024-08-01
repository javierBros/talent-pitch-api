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

type MockChallengeService struct {
	mock.Mock
}

func (m *MockChallengeService) CreateChallenge(challenge *entities.Challenge) error {
	args := m.Called(challenge)
	return args.Error(0)
}

func (m *MockChallengeService) GetChallengeByID(id int) (*entities.Challenge, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Challenge), args.Error(1)
}

func (m *MockChallengeService) ListChallenges(limit, offset int) ([]entities.Challenge, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]entities.Challenge), args.Error(1)
}

func (m *MockChallengeService) DeleteChallenge(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateChallenge(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/challenges", strings.NewReader(`{"title":"Challenge 1","description":"Description 1","difficulty":1,"userId":1}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockChallengeService)
	controller := &controller.ChallengeController{Service: mockService}

	challenge := &entities.Challenge{Title: "Challenge 1", Description: "Description 1", Difficulty: 1, UserID: 1}
	mockService.On("CreateChallenge", challenge).Return(nil)

	if assert.NoError(t, controller.CreateChallenge(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var response domain.ChallengeResponse
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Challenge 1", response.Title)
	}
}

func TestCreateChallenge_InvalidData(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/challenges", strings.NewReader(`{"title":"","description":"","difficulty":0,"user_id":0}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockChallengeService)
	controller := &controller.ChallengeController{Service: mockService}

	if assert.NoError(t, controller.CreateChallenge(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
	}
}

func TestGetChallenge(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/challenges/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockService := new(MockChallengeService)
	controller := &controller.ChallengeController{Service: mockService}

	challenge := &entities.Challenge{ID: 1, Title: "Challenge 1", Description: "Description 1", Difficulty: 1, UserID: 1}
	mockService.On("GetChallengeByID", 1).Return(challenge, nil)

	if assert.NoError(t, controller.GetChallenge(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response domain.ChallengeResponse
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Challenge 1", response.Title)
	}
}

func TestGetChallenge_InvalidID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/challenges/invalid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	mockService := new(MockChallengeService)
	controller := &controller.ChallengeController{Service: mockService}

	if assert.NoError(t, controller.GetChallenge(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Invalid challenge ID", response["message"])
	}
}

func TestGetChallenge_NotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/challenges/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockService := new(MockChallengeService)
	controller := &controller.ChallengeController{Service: mockService}

	mockService.On("GetChallengeByID", 1).Return(&entities.Challenge{}, assert.AnError)
	if assert.NoError(t, controller.GetChallenge(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Challenge not found", response["message"])
	}
}

func TestListChallenges(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/challenges?limit=10&offset=0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockChallengeService)
	controller := &controller.ChallengeController{Service: mockService}

	challenges := []entities.Challenge{
		{ID: 1, Title: "Challenge 1", Description: "Description 1", Difficulty: 1, UserID: 1},
		{ID: 2, Title: "Challenge 2", Description: "Description 2", Difficulty: 2, UserID: 2},
	}
	mockService.On("ListChallenges", 10, 0).Return(challenges, nil)

	if assert.NoError(t, controller.ListChallenges(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response []domain.ChallengeResponse
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Len(t, response, 2)
	}
}

func TestListChallenges_Error(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/challenges?limit=10&offset=0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockChallengeService)
	controller := &controller.ChallengeController{Service: mockService}

	challenges := make([]entities.Challenge, 0)
	mockService.On("ListChallenges", 10, 0).Return(challenges, assert.AnError)

	if assert.NoError(t, controller.ListChallenges(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Failed to list challenges", response["message"])
	}
}

func TestDeleteChallenge(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/challenges/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockService := new(MockChallengeService)
	controller := &controller.ChallengeController{Service: mockService}

	mockService.On("DeleteChallenge", 1).Return(nil)

	if assert.NoError(t, controller.DeleteChallenge(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func TestDeleteChallenge_NotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/challenges/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockService := new(MockChallengeService)
	controller := &controller.ChallengeController{Service: mockService}

	mockService.On("DeleteChallenge", 1).Return(assert.AnError)

	if assert.NoError(t, controller.DeleteChallenge(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Challenge not found", response["message"])
	}
}
