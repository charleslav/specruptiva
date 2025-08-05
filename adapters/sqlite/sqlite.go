package sqlite

import (
	"disruptiva.org/specruptiva/pkg/core/domain"
	"disruptiva.org/specruptiva/pkg/core/port"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strconv"
)

type SqliteConfig struct {
	LogMode bool
	DbFile  string
}

type GormData struct {
	Id   int    `gorm:"primary_key;AUTO_INCREMENT"`
	Data string `gorm:"not null"`
}

type GormSchema struct {
	Id     int    `gorm:"primary_key;AUTO_INCREMENT" form:"id" json:"id"`
	Schema string `gorm:"not null" form:"schema" json:"schema"`
}

type DataStore struct {
	db     *gorm.DB
	config SqliteConfig
}

type SchemaStore struct {
	db     *gorm.DB
	config SqliteConfig
}

func InitDb(config SqliteConfig) *gorm.DB {
	db, err := gorm.Open("sqlite3", config.DbFile)
	if err != nil {
		panic(err)
	}
	db.LogMode(config.LogMode)
	return db
}

func NewDataStore(config SqliteConfig) (ports.DataStore, error) {
	db := InitDb(config)
	db.AutoMigrate(&GormData{})
	return &DataStore{db: db, config: config}, nil
}

func NewSchemaStore(config SqliteConfig) (ports.SchemaStore, error) {
	db := InitDb(config)
	db.AutoMigrate(&GormSchema{})
	return &SchemaStore{db: db, config: config}, nil
}

func (s *DataStore) List() (domain.Datas, error) {
	var datas []GormData
	if err := s.db.Find(&datas).Error; err != nil {
		return nil, err
	}

	out := make(domain.Datas, 0, len(datas))
	for _, d := range datas {
		out = append(out, domain.Data{
			Id:   strconv.Itoa(d.Id),
			Data: d.Data,
		})
	}
	return out, nil
}

func (s *DataStore) Create(data string) (domain.Success, error) {
	if data == "" {
		return domain.Success{}, errors.New("data field is empty")
	}

	gd := GormData{Data: data}
	if err := s.db.Create(&gd).Error; err != nil {
		return domain.Success{}, err
	}

	return domain.Success{
		Id:      strconv.Itoa(gd.Id),
		Message: "data created",
	}, nil
}

func (s *DataStore) Read(id string) (domain.Data, error) {
	var data GormData
	if err := s.db.First(&data, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return domain.Data{}, errors.New("data not found")
		}
		return domain.Data{}, err
	}

	return domain.Data{
		Id:   strconv.Itoa(data.Id),
		Data: data.Data,
	}, nil
}

func (s *DataStore) Update(id string, data string) (domain.Success, error) {
	if data == "" {
		return domain.Success{}, errors.New("data field is empty")
	}

	var gormData GormData
	if err := s.db.First(&gormData, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return domain.Success{}, errors.New("data not found")
		}
		return domain.Success{}, err
	}

	gormData.Data = data
	if err := s.db.Save(&gormData).Error; err != nil {
		return domain.Success{}, err
	}

	return domain.Success{
		Id:      id,
		Message: "data updated",
	}, nil
}

func (s *DataStore) Delete(id string) (domain.Success, error) {
	var data GormData
	if err := s.db.First(&data, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return domain.Success{}, errors.New("data not found")
		}
		return domain.Success{}, err
	}

	if err := s.db.Delete(&data).Error; err != nil {
		return domain.Success{}, err
	}

	return domain.Success{
		Id:      strconv.Itoa(data.Id),
		Message: "data deleted",
	}, nil
}

func (s *SchemaStore) List() (domain.Schemas, error) {
	var schemas []GormSchema
	if err := s.db.Find(&schemas).Error; err != nil {
		return nil, err
	}

	out := make(domain.Schemas, 0, len(schemas))
	for _, schema := range schemas {
		out = append(out, domain.Schema{
			Id:     strconv.Itoa(schema.Id),
			Schema: schema.Schema,
		})
	}
	return out, nil
}

func (s *SchemaStore) Create(schema string) (domain.Success, error) {
	if schema == "" {
		return domain.Success{}, errors.New("schema field is empty")
	}

	gs := GormSchema{Schema: schema}
	if err := s.db.Create(&gs).Error; err != nil {
		return domain.Success{}, err
	}

	return domain.Success{
		Id:      strconv.Itoa(gs.Id),
		Message: "schema created",
	}, nil
}

func (s *SchemaStore) Read(id string) (domain.Schema, error) {
	var schema GormSchema
	if err := s.db.First(&schema, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return domain.Schema{}, errors.New("schema not found")
		}
		return domain.Schema{}, err
	}

	return domain.Schema{
		Id:     strconv.Itoa(schema.Id),
		Schema: schema.Schema,
	}, nil
}

func (s *SchemaStore) Update(id string, schema string) (domain.Success, error) {
	if schema == "" {
		return domain.Success{}, errors.New("schema field is empty")
	}

	var gormSchema GormSchema
	if err := s.db.First(&gormSchema, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return domain.Success{}, errors.New("schema not found")
		}
		return domain.Success{}, err
	}

	gormSchema.Schema = schema
	if err := s.db.Save(&gormSchema).Error; err != nil {
		return domain.Success{}, err
	}

	return domain.Success{
		Id:      id,
		Message: "schema updated",
	}, nil
}

func (s *SchemaStore) Delete(id string) (domain.Success, error) {
	var schema GormSchema
	if err := s.db.First(&schema, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return domain.Success{}, errors.New("schema not found")
		}
		return domain.Success{}, err
	}

	if err := s.db.Delete(&schema).Error; err != nil {
		return domain.Success{}, err
	}

	return domain.Success{
		Id:      strconv.Itoa(schema.Id),
		Message: "schema deleted",
	}, nil
}
