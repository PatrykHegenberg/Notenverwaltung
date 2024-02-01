package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetStudentsHandler godoc
// @Summary get all students
// @Description get all students from db.
// @Tags student
// @Accept application/json
// @Produce json
// @Success 200 {object} []model.Student
// @Failure 500 {object} ErrorResponse "Fehler beim Abrufen der Studenten"
// @Router /students [get]
func GetStudentsHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var students []model.Student
	if err := db.Preload("Address").Find(&students).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Studenten"})
	}

	return c.JSON(http.StatusOK, students)
}

// GetStudentHandler godoc
// @Summary get one student by id
// @Description get one student from db by ID.
// @Tags student
// @Accept application/json
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} model.Student
// @Failure 400 {object} ErrorResponse "Ungültige Studenten-ID"
// @Failure 404 {object} ErrorResponse "Student nicht gefunden"
// @Router /students/:id [get]
func GetStudentHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Studenten-ID"})
	}

	var student model.Student
	if err := db.Preload("Address").Find(&student, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Student nicht gefunden"})
	}

	return c.JSON(http.StatusOK, student)
}

// GetStudentByClassHandler godoc
// @Summary get all students by a specific class id
// @Description get all students from db by class ID.
// @Tags student
// @Accept application/json
// @Produce json
// @Param id path int true "Class ID"
// @Success 200 {object} []model.Student
// @Failure 400 {object} ErrorResponse "Ungültige Class-ID"
// @Failure 404 {object} ErrorResponse "Keine Students gefunden"
// @Router /students/class/:id [get]
func GetStudentsByClassHandler(c echo.Context) error {
	db := DB.GetDBInstance()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Class-ID"})
	}

	var class model.Class
	if err := db.Model(&model.Class{}).Preload("Students").First(&class, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Class nicht gefunden"})
	}

	return c.JSON(http.StatusOK, class.Students)
}

// CreateStudentHandler godoc
// @Summary create student
// @Description create a new student
// @Tags student
// @Accept application/json
// @Produce json
// @Param student body model.Student true "Student object that needs to be added"
// @Success 201 {object} model.Student
// @Failure 500 {object} ErrorResponse "Fehler beim Erstellen des Studenten"
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Router /students [post]
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

// UpdateStudentHandler godoc
// @Summary update student
// @Description update an existing student
// @Tags student
// @Accept application/json
// @Produce json
// @Param id path int true "Student ID"
// @Param student body model.Student true "Updated student object"
// @Success 200 {object} model.Student
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Failure 404 {object} ErrorResponse "Student nicht gefunden"
// @Router /students/:id [put]
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

// DeleteStudentHandler godoc
// @Summary delete student
// @Description delete an existing student
// @Tags student
// @Accept application/json
// @Produce json
// @Param id path int true "Student ID"
// @Success 204
// @Failure 400 {object} ErrorResponse "Ungültige Studenten-ID"
// @Failure 500 {object} ErrorResponse "Fehler beim Löschen des Studenten"
// @Router /students/:id [delete]
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
