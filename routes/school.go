package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

// GetSchoolsHandler gibt alle Schoolen zurück
func GetSchoolsHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var schools []model.School
	// if err := db.Model(&model.School{}).Preload("Address").Preload("SchoolForm").Preload("Classes").Preload("Users").Find(&schools).Error; err != nil {
	if err := db.Model(&model.School{}).Preload(clause.Associations).Preload("Classes.Students").Preload("Classes.Students.Address").Preload("Users.Address").Find(&schools).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Schoolen"})
	}

	return c.JSON(http.StatusOK, schools)
}

// GetSchoolHandler gibt einen bestimmten Schoolen anhand der ID zurück
func GetSchoolHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Schoolen-ID"})
	}

	var school model.School
	if err := db.Model(&model.School{}).Preload(clause.Associations).Preload("Classes.Students").Preload("Classes.Students.Address").Preload("Users.Address").Find(&school, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "School nicht gefunden"})
	}

	return c.JSON(http.StatusOK, school)
}

// CreateSchoolHandler erstellt einen neuen Schoolen
func CreateSchoolHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var school model.School
	if err := c.Bind(&school); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&school).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des Schoolen"})
	}

	return c.JSON(http.StatusCreated, school)
}

// UpdateSchoolHandler aktualisiert einen vorhandenen Schoolen
func UpdateSchoolHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Schoolen-ID"})
	}

	var existingSchool model.School
	if err := db.First(&existingSchool, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "School nicht gefunden"})
	}

	if err := c.Bind(&existingSchool); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingSchool).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des Schoolen"})
	}

	return c.JSON(http.StatusOK, existingSchool)
}

// DeleteSchoolHandler löscht einen Schoolen anhand der ID
func DeleteSchoolHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Schoolen-ID"})
	}

	if err := db.Delete(&model.School{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des Schoolen"})
	}

	return c.NoContent(http.StatusNoContent)
}
