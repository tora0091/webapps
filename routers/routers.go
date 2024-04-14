package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/tora0091/webapps/environments"
	"github.com/tora0091/webapps/handlers"
)

type Routers struct {
	Env *environments.Env
	Db  *gorm.DB
}

func NewRouters(env *environments.Env, db *gorm.DB) *Routers {
	return &Routers{
		Env: env,
		Db:  db,
	}
}

func (r *Routers) RouterUp() {
	router := gin.Default()

	// movie system
	m := router.Group("/movie")
	{
		movieHandler := handlers.NewMovieHandler(r.Env, r.Db)
		m.GET("/list", movieHandler.GetList)
		m.GET("/movie/list/:target")
		m.GET("/movie/detail/:movie_id", movieHandler.GetDetail)
	}

	// face system
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
