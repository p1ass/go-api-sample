package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki-kishi/go-api-sample/domain/model"
	"github.com/naoki-kishi/go-api-sample/usecase/repository"
	"net/http"
	"strconv"
)

type tagHandler struct {
	tagRepository repository.TagRepository
}

type TagHandler interface {
	CreateTag(c *gin.Context)
	GetTag(c *gin.Context)
	GetTags(c *gin.Context)
}

func NewTagHandler(eR repository.TagRepository) TagHandler {
	return &tagHandler{tagRepository: eR}
}

func (tH *tagHandler) GetTag(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	t, err := tH.tagRepository.FindByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}

func (tH *tagHandler) GetTags(c *gin.Context) {

	t, err := tH.tagRepository.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}

func (tH *tagHandler) CreateTag(c *gin.Context) {
	t := &model.Tag{}

	if err := c.Bind(t); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := tH.tagRepository.Store(t); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}
