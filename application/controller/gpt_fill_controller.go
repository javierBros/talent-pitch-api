package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project/application/services"
)

type GPTFillController struct {
	service *services.GPTFillService
}

func RegisterGPTFillRoutes(e *echo.Echo, service *services.GPTFillService) {
	handler := &GPTFillController{service}
	e.POST("/fill-tables", handler.FillTables)
}

func (h *GPTFillController) FillTables(c echo.Context) error {
	if err := h.service.FillTables(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to fill tables"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "tables filled successfully"})
}
