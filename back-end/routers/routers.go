package routers

import (
	v1 "gin-vue-element-ui/web/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/static", http.Dir("templates/static"))

	r.LoadHTMLFiles("templates/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/users/list", v1.List)
		//apiv1.POST("/users", nil)
		//apiv1.PUT("/users/:id", nil)
		//apiv1.DELETE("/users/:id", nil)
	}

	return r
}
