package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"net/http"
	"project/application/core/domain"
	"project/application/mappers"
	"project/application/services"
	"strconv"
)

type VideoController struct {
	service *services.VideoService
}

func RegisterVideoRoutes(e *echo.Echo, service *services.VideoService) {
	handler := &VideoController{service}
	e.POST("/videos", handler.CreateVideo)
	e.GET("/videos/:id", handler.GetVideo)
	e.GET("/videos", handler.ListVideos)
}

func (h *VideoController) CreateVideo(c echo.Context) error {
	req := new(domain.CreateVideoRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	video := mappers.ToVideoEntity(req)
	if err := h.service.CreateVideo(video); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, mappers.ToVideoResponse(video))
}

func (h *VideoController) GetVideo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid video ID"})
	}

	video, err := h.service.GetVideoByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Video not found"})
	}

	return c.JSON(http.StatusOK, mappers.ToVideoResponse(video))
}

func (h *VideoController) ListVideos(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	videos, err := h.service.ListVideos(limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to list videos"})
	}

	return c.JSON(http.StatusOK, mappers.ToVideoResponses(videos))
}
