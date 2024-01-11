package templates

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func HeroIndex() elem.Node {
	heroSection := elem.Section(
		attrs.Props{
			attrs.Class: "hero is-info is-fullheight",
		},
		elem.Div(
			attrs.Props{
				attrs.Class: "hero-body has-text-centered",
			},
			elem.Div(attrs.Props{
				attrs.Class: "container",
			},
				elem.H1(attrs.Props{attrs.Class: "title"}, elem.Text("Willkommen bei der Notenverwaltung")),
				elem.H2(attrs.Props{attrs.Class: "subtitle"}, elem.Text("Diese WebApp ermöglicht die interaktive Verwaltung von Klassen, Noten und Schülern.")),
			),
		),
	)
	return heroSection
}
