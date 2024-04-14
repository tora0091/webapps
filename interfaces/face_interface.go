package interfaces

type FaceResponse struct {
	ID          int     `json:"id"`
	FileId      string  `json:"file_id"`
	FaceId      string  `json:"face_id"`
	AlikeFaceId string  `json:"alike_face_id"`
	Score       float32 `json:"score"`
}
