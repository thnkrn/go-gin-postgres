package services

import (
	models "github.com/thnkrn/go-gin-postgres/models"
	repository "github.com/thnkrn/go-gin-postgres/repository"
)

type CampaignService struct {
	CampaignRepository repository.CampaignRepository
}

func ProvideCampaignService(p repository.CampaignRepository) CampaignService {
	return CampaignService{CampaignRepository: p}
}

func (p *CampaignService) FindAll() []models.Campaigns {
	return p.CampaignRepository.FindAll()
}

func (p *CampaignService) FindByID(id uint) models.Campaigns {
	return p.CampaignRepository.FindByID(id)
}

func (p *CampaignService) Save(campaign models.Campaigns) models.Campaigns {
	p.CampaignRepository.Save(campaign)

	return campaign
}

func (p *CampaignService) Delete(campaign models.Campaigns) {
	p.CampaignRepository.Delete(campaign)
}
