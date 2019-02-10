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
	r.POST("/entries", eH.CreateEntry)
	r.GET("/entries/:id", eH.GetEntry)
	r.PUT("/entries/:id", eH.UpdateEntry)
	r.DELETE("/entries/:id", eH.DeleteEntry)

	r.GET("/tags", tH.GetTags)
	r.POST("/tags", tH.CreateTag)
	r.GET("/tags/:id", tH.GetTag)
	r.PUT("/tags/:id", tH.UpdateTag)
	r.DELETE("/tags/:id", tH.DeleteTag)

	r.Run(":8080")
}
