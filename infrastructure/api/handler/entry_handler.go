package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki-kishi/go-api-sample/domain/model"
	"github.com/naoki-kishi/go-api-sample/usecase/repository"
	"net/http"
	"strconv"
)

type entryHandler struct {
	entryRepository repository.EntryRepository
	tagRepository   repository.TagRepository
}

type EntryHandler interface {
	CreateEntry(c *gin.Context)
	GetEntries(c *gin.Context)
	GetEntry(c *gin.Context)
	UpdateEntry(c *gin.Context)
	DeleteEntry(c *gin.Context)
}

func NewEntryHandler(eR repository.EntryRepository, tR repository.TagRepository) EntryHandler {
	return &entryHandler{entryRepository: eR, tagRepository: tR}
}

func (eH *entryHandler) GetEntry(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	e, err := eH.entryRepository.FindByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

func (eH *entryHandler) GetEntries(c *gin.Context) {

	e, err := eH.entryRepository.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

func (eH *entryHandler) CreateEntry(c *gin.Context) {
	e := &model.Entry{}

	if err := c.Bind(e); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := eH.tagRepository.FindOrCreateAll(e.Tags); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := eH.entryRepository.Store(e); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

func (eH *entryHandler) UpdateEntry(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	e := &model.Entry{ID: id}

	if err := c.Bind(e); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := eH.tagRepository.FindOrCreateAll(e.Tags); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := eH.entryRepository.Update(e); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

func (eH *entryHandler) DeleteEntry(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := eH.entryRepository.Delete(&model.Entry{ID: id}); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "ok"})
}
