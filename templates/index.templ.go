package templates

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func RenderIndex() string {
	// Hero Section
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

	// HTML Dokument
	doc := elem.Html(nil,
		elem.Head(nil,
			elem.Meta(attrs.Props{
				attrs.Charset: "UTF-8",
			},
			),
			elem.Meta(
				attrs.Props{
					attrs.Name:    "viewport",
					attrs.Content: "width=device-width, initial-scale=1",
				},
			),
			elem.Title(nil, elem.Text("Notenverwaltung")),
			// Bulma CSS
			elem.Link(
				attrs.Props{
					attrs.Rel:  "stylesheet",
					attrs.Href: "https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css",
				},
			),
			// htmx JS
			elem.Script(
				attrs.Props{
					attrs.Src:         "https://unpkg.com/htmx.org@1.9.10",
					attrs.Integrity:   "sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC",
					attrs.Crossorigin: "anonymous",
				},
			),
		),
		elem.Body(nil,
			GetNavbar(),
			elem.Main(
				attrs.Props{
					attrs.Class: "main",
				},
				elem.Div(
					attrs.Props{
						attrs.ID: "content-div",
					},
					heroSection,
				),
			),
			GetFooter(),
		),
	)

	// Ausgabe
	return doc.Render()
}
