// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/thnkrn/go-gin-postgres/controllers"
	"github.com/thnkrn/go-gin-postgres/repository"
	"github.com/thnkrn/go-gin-postgres/services"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitCampaignAPI(db *gorm.DB) controllers.CampaignAPI {
	campaignRepository := repository.ProvideCampaignRepository(db)
	campaignService := services.ProvideCampaignService(campaignRepository)
	campaignAPI := controllers.ProvideCampaignAPI(campaignService)
	return campaignAPI
}
