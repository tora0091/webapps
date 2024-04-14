package interfaces

type MovieResponse struct {
	ID        int    `json:"id"`
	FileId    string `json:"file_id"`
	FileName  string `json:"file_name"`
	Directory string `json:"directory"`
	Type      string `json:"type"`
}
