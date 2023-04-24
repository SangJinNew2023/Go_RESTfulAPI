package router

import (
	"CRUD02/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.TagsController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcom home")
	})

	baseRouter := router.Group("/api")
	tagRotuer := baseRouter.Group("/tags")
	tagRotuer.GET("", tagsController.FindAll)
	tagRotuer.GET("/:tagId", tagsController.FindById)
	tagRotuer.POST("", tagsController.Create)
	tagRotuer.PATCH("/:tagId", tagsController.Update)
	tagRotuer.GET("/:tagId", tagsController.Delete)

	return router

}
