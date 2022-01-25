package controllers

import (
	"net/http"
	"strconv"

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

	c.JSON(http.StatusOK, &campaigns)
}

func (db *DBController) GetCampaign(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	var campaign models.Campaigns

	db.Database.Find(&campaign, id)

	if int(campaign.ID) == id {
		c.JSON(http.StatusOK, &campaign)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
}

func (db *DBController) BookCampaign(c *gin.Context) {
	var campaign models.Campaigns

	if err := c.BindJSON(&campaign); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if result := db.Database.Create(&campaign); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusCreated, &campaign)
}

func (db *DBController) DeleteCampaign(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	if isValid := db.Database.First(&models.Campaigns{}, id); isValid.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Campaign is not booking yet",
		})
		return
	}

	if result := db.Database.Delete(&models.Campaigns{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campaign is deleted successfully"})
}

func (db *DBController) UpdateCampaign(c *gin.Context) {
	type request struct {
		CampaignName *string `json:"campaignName"`
		Date         *string `json:"date"`
	}

	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	if isValid := db.Database.First(&models.Campaigns{}, id); isValid.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Campaign is not booking yet",
		})
		return
	}

	campaign := new(models.Campaigns)
	db.Database.First(&campaign, id)

	if err := c.BindJSON(&campaign); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var body request

	if body.CampaignName != nil {
		campaign.CampaignName = *body.CampaignName
	}

	if body.Date != nil {
		campaign.Date = *body.Date
	}

	if result := db.Database.Save(&campaign); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &campaign)
}
