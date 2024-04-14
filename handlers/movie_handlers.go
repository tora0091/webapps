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
		Rep: repositories.NewMovieRepository(),
	}
}

func (h *MovieHandler) GetList(c *gin.Context) {
	err := h.Rep.GetList(h.Db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfaces.ErrorResponse{
			ErrorCode:    http.StatusInternalServerError,
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

func (h *MovieHandler) GetDetail(c *gin.Context) {

}
