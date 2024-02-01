package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

// GetSchoolsHandler godoc
// @Summary get all schools
// @Description get all schools from db.
// @Tags school
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Success 200 {object} []model.School
// @Failure 500 {object} ErrorResponse "Fehler beim Abrufen der Schools"
// @Router /schools [get]
func GetSchoolsHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var schools []model.School
	// if err := db.Model(&model.School{}).Preload("Address").Preload("SchoolForm").Preload("Classes").Preload("Users").Find(&schools).Error; err != nil {
	if err := db.Model(&model.School{}).Preload(clause.Associations).Preload("Classes.Students").Preload("Classes.Students.Address").Preload("Users.Address").Find(&schools).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Schoolen"})
	}

	return c.JSON(http.StatusOK, schools)
}

// GetSchoolHandler godoc
// @Summary get one school by id
// @Description get one school from db by ID.
// @Tags school
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Param id path int true "School ID"
// @Success 200 {object} model.School
// @Failure 400 {object} ErrorResponse "Ungültige School-ID"
// @Failure 404 {object} ErrorResponse "School nicht gefunden"
// @Router /schools/:id [get]
func GetSchoolHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Schoolen-ID"})
	}

	var school model.School
	if err := db.Model(&model.School{}).Preload(clause.Associations).Preload("Classes.Students").Preload("Classes.Students.Address").Preload("Users.Address").Find(&school, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "School nicht gefunden"})
	}

	return c.JSON(http.StatusOK, school)
}

// CreateSchoolHandler godoc
// @Summary create school
// @Description create a new school
// @Tags school
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Param school body model.School true "School object that needs to be added"
// @Success 201 {object} model.School
// @Failure 500 {object} ErrorResponse "Fehler beim Erstellen des Schools"
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Router /schools [post]
func CreateSchoolHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var school model.School
	if err := c.Bind(&school); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&school).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des Schoolen"})
	}

	return c.JSON(http.StatusCreated, school)
}

// UpdateSchoolHandler godoc
// @Summary update school
// @Description update an existing school
// @Tags school
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Param id path int true "School ID"
// @Param school body model.School true "Updated school object"
// @Success 200 {object} model.School
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Failure 404 {object} ErrorResponse "School nicht gefunden"
// @Router /schools/:id [put]
func UpdateSchoolHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Schoolen-ID"})
	}

	var existingSchool model.School
	if err := db.First(&existingSchool, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "School nicht gefunden"})
	}

	if err := c.Bind(&existingSchool); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingSchool).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des Schoolen"})
	}

	return c.JSON(http.StatusOK, existingSchool)
}

// DeleteSchoolHandler godoc
// @Summary delete school
// @Description delete an existing school
// @Tags school
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Param id path int true "School ID"
// @Success 204
// @Failure 400 {object} ErrorResponse "Ungültige School-ID"
// @Failure 500 {object} ErrorResponse "Fehler beim Löschen des Schools"
// @Router /schools/:id [delete]
func DeleteSchoolHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Schoolen-ID"})
	}

	if err := db.Delete(&model.School{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des Schoolen"})
	}

	return c.NoContent(http.StatusNoContent)
}
