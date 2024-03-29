package main

import (
	// "net/http"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/routes"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"embed"

	_ "github.com/PatrykHegenberg/Notenverwaltung/docs"
)

var (
	//go:embed frontend
	fs embed.FS
)

// @title Notenverwaltung API
// @version 0.0.8
// @descritption This is a simple API server for Notenverwaltung
// @termsOfService http://swagger.io/terms

// @contact.name API support
// @contact.url http://swagger.io/support
// @contact.email patrykhegenberg@gmail.com

// @license.name Apache 2.0
// @license.url http://apache.com/licenses/LICENSE-2.0.html

// @securityDefinitions.basic BasicAuth
// @host localhost:1323
// @BasePath /api/v1
func main() {
	DB.AutoMigrate()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("my-secret"))))
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	// 	HTML5:      true,
	// 	Root:       "frontend",
	// 	Filesystem: http.FS(fs),
	// }))
	e.POST("/auth", routes.AuthenticateUserHandler)
	e.POST("/signup", routes.CreateUserHandler)

	// API Routes
	apiGroup := e.Group("/api/v1")
	apiGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		return checkCredentials(username, password), nil
	}))
	studentGroup := apiGroup.Group("/students")
	configureStudentRoutes(studentGroup)

	userGroup := apiGroup.Group("/users")
	configureUserRoutes(userGroup)

	examGroup := apiGroup.Group("/exams")
	configureExamRoutes(examGroup)

	classGroup := apiGroup.Group("/classes")
	configureClassRoutes(classGroup)

	scoreGroup := apiGroup.Group("/scores")
	configureScoreRoutes(scoreGroup)

	addressesGroup := apiGroup.Group("/addresses")
	configureAddressRoutes(addressesGroup)

	gradesGroup := apiGroup.Group("/grades")
	configureGradeRoutes(gradesGroup)

	examTypeGroup := apiGroup.Group("/examTypes")
	configureExamTypeRoutes(examTypeGroup)

	schoolGroup := apiGroup.Group("/schools")
	configureSchoolRoutes(schoolGroup)

	subjectGroup := apiGroup.Group("/subjects")
	configureSubjectRoutes(subjectGroup)

	subjectUserClassExamGroup := apiGroup.Group("/suce")
	configureSubjectUserClassEexamRoutes(subjectUserClassExamGroup)

	e.Logger.Fatal(e.Start(":1323"))
}

func configureStudentRoutes(g *echo.Group) {
	g.GET("", routes.GetStudentsHandler)
	g.GET("/:id", routes.GetStudentHandler)
	g.GET("/class/:id", routes.GetStudentsByClassHandler)
	g.POST("", routes.CreateStudentHandler)
	g.PUT("/:id", routes.UpdateStudentHandler)
	g.DELETE("/:id", routes.DeleteStudentHandler)
}

func configureUserRoutes(g *echo.Group) {
	g.GET("", routes.GetUsersHandler)
	g.GET("/:id", routes.GetUserHandler)
	g.POST("", routes.CreateUserHandler)
	g.PUT("/:id", routes.UpdateUserHandler)
	g.DELETE("/:id", routes.DeleteUserHandler)
}

func configureExamRoutes(g *echo.Group) {
	g.GET("", routes.GetExamsHandler)
	g.GET("/:id", routes.GetExamHandler)
	g.POST("", routes.CreateExamHandler)
	g.PUT("/:id", routes.UpdateExamHandler)
	g.DELETE("/:id", routes.DeleteExamHandler)
}

func configureClassRoutes(g *echo.Group) {
	g.GET("", routes.GetClasssHandler)
	g.GET("/:id", routes.GetClassHandler)
	g.POST("", routes.CreateClassHandler)
	g.PUT("/:id", routes.UpdateClassHandler)
	g.DELETE("/:id", routes.DeleteClassHandler)
	g.GET("/schools/:id", routes.GetClassBySchoolHandler)
}

func configureScoreRoutes(g *echo.Group) {
	g.GET("", routes.GetScoresHandler)
	g.GET("/:id", routes.GetScoreHandler)
	g.POST("", routes.CreateScoreHandler)
	g.PUT("/:id", routes.UpdateScoreHandler)
	g.DELETE("/:id", routes.DeleteScoreHandler)
}

func configureAddressRoutes(g *echo.Group) {
	g.GET("", routes.GetAddressesHandler)
	g.GET("/:id", routes.GetAddressHandler)
	g.POST("", routes.CreateAddressHandler)
	g.PUT("/:id", routes.UpdateAddressHandler)
	g.DELETE("/:id", routes.DeleteAddressHandler)
}

func configureGradeRoutes(g *echo.Group) {
	g.GET("", routes.GetGradesHandler)
	g.GET("/:id", routes.GetGradeHandler)
	g.POST("", routes.CreateGradeHandler)
	g.PUT("/:id", routes.UpdateGradeHandler)
	g.DELETE("/:id", routes.DeleteGradeHandler)
}

func configureExamTypeRoutes(g *echo.Group) {
	g.GET("", routes.GetExamTypesHandler)
	g.GET("/:id", routes.GetExamTypeHandler)
	g.POST("", routes.CreateExamTypeHandler)
	g.PUT("/:id", routes.UpdateExamTypeHandler)
	g.DELETE("/:id", routes.DeleteExamTypeHandler)
}

func configureSchoolRoutes(g *echo.Group) {
	g.GET("", routes.GetSchoolsHandler)
	g.GET("/:id", routes.GetSchoolHandler)
	g.POST("", routes.CreateSchoolHandler)
	g.PUT("/:id", routes.UpdateSchoolHandler)
	g.DELETE("/:id", routes.DeleteSchoolHandler)
}

func configureSubjectRoutes(g *echo.Group) {
	g.GET("", routes.GetSubjectsHandler)
	g.GET("/:id", routes.GetSubjectHandler)
	g.POST("", routes.CreateSubjectHandler)
	g.PUT("/:id", routes.UpdateSubjectHandler)
	g.DELETE("/:id", routes.DeleteSubjectHandler)
}

func configureSubjectUserClassEexamRoutes(g *echo.Group) {
	g.GET("", routes.GetSubjectUserClassExamsHandler)
	g.GET("/:id", routes.GetSubjectUserClassExamHandler)
	g.POST("", routes.CreateSubjectUserClassExamHandler)
	g.PUT("/:id", routes.UpdateSubjectUserClassExamHandler)
	g.DELETE("/:id", routes.DeleteSubjectUserClassExamHandler)
}

func checkCredentials(username, password string) bool {
	user, err := DB.GetUserByName(username)
	if err != nil {
		return false
	}
	if user != nil && user.Password == password {
		return true
	}
	return false
}
