package templates

import (
	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/htmx"
	"github.com/labstack/gommon/log"
)

func GetRegister() elem.Node {
	var (
		schools []model.School
	)

	db := DB.GetDBInstance()

	if err := db.Find(&schools).Error; err != nil {
		log.Errorf("error getting schools from db")
	}

	schoolOptions := elem.TransformEach(schools, func(school model.School) elem.Node {
		return elem.Option(nil, elem.Text(school.Name))
	})
	schoolSelect := elem.Select(
		attrs.Props{
			attrs.Class:       "is-1",
			attrs.Name:        "school",
			attrs.ID:          "school",
			attrs.Placeholder: "Wählen Sie ihre Schule.",
		}, schoolOptions...)

	register := elem.Section(attrs.Props{attrs.Class: "hero is-info is-fullheight"},
		elem.Div(attrs.Props{attrs.Class: "hero-body"},
			elem.Div(attrs.Props{attrs.Class: "container column is-centered"},
				elem.Div(attrs.Props{attrs.Class: "tile box is-ancestor"},
					elem.Div(attrs.Props{attrs.Class: "tile is-parent"},
						elem.Div(attrs.Props{attrs.Class: "tile is-child"},
							elem.P(attrs.Props{attrs.Class: "title has-text-black is-3"},
								elem.Text("Notenverwaltung")),
							elem.P(attrs.Props{attrs.Class: "subtitle has-text-black is-5"},
								elem.Text("Zentralisierte Verwaltung der Noten aller Schüler."),
							),
						),
						elem.Div(attrs.Props{attrs.Class: "tile is-child has-text-centered"},
							elem.H1(nil, elem.Text("Registriere dich noch heute.")),
							elem.Form(
								attrs.Props{
									htmx.HXPost:    "/users",
									htmx.HXTrigger: "click",
								},
								elem.Div(attrs.Props{attrs.Class: "field"},
									elem.Div(attrs.Props{attrs.Class: "control"},
										elem.Input(attrs.Props{
											attrs.Class:       "input",
											attrs.ID:          "username",
											attrs.Name:        "username",
											attrs.Type:        "text",
											attrs.Placeholder: "Username",
										},
										),
									),
								),
								elem.Div(attrs.Props{attrs.Class: "field"},
									elem.Div(attrs.Props{attrs.Class: "control"},
										elem.Input(attrs.Props{
											attrs.Class:       "input",
											attrs.Type:        "email",
											attrs.ID:          "email",
											attrs.Name:        "email",
											attrs.Placeholder: "Email",
										},
										),
									),
								),
								elem.Div(attrs.Props{attrs.Class: "field"},
									elem.Div(attrs.Props{attrs.Class: "control"},
										elem.Input(attrs.Props{
											attrs.Class:       "input",
											attrs.Type:        "password",
											attrs.Placeholder: "Passwort",
											attrs.ID:          "password",
											attrs.Name:        "password",
										},
										),
									),
								),
								elem.Div(attrs.Props{attrs.Class: "field"},
									elem.Div(attrs.Props{attrs.Class: "control"},
										elem.Input(attrs.Props{
											attrs.Class:       "input",
											attrs.Type:        "password",
											attrs.Placeholder: "Passwort bestätigen",
										},
										),
									),
								),
								elem.Div(attrs.Props{attrs.Class: "field"},
									elem.Div(attrs.Props{attrs.Class: "control"},
										elem.Div(attrs.Props{attrs.Class: "select"}, schoolSelect),
									),
								),
								elem.Button(attrs.Props{attrs.Class: "button is-block is-info is-fullwidth is-medium"},
									elem.Text("Registrieren"),
								),
							),
						),
					),
				),
			),
		),
	)
	log.Print(register)
	return register
}

func RenderRegister(register elem.Node) string {
	return register.Render()
}
