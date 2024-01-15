package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetClassesHandler godoc
// @Summary get all classes
// @Description get all classes from db.
// @Tags class
// @Accept application/json
// @Produce json
// @Success 200 {object} []model.Class
// @Failure 500 {object} ErrorResponse "Fehler beim Abrufen der Classes"
// @Router /classes [get]
func GetClasssHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var class []model.Class
	if err := db.Model(&model.Class{}).Preload("Students").Find(&class).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Classs"})
	}

	return c.JSON(http.StatusOK, class)
}

// GetClassHandler godoc
// @Summary get one class by id
// @Description get one class from db by ID.
// @Tags class
// @Accept application/json
// @Produce json
// @Param id path int true "Class ID"
// @Success 200 {object} model.Class
// @Failure 400 {object} ErrorResponse "Ungültige Class-ID"
// @Failure 404 {object} ErrorResponse "Class nicht gefunden"
// @Router /classes/:id [get]
func GetClassHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Class-ID"})
	}

	var class model.Class
	if err := db.Model(&model.Class{}).Preload("Students").First(&class, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Class nicht gefunden"})
	}

	return c.JSON(http.StatusOK, class)
}

// CreateClassHandler godoc
// @Summary create class
// @Description create a new class
// @Tags class
// @Accept application/json
// @Produce json
// @Param class body model.Class true "Class object that needs to be added"
// @Success 201 {object} model.Class
// @Failure 500 {object} ErrorResponse "Fehler beim Erstellen des Classes"
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Router /classes [post]
func CreateClassHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var class model.Class
	if err := c.Bind(&class); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&class).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des Classs"})
	}

	return c.JSON(http.StatusCreated, class)
}

// UpdateClassHandler godoc
// @Summary update class
// @Description update an existing class
// @Tags class
// @Accept application/json
// @Produce json
// @Param id path int true "Class ID"
// @Param class body model.Class true "Updated class object"
// @Success 200 {object} model.Class
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Failure 404 {object} ErrorResponse "Class nicht gefunden"
// @Router /classes/:id [put]
func UpdateClassHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Class-ID"})
	}

	var existingClass model.Class
	if err := db.First(&existingClass, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Class nicht gefunden"})
	}

	if err := c.Bind(&existingClass); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingClass).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des Classs"})
	}

	return c.JSON(http.StatusOK, existingClass)
}

// DeleteClassHandler godoc
// @Summary delete class
// @Description delete an existing class
// @Tags class
// @Accept application/json
// @Produce json
// @Param id path int true "Class ID"
// @Success 204
// @Failure 400 {object} ErrorResponse "Ungültige Class-ID"
// @Failure 500 {object} ErrorResponse "Fehler beim Löschen des Class"
// @Router /classes/:id [delete]
func DeleteClassHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Class-ID"})
	}

	if err := db.Delete(&model.Class{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des Class"})
	}

	return c.NoContent(http.StatusNoContent)
}
