package templates

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func GetClasses() elem.Node {
	return elem.Div(attrs.Props{attrs.Class: "title"},
		elem.Text("Dies ist der Classesbereich."),
	)
}
