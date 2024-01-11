package routes

import (
	"github.com/PatrykHegenberg/Notenverwaltung/templates"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetIndexHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, templates.RenderIndex())
}
