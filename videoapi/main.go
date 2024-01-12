package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/videoapi/controller"
	middleware2 "gitlab.com/pragmaticreviews/golang-gin-poc/videoapi/middleware"
	"gitlab.com/pragmaticreviews/golang-gin-poc/videoapi/service"
	"io"
	"os"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()
	server.Static("/css", ".templates/css")
	server.LoadHTMLGlob("templates/*.html")
	server.Use(gin.Recovery(), middleware2.Logger(), middleware2.BasicAuth())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/posts", func(context *gin.Context) {
			context.JSON(200, videoController.FindAll(context))
		})

		apiRoutes.POST("/posts", func(context *gin.Context) {
			videoController.Save(context)
		})
	}

	viewRoutes := server.Group("/view") //Muestra los videos publicados
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8080")

}
