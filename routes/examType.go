package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetExamTypesHandler godoc
// @Summary get all exam types
// @Description get all exam types from db.
// @Tags examtype
// @Accept application/json
// @Produce json
// @Success 200 {object} []model.ExamType
// @Failure 500 {object} ErrorResponse "Fehler beim Abrufen der ExamTypeen"
// @Router /examtypes [get]
func GetExamTypesHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var examTypes []model.ExamType
	if err := db.Find(&examTypes).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der ExamTypeen"})
	}

	return c.JSON(http.StatusOK, examTypes)
}

// GetExamTypeHandler godoc
// @Summary get one exam type by id
// @Description get one exam type from db by ID.
// @Tags examtype
// @Accept application/json
// @Produce json
// @Param id path int true "ExamType ID"
// @Success 200 {object} model.ExamType
// @Failure 400 {object} ErrorResponse "Ungültige ExamTypeen-ID"
// @Failure 404 {object} ErrorResponse "ExamType nicht gefunden"
// @Router /examtypes/:id [get]
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

// CreateExamTypeHandler godoc
// @Summary create exam type
// @Description create a new exam type
// @Tags examtype
// @Accept application/json
// @Produce json
// @Param examType body model.ExamType true "ExamType object that needs to be added"
// @Success 201 {object} model.ExamType
// @Failure 500 {object} ErrorResponse "Fehler beim Erstellen des ExamTypeen"
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Router /examtypes [post]
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

// UpdateExamTypeHandler godoc
// @Summary update exam type
// @Description update an existing exam type
// @Tags examtype
// @Accept application/json
// @Produce json
// @Param id path int true "ExamType ID"
// @Param examType body model.ExamType true "Updated exam type object"
// @Success 200 {object} model.ExamType
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Failure 404 {object} ErrorResponse "ExamType nicht gefunden"
// @Router /examtypes/:id [put]
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

// DeleteExamTypeHandler godoc
// @Summary delete exam type
// @Description delete an existing exam type
// @Tags examtype
// @Accept application/json
// @Produce json
// @Param id path int true "ExamType ID"
// @Success 204
// @Failure 400 {object} ErrorResponse "Ungültige ExamTypeen-ID"
// @Failure 500 {object} ErrorResponse "Fehler beim Löschen des ExamTypeen"
// @Router /examtypes/:id [delete]
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
