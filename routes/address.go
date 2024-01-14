package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetAddressesHandler is a function that handles the request to get addresses.
//
// It takes a parameter of type `echo.Context`.
// It returns an error.
func GetAddressesHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var addresses []model.Address
	if err := db.Find(&addresses).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Adressen"})
	}
	return c.JSON(http.StatusOK, addresses)
}

// GetAddressHandler handles the request to get an address.
//
// It takes a `c` object of type `echo.Context` as a parameter.
// It returns an error.
func GetAddressHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Adressen-ID"})
	}

	var address model.Address
	if err := db.Find(&address, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Adresse nicht gefunden"})
	}

	return c.JSON(http.StatusOK, address)
}

// CreateAddressHandler handles the creation of an address in the application.
//
// It takes a `c` parameter of type `echo.Context` which represents the current HTTP request context.
// The function returns an error.
func CreateAddressHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var address model.Address
	if err := c.Bind(&address); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&address).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des Addressen"})
	}

	return c.JSON(http.StatusCreated, address)
}

// UpdateAddressHandler updates an address in the database.
//
// The function takes a parameter of type echo.Context, which represents the
// context of the HTTP request. The function returns an error.
func UpdateAddressHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Addressen-ID"})
	}

	var existingAddress model.Address
	if err := db.First(&existingAddress, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Address nicht gefunden"})
	}

	if err := c.Bind(&existingAddress); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingAddress).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des Addressen"})
	}

	return c.JSON(http.StatusOK, existingAddress)
}

func DeleteAddressHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Addressen-ID"})
	}

	if err := db.Delete(&model.Address{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des Addressen"})
	}

	return c.NoContent(http.StatusNoContent)
}
