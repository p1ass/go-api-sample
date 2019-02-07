package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki-kishi/go-api-sample/domain/model"
	"github.com/naoki-kishi/go-api-sample/infrastructure/persistance"
	"github.com/naoki-kishi/go-api-sample/usecase/repository"
	"net/http"
)

type tagHander struct {
	tagRepository repository.TagRepository
}

type TagHander interface {
	CreateTag(c *gin.Context)
	GetTags(c *gin.Context)
}

func NewTagHandler(eR repository.TagRepository) TagHander {
	return &tagHander{tagRepository: eR}
}

func (eH *tagHander) GetTags(c *gin.Context) {
	conn := persistance.NewSqlDB()
	tagRepo := persistance.NewTagRepository(conn)

	u := []*model.Tag{}

	us, err := tagRepo.FindAll(u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, us)

}

func (eH *tagHander) CreateTag(c *gin.Context) {
	return
}
