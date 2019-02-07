package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki-kishi/go-api-sample/domain/model"
	"github.com/naoki-kishi/go-api-sample/infrastructure/persistance"
	"github.com/naoki-kishi/go-api-sample/usecase/repository"
	"net/http"
)

type entryHander struct {
	entryRepository repository.EntryRepository
}

type EntryHander interface {
	CreateEntry(c *gin.Context)
	GetEntries(c *gin.Context)
}

func NewEntryHandler(eR repository.EntryRepository) EntryHander {
	return &entryHander{entryRepository: eR}
}

func (eH *entryHander) GetEntries(c *gin.Context) {
	conn := persistance.NewSqlDB()
	entryRepo := persistance.NewEntryRepository(conn)

	u := []*model.Entry{}

	us, err := entryRepo.FindAll(u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, us)

}

func (eH *entryHander) CreateEntry(c *gin.Context) {
	return
}
