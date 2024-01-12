package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/youtubeapi/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/youtubeapi/middleware"
	"gitlab.com/pragmaticreviews/golang-gin-poc/youtubeapi/service"
)

var (
	videoService    service.VideoService       = service.NewService()
	videoController controller.VideoController = controller.NewController(videoService)
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), middleware.BasicAuth())
	server.GET("/video", func(ctx *gin.Context) {
		videoController.FindVideo(ctx)
	})
	server.GET("/youtubeapi", func(ctx *gin.Context) {
		videoController.FindVideoApi(ctx)
	})

	server.Run(":8080")
}
