package templates

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func GetLogin() elem.Node {
	return elem.Section(
		attrs.Props{
			attrs.Class: "hero is-info is-fullheight is-widescreen",
		},
		elem.Div(
			attrs.Props{
				attrs.Class: "hero-body",
			},
			elem.Div(
				attrs.Props{
					attrs.Class: "container",
				},
				elem.Div(
					attrs.Props{
						attrs.Class: "columns is-centered",
					},
					elem.Div(
						attrs.Props{
							attrs.Class: "column is-5-tablet is-4-desktop is-3-widescreen",
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
										attrs.Class: "tile is-child has-text-centered",
									},
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
														attrs.Type:        "email",
														attrs.Placeholder: "Email",
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
														attrs.Type:        "password",
														attrs.Placeholder: "Password",
													},
												),
											),
										),
										elem.Button(
											attrs.Props{
												attrs.Class: "button is-block is-info is-fullwidth is-medium",
											},
											elem.Text("anmelden"),
										),
									),
								),
							),
						),
					),
				),
			),
		),
	)
}

func RenderLogin(login elem.Node) string {
	return login.Render()
}
