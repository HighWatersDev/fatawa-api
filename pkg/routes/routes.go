package routes

import (
	"fatawa-api/pkg/controllers"
	"github.com/gin-gonic/gin"
)

type FatwaRouteController struct {
	fatwaController controllers.FatwaController
}

func NewFatwaControllerRoute(fatwaController controllers.FatwaController) FatwaRouteController {
	return FatwaRouteController{fatwaController}
}

func (r *FatwaRouteController) FatwaRoute(rg *gin.RouterGroup) {
	router := rg.Group("/fatawa")

	router.GET("/", r.fatwaController.FindFatawa)
	router.GET("/:fatwaId", r.fatwaController.FindFatwaById)
	router.POST("/", r.fatwaController.CreateFatwa)
	router.PATCH("/:fatwaId", r.fatwaController.UpdateFatwa)
	router.DELETE("/:fatwaId", r.fatwaController.DeleteFatwa)
}
