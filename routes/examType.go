package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetExamTypesHandler is a function that handles the request to retrieve exam types.
//
// It takes in a parameter of type echo.Context.
// It returns an error.
func GetExamTypesHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var examTypes []model.ExamType
	if err := db.Find(&examTypes).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der ExamTypeen"})
	}

	return c.JSON(http.StatusOK, examTypes)
}

// GetExamTypeHandler is a function that handles the API endpoint for retrieving an ExamType.
//
// It takes a parameter `c` of type `echo.Context`.
// It returns an error.
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

// CreateExamTypeHandler is responsible for handling the creation of an ExamType.
//
// Parameters:
// - c: echo.Context - the context object for the HTTP request.
//
// Returns:
// - error: the error that occurred during the execution of the function.
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

// UpdateExamTypeHandler is a function that handles the update of an ExamType in the database.
//
// It takes in a parameter `c` of type `echo.Context` which represents the current HTTP context.
// It returns an error.
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

// DeleteExamTypeHandler is a function that handles the deletion of an ExamType.
//
// It takes a parameter of type `echo.Context` which represents the HTTP request context.
// It returns an error.
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
