package http

import (

	"log"
	"disruptiva.org/specruptiva/pkg/core/service"
//	"disruptiva.org/specruptiva/pkg/core/port"
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
	if !db.HasTable(&inputSchema{}) {
		db.CreateTable(&inputSchema{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&inputSchema{})
	}

	return db
}

// todo: supprimer CueSchema
type inputSchema struct {
	Schema string `form:"schema" json:"schema"`
}

type SchemaHandler struct {
	service service.SchemaService
}

func NewSchemaHandler(service service.SchemaService) (*SchemaHandler) {
	return &SchemaHandler{ service: service }
}

func (h *SchemaHandler) Create (c *gin.Context) {

	var inSchema inputSchema
	c.Bind(&inSchema)

  log.Println(inSchema.Schema)

	success, err:=h.service.Create(inSchema.Schema)


	if err == nil {

		c.JSON(201, success)
	} else {
		c.JSON(422, gin.H{"error": err.Error()})
	}
}
func (h *SchemaHandler) List (c *gin.Context) {
   schemas, err:= h.service.List()
	 if err !=nil {
		 c.JSON(500,gin.H{"error": "Internal error while listing schemas"})
	 }else{
		 c.JSON(200,schemas)
	 }
}

func (h *SchemaHandler) Read (c *gin.Context){


//	db := InitDb()

//	defer db.Close()

	id := c.Params.ByName("id")
	
	schema, err:= h.service.Read(id)
  if err !=nil {
    c.JSON(500,gin.H{"error": "Internal error while reading schemas"})
  }else if schema.Schema == "" {
   c.JSON(404, gin.H{"error": "Schema not found"}) 
	}else{
		 c.JSON(200,schema)
	}
}
func (h *SchemaHandler) Update (c *gin.Context) {
  c.JSON(200, gin.H{"state": "Not Implemented"})


	//db := InitDb()

//	defer db.Close()

//	id := c.Params.ByName("id")
//	var schema inputSchema
//	db.First(&schema, id)

//	if schema.Cuelang != "" {

//		if schema.ID != 0 {
//			var newSchema inputSchema
//			c.Bind(&newSchema)
//
//			result := inputSchema{
//				ID:      schema.ID,
//				Cuelang: newSchema.Cuelang,
//			}

//			db.Save(&result)
//			c.JSON(200, gin.H{"success": result})
//		} else {
//			c.JSON(404, gin.H{"error": "Schema not found"})
//		}

//	} else {
//		c.JSON(422, gin.H{"error": "Fields are empty"})
//	}
}
func (h *SchemaHandler) Delete (c *gin.Context) {

  c.JSON(200, gin.H{"state": "Not Implemented"})

//	db := InitDb()

//	defer db.Close()

//	id := c.Params.ByName("id")
//	var schema CueSchema

//	db.First(&schema, id)

//	if schema.ID != 0 {

//		db.Delete(&schema)

//		c.JSON(200, gin.H{"success": "Schema #" + id + " deleted"})
//	} else {

	//	c.JSON(404, gin.H{"error": "Schema not found"})
//	}

}

