package templates

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func GetRegister() elem.Node {
	register := elem.Section(
		attrs.Props{
			attrs.Class: "hero is-info is-fullheight",
		},
		elem.Div(
			attrs.Props{
				attrs.Class: "container column is-centered",
			},
			elem.Div(
				attrs.Props{
					attrs.Class: "tile box is-ancestor",
				},
				elem.Div(
					attrs.Props{
						attrs.Class: "tile is-parent",
					},
					elem.Div(
						attrs.Props{
							attrs.Class: "tile is-child",
						},
						elem.H1(nil,
							elem.Text("Notenverwaltung"),
						),
						elem.H2(nil,
							elem.Text("Zentralisierte Verwaltung der Noten aller Schüler."),
						),
					),
					elem.Div(
						attrs.Props{
							attrs.Class: "tile is-child has-text-centered",
						},
						elem.H1(nil,
							elem.Text("Registriere dich noch heute."),
						),
						elem.Form(nil,
							elem.Div(
								attrs.Props{
									attrs.Class: "field",
								},
								elem.Div(
									attrs.Props{
										attrs.Class: "control",
									},
									elem.Input(
										attrs.Props{
											attrs.Class:       "input",
											attrs.Type:        "text",
											attrs.Placeholder: "Name",
										},
									),
								),
							),
							elem.Div(
								attrs.Props{
									attrs.Class: "field",
								},
								elem.Div(
									attrs.Props{
										attrs.Class: "control",
									},
									elem.Input(
										attrs.Props{
											attrs.Class:       "input",
											attrs.Type:        "email",
											attrs.Placeholder: "Email",
										},
									),
								),
							),
							elem.Button(
								attrs.Props{
									attrs.Class: "button is-block is-info is-fullwidth is-medium",
								},
								elem.Text("Registrieren"),
							),
						),
					),
				),
			),
		),
	)
	return register
}

func RenderRegister(register elem.Node) string {
	return register.Render()
}
