package templates

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func RenderIndex(loggedIn bool) string {
	doc := elem.Html(nil,
		elem.Head(nil,
			elem.Meta(attrs.Props{attrs.Charset: "UTF-8"}),
			elem.Meta(attrs.Props{attrs.Name: "viewport", attrs.Content: "width=device-width, initial-scale=1"}),
			elem.Title(nil, elem.Text("Notenverwaltung")),
			// Bulma CSS
			elem.Link(attrs.Props{attrs.Rel: "stylesheet", attrs.Href: "https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"}),
			// htmx JS
			elem.Script(attrs.Props{attrs.Src: "https://unpkg.com/htmx.org@1.9.10", attrs.Integrity: "sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC", attrs.Crossorigin: "anonymous"}),
		),
		elem.Body(nil,
			elem.Div(attrs.Props{attrs.ID: "outer"},
				GetNavbar(loggedIn),
				elem.Main(attrs.Props{attrs.Class: "main"},
					elem.Div(attrs.Props{attrs.ID: "content-div"},
						HeroIndex(),
					),
				),
				GetFooter(),
			),
		),
	)

	// Ausgabe
	return doc.Render()
}
