package controller_test

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/talent-pitch-api/application/controller"
	"github.com/talent-pitch-api/application/core/domain"
	"github.com/talent-pitch-api/application/core/entities"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockVideoService struct {
	mock.Mock
}

func (m *MockVideoService) CreateVideo(video *entities.Video) error {
	args := m.Called(video)
	return args.Error(0)
}

func (m *MockVideoService) GetVideoByID(id int) (*entities.Video, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Video), args.Error(1)
}

func (m *MockVideoService) ListVideos(limit, offset int) ([]entities.Video, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]entities.Video), args.Error(1)
}

func (m *MockVideoService) DeleteVideo(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateVideo(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/videos", strings.NewReader(`{"title":"Video 1","description":"Description 1","url":"http://example.com/video1","userId":1}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockVideoService)
	controller := &controller.VideoController{Service: mockService}

	video := &entities.Video{Title: "Video 1", Description: "Description 1", URL: "http://example.com/video1", UserID: 1}
	mockService.On("CreateVideo", video).Return(nil)

	if assert.NoError(t, controller.CreateVideo(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var response domain.VideoResponse
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Video 1", response.Title)
	}
}

func TestCreateVideo_InvalidData(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/videos", strings.NewReader(`{"title":"","description":"","url":"","user_id":0}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockVideoService)
	controller := &controller.VideoController{Service: mockService}

	if assert.NoError(t, controller.CreateVideo(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
	}
}

func TestGetVideo(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/videos/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockService := new(MockVideoService)
	controller := &controller.VideoController{Service: mockService}

	video := &entities.Video{ID: 1, Title: "Video 1", Description: "Description 1", URL: "http://example.com/video1", UserID: 1}
	mockService.On("GetVideoByID", 1).Return(video, nil)

	if assert.NoError(t, controller.GetVideo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response domain.VideoResponse
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Video 1", response.Title)
	}
}

func TestGetVideo_InvalidID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/videos/invalid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	mockService := new(MockVideoService)
	controller := &controller.VideoController{Service: mockService}

	if assert.NoError(t, controller.GetVideo(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Invalid video ID", response["message"])
	}
}

func TestGetVideo_NotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/videos/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockService := new(MockVideoService)
	controller := &controller.VideoController{Service: mockService}

	mockService.On("GetVideoByID", 1).Return(&entities.Video{}, assert.AnError)
	if assert.NoError(t, controller.GetVideo(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Video not found", response["message"])
	}
}

func TestListVideos(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/videos?limit=10&offset=0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockVideoService)
	controller := &controller.VideoController{Service: mockService}

	videos := []entities.Video{
		{ID: 1, Title: "Video 1", Description: "Description 1", URL: "http://example.com/video1", UserID: 1},
		{ID: 2, Title: "Video 2", Description: "Description 2", URL: "http://example.com/video2", UserID: 2},
	}
	mockService.On("ListVideos", 10, 0).Return(videos, nil)

	if assert.NoError(t, controller.ListVideos(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response []domain.VideoResponse
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Len(t, response, 2)
	}
}

func TestListVideos_Error(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/videos?limit=10&offset=0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockVideoService)
	controller := &controller.VideoController{Service: mockService}

	videos := make([]entities.Video, 0)
	mockService.On("ListVideos", 10, 0).Return(videos, assert.AnError)

	if assert.NoError(t, controller.ListVideos(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Failed to list videos", response["message"])
	}
}

func TestDeleteVideo(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/v1/videos/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockService := new(MockVideoService)
	controller := &controller.VideoController{Service: mockService}

	mockService.On("DeleteVideo", 1).Return(nil)

	if assert.NoError(t, controller.DeleteVideo(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func TestDeleteVideo_NotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/v1/videos/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockService := new(MockVideoService)
	controller := &controller.VideoController{Service: mockService}

	mockService.On("DeleteVideo", 1).Return(assert.AnError)

	if assert.NoError(t, controller.DeleteVideo(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Video not found", response["message"])
	}
}
