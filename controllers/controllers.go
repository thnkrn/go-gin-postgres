package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/thnkrn/go-gin-postgres/models"
	"gorm.io/gorm"
)

type DBController struct {
	Database *gorm.DB
}

func (db *DBController) GetCampaigns(c *gin.Context) {
	var campaigns []models.Campaigns
	db.Database.Find(&campaigns)

	c.JSON(http.StatusOK, gin.H{"message": &campaigns})
}

func (db *DBController) BookCampaign(c *gin.Context) {
	var campaign models.Campaigns

	if err := c.BindJSON(&campaign); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	if result := db.Database.Create(&campaign); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	}

	c.JSON(http.StatusCreated, &campaign)
}

func (db *DBController) DeleteCampaign(c *gin.Context) {
	id := c.Param("id")

	if result := db.Database.Delete(&models.Campaigns{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campaign is deleted successfully"})
}
