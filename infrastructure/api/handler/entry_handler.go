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
}

type EntryHandler interface {
	CreateEntry(c *gin.Context)
	GetEntries(c *gin.Context)
	GetEntry(c *gin.Context)
}

func NewEntryHandler(eR repository.EntryRepository) EntryHandler {
	return &entryHandler{entryRepository: eR}
}

func (eH *entryHandler) GetEntry(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	es, err := eH.entryRepository.FindByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, es)
}

func (eH *entryHandler) GetEntries(c *gin.Context) {
	e := []*model.Entry{}

	es, err := eH.entryRepository.FindAll(e)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, es)
}

func (eH *entryHandler) CreateEntry(c *gin.Context) {
	entry := &model.Entry{}

	if err := c.Bind(entry); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := eH.entryRepository.Store(entry); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, entry)
}
