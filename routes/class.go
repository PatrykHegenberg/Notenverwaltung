package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetClasssHandler is a function that handles the retrieval of class data.
//
// It takes an echo.Context as a parameter.
// It returns an error.
func GetClasssHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var class []model.Class
	if err := db.Model(&model.Class{}).Preload("Students").Find(&class).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Classs"})
	}

	return c.JSON(http.StatusOK, class)
}

// GetClassHandler is a function that handles the retrieval of a class.
//
// It takes a parameter `c` of type `echo.Context` which represents the context
// of the HTTP request.
// It returns an error.
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

// CreateClassHandler handles the creation of a class.
//
// It takes a `c` parameter of type `echo.Context` which represents the Echo context.
// It returns an error.
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

// UpdateClassHandler updates a class in the database.
//
// c: The echo context.
// Returns an error if any.
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

// DeleteClassHandler is a function that handles the deletion of a class.
//
// It takes a `c` parameter of type `echo.Context` that represents the current HTTP request context.
// The function returns an error type.
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
