package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thnkrn/go-gin-postgres/models"
	"gorm.io/gorm"
)

type DBController struct {
	Database *gorm.DB
}

func (db *DBController) GetCampaigns(c *gin.Context) {
	var campaigns []models.Campaigns
	db.Database.Find(&campaigns)

	c.JSON(http.StatusOK, gin.H{"results": &campaigns})
}
