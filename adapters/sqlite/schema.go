package sqlite

import (
	"disruptiva.org/specruptiva/pkg/core/domain"
	"disruptiva.org/specruptiva/pkg/core/port"
	"errors"
	"github.com/jinzhu/gorm"
	"strconv"
)

type GormSchema struct {
	Id     int    `gorm:"primary_key;AUTO_INCREMENT" form:"id" json:"id"`
	Schema string `gorm:"not null" form:"schema" json:"schema"`
}

type SchemaStore struct {
	db     *gorm.DB
	config SqliteConfig
}

func NewSchemaStore(config SqliteConfig) (ports.SchemaStore, error) {
	db := InitDb(config)
	db.AutoMigrate(&GormSchema{})
	return &SchemaStore{db: db, config: config}, nil
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
