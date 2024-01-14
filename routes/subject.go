package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetSubjectsHandler gibt alle Subjecten zurück
func GetSubjectsHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var subjects []model.Subject
	if err := db.Find(&subjects).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Subjecten"})
	}

	return c.JSON(http.StatusOK, subjects)
}

// GetSubjectHandler gibt einen bestimmten Subjecten anhand der ID zurück
func GetSubjectHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Subjecten-ID"})
	}

	var subject model.Subject
	if err := db.Find(&subject, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Subject nicht gefunden"})
	}

	return c.JSON(http.StatusOK, subject)
}

// CreateSubjectHandler erstellt einen neuen Subjecten
func CreateSubjectHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var subject model.Subject
	if err := c.Bind(&subject); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&subject).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des Subjecten"})
	}

	return c.JSON(http.StatusCreated, subject)
}

// UpdateSubjectHandler aktualisiert einen vorhandenen Subjecten
func UpdateSubjectHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Subjecten-ID"})
	}

	var existingSubject model.Subject
	if err := db.First(&existingSubject, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Subject nicht gefunden"})
	}

	if err := c.Bind(&existingSubject); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingSubject).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des Subjecten"})
	}

	return c.JSON(http.StatusOK, existingSubject)
}

// DeleteSubjectHandler löscht einen Subjecten anhand der ID
func DeleteSubjectHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Subjecten-ID"})
	}

	if err := db.Delete(&model.Subject{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des Subjecten"})
	}

	return c.NoContent(http.StatusNoContent)
}
