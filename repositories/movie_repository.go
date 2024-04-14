package repositories

import (
	"github.com/tora0091/webapps/models"
	"gorm.io/gorm"
)

type MovieRepository struct {
	Dd     *gorm.DB
	Models []models.MovieModel
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{
		Dd:     db,
		Models: []models.MovieModel{},
	}
}

func (r *MovieRepository) GetTarget() error {
	result := r.Dd.Select("directory").Group("directory").Find(&r.Models)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *MovieRepository) GetMovieList(target string) error {
	result := r.Dd.Where("directory = ?", target).Find(&r.Models)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *MovieRepository) GetMovieDetail(movieId string) error {
	result := r.Dd.Where("file_id = ?", movieId).Find(&r.Models)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
