package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetScoresHandler gibt alle Score zurück
func GetScoresHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var score []model.Score
	if err := db.Preload("Grade").Find(&score).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Scores"})
	}

	return c.JSON(http.StatusOK, score)
}

// GetScoreHandler gibt einen bestimmten Score anhand der ID zurück
func GetScoreHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Score-ID"})
	}

	var score model.Score
	if err := db.Preload("Grade").First(&score, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Score nicht gefunden"})
	}

	return c.JSON(http.StatusOK, score)
}

// CreateScoreHandler erstellt einen neuen Scores
func CreateScoreHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	var score model.Score
	if err := c.Bind(&score); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Create(&score).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Erstellen des Scores"})
	}

	return c.JSON(http.StatusCreated, score)
}

// UpdateScoreHandler aktualisiert einen vorhandenen Score
func UpdateScoreHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Score-ID"})
	}

	var existingScore model.Score
	if err := db.First(&existingScore, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Score nicht gefunden"})
	}

	if err := c.Bind(&existingScore); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Anfrage"})
	}

	if err := db.Save(&existingScore).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Aktualisieren des Scores"})
	}

	return c.JSON(http.StatusOK, existingScore)
}

// DeleteScoreHandler löscht einen Score anhand der ID
func DeleteScoreHandler(c echo.Context) error {
	db := DB.GetDBInstance()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ungültige Score-ID"})
	}

	if err := db.Delete(&model.Score{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Löschen des Score"})
	}

	return c.NoContent(http.StatusNoContent)
}
