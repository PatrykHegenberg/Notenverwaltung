package main

import (
	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/routes"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	DB.AutoMigrate()
	e := echo.New()

	// HTMX Frontend Routes
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("my-secret"))))
	e.GET("/", routes.GetIndexHandler)
	e.GET("/register", routes.GetRegisterHandler)
	e.GET("/login", routes.GetLoginHandler)
	e.GET("/dashboard", routes.GetDashboardHandler)
	e.GET("/logout", routes.LogoutHXUserHandler)
	e.POST("/authenticate", routes.AuthenticateHXUserHandler)

	// API Routes
	apiGroup := e.Group("/api/v1")
	apiGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		return checkCredentials(username, password), nil
	}))
	studentGroup := apiGroup.Group("/student")
	configureStudentRoutes(studentGroup)

	userGroup := apiGroup.Group("/users")
	configureUserRoutes(userGroup)

	roleGroup := apiGroup.Group("/roles")
	configureRoleRoutes(roleGroup)

	examGroup := apiGroup.Group("/exams")
	configureExamRoutes(examGroup)

	classGroup := apiGroup.Group("/class")
	configureClassRoutes(classGroup)

	teacherGroup := apiGroup.Group("/teachers")
	configureTeacherRoutes(teacherGroup)

	scoreGroup := apiGroup.Group("/scores")
	configureScoreRoutes(scoreGroup)

	e.Logger.Fatal(e.Start(":1323"))
}

func configureStudentRoutes(g *echo.Group) {
	g.GET("", routes.GetStudentsHandler)
	g.GET("/:id", routes.GetStudentHandler)
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

func configureRoleRoutes(g *echo.Group) {
	g.GET("", routes.GetRolesHandler)
	g.GET("/:id", routes.GetRoleHandler)
	g.POST("", routes.CreateRoleHandler)
	g.PUT("/:id", routes.UpdateRoleHandler)
	g.DELETE("/:id", routes.DeleteRoleHandler)
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
}

func configureTeacherRoutes(g *echo.Group) {
	g.GET("", routes.GetTeachersHandler)
	g.GET("/:id", routes.GetTeacherHandler)
	g.POST("", routes.CreateTeacherHandler)
	g.PUT("/:id", routes.UpdateTeacherHandler)
	g.DELETE("/:id", routes.DeleteTeacherHandler)
}

func configureScoreRoutes(g *echo.Group) {
	g.GET("", routes.GetScoresHandler)
	g.GET("/:id", routes.GetScoreHandler)
	g.POST("", routes.CreateScoreHandler)
	g.PUT("/:id", routes.UpdateScoreHandler)
	g.DELETE("/:id", routes.DeleteScoreHandler)
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
