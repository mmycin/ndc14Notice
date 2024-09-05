package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mmycin/ndc14notice/controller"
)

func Router() *gin.Engine {
	r := gin.Default()
	
	r.GET("/notice", controller.GetAllNotice)
	return r
}