package repository

import "github.com/guiceolin/minigo/models"

type UrlRepository interface {
	GetById(id uint64) (*models.Url, error)
	Store(url models.Url) uint
}
