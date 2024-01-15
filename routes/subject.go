package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetSubjectsHandler godoc
// @Summary get all subjects
// @Description get all subjects from db.
// @Tags subject
// @Accept application/json
// @Produce json
// @Success 200 {object} []model.Subject
// @Failure 500 {object} ErrorResponse "Fehler beim Abrufen der Subjecten"
// @Router /subjects [get]
func GetSubjectsHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var subjects []model.Subject
	if err := db.Find(&subjects).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Subjecten"})
	}

	return c.JSON(http.StatusOK, subjects)
}

// GetSubjectHandler godoc
// @Summary get one subject by id
// @Description get one subject from db by ID.
// @Tags subject
// @Accept application/json
// @Produce json
// @Param id path int true "Subject ID"
// @Success 200 {object} model.Subject
// @Failure 400 {object} ErrorResponse "Ungültige Subjecten-ID"
// @Failure 404 {object} ErrorResponse "Subject nicht gefunden"
// @Router /subjects/:id [get]
func GetSubjectHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Subjecten-ID"})
	}

	var subject model.Subject
	if err := db.Find(&subject, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Subject nicht gefunden"})
	}

	return c.JSON(http.StatusOK, subject)
}

// CreateSubjectHandler godoc
// @Summary create subject
// @Description create a new subject
// @Tags subject
// @Accept application/json
// @Produce json
// @Param subject body model.Subject true "Subject object that needs to be added"
// @Success 201 {object} model.Subject
// @Failure 500 {object} ErrorResponse "Fehler beim Erstellen des Subjecten"
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Router /subjects [post]
func CreateSubjectHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var subject model.Subject
	if err := c.Bind(&subject); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&subject).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des Subjecten"})
	}

	return c.JSON(http.StatusCreated, subject)
}

// UpdateSubjectHandler godoc
// @Summary update subject
// @Description update an existing subject
// @Tags subject
// @Accept application/json
// @Produce json
// @Param id path int true "Subject ID"
// @Param subject body model.Subject true "Updated subject object"
// @Success 200 {object} model.Subject
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Failure 404 {object} ErrorResponse "Subject nicht gefunden"
// @Router /subjects/:id [put]
func UpdateSubjectHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Subjecten-ID"})
	}

	var existingSubject model.Subject
	if err := db.First(&existingSubject, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Subject nicht gefunden"})
	}

	if err := c.Bind(&existingSubject); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingSubject).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des Subjecten"})
	}

	return c.JSON(http.StatusOK, existingSubject)
}

// DeleteSubjectHandler godoc
// @Summary delete subject
// @Description delete an existing subject
// @Tags subject
// @Accept application/json
// @Produce json
// @Param id path int true "Subject ID"
// @Success 204
// @Failure 400 {object} ErrorResponse "Ungültige Subjecten-ID"
// @Failure 500 {object} ErrorResponse "Fehler beim Löschen des Subjecten"
// @Router /subjects/:id [delete]
func DeleteSubjectHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Subjecten-ID"})
	}

	if err := db.Delete(&model.Subject{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des Subjecten"})
	}

	return c.NoContent(http.StatusNoContent)
}
