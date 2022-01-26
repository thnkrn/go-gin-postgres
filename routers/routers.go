package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/thnkrn/go-gin-postgres/controllers"
)

func SetCampaignsRoutes(router *gin.RouterGroup, campaignAPI controllers.CampaignAPI) {
	router.GET("campaigns", campaignAPI.FindAll)
	router.GET("campaigns/:id", campaignAPI.FindByID)
	router.POST("campaigns", campaignAPI.Create)
	router.DELETE("campaigns/:id", campaignAPI.Delete)
	router.PUT("campaigns/:id", campaignAPI.Update)

	// router.GET("campaigns", ctrls.GetCampaigns)
	// router.GET("campaigns/:id", ctrls.GetCampaign)
	// router.POST("campaigns", ctrls.BookCampaign)
	// router.DELETE("campaigns/:id", ctrls.DeleteCampaign)
	// router.PATCH("campaigns/:id", ctrls.UpdateCampaign)
}
