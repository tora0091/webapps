package models

func (FaceModel) TableName() string {
	return "face_recognize"
}

type FaceModel struct {
	ID          int
	FileId      string
	FaceId      string
	AlikeFaceId string
	Score       float32
}
