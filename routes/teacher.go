package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetTeachersHandler gibt alle Teacher zurück
func GetTeachersHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var teacher []model.Teacher
	if err := db.Preload("User").Preload("User.Role").Preload("User.School").Find(&teacher).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Teachers"})
	}

	return c.JSON(http.StatusOK, teacher)
}

// GetTeacherHandler gibt einen bestimmten Teacher anhand der ID zurück
func GetTeacherHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Teacher-ID"})
	}

	var teacher model.Teacher
	if err := db.Preload("User").Preload("User.Role").Preload("User.School").First(&teacher, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Teacher nicht gefunden"})
	}

	return c.JSON(http.StatusOK, teacher)
}

// CreateTeacherHandler erstellt einen neuen Teachers
func CreateTeacherHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var teacher model.Teacher
	if err := c.Bind(&teacher); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&teacher).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des Teachers"})
	}

	return c.JSON(http.StatusCreated, teacher)
}

// UpdateTeacherHandler aktualisiert einen vorhandenen Teacher
func UpdateTeacherHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Teacher-ID"})
	}

	var existingTeacher model.Teacher
	if err := db.First(&existingTeacher, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Teacher nicht gefunden"})
	}

	if err := c.Bind(&existingTeacher); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingTeacher).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des Teachers"})
	}

	return c.JSON(http.StatusOK, existingTeacher)
}

// DeleteTeacherHandler löscht einen Teacher anhand der ID
func DeleteTeacherHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Teacher-ID"})
	}

	if err := db.Delete(&model.Teacher{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des Teacher"})
	}

	return c.NoContent(http.StatusNoContent)
}
