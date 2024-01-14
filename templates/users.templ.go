package templates

import (
	"fmt"

	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func GetUsers(users []model.User) elem.Node {
	usersTable := elem.TransformEach(users, func(user model.User) elem.Node {
		return elem.Tr(nil,
			elem.Td(nil, elem.Text(user.Username)),
			elem.Td(nil, elem.Text(user.Email)),
			elem.Td(nil, elem.Text(fmt.Sprintf("%v", user.UpdatedAt))),
			elem.Td(nil, elem.Text(fmt.Sprintf("%v", user.DeletedAt))),
			elem.Td(nil, elem.Button(attrs.Props{attrs.Class: "button"}, elem.Text("Beartbeiten"))),
			elem.Td(nil, elem.Button(attrs.Props{attrs.Class: "button"}, elem.Text("Löschen"))),
		)
	})
	return elem.Div(attrs.Props{attrs.Class: "tile is-parent"},
		elem.Div(attrs.Props{attrs.Class: "tile is-child is-ancestor"},
			elem.Div(attrs.Props{attrs.Class: "tile is-parent"}),
			elem.Table(attrs.Props{attrs.Class: "table"},
				elem.THead(nil,
					elem.Tr(nil,
						elem.Th(nil, elem.Text("Username")),
						elem.Th(nil, elem.Text("Email")),
						elem.Th(nil, elem.Text("UpdatedAt")),
						elem.Th(nil, elem.Text("DeletedAt")),
						elem.Th(nil, elem.Text("")),
						elem.Th(nil, elem.Text("")),
					),
				),
				elem.TBody(nil, usersTable...),
			),
		),
	)
}
