package dto

// import "time"

type CampaignDTO struct {
	ID           uint   `json:"id,string,omitempty"`
	CampaignName string `json:"campaignName" binding:"required"`
	Date         string `json:"date" binding:"required"`
	// NOTE: Optional
	// CreatedAt    time.Time
	// UpdatedAt    time.Time
}
