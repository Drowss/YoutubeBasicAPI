package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gitlab.com/pragmaticreviews/golang-gin-poc/videoapi/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/videoapi/validators"
	"net/http"
)

type VideoService interface {
	Save(ctx *gin.Context) entity.Video
	FindAll(ctx *gin.Context) []entity.Video
}

type videoService struct {
	videos []entity.Video
}

var validate *validator.Validate

func New() VideoService {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoService{}
}

func (service *videoService) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return video
	}
	err = validate.Struct(video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return video
	}
	service.videos = append(service.videos, video)
	ctx.JSON(http.StatusOK, video)
	return video
}

func (service *videoService) FindAll(ctx *gin.Context) []entity.Video {
	//ctx.JSON(http.StatusOK, service.videos)
	return service.videos
}
