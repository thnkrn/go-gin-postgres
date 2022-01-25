package models

import (
	"gorm.io/gorm"
)

type Campaigns struct {
	gorm.Model
	CampaignName string `json:"campaignName" binding:"required"`
	Date         string `json:"date" binding:"required"`
}
