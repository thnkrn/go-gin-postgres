package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/thnkrn/go-gin-postgres/controllers"
	"gorm.io/gorm"
)

func SetCollectionRoutes(router *gin.RouterGroup, db *gorm.DB) {
	ctrls := controllers.DBController{Database: db}

	router.GET("campaigns", ctrls.GetCampaigns)
	router.GET("campaigns/:id", ctrls.GetCampaign)
	router.POST("campaigns", ctrls.BookCampaign)
	router.DELETE("campaigns/:id", ctrls.DeleteCampaign)
	router.PATCH("campaigns/:id", ctrls.UpdateCampaign)
}
