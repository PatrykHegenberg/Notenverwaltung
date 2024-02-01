package routes

import (
	"net/http"
	"strconv"

	DB "github.com/PatrykHegenberg/Notenverwaltung/database"
	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"github.com/labstack/echo/v4"
)

// GetScoresHandler godoc
// @Summary get all scores
// @Description get all scores from db.
// @Tags score
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Success 200 {object} []model.Score
// @Failure 500 {object} ErrorResponse "Fehler beim Abrufen der Scores"
// @Router /scores [get]
func GetScoresHandler(c echo.Context) error {
	db := DB.GetDBInstance() // Funktion zum Abrufen der Datenbankinstanz

	var score []model.Score
	if err := db.Preload("Grade").Find(&score).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Fehler beim Abrufen der Scores"})
	}

	return c.JSON(http.StatusOK, score)
}

// GetScoreHandler godoc
// @Summary get one score by id
// @Description get one score from db by ID.
// @Tags score
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Param id path int true "Score ID"
// @Success 200 {object} model.Score
// @Failure 400 {object} ErrorResponse "Ungültige Score-ID"
// @Failure 404 {object} ErrorResponse "Score nicht gefunden"
// @Router /scores/:id [get]
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

// CreateScoreHandler godoc
// @Summary create score
// @Description create a new score
// @Tags score
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Param score body model.Score true "Score object that needs to be added"
// @Success 201 {object} model.Score
// @Failure 500 {object} ErrorResponse "Fehler beim Erstellen des Scores"
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Router /scores [post]
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

// UpdateScoreHandler godoc
// @Summary update score
// @Description update an existing score
// @Tags score
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Param id path int true "Score ID"
// @Param score body model.Score true "Updated score object"
// @Success 200 {object} model.Score
// @Failure 400 {object} ErrorResponse "Ungültige Anfrage"
// @Failure 404 {object} ErrorResponse "Score nicht gefunden"
// @Router /scores/:id [put]
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

// DeleteScoreHandler godoc
// @Summary delete score
// @Description delete an existing score
// @Tags score
// @securityDefinitions.basic BasicAuth
// @Accept application/json
// @Produce json
// @Param id path int true "Score ID"
// @Success 204
// @Failure 400 {object} ErrorResponse "Ungültige Score-ID"
// @Failure 500 {object} ErrorResponse "Fehler beim Löschen des Scores"
// @Router /scores/:id [delete]
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
