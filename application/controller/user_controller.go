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

type UserController struct {
	service *services.UserService
}

func RegisterUserRoutes(e *echo.Echo, service *services.UserService) {
	handler := &UserController{service}
	e.POST("/users", handler.CreateUser)
	e.GET("/users/:id", handler.GetUser)
	e.GET("/users", handler.ListUsers)
}

func (h *UserController) CreateUser(c echo.Context) error {
	req := new(domain.CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	user := mappers.ToUserEntity(req)
	if err := h.service.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, mappers.ToUserResponse(user))
}

func (h *UserController) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}

	return c.JSON(http.StatusOK, mappers.ToUserResponse(user))
}

func (h *UserController) ListUsers(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	users, err := h.service.ListUsers(limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to list users"})
	}

	return c.JSON(http.StatusOK, mappers.ToUserResponses(users))
}
