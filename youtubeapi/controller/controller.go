package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/youtubeapi/service"
)

type VideoController interface {
	FindVideo(ctx *gin.Context)
	FindVideoApi(ctx *gin.Context)
}

type Controller struct {
	service service.VideoService
}

func NewController(service service.VideoService) service.VideoService {
	return &Controller{
		service: service,
	}
}

func (c *Controller) FindVideo(ctx *gin.Context) {
	c.service.FindVideo(ctx)
}

func (c *Controller) FindVideoApi(ctx *gin.Context) {
	c.service.FindVideoApi(ctx)
}
