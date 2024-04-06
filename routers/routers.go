package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routers struct{}

func NewRouters() *Routers {
	return &Routers{}
}

func (r *Routers) RouterUp() {
	router := gin.Default()

	m := router.Group("/movie")
	{
		m.GET("/list/:target")
		m.GET("/detail/:movie_id")
	}

	f := router.Group("/face")
	{
		f.GET("/list/:face_id")
	}

	router.GET("/health", health)
	router.Run("localhost:8000")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, "hello, this is sample apps.")
}
