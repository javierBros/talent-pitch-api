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

type ChallengeController struct {
	Service ports.IChallengeService
}

// RegisterChallengeRoutes registra las rutas para los endpoints de desafío.
func RegisterChallengeRoutes(e *echo.Echo, service ports.IChallengeService) {
	handler := &ChallengeController{service}
	e.POST("/challenges", handler.CreateChallenge)
	e.GET("/challenges/:id", handler.GetChallenge)
	e.GET("/challenges", handler.ListChallenges)
	e.DELETE("/challenges/:id", handler.DeleteChallenge)
}

// CreateChallenge crea un nuevo desafío.
func (h *ChallengeController) CreateChallenge(c echo.Context) error {
	req := new(domain.CreateChallengeRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	challenge := mappers.ToChallengeEntity(req)
	if err := h.Service.CreateChallenge(challenge); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, mappers.ToChallengeResponse(challenge))
}

// GetChallenge obtiene un desafío por ID.
func (h *ChallengeController) GetChallenge(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid challenge ID"})
	}

	challenge, err := h.Service.GetChallengeByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Challenge not found"})
	}

	return c.JSON(http.StatusOK, mappers.ToChallengeResponse(challenge))
}

// ListChallenges lista los desafíos con paginación.
func (h *ChallengeController) ListChallenges(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	challenges, err := h.Service.ListChallenges(limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to list challenges"})
	}

	return c.JSON(http.StatusOK, mappers.ToChallengeResponses(challenges))
}

// DeleteChallenge elimina un desafío por ID.
func (h *ChallengeController) DeleteChallenge(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid challenge ID"})
	}

	if err := h.Service.DeleteChallenge(id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Challenge not found"})
	}

	return c.NoContent(http.StatusNoContent)
}
