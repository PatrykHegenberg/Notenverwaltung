package routes

import (
	"net/http"

	"github.com/PatrykHegenberg/Notenverwaltung/templates"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// GetIndexHandler handles the request for the index page.
//
// It receives a context object (c) of type `echo.Context` as a parameter.
// It returns an error.
func GetIndexHandler(c echo.Context) error {
	sess, _ := session.Get("authenticate-session", c)
	if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
		return c.HTML(http.StatusOK, templates.RenderIndex(false, c))
	} else {
		return c.HTML(http.StatusOK, templates.RenderIndex(true, c))
	}
}
