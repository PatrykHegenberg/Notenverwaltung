package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetGradesHandler godoc
// @Summary get all grades
// @Description get all grades from db.
// @Tags grade
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Success 200 {object} []model.Grade
// @Failure 500 {object} ErrorResponse "Fehler beim Abrufen der Gradeen"
// @Router /grades [get]
func GetGradesHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var grades []model.Grade
	if err := db.Find(&grades).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Gradeen"})
	}

	return c.JSON(http.StatusOK, grades)
}

// GetGradeHandler godoc
// @Summary get one grade by id
// @Description get one grade from db by ID.
// @Tags grade
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Param id path int true "Grade ID"
// @Success 200 {object} model.Grade
// @Failure 400 {object} ErrorResponse "Ungültige Gradeen-ID"
// @Failure 404 {object} ErrorResponse "Grade nicht gefunden"
// @Router /grades/:id [get]
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

// CreateGradeHandler godoc
// @Summary create grade
// @Description create a new grade
// @Tags grade
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Param grade body model.Grade true "Grade object that needs to be added"
// @Success 201 {object} model.Grade
// @Failure 500 {object} ErrorResponse "Fehler beim Erstellen des Gradeen"
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Router /grades [post]
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

// UpdateGradeHandler godoc
// @Summary update grade
// @Description update an existing grade
// @Tags grade
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Param id path int true "Grade ID"
// @Param grade body model.Grade true "Updated grade object"
// @Success 200 {object} model.Grade
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Failure 404 {object} ErrorResponse "Grade nicht gefunden"
// @Router /grades/:id [put]
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

// DeleteGradeHandler godoc
// @Summary delete grade
// @Description delete an existing grade
// @Tags grade
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Param id path int true "Grade ID"
// @Success 204
// @Failure 400 {object} ErrorResponse "Ungültige Gradeen-ID"
// @Failure 500 {object} ErrorResponse "Fehler beim Löschen des Gradeen"
// @Router /grades/:id [delete]
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
