//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	controllers "github.com/thnkrn/go-gin-postgres/controllers"
	repository "github.com/thnkrn/go-gin-postgres/repository"
	services "github.com/thnkrn/go-gin-postgres/services"
	"gorm.io/gorm"
)

func InitCampaignAPI(db *gorm.DB) controllers.CampaignAPI {
	wire.Build(repository.ProvideCampaignRepository, services.ProvideCampaignService, controllers.ProvideCampaignAPI)

	return controllers.CampaignAPI{}
}
