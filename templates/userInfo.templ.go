package templates

import (
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func GetUserInfo(user model.User, role model.Role, school model.School) elem.Node {
	return elem.Div(attrs.Props{attrs.Class: "tile is-parent"},
		elem.Div(attrs.Props{attrs.Class: "tile is-child is-ancestor"},
			elem.Div(attrs.Props{attrs.Class: "tile is-parent"},
				elem.P(attrs.Props{attrs.Class: "tile is-child"}, elem.Text("Username: ")),
				elem.P(attrs.Props{attrs.Class: "tile is-child"}, elem.Text(user.Username)),
			),
			elem.Div(attrs.Props{attrs.Class: "tile is-parent"},
				elem.P(attrs.Props{attrs.Class: "tile is-child"}, elem.Text("E-Mail: ")),
				elem.P(attrs.Props{attrs.Class: "tile is-child"}, elem.Text(user.Email)),
			),
			elem.Div(attrs.Props{attrs.Class: "tile is-parent"},
				elem.P(attrs.Props{attrs.Class: "tile is-child"}, elem.Text("Rolle: ")),
				elem.P(attrs.Props{attrs.Class: "tile is-child"}, elem.Text(role.Name)),
			),
			elem.Div(attrs.Props{attrs.Class: "tile is-parent"},
				elem.P(attrs.Props{attrs.Class: "tile is-child"}, elem.Text("Schule: ")),
				elem.P(attrs.Props{attrs.Class: "tile is-child"}, elem.Text(school.Name)),
			),
		),
	)
}
