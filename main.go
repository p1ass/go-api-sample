package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki-kishi/go-api-sample/domain/model"
	"github.com/naoki-kishi/go-api-sample/infrastructure/datastore"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		conn := datastore.NewSqlDB()
		entryRepo := datastore.NewEntryRepository(conn)

		u := []*model.Entry{}

		us, err := entryRepo.FindAll(u)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, us)
	})
	r.Run(":8080")
}
