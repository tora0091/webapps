package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tora0091/webapps/environments"
)

type Routers struct {
	Env *environments.Env
}

func NewRouters(env *environments.Env) *Routers {
	return &Routers{
		Env: env,
	}
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
	router.Run(r.Env.HostName() + ":" + r.Env.Port())
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, "hello, this is sample apps.")
}
