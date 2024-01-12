package service

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/youtubeapi/entity"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"net/http"
)

type VideoService interface {
	FindVideo(ctx *gin.Context)
	FindVideoApi(ctx *gin.Context)
}

type Service struct {
	video entity.Video
}

func NewService() VideoService {
	return &Service{}
}

func (s *Service) FindVideo(ctx *gin.Context) {
	videoId := ctx.Query("id")
	part := ctx.Query("part")
	apiKey := ctx.Query("key")

	resp, err := http.Get("https://www.googleapis.com/youtube/v3/videos?id=" + videoId + "&part=" + part + "&key=" + apiKey)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	var video entity.Video

	if err := json.NewDecoder(resp.Body).Decode(&video); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error"})
		return
	}

	ctx.JSON(http.StatusOK, video)
}

func (s *Service) FindVideoApi(ctx *gin.Context) {
	youtubeService, err := youtube.NewService(context.Background(), option.WithAPIKey("AIzaSyA2WdKwaT8ysUGiJhwj9XjzhjxsSg15uSk"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se creo el servicio"})
		return
	}
	videoResponse, err := youtubeService.Videos.List([]string{"snippet"}).Id(ctx.Query("id")).Do()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se encuentra el video"})
		return
	}
	video := videoResponse.Items[0]
	title := video.Snippet.Title
	description := video.Snippet.Description

	ctx.JSON(http.StatusOK, gin.H{
		"title":       title,
		"description": description,
	})
}
