package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetAddressesHandler godoc
// @Summary get all addresses
// @Description get all addresses from db.
// @Tags address
// @Accept application/json
// @Produce json
// @Success 200 {object} []model.Address
// @Failure 500 {object} ErrorResponse "Fehler beim Abrufen der Adressen"
// @Router /addresses [get]
func GetAddressesHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var addresses []model.Address
	if err := db.Find(&addresses).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Adressen"})
	}
	return c.JSON(http.StatusOK, addresses)
}

// GetAddressHandler godoc
// @Summary get one address by id
// @Description get one address from db by ID.
// @Tags address
// @Accept application/json
// @Produce json
// @Param id path int true "Address ID"
// @Success 200 {object} model.Address
// @Failure 400 {object} ErrorResponse "Ungültige Adressen-ID"
// @Failure 404 {object} ErrorResponse "Adresse nicht gefunden"
// @Router /addresses/:id [get]
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

// CreateAddressHandler godoc
// @Summary create address
// @Description create a new address
// @Tags address
// @Accept application/json
// @Produce json
// @Param address body model.Address true "Address object that needs to be added"
// @Success 201 {object} model.Address
// @Failure 500 {object} ErrorResponse "Fehler beim Erstellen des Addressen"
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Router /addresses [post]
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

// UpdateAddressHandler godoc
// @Summary update address
// @Description update an existing address
// @Tags address
// @Accept application/json
// @Produce json
// @Param id path int true "Address ID"
// @Param address body model.Address true "Updated address object"
// @Success 200 {object} model.Address
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Failure 404 {object} ErrorResponse "Address nicht gefunden"
// @Router /addresses/:id [put]
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

// DeleteAddressHandler godoc
// @Summary delete address
// @Description delete an existing address
// @Tags address
// @Accept application/json
// @Produce json
// @Param id path int true "Address ID"
// @Success 204
// @Failure 400 {object} ErrorResponse "Ungültige Addressen-ID"
// @Failure 500 {object} ErrorResponse "Fehler beim Löschen des Addressen"
// @Router /addresses/:id [delete]
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
