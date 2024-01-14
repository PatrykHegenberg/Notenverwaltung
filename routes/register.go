package routes

import (
	"net/http"

	"github.com/PatrykHegenberg/Notenverwaltung/templates"
	"github.com/labstack/echo/v4"
)

// GetRegisterHandler is a function that handles the register endpoint.
//
// It takes in a context object of the echo framework.
// It returns an error object.
func GetRegisterHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, templates.RenderRegister(templates.GetRegister()))
}
