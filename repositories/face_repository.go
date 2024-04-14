package repositories

import (
	"github.com/tora0091/webapps/models"
	"gorm.io/gorm"
)

type FaceRepository struct {
	Db     *gorm.DB
	Models []models.FaceModel
}

func NewFaceRepository(db *gorm.DB) *FaceRepository {
	return &FaceRepository{
		Db:     db,
		Models: []models.FaceModel{},
	}
}

func (r *FaceRepository) GetFaceList(movieId string) error {
	result := r.Db.Select("file_id", "face_id").
		Where("file_id = ?", movieId).Group("file_id").Group("face_id").Find(&r.Models)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *FaceRepository) GetLookLikeFaceList(faceId string) error {
	result := r.Db.Select("id", "file_id", "face_id", "alike_face_id", "score").
		Where("face_id = ?", faceId).Order("score DESC").Find(&r.Models)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
