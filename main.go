package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki-kishi/go-api-sample/infrastructure/api/handler"
	"github.com/naoki-kishi/go-api-sample/infrastructure/datastore"
)

func main() {
	r := gin.Default()

	conn := datastore.NewSqlDB()
	entryRepo := datastore.NewEntryRepository(conn)
	tagRepo := datastore.NewTagRepository(conn)

	eH := handler.NewEntryHandler(entryRepo)
	tH := handler.NewTagHandler(tagRepo)

	r.GET("/entries", eH.GetEntries)

	r.GET("/tags", tH.GetTags)

	r.Run(":8080")
}
