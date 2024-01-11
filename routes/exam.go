package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetExamsHandler gibt alle Exam zurück
func GetExamsHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var exam []model.Exam
	if err := db.Preload("ExamType").Find(&exam).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Exams"})
	}

	return c.JSON(http.StatusOK, exam)
}

// GetExamHandler gibt einen bestimmten Exam anhand der ID zurück
func GetExamHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Exam-ID"})
	}

	var exam model.Exam
	if err := db.Preload("ExamType").First(&exam, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Exam nicht gefunden"})
	}

	return c.JSON(http.StatusOK, exam)
}

// CreateExamHandler erstellt einen neuen Exams
func CreateExamHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var exam model.Exam
	if err := c.Bind(&exam); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&exam).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des Exams"})
	}

	return c.JSON(http.StatusCreated, exam)
}

// UpdateExamHandler aktualisiert einen vorhandenen Exam
func UpdateExamHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Exam-ID"})
	}

	var existingExam model.Exam
	if err := db.First(&existingExam, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Exam nicht gefunden"})
	}

	if err := c.Bind(&existingExam); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingExam).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des Exams"})
	}

	return c.JSON(http.StatusOK, existingExam)
}

// DeleteExamHandler löscht einen Exam anhand der ID
func DeleteExamHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Exam-ID"})
	}

	if err := db.Delete(&model.Exam{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des Exam"})
	}

	return c.NoContent(http.StatusNoContent)
}
