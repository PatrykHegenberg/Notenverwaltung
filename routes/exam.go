package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetExamsHandler is a Go function that handles the retrieval of exams.
//
// It takes a parameter of type `echo.Context`.
// It returns an error.
func GetExamsHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var exam []model.Exam
	if err := db.Preload("ExamType").Find(&exam).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Exams"})
	}

	return c.JSON(http.StatusOK, exam)
}

// GetExamHandler is a function that handles the GET request for retrieving an exam.
//
// It takes a parameter of type echo.Context, which represents the HTTP request and response context.
// The function returns an error.
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

// CreateExamHandler handles the creation of an exam.
//
// It takes a context object as a parameter and returns an error.
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

// UpdateExamHandler handles the update of an exam.
//
// The function takes a `c` parameter of type `echo.Context` which represents the HTTP request and response context.
// The function returns an error.
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

// DeleteExamHandler is a function that handles the deletion of an exam.
//
// It takes in a parameter of type echo.Context, which represents the context of the HTTP request.
// The function returns an error.
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
