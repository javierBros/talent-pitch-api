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

type UserController struct {
	UserService ports.IUserService
}

// RegisterUserRoutes registra las rutas para los endpoints de usuario.
func RegisterUserRoutes(e *echo.Echo, service ports.IUserService) {
	handler := &UserController{service}
	e.POST("/users", handler.CreateUser)
	e.GET("/users/:id", handler.GetUser)
	e.GET("/users", handler.ListUsers)
	e.DELETE("/users/:id", handler.DeleteUser) // Endpoint DELETE
}

// CreateUser crea un nuevo usuario.
func (h *UserController) CreateUser(c echo.Context) error {
	req := new(domain.CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	user := mappers.ToUserEntity(req)
	if err := h.UserService.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, mappers.ToUserResponse(user))
}

// GetUser obtiene un usuario por ID.
func (h *UserController) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}

	return c.JSON(http.StatusOK, mappers.ToUserResponse(user))
}

// ListUsers lista los usuarios con paginación.
func (h *UserController) ListUsers(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	users, err := h.UserService.ListUsers(limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to list users"})
	}

	return c.JSON(http.StatusOK, mappers.ToUserResponses(users))
}

// DeleteUser elimina un usuario por ID.
func (h *UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	if err := h.UserService.DeleteUser(id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}

	return c.NoContent(http.StatusNoContent)
}
