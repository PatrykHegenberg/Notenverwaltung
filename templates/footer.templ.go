package templates

import (
	"time"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func GetFooter() elem.Node {
	footer := elem.Footer(attrs.Props{attrs.Class: "footer"},
		elem.Div(attrs.Props{attrs.Class: "content has-text-centered"},
			elem.P(nil, elem.Strong(nil, elem.Text("Notenverwaltung")),
				elem.Text(" © "+time.Now().Format("2006")+" by Rudi Gola, Patryk Hegenberg, Manuel Keidel, Alina Kudrina.\n Alle Rechte vorbehalten."),
			),
		),
	)
	return footer
}
