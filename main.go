package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki-kishi/go-api-sample/infrastructure/api/handler"
	"github.com/naoki-kishi/go-api-sample/infrastructure/persistance"
)

func main() {
	r := gin.Default()

	conn := persistance.NewSqlDB()
	entryRepo := persistance.NewEntryRepository(conn)
	tagRepo := persistance.NewTagRepository(conn)

	eH := handler.NewEntryHandler(entryRepo)
	tH := handler.NewTagHandler(tagRepo)

	r.GET("/entries", eH.GetEntries)

	r.GET("/tags", tH.GetTags)

	r.Run(":8080")
}
