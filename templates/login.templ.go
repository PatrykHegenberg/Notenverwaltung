package templates

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/htmx"
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
									elem.Form(
										attrs.Props{
											attrs.Method: "post",
											attrs.Action: "/authenticate",
										},
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
														attrs.Name:        "email",
														attrs.ID:          "email",
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
														attrs.Name:        "password",
														attrs.ID:          "password",
													},
												),
											),
										),
										elem.Button(
											attrs.Props{
												attrs.Type:     "submit",
												attrs.Class:    "button is-block is-info is-fullwidth is-medium",
												attrs.Value:    "Anmelden",
												htmx.HXPost:    "/authenticate",
												htmx.HXTrigger: "onClick",
											},
											elem.Text("Anmelden"),
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
