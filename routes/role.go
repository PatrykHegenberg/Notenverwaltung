package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetRolesHandler gibt alle Role zurück
func GetRolesHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var role []model.Role
	if err := db.Find(&role).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Roles"})
	}

	return c.JSON(http.StatusOK, role)
}

// GetRoleHandler gibt einen bestimmten Role anhand der ID zurück
func GetRoleHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Role-ID"})
	}

	var role model.Role
	if err := db.First(&role, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Role nicht gefunden"})
	}

	return c.JSON(http.StatusOK, role)
}

// CreateRoleHandler erstellt einen neuen Roles
func CreateRoleHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var role model.Role
	if err := c.Bind(&role); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&role).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des Roles"})
	}

	return c.JSON(http.StatusCreated, role)
}

// UpdateRoleHandler aktualisiert einen vorhandenen Role
func UpdateRoleHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Role-ID"})
	}

	var existingRole model.Role
	if err := db.First(&existingRole, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Role nicht gefunden"})
	}

	if err := c.Bind(&existingRole); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingRole).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des Roles"})
	}

	return c.JSON(http.StatusOK, existingRole)
}

// DeleteRoleHandler löscht einen Role anhand der ID
func DeleteRoleHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Role-ID"})
	}

	if err := db.Delete(&model.Role{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des Role"})
	}

	return c.NoContent(http.StatusNoContent)
}
