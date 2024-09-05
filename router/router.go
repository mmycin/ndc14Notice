package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mmycin/ndc14notice/controller"
	"github.com/gin-contrib/cors"
)

func Router() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
    config.ExposeHeaders = []string{"Content-Length"}
	r.Use(cors.New(config))
    
	r.GET("/notice", controller.GetAllNotice)
	return r
}