package sqlite

import (
	"errors"
	"disruptiva.org/specruptiva/pkg/core/port"
	"disruptiva.org/specruptiva/pkg/core/domain"
  "github.com/jinzhu/gorm"
)

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

func (s *SqliteStore) Init() error {
	db, err:= gorm.Open("sqlite3", s.dbfile)
	s.db=db
	s.db.LogMode(true) // todo: pass sqlite/gorm options
	if err != nil {
		return err
	}

	if !db.HasTable(&GormSchema{}){
		s.db.CreateTable(&GormSchema{})
		s.db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&GormSchema{})
	}
  return nil
}

func (s* SqliteStore)List()(domain.Schemas, error){

	return nil, nil
}

func (s* SqliteStore)Create(schema string)(domain.Success, error){
   s.Init()

	 var gormSchema= GormSchema{Schema: schema}

	 if gormSchema.Schema != "" {
		 result:= s.db.Create(&gormSchema)
     if result.Error != nil {
			 return domain.Success{}, result.Error
		 }

		 return domain.Success{
			 Id: string(gormSchema.Id),
			 Message: "created with success",
		 }, nil

	 }else{ 
		 return domain.Success{}, errors.New("(create) Fields are empty")
	 }
}

func (s* SqliteStore)Read(id string)(domain.Schema, error){
	return domain.Schema{}, nil
}
func (s* SqliteStore)Update(id string, schema string)(domain.Success, error){
	return domain.Success{}, nil
}
func (s* SqliteStore)Delete(id string)(domain.Success, error){
	return domain.Success{}, nil
}
