package endpoints

import (
	"github.com/Venukishore-R/CODECRAFT_BW_02/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func Endpoints(r *gin.Engine, handlers *handlers.Handler) {
	r.POST("/create", handlers.Create)
	r.GET("/users", handlers.FindAll)
	r.GET("/users/:email", handlers.FindByEmail)
	r.PUT("/users/:email", handlers.Update)
	r.DELETE("/users/:email", handlers.Delete)
}
