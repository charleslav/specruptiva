package main

import (
	"disruptiva.org/specruptiva/pkg/core/service"
	"disruptiva.org/specruptiva/adapters/sqlite"
	"disruptiva.org/specruptiva/adapters/http"
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
		DbFile: "./data.db",   // todo: retrieve from env var
		LogMode: true,
	 }
  var port = "9000"        // todo: retrieve from env var

	var schemaStore = sqlite.NewSchemaStore(config)
	var schemaService = service.NewSchemaService(schemaStore)
	var schemaHandler = http.NewSchemaHandler(*schemaService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.POST("/schemas", schemaHandler.Create)
		v1.GET("/schemas", schemaHandler.List)
		v1.GET("/schemas/:id", schemaHandler.Read)
		v1.PUT("/schemas/:id", schemaHandler.Update)
		v1.DELETE("/schemas/:id", schemaHandler.Delete)
	}

	r.Run(":" + port)
}
