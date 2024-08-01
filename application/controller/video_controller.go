package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"net/http"
	"project/application/core/domain"
	"project/application/core/ports"
	"project/application/mappers"
	"strconv"
)

type VideoController struct {
	Service ports.IVideoService
}

// RegisterVideoRoutes registra las rutas para los endpoints de video.
func RegisterVideoRoutes(e *echo.Echo, service ports.IVideoService) {
	handler := &VideoController{service}
	e.POST("/videos", handler.CreateVideo)
	e.GET("/videos/:id", handler.GetVideo)
	e.GET("/videos", handler.ListVideos)
	e.DELETE("/videos/:id", handler.DeleteVideo) // Endpoint DELETE
}

// CreateVideo crea un nuevo video.
func (h *VideoController) CreateVideo(c echo.Context) error {
	req := new(domain.CreateVideoRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	video := mappers.ToVideoEntity(req)
	if err := h.Service.CreateVideo(video); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, mappers.ToVideoResponse(video))
}

// GetVideo obtiene un video por ID.
func (h *VideoController) GetVideo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid video ID"})
	}

	video, err := h.Service.GetVideoByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Video not found"})
	}

	return c.JSON(http.StatusOK, mappers.ToVideoResponse(video))
}

// ListVideos lista los videos con paginación.
func (h *VideoController) ListVideos(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	videos, err := h.Service.ListVideos(limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to list videos"})
	}

	return c.JSON(http.StatusOK, mappers.ToVideoResponses(videos))
}

// DeleteVideo elimina un video por ID.
func (h *VideoController) DeleteVideo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid video ID"})
	}

	if err := h.Service.DeleteVideo(id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Video not found"})
	}

	return c.NoContent(http.StatusNoContent)
}
