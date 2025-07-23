package main

import (
	"github.com/gin-gonic/gin"
)

func PostSchema(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var cuelangSchema CueSchema
	c.Bind(&cuelangSchema)

	if cuelangSchema.Cuelang != "" {
		db.Create(&cuelangSchema)

		c.JSON(201, gin.H{"success": cuelangSchema})
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

}

func GetSchemas(c *gin.Context) {

	db := InitDb()

	defer db.Close()

	var schema []CueSchema
	db.Find(&schema)

	c.JSON(200, schema)
}

func GetSchema(c *gin.Context) {

	db := InitDb()

	defer db.Close()

	id := c.Params.ByName("id")
	var schema CueSchema

	db.First(&schema, id)

	if schema.ID != 0 {
		c.JSON(200, schema)
	} else {
		c.JSON(404, gin.H{"error": "Schema not found"})
	}
}

func UpdateSchema(c *gin.Context) {

	db := InitDb()

	defer db.Close()

	id := c.Params.ByName("id")
	var schema CueSchema
	db.First(&schema, id)

	if schema.Cuelang != "" {

		if schema.ID != 0 {
			var newSchema CueSchema
			c.Bind(&newSchema)

			result := CueSchema{
				ID:      schema.ID,
				Cuelang: newSchema.Cuelang,
			}

			db.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": "Schema not found"})
		}

	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func DeleteSchema(c *gin.Context) {

	db := InitDb()

	defer db.Close()

	id := c.Params.ByName("id")
	var schema CueSchema

	db.First(&schema, id)

	if schema.ID != 0 {

		db.Delete(&schema)

		c.JSON(200, gin.H{"success": "Schema #" + id + " deleted"})
	} else {

		c.JSON(404, gin.H{"error": "Schema not found"})
	}

}

func OptionsSchema(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
