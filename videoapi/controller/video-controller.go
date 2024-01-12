package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/videoapi/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/videoapi/service"
	"net/http"
)

type VideoController interface {
	FindAll(ctx *gin.Context) []entity.Video
	Save(ctx *gin.Context) entity.Video
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll(ctx *gin.Context) []entity.Video {
	return c.service.FindAll(ctx)
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	return c.service.Save(ctx)
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll(ctx)
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
