package routes

import (
	"net/http"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/templates"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetDashboardHandler(c echo.Context) error {
	sess, _ := session.Get("authenticate-session", c)
	if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
		return c.Redirect(http.StatusSeeOther, "/")
	} else {
		username := sess.Values["username"].(string)
		user, err := DB.GetUserByName(username)
		if err != nil {
			log.Error("Fehler")
		}
		school, err := DB.GetSchoolById(user.SchoolID)
		if err != nil {
			log.Error("Fehler")
		}
		users, err := DB.GetAllUsers()
		if err != nil {
			log.Error(err)
		}
		dashboardHTML := templates.GetDashboard(*user, *school, users).Render()
		log.Print(dashboardHTML)
		return c.HTML(http.StatusOK, dashboardHTML)
	}
}
