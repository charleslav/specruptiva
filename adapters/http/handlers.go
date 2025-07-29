package http

import (
	"log"
	"disruptiva.org/specruptiva/pkg/core/service"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&CueSchema{}) {
		db.CreateTable(&CueSchema{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&CueSchema{})
	}

	return db
}

// todo: supprimer CueSchema
type CueSchema struct {
	ID      int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Cuelang string `gorm:"not null" form:"Cuelang" json:"Cuelang"`
}

type SchemaHandler struct {
	service service.SchemaService
}

func NewSchemaHandler(service service.SchemaService) (*SchemaHandler) {
	return &SchemaHandler{ service: service }
}

func (h *SchemaHandler) Create (c *gin.Context) {
	log.Println("Create handler")
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
func (h *SchemaHandler) List (c *gin.Context) {
	db := InitDb()

	defer db.Close()

	var schema []CueSchema
	db.Find(&schema)

	c.JSON(200, schema)
}
func (h *SchemaHandler) Read (c *gin.Context){


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
func (h *SchemaHandler) Update (c *gin.Context) {


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
func (h *SchemaHandler) Delete (c *gin.Context) {


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

