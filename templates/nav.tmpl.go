package templates

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/htmx"
)

func GetNavbar() elem.Node {
	navBar := elem.Nav(
		attrs.Props{
			attrs.Class:     "navbar",
			attrs.Role:      "navigation",
			attrs.AriaLabel: "main navigation",
		},
		elem.Div(
			attrs.Props{
				attrs.Class: "navbar-brand",
			},
			elem.A(
				attrs.Props{
					attrs.Class: "navbar-item",
					attrs.Href:  "",
				},
				elem.Img(attrs.Props{
					attrs.Src:    "https://bulma.io/images/bulma-logo.png",
					attrs.Width:  "112",
					attrs.Height: "28",
				},
				),
				elem.A(
					attrs.Props{
						attrs.Role:         "button",
						attrs.Class:        "navbar-burger",
						attrs.AriaLabel:    "menu",
						attrs.AriaExpanded: "false",
					},
					elem.Span(
						attrs.Props{
							attrs.AriaHidden: "true",
						},
					),
					elem.Span(
						attrs.Props{
							attrs.AriaHidden: "true",
						},
					),
					elem.Span(
						attrs.Props{
							attrs.AriaHidden: "true",
						},
					),
				),
			),
		),
		elem.Div(
			attrs.Props{
				attrs.ID:    "navbarBasicExample",
				attrs.Class: "navbar-menu",
			},
			elem.Div(
				attrs.Props{
					attrs.Class: "navbar-start",
				},
				elem.A(attrs.Props{
					attrs.Class: "navbar-item",
				},
					elem.Text("Home")),
				elem.A(attrs.Props{
					attrs.Class: "navbar-item",
				},
					elem.Text("Documentation")),
				elem.A(attrs.Props{
					attrs.Class: "navbar-item",
				},
					elem.Text("Meine Klassen")),
			),
			elem.Div(
				attrs.Props{
					attrs.Class: "navbar-end",
				},
				elem.Div(
					attrs.Props{
						attrs.Class: "navbar-item",
					},
					elem.Div(
						attrs.Props{
							attrs.Class: "buttons",
						},
						elem.A(
							attrs.Props{
								attrs.Class:   "button is-info",
								htmx.HXGet:    "/register",
								htmx.HXTarget: "#content-div",
								htmx.HXSwap:   "outerHtml",
							},
							elem.Strong(nil, elem.Text("Registrieren")),
						),
						elem.A(
							attrs.Props{
								attrs.Class: "button",
							},
							elem.Text("Anmelden"),
						),
					),
				),
			),
		),
	)
	return navBar
}
