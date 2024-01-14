package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

// GetSchoolsHandler handles the request to get the schools.
//
// c: the echo context object.
// Returns an error if there was a problem with the request.
func GetSchoolsHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var schools []model.School
	// if err := db.Model(&model.School{}).Preload("Address").Preload("SchoolForm").Preload("Classes").Preload("Users").Find(&schools).Error; err != nil {
	if err := db.Model(&model.School{}).Preload(clause.Associations).Preload("Classes.Students").Preload("Classes.Students.Address").Preload("Users.Address").Find(&schools).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Schoolen"})
	}

	return c.JSON(http.StatusOK, schools)
}

// GetSchoolHandler is a function that handles the request to get a school.
//
// It takes a `c` parameter of type `echo.Context`, which represents the HTTP request context.
// It returns an error.
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

// CreateSchoolHandler handles the creation of a school.
//
// It takes a `c` parameter of type `echo.Context` which represents the HTTP request context.
// It returns an error.
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

// UpdateSchoolHandler is a function that handles the update of a school in the API.
//
// It takes a context object as a parameter.
// It returns an error.
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

// DeleteSchoolHandler is a function that handles the deletion of a school.
//
// It takes a parameter of type echo.Context, which represents the HTTP request and response context.
// It returns an error type.
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
