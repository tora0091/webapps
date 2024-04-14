package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/tora0091/webapps/environments"
	"github.com/tora0091/webapps/interfaces"
	"github.com/tora0091/webapps/repositories"
)

type MovieHandler struct {
	Env *environments.Env
	Db  *gorm.DB
	Rep *repositories.MovieRepository
}

func NewMovieHandler(env *environments.Env, db *gorm.DB) *MovieHandler {
	return &MovieHandler{
		Env: env,
		Db:  db,
		Rep: repositories.NewMovieRepository(db),
	}
}

func (h *MovieHandler) GetTarget(c *gin.Context) {
	err := h.Rep.GetTarget()
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfaces.ErrorResponse{
			ErrorMessage: err.Error(),
		})
	}

	response := []interfaces.MovieResponse{}
	for _, movie := range h.Rep.Models {
		response = append(response, interfaces.MovieResponse{
			Directory: movie.Directory,
		})
	}
	c.JSON(http.StatusOK, response)
}

func (h *MovieHandler) GetMovieList(c *gin.Context) {
	target := c.Param("target")
	if target == "" {
		c.JSON(http.StatusBadRequest, interfaces.ErrorResponse{
			ErrorMessage: "empty parameter.",
		})
	}

	err := h.Rep.GetMovieList(target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfaces.ErrorResponse{
			ErrorMessage: err.Error(),
		})
	}

	response := []interfaces.MovieResponse{}
	for _, movie := range h.Rep.Models {
		response = append(response, interfaces.MovieResponse{
			ID:        movie.ID,
			FileId:    movie.FileId,
			FileName:  movie.FileName,
			Directory: movie.Directory,
			Type:      movie.Type,
		})
	}
	c.JSON(http.StatusOK, response)
}

func (h *MovieHandler) GetMovieDetail(c *gin.Context) {
	movieId := c.Param("movie_id")
	if movieId == "" {
		c.JSON(http.StatusBadRequest, interfaces.ErrorResponse{
			ErrorMessage: "empty parameter.",
		})
	}

	err := h.Rep.GetMovieDetail(movieId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfaces.ErrorResponse{
			ErrorMessage: err.Error(),
		})
	}

	response := []interfaces.MovieResponse{}
	for _, movie := range h.Rep.Models {
		response = append(response, interfaces.MovieResponse{
			ID:        movie.ID,
			FileId:    movie.FileId,
			FileName:  movie.FileName,
			Directory: movie.Directory,
			Type:      movie.Type,
		})
	}
	c.JSON(http.StatusOK, response)
}
