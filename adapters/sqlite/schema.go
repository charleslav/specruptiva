package sqlite

import ( 
	"strconv"
	"log"
	"errors"
	"disruptiva.org/specruptiva/pkg/core/port"
	"disruptiva.org/specruptiva/pkg/core/domain"
  "github.com/jinzhu/gorm"
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
	dbfile string
	db *gorm.DB
}

func NewSchemaStore (dbfile string ) ports.SchemaStore {
	return &SqliteStore{dbfile: dbfile}
}

//func (s *SqliteStore) Init() error {
//	db, err:= gorm.Open("sqlite3", s.dbfile)
//	s.db=db
//	s.db.LogMode(true) // todo: pass sqlite/gorm options
//	if err != nil {
//		return err
//	}

//	if !db.HasTable(&GormSchema{}){
//		s.db.CreateTable(&GormSchema{})
//		s.db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&GormSchema{})
//	}
//  return nil
//}

func (s* SqliteStore)List()(domain.Schemas, error){

	db := InitDb()

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

	log.Println("create SQLITE Adapter")
	db := InitDb()
	defer db.Close()


	if schema != "" {
		gormSchema:= GormSchema{Schema: schema,}
		result:= db.Create(&gormSchema)
		log.Println(gormSchema)
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
  db := InitDb()
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
	return domain.Success{}, nil
}
func (s* SqliteStore)Delete(id string)(domain.Success, error){
	return domain.Success{}, nil
}
