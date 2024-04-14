package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/tora0091/webapps/environments"
	"github.com/tora0091/webapps/interfaces"
	"github.com/tora0091/webapps/repositories"
)

type FaceHandler struct {
	Env *environments.Env
	Db  *gorm.DB
	Rep *repositories.FaceRepository
}

func NewFaceHandler(env *environments.Env, db *gorm.DB) *FaceHandler {
	return &FaceHandler{
		Env: env,
		Db:  db,
		Rep: repositories.NewFaceRepository(db),
	}
}

func (h *FaceHandler) GetFaceList(c *gin.Context) {
	movieId := c.Param("movie_id")
	if movieId == "" {
		c.JSON(http.StatusBadRequest, interfaces.ErrorResponse{
			ErrorMessage: "empty parameter.",
		})
	}

	err := h.Rep.GetFaceList(movieId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfaces.ErrorResponse{
			ErrorMessage: err.Error(),
		})
	}

	response := []interfaces.FaceResponse{}
	for _, face := range h.Rep.Models {
		response = append(response, interfaces.FaceResponse{
			FileId: face.FileId,
			FaceId: face.FaceId,
		})
	}
	c.JSON(http.StatusOK, response)
}

func (h *FaceHandler) GetLookLikeFaceList(c *gin.Context) {
	faceId := c.Param("face_id")
	if faceId == "" {
		c.JSON(http.StatusBadRequest, interfaces.ErrorResponse{
			ErrorMessage: "empty parameter.",
		})
	}

	err := h.Rep.GetLookLikeFaceList(faceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfaces.ErrorResponse{
			ErrorMessage: err.Error(),
		})
	}

	response := []interfaces.FaceResponse{}
	for _, face := range h.Rep.Models {
		response = append(response, interfaces.FaceResponse{
			ID:          face.ID,
			FileId:      face.FileId,
			FaceId:      face.FaceId,
			AlikeFaceId: face.AlikeFaceId,
			Score:       face.Score,
		})
	}
	c.JSON(http.StatusOK, response)
}
