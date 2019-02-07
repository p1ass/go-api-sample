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

func (eH *tagHandler) GetTags(c *gin.Context) {
	u := []*model.Tag{}

	us, err := eH.tagRepository.FindAll(u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, us)
}

func (eH *tagHandler) CreateTag(c *gin.Context) {
	return
}
