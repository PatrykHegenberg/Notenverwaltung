package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

// GetSubjectUserClassExamsHandler godoc
// @Summary get all subject user class exams
// @Description get all subject user class exams from db.
// @Tags subjectuserclassexam
// @Accept application/json
// @Produce json
// @Success 200 {object} []model.SubjectUserClassExam
// @Failure 500 {object} ErrorResponse "Fehler beim Abrufen der SubjectUserClassExamen"
// @Router /sucs [get]
func GetSubjectUserClassExamsHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var subjectUserClassExams []model.SubjectUserClassExam
	if err := db.Preload(clause.Associations).Preload("Class.Students").Preload("Exam.Scores").Preload("Exam.Scores.Grade").Find(&subjectUserClassExams).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der SubjectUserClassExamen"})
	}

	return c.JSON(http.StatusOK, subjectUserClassExams)
}

// GetSubjectUserClassExamHandler godoc
// @Summary get one subject user class exam by id
// @Description get one subject user class exam from db by ID.
// @Tags subjectuserclassexam
// @Accept application/json
// @Produce json
// @Param id path int true "SubjectUserClassExam ID"
// @Success 200 {object} model.SubjectUserClassExam
// @Failure 400 {object} ErrorResponse "Ungültige SubjectUserClassExamen-ID"
// @Failure 404 {object} ErrorResponse "SubjectUserClassExam nicht gefunden"
// @Router /sucs/:id [get]
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

// CreateSubjectUserClassExamHandler godoc
// @Summary create subject user class exam
// @Description create a new subject user class exam
// @Tags subjectuserclassexam
// @Accept application/json
// @Produce json
// @Param subjectUserClassExam body model.SubjectUserClassExam true "SubjectUserClassExam object that needs to be added"
// @Success 201 {object} model.SubjectUserClassExam
// @Failure 500 {object} ErrorResponse "Fehler beim Erstellen des SubjectUserClassExamen"
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Router /sucs [post]
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

// UpdateSubjectUserClassExamHandler godoc
// @Summary update subject user class exam
// @Description update an existing subject user class exam
// @Tags subjectuserclassexam
// @Accept application/json
// @Produce json
// @Param id path int true "SubjectUserClassExam ID"
// @Param subjectUserClassExam body model.SubjectUserClassExam true "Updated subject user class exam object"
// @Success 200 {object} model.SubjectUserClassExam
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Failure 404 {object} ErrorResponse "SubjectUserClassExam nicht gefunden"
// @Router /sucs/:id [put]
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

// DeleteSubjectUserClassExamHandler godoc
// @Summary delete subject user class exam
// @Description delete an existing subject user class exam
// @Tags subjectuserclassexam
// @Accept application/json
// @Produce json
// @Param id path int true "SubjectUserClassExam ID"
// @Success 204
// @Failure 400 {object} ErrorResponse "Ungültige SubjectUserClassExamen-ID"
// @Failure 500 {object} ErrorResponse "Fehler beim Löschen des SubjectUserClassExamen"
// @Router /sucs/:id [delete]
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
