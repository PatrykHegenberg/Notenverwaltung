package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetExamTypesHandler gibt alle ExamTypeen zurück
func GetExamTypesHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var examTypes []model.ExamType
	if err := db.Find(&examTypes).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der ExamTypeen"})
	}

	return c.JSON(http.StatusOK, examTypes)
}

// GetExamTypeHandler gibt einen bestimmten ExamTypeen anhand der ID zurück
func GetExamTypeHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige ExamTypeen-ID"})
	}

	var examType model.ExamType
	if err := db.Find(&examType, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "ExamType nicht gefunden"})
	}

	return c.JSON(http.StatusOK, examType)
}

// CreateExamTypeHandler erstellt einen neuen ExamTypeen
func CreateExamTypeHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var examType model.ExamType
	if err := c.Bind(&examType); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&examType).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des ExamTypeen"})
	}

	return c.JSON(http.StatusCreated, examType)
}

// UpdateExamTypeHandler aktualisiert einen vorhandenen ExamTypeen
func UpdateExamTypeHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige ExamTypeen-ID"})
	}

	var existingExamType model.ExamType
	if err := db.First(&existingExamType, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "ExamType nicht gefunden"})
	}

	if err := c.Bind(&existingExamType); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingExamType).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des ExamTypeen"})
	}

	return c.JSON(http.StatusOK, existingExamType)
}

// DeleteExamTypeHandler löscht einen ExamTypeen anhand der ID
func DeleteExamTypeHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige ExamTypeen-ID"})
	}

	if err := db.Delete(&model.ExamType{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des ExamTypeen"})
	}

	return c.NoContent(http.StatusNoContent)
}
