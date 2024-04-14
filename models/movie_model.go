package models

func (MovieModel) TableName() string {
	return "movie_list"
}

type MovieModel struct {
	ID        int
	FileId    string
	FileName  string
	Directory string
	Type      string
}
