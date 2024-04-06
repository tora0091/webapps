package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/tora0091/webapps/environments"
)

type MovieHandler struct {
	Env *environments.Env
	Db  *gorm.DB
}

func NewMovieHandler(env *environments.Env, db *gorm.DB) *MovieHandler {
	return &MovieHandler{
		Env: env,
		Db:  db,
	}
}

func (h *MovieHandler) GetList(c *gin.Context) {

}

func (h *MovieHandler) GetDetail(c *gin.Context) {

}
