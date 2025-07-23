package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Cors is the additional header on every request to server
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.POST("/schemas", PostSchema)
		v1.GET("/schemas", GetSchemas)
		v1.GET("/schemas/:id", GetSchema)
		v1.PUT("/schemas/:id", UpdateSchema)
		v1.DELETE("/schemas/:id", DeleteSchema)
	}

	r.Run(":8080")
}
