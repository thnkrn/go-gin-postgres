package models

import (
	"gorm.io/gorm"
)

type Campaigns struct {
	gorm.Model
	CampaignName string `json:"campaignName"`
	Date         string `json:"date"`
}
