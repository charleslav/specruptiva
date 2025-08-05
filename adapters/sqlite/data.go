package sqlite

import (
	"disruptiva.org/specruptiva/pkg/core/domain"
	"disruptiva.org/specruptiva/pkg/core/port"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strconv"
)

type GormData struct {
	Id   int    `gorm:"primary_key;AUTO_INCREMENT"`
	Data string `gorm:"not null"`
}

type DataStore struct {
	db     *gorm.DB
	config SqliteConfig
}

func NewDataStore(config SqliteConfig) (ports.DataStore, error) {
	db := InitDb(config)
	db.AutoMigrate(&GormData{})
	return &DataStore{db: db, config: config}, nil
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
