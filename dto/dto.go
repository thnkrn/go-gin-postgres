package dto

type CampaignDTO struct {
	ID           uint   `json:"id,string,omitempty"`
	CampaignName string `json:"campaignName" binding:"required"`
	Date         string `json:"date" binding:"required"`
}
