package postgres

import (
	"github.com/guiceolin/minigo/models"
	"github.com/jinzhu/gorm"
)

type PostgresUrlRepository struct {
	DB *gorm.DB
}

func (p PostgresUrlRepository) GetById(id uint64) (*models.Url, error) {
	url := &models.Url{}
	p.DB.Find(&url, id)
	return url, nil
}

func (p PostgresUrlRepository) Store(url models.Url) uint {
	p.DB.Create(&url)
	return url.ID
}
