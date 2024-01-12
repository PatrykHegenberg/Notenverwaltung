package templates

import (
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func GetDashboard(user model.User, role model.Role, school model.School, users []model.User) elem.Node {
	dashboard := elem.Section(attrs.Props{attrs.Class: "hero is-info is-fullheight"},
		elem.Div(attrs.Props{attrs.Class: "hero-body tile is-ancestor is-vertical"},
			elem.Div(attrs.Props{attrs.Class: "box"},
				GetUserInfo(user, role, school),
				GetUsers(users),
				GetClasses(),
				GetSchools(),
			),
		),
	)
	return dashboard
}

func RenderDashboard(dashboard elem.Node) string {
	return dashboard.Render()
}
