package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/thnkrn/go-gin-postgres/dto"
	mapper "github.com/thnkrn/go-gin-postgres/mapper"
	models "github.com/thnkrn/go-gin-postgres/models"
	services "github.com/thnkrn/go-gin-postgres/services"
)

// type DBController struct {
// 	Database *gorm.DB
// }

type CampaignAPI struct {
	CampaignService services.CampaignService
}

func ProvideCampaignAPI(p services.CampaignService) CampaignAPI {
	return CampaignAPI{CampaignService: p}
}

func (p *CampaignAPI) FindAll(c *gin.Context) {
	campaigns := p.CampaignService.FindAll()

	c.JSON(http.StatusOK, mapper.ToCampaignDTOs(campaigns))
}

// func (db *DBController) GetCampaigns(c *gin.Context) {
// 	var campaigns []models.Campaigns
// 	db.Database.Find(&campaigns)

// 	c.JSON(http.StatusOK, &campaigns)
// }

func (p *CampaignAPI) FindByID(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	campaign := p.CampaignService.FindByID(uint(id))

	if int(campaign.ID) == id {
		c.JSON(http.StatusOK, mapper.ToCampaignDTO(campaign))
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
}

// func (db *DBController) GetCampaign(c *gin.Context) {
// 	paramsId := c.Param("id")
// 	id, err := strconv.Atoi(paramsId)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "cannot parse id",
// 		})
// 		return
// 	}

// 	var campaign models.Campaigns

// 	db.Database.Find(&campaign, id)

// 	if int(campaign.ID) == id {
// 		c.JSON(http.StatusOK, &campaign)
// 		return
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
// }

func (p *CampaignAPI) Create(c *gin.Context) {
	var campaignDTO dto.CampaignDTO
	err := c.BindJSON(&campaignDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdCampaign := p.CampaignService.Save(mapper.ToCampaign(campaignDTO))

	c.JSON(http.StatusOK, mapper.ToCampaignDTO(createdCampaign))
}

// func (db *DBController) BookCampaign(c *gin.Context) {
// 	var campaign models.Campaigns

// 	if err := c.BindJSON(&campaign); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}

// 	if result := db.Database.Create(&campaign); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, &campaign)
// }

func (p *CampaignAPI) Delete(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	campaign := p.CampaignService.FindByID(uint(id))

	if campaign == (models.Campaigns{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Campaign is not booking yet",
		})
		return
	}

	p.CampaignService.Delete(campaign)

	c.JSON(http.StatusOK, gin.H{"message": "Campaign is deleted successfully"})
}

// func (db *DBController) DeleteCampaign(c *gin.Context) {
// 	paramsId := c.Param("id")
// 	id, err := strconv.Atoi(paramsId)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "cannot parse id",
// 		})
// 		return
// 	}

// 	if isValid := db.Database.First(&models.Campaigns{}, id); isValid.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Campaign is not booking yet",
// 		})
// 		return
// 	}

// 	if result := db.Database.Delete(&models.Campaigns{}, id); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Campaign is deleted successfully"})
// }

func (p *CampaignAPI) Update(c *gin.Context) {
	var campaignDTO dto.CampaignDTO
	err := c.BindJSON(&campaignDTO)

	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	product := p.CampaignService.FindByID(uint(id))

	if product == (models.Campaigns{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Campaign is not booking yet",
		})
		return
	}

	product.CampaignName = campaignDTO.CampaignName
	product.Date = campaignDTO.Date
	p.CampaignService.Save(product)

	c.JSON(http.StatusOK, product)
}

// func (db *DBController) UpdateCampaign(c *gin.Context) {
// 	type request struct {
// 		CampaignName *string `json:"campaignName"`
// 		Date         *string `json:"date"`
// 	}

// 	paramsId := c.Param("id")
// 	id, err := strconv.Atoi(paramsId)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "cannot parse id",
// 		})
// 		return
// 	}

// 	if isValid := db.Database.First(&models.Campaigns{}, id); isValid.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Campaign is not booking yet",
// 		})
// 		return
// 	}

// 	campaign := new(models.Campaigns)
// 	db.Database.First(&campaign, id)

// 	if err := c.BindJSON(&campaign); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}

// 	var body request

// 	if body.CampaignName != nil {
// 		campaign.CampaignName = *body.CampaignName
// 	}

// 	if body.Date != nil {
// 		campaign.Date = *body.Date
// 	}

// 	if result := db.Database.Save(&campaign); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, &campaign)
// }
