package routes

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/PatrykHegenberg/Notenverwaltung/utils"
	"github.com/PatrykHegenberg/Notenverwaltung/views"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// GetUsersHandler godoc
// @Summary get all users
// @Description get all users from db.
// @Tags user
// @Accept application/json
// @Produce json
// @Success 200 {object} []model.User
// @Failure 500 {object} ErrorResponse "Fehler beim Abrufen der Users"
// @Router /users [get]
func GetUsersHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var users []model.User
	if err := db.Preload("Address").Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Users"})
	}

	return c.JSON(http.StatusOK, users)
}

// GetUserHandler godoc
// @Summary get one user by id
// @Description get one user from db by ID.
// @Tags user
// @Accept application/json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} ErrorResponse "Ungültige User-ID"
// @Failure 404 {object} ErrorResponse "User nicht gefunden"
// @Router /users/:id [get]
func GetUserHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige User-ID"})
	}

	var user model.User
	if err := db.Preload("Address").First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User nicht gefunden"})
	}

	return c.JSON(http.StatusOK, user)
}

// CreateUserHandler godoc
// @Summary create user
// @Description create a new user
// @Tags user
// @Accept application/json
// @Produce json
// @Param user body model.User true "User object that needs to be added"
// @Success 201 {object} model.User
// @Failure 500 {object} ErrorResponse "Fehler beim Erstellen des Users"
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Router /users [post]
func CreateUserHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des Users"})
	}

	return c.JSON(http.StatusCreated, user)
}

// CreateHXUserHandler godoc
// @Summary create user with HX
// @Description create a new user with HX
// @Tags user
// @Accept application/json
// @Produce json
// @Param username formData string true "Username"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Param school formData string true "School"
// @Success 302
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Fehler beim Erstellen des Users"
// @Router /users/hx [post]
func CreateHXUserHandler(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	school := c.FormValue("school")
	db := DB.GetDBInstance()

	schoolID, err := DB.GetSchoolIDByName(school)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Schule nicht gefunden"})
	}

	var user = model.User{
		Username: username,
		Email:    email,
		Password: password,
		SchoolID: schoolID,
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}

	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des Users"})
	}

	return c.Redirect(http.StatusCreated, "/")
}

// UpdateUserHandler godoc
// @Summary update user
// @Description update an existing user
// @Tags user
// @Accept application/json
// @Produce json
// @Param id path int true "User ID"
// @Param user body model.User true "Updated user object"
// @Success 200 {object} model.User
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Failure 404 {object} ErrorResponse "User nicht gefunden"
// @Router /users/:id [put]
func UpdateUserHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige User-ID"})
	}

	var existingUser model.User
	if err := db.First(&existingUser, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User nicht gefunden"})
	}

	log.Printf("%v\n", existingUser)
	if err := c.Bind(&existingUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}
	log.Printf("%v\n", existingUser)

	if err := db.Save(&existingUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des Users"})
	}
	log.Printf("%v\n", existingUser)

	return c.JSON(http.StatusOK, existingUser)
}

// DeleteUserHandler godoc
// @Summary delete user
// @Description delete an existing user
// @Tags user
// @Accept application/json
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Failure 400 {object} ErrorResponse "Ungültige User-ID"
// @Failure 500 {object} ErrorResponse "Fehler beim Löschen des Users"
// @Router /users/:id [delete]
func DeleteUserHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige User-ID"})
	}

	if err := db.Delete(&model.User{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des User"})
	}

	return c.NoContent(http.StatusNoContent)
}

// AuthenticateHXUserHandler godoc
// @Summary authenticate HX user
// @Description authenticate a user with HX credentials
// @Tags user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 303
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 401 {object} ErrorResponse "Ungültige Anmeldeinformationen"
// @Router /users/authenticate/hx [post]
func AuthenticateHXUserHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	email := c.FormValue("email")
	password := c.FormValue("password")

	var user model.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Ungültige Anmeldeinformationen"})
	}

	if user.Password != password {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Ungültige Anmeldeinformationen"})
	} else {
		sess, _ := session.Get("authenticate-session", c)
		sess.Options = &sessions.Options{
			Path:   "/",
			MaxAge: 86400 * 7,
		}
		sess.Values["authenticated"] = true
		sess.Values["username"] = user.Username
		sess.Save(c.Request(), c.Response())
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

// AuthenticateUserHandler godoc
// @Summary authenticate user
// @Description authenticate a user with basic authentication
// @Tags user
// @Accept application/json
// @Produce plain
// @Param Authorization header string true "Basic Authentication"
// @Success 200 {string} string "Erfolgreich authentifiziert"
// @Failure 400 {object} ErrorResponse "Ungültige Anmeldeinformationen"
// @Failure 401 {object} ErrorResponse "Ungültiges Authentifizierungsschema"
// @Router /users/authenticate [get]
func AuthenticateUserHandler(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authentifizierung erforderlich"})
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Basic" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Ungültiges Authentifizierungsschema"})
	}

	credentials, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Ungültige Base64-Codierung"})
	}

	userPass := strings.SplitN(string(credentials), ":", 2)
	if len(userPass) != 2 {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Ungültige Benutzerinformationen"})
	}

	username := userPass[0]
	password := userPass[1]

	db := DB.GetDBInstance()
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Ungültige Anmeldeinformationen"})
	}

	if validCredentials(username, password) {
		return c.JSON(http.StatusOK, user)
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Ungültige Anmeldeinformationen"})
}

func validCredentials(username, password string) bool {
	db := DB.GetDBInstance()
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return false
	}
	return username == user.Username && password == user.Password
}

// LogoutHXUserHandler godoc
// @Summary logout HX user
// @Description logout an HX authenticated user
// @Tags user
// @Accept application/json
// @Produce json
// @Success 303
// @Router /users/logout/hx [get]
func LogoutHXUserHandler(c echo.Context) error {
	sess, _ := session.Get("authenticate-session", c)
	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400 * -1,
	}
	sess.Values["authenticated"] = false
	sess.Values["username"] = ""
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/")
}

type JwtCustomClaims struct {
	Name   string `json:"name"`
	Admin  bool   `json:"admin"`
	School uint   `json:"school"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	fmt.Println(username)
	fmt.Println(password)
	user, err := checkCredentials(username, password)
	if err != nil {
		return echo.ErrUnauthorized
	}

	claims := &JwtCustomClaims{
		user.Username,
		user.IsAdmin,
		user.SchoolID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = t
	cookie.Expires = time.Now().Add(72 * time.Hour)

	c.SetCookie(cookie)

	// return utils.Render(c, http.StatusOK, views.Dashboard())
	return c.Redirect(http.StatusSeeOther, "/restricted/dashboard")
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func HomeHandler(c echo.Context) error {
	return utils.Render(c, http.StatusOK, views.Root())
}

func checkCredentials(username, password string) (*model.User, error) {
	user, err := DB.GetUserByName(username)
	if err != nil {
		return nil, err
	}
	if user != nil && user.Password == password {
		return user, nil
	}
	return nil, fmt.Errorf("error: wrong credentials")
}
