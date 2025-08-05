package sqlite

import ( 
	"strconv"
	"errors"
	"disruptiva.org/specruptiva/pkg/core/port"
	"disruptiva.org/specruptiva/pkg/core/domain"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SqliteConfig struct {
	LogMode bool
	DbFile string
}


func InitDb(config SqliteConfig) *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", config.DbFile)
	// Display SQL queries
	db.LogMode(config.LogMode)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&GormSchema{}) {
		db.CreateTable(&GormSchema{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&GormSchema{})
	}

	return db
}

type GormSchema struct {
	Id  int  `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Schema string `gorm:"no null" form:"schema" json:"schema"`
}

type SqliteStore struct {
	db *gorm.DB
	config SqliteConfig
}

func NewSchemaStore (config SqliteConfig) ports.SchemaStore {
	return &SqliteStore{config: config}
}

func (s* SqliteStore)List()(domain.Schemas, error){

	db := InitDb(s.config)

	defer db.Close()

	var schemas []GormSchema
	db.Find(&schemas)

	var outschema = domain.Schemas{}

	// casting ...
	for _, value := range schemas {
		outschema = append(outschema, domain.Schema{
			Id: strconv.Itoa(value.Id),
		  Schema: value.Schema,
		})
	}

	return outschema, nil
}

func (s* SqliteStore)Create(schema string)(domain.Success, error){

	db := InitDb(s.config)
	defer db.Close()


	if schema != "" {
		gormSchema:= GormSchema{Schema: schema,}
		result:= db.Create(&gormSchema)
    if result.Error != nil {
			return domain.Success{}, result.Error
		}
		return domain.Success{
			Id:  strconv.Itoa(gormSchema.Id),
			Message: "schema created",
		  }, nil
		}
	return domain.Success{}, errors.New("Fields are empty")
}

func (s* SqliteStore)Read(id string)(domain.Schema, error){
  db := InitDb(s.config)
	defer db.Close()

	var schema GormSchema

	db.First(&schema, id)

	if schema.Id != 0 {
		return domain.Schema{
			Id:  strconv.Itoa(schema.Id),
			Schema: schema.Schema,
		}, nil
	}
	return domain.Schema{}, nil
}
func (s* SqliteStore)Update(id string, schema string)(domain.Success, error){

	db := InitDb(s.config)
	defer db.Close()

	var gormSchema GormSchema
	result:= db.First(&gormSchema, id)
  if result.Error != nil {
		return domain.Success{}, result.Error
	}

	if gormSchema.Id != 0 {

		idi, err:= strconv.Atoi(id)
		if err != nil {
			return domain.Success{}, err
		}
		newGormSchema := GormSchema{
			Id:     idi, 
			Schema: schema,
		}

		result:= db.Save(&newGormSchema)
    if result.Error != nil {
		  return domain.Success{}, result.Error
  	}
		  return domain.Success{
				Id: id,
				Message: "schema updated",
			}, nil
		} else {
			return domain.Success{}, nil // todo: comportement ambigue... revoir ça
		}

}
func (s* SqliteStore)Delete(id string)(domain.Success, error){
	db := InitDb(s.config)
	defer db.Close()

	var schema GormSchema
	result := db.First(&schema, id)
  if result.Error != nil {
		return domain.Success{}, nil // todo: retourner ce message (result.Error.Error())
 	}
	if schema.Id != 0 {  
		result:= db.Delete(&schema) 
    if result.Error != nil {
	    return domain.Success{}, result.Error
 	  }
		return domain.Success{
			Id: strconv.Itoa(schema.Id),
			Message: "schema deleted",
		},nil
	}
	return domain.Success{}, nil
}
