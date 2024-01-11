package routes

import (
	"fmt"
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetStudentsHandler gibt alle Studenten zurück
func GetStudentsHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var students []model.Student
	if err := db.Find(&students).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Studenten"})
	}

	return c.JSON(http.StatusOK, students)
}

// GetStudentHandler gibt einen bestimmten Studenten anhand der ID zurück
func GetStudentHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Studenten-ID"})
	}

	var student model.Student
	if err := db.Find(&student, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Student nicht gefunden"})
	}

	fmt.Println(student)
	return c.JSON(http.StatusOK, student)
}

// CreateStudentHandler erstellt einen neuen Studenten
func CreateStudentHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var student model.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&student).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des Studenten"})
	}

	return c.JSON(http.StatusCreated, student)
}

// UpdateStudentHandler aktualisiert einen vorhandenen Studenten
func UpdateStudentHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Studenten-ID"})
	}

	var existingStudent model.Student
	if err := db.First(&existingStudent, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Student nicht gefunden"})
	}

	if err := c.Bind(&existingStudent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingStudent).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des Studenten"})
	}

	return c.JSON(http.StatusOK, existingStudent)
}

// DeleteStudentHandler löscht einen Studenten anhand der ID
func DeleteStudentHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Studenten-ID"})
	}

	if err := db.Delete(&model.Student{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des Studenten"})
	}

	return c.NoContent(http.StatusNoContent)
}
