package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

// GetSubjectUserClassExamsHandler gibt alle SubjectUserClassExamen zurück
func GetSubjectUserClassExamsHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var subjectUserClassExams []model.SubjectUserClassExam
	if err := db.Preload(clause.Associations).Preload("Class.Students").Preload("Exam.Scores").Preload("Exam.Scores.Grade").Find(&subjectUserClassExams).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der SubjectUserClassExamen"})
	}

	return c.JSON(http.StatusOK, subjectUserClassExams)
}

// GetSubjectUserClassExamHandler gibt einen bestimmten SubjectUserClassExamen anhand der ID zurück
func GetSubjectUserClassExamHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige SubjectUserClassExamen-ID"})
	}

	var subjectUserClassExam model.SubjectUserClassExam
	if err := db.Preload(clause.Associations).Preload("Classes.Students").Preload("Exam.Scores").Preload("Exam.Scores.Grade").Find(&subjectUserClassExam, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "SubjectUserClassExam nicht gefunden"})
	}

	return c.JSON(http.StatusOK, subjectUserClassExam)
}

// CreateSubjectUserClassExamHandler erstellt einen neuen SubjectUserClassExamen
func CreateSubjectUserClassExamHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var subjectUserClassExam model.SubjectUserClassExam
	if err := c.Bind(&subjectUserClassExam); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&subjectUserClassExam).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des SubjectUserClassExamen"})
	}

	return c.JSON(http.StatusCreated, subjectUserClassExam)
}

// UpdateSubjectUserClassExamHandler aktualisiert einen vorhandenen SubjectUserClassExamen
func UpdateSubjectUserClassExamHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige SubjectUserClassExamen-ID"})
	}

	var existingSubjectUserClassExam model.SubjectUserClassExam
	if err := db.First(&existingSubjectUserClassExam, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "SubjectUserClassExam nicht gefunden"})
	}

	if err := c.Bind(&existingSubjectUserClassExam); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingSubjectUserClassExam).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des SubjectUserClassExamen"})
	}

	return c.JSON(http.StatusOK, existingSubjectUserClassExam)
}

// DeleteSubjectUserClassExamHandler löscht einen SubjectUserClassExamen anhand der ID
func DeleteSubjectUserClassExamHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige SubjectUserClassExamen-ID"})
	}

	if err := db.Delete(&model.SubjectUserClassExam{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des SubjectUserClassExamen"})
	}

	return c.NoContent(http.StatusNoContent)
}
