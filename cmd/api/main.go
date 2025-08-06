package main

import (
	"disruptiva.org/specruptiva/adapters/http"
	"disruptiva.org/specruptiva/adapters/sqlite"
	"disruptiva.org/specruptiva/pkg/core/service"
	"github.com/gin-gonic/gin"
)

// Cors is the additional header on every request to server
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {

	var config = sqlite.SqliteConfig{
		DbFile:  "./data.db", // todo: retrieve from env var
		LogMode: true,
	}
	var port = "9000" // todo: retrieve from env var

	var schemaStore, _ = sqlite.NewSchemaStore(config)
	var schemaService = service.NewSchemaService(schemaStore)
	var schemaHandler = http.NewSchemaHandler(*schemaService)

	var dataStore, _ = sqlite.NewDataStore(config)
	var dataService = service.NewDataService(dataStore)
	var dataHandler = http.NewDataHandler(*dataService)

	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.POST("/schemas", schemaHandler.Create)
		v1.GET("/schemas", schemaHandler.List)
		v1.GET("/schemas/:id", schemaHandler.Read)
		v1.PUT("/schemas/:id", schemaHandler.Update)
		v1.DELETE("/schemas/:id", schemaHandler.Delete)
		v1.POST("/data", dataHandler.Create)
		v1.GET("/data", dataHandler.List)
		v1.GET("/data/:id", dataHandler.Read)
		v1.PUT("/data/:id", dataHandler.Update)
		v1.DELETE("/data/:id", dataHandler.Delete)
	}

	r.Run(":" + port)
}
