package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki-kishi/go-api-sample/domain/model"
	"github.com/naoki-kishi/go-api-sample/usecase/repository"
	"net/http"
)

type entryHandler struct {
	entryRepository repository.EntryRepository
}

type EntryHandler interface {
	CreateEntry(c *gin.Context)
	GetEntries(c *gin.Context)
}

func NewEntryHandler(eR repository.EntryRepository) EntryHandler {
	return &entryHandler{entryRepository: eR}
}

func (eH *entryHandler) GetEntries(c *gin.Context) {
	u := []*model.Entry{}

	us, err := eH.entryRepository.FindAll(u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, us)
}

func (eH *entryHandler) CreateEntry(c *gin.Context) {
	return
}
