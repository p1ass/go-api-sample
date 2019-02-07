package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki-kishi/go-api-sample/domain/model"
	"github.com/naoki-kishi/go-api-sample/infrastructure/datastore"
	"net/http"
)

type entryHander struct {
	entryRepository datastore.EntryRepository
}

type EntryHander interface {
	CreateEntry(c *gin.Context)
	GetEntries(c *gin.Context)
}

func NewEntryHandler(eR datastore.EntryRepository) EntryHander {
	return &entryHander{entryRepository: eR}
}

func (eH *entryHander) GetEntries(c *gin.Context) {
	conn := datastore.NewSqlDB()
	entryRepo := datastore.NewEntryRepository(conn)

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
