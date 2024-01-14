package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetGradesHandler is a function that handles the request to get grades.
//
// It takes an echo.Context parameter.
// It returns an error.
func GetGradesHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var grades []model.Grade
	if err := db.Find(&grades).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Gradeen"})
	}

	return c.JSON(http.StatusOK, grades)
}

// GetGradeHandler is a function that handles the request to get a grade.
//
// It takes in a `c` parameter of type `echo.Context` which represents the request context.
// It returns an error.
func GetGradeHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Gradeen-ID"})
	}

	var grade model.Grade
	if err := db.Find(&grade, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Grade nicht gefunden"})
	}

	return c.JSON(http.StatusOK, grade)
}

// CreateGradeHandler handles the creation of a grade.
//
// It takes a parameter of type `echo.Context` and returns an error.
func CreateGradeHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var grade model.Grade
	if err := c.Bind(&grade); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&grade).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des Gradeen"})
	}

	return c.JSON(http.StatusCreated, grade)
}

// UpdateGradeHandler updates a grade in the database.
//
// It takes a `c` parameter of type `echo.Context`.
// It returns an error.
func UpdateGradeHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Gradeen-ID"})
	}

	var existingGrade model.Grade
	if err := db.First(&existingGrade, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Grade nicht gefunden"})
	}

	if err := c.Bind(&existingGrade); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingGrade).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des Gradeen"})
	}

	return c.JSON(http.StatusOK, existingGrade)
}

// DeleteGradeHandler is a function that handles the deletion of a grade.
//
// It takes a parameter of type echo.Context and returns an error.
func DeleteGradeHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Gradeen-ID"})
	}

	if err := db.Delete(&model.Grade{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des Gradeen"})
	}

	return c.NoContent(http.StatusNoContent)
}
