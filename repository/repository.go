package repository

import (
	models "github.com/thnkrn/go-gin-postgres/models"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	DB *gorm.DB
}

func ProvideCampaignRepostiory(DB *gorm.DB) CampaignRepository {
	return CampaignRepository{DB: DB}
}

func (p *CampaignRepository) FindAll() []models.Campaigns {
	var campaigns []models.Campaigns
	p.DB.Find(&campaigns)

	return campaigns
}

func (p *CampaignRepository) FindByID(id uint) models.Campaigns {
	var campaign models.Campaigns
	p.DB.First(&campaign, id)

	return campaign
}

func (p *CampaignRepository) Save(campaign models.Campaigns) models.Campaigns {
	p.DB.Save(&campaign)

	return campaign
}

func (p *CampaignRepository) Delete(campaign models.Campaigns) {
	p.DB.Delete(&campaign)
}
