package routes

import (
	"net/http"

	"github.com/PatrykHegenberg/Notenverwaltung/templates"
	"github.com/labstack/echo/v4"
)

// GetLoginHandler is a Go function that handles the login request.
//
// It takes a parameter "c" of type echo.Context which represents the request context.
// It returns an error.
func GetLoginHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, templates.RenderLogin(templates.GetLogin()))
}
