package routes

import (
	"net/http"

	"github.com/PatrykHegenberg/Notenverwaltung/templates"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetIndexHandler(c echo.Context) error {
	sess, _ := session.Get("authenticate-session", c)
	if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
		return c.HTML(http.StatusOK, templates.RenderIndex(false))
	} else {
		return c.HTML(http.StatusOK, templates.RenderIndex(true))
	}
}
