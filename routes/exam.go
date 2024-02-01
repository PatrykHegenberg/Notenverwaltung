package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// GetExamsHandler godoc
// @Summary get all exams
// @Description get all exams from db.
// @Tags exam
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Success 200 {object} []model.Exam
// @Failure 500 {object} ErrorResponse "Fehler beim Abraufen der Exams"
// @Router /exams [get]
func GetExamsHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var exam []model.Exam
	if err := db.Preload("ExamType").Find(&exam).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Exams"})
	}

	return c.JSON(http.StatusOK, exam)
}

// GetExamHandler godoc
// @Summary get one exams by id
// @Description get all exams from db.
// @Tags exam
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Success 200 {object} model.Exam
// @Failure 500 {object} ErrorResponse "ungueltige Exam-ID"
// @Failure 404 {object} ErrorResponse "Exam nicht gefunden"
// @Router /exams/:id [get]
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

// CreateExamHandler godoc
// @Summary create exam
// @Description create a new exam
// @Tags exam
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Success 203 {object} model.Exam
// @Failure 500 {object} ErrorResponse "Fehler beim Erstellen des Exams"
// @Failure 400 {object} ErrorResponse "ungueltige anfrage"
// @Router /exams [post]
// @securityDefinitions.basic BasicAuth
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

// UpdateExamHandler godoc
// @Summary update exam
// @Description update an existing exam
// @Tags exam
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Success 200 {object} model.Exam
// @Failure 500 {object} ErrorResponse "Fehler beim aktualisieren des Exams"
// @Failure 400 {object} ErrorResponse "ungueltige anfrage"
// @Failure 404 {object} ErrorResponse "Exam nicht gefunden"
// @Failure 400 {object} ErrorResponse "ungueltige exam-id"
// @Router /exams/:id [put]
// @securityDefinitions.basic BasicAuth
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

// DeleteExamHandler godoc
// @Summary delete exam
// @Description delete an existing exam
// @Tags exam
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Success 204
// @Failure 500 {object} ErrorResponse "Fehler beim loeschen des Exams"
// @Failure 400 {object} ErrorResponse "ungueltige exam-id"
// @Router /exams/:id [delete]
// @securityDefinitions.basic BasicAuth
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
