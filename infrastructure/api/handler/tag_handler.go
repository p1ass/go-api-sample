package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki-kishi/go-api-sample/domain/model"
	"github.com/naoki-kishi/go-api-sample/usecase/repository"
	"net/http"
)

type tagHandler struct {
	tagRepository repository.TagRepository
}

type TagHandler interface {
	CreateTag(c *gin.Context)
	GetTags(c *gin.Context)
}

func NewTagHandler(eR repository.TagRepository) TagHandler {
	return &tagHandler{tagRepository: eR}
}

func (tH *tagHandler) GetTags(c *gin.Context) {
	t := []*model.Tag{}

	ts, err := tH.tagRepository.FindAll(t)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ts)
}

func (tH *tagHandler) CreateTag(c *gin.Context) {
	tag := &model.Tag{}

	if err := c.Bind(tag); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := tH.tagRepository.Store(tag); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tag)
}
