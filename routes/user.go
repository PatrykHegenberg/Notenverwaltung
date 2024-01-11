package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetUsersHandler gibt alle User zurück
func GetUsersHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var users []model.User
	if err := db.Preload("Role").Preload("School").Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Users"})
	}

	return c.JSON(http.StatusOK, users)
}

// GetUserHandler gibt einen bestimmten User anhand der ID zurück
func GetUserHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige User-ID"})
	}

	var user model.User
	if err := db.Preload("Role").Preload("School").First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User nicht gefunden"})
	}

	return c.JSON(http.StatusOK, user)
}

// CreateUserHandler erstellt einen neuen Users
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

// UpdateUserHandler aktualisiert einen vorhandenen User
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

	if err := c.Bind(&existingUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des Users"})
	}

	return c.JSON(http.StatusOK, existingUser)
}

// DeleteUserHandler löscht einen User anhand der ID
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
