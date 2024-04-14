package repositories

import (
	"github.com/tora0091/webapps/models"
	"gorm.io/gorm"
)

type MovieRepository struct {
	Models []models.MovieModel
}

func NewMovieRepository() *MovieRepository {
	return &MovieRepository{
		Models: []models.MovieModel{},
	}
}

func (r *MovieRepository) GetList(db *gorm.DB) error {
	result := db.Select("directory").Group("directory").Find(&r.Models)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
