package mapper

import (
	dto "github.com/thnkrn/go-gin-postgres/dto"
	models "github.com/thnkrn/go-gin-postgres/models"
)

func ToCampaign(campaignDTO dto.CampaignDTO) models.Campaigns {
	return models.Campaigns{CampaignName: campaignDTO.CampaignName, Date: campaignDTO.Date}
}

func ToCampaignDTO(campaign models.Campaigns) dto.CampaignDTO {
	return dto.CampaignDTO{ID: campaign.ID, CampaignName: campaign.CampaignName, Date: campaign.Date}
}

func ToCampaignDTOs(campaigns []models.Campaigns) []dto.CampaignDTO {
	campaigndtos := make([]dto.CampaignDTO, len(campaigns))

	for i, itm := range campaigns {
		campaigndtos[i] = ToCampaignDTO(itm)
	}

	return campaigndtos
}
