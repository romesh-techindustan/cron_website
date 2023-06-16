package mappings

import (
	"goapi/controller"
	"github.com/gin-gonic/gin"
)
var Router *gin.Engine
func CreateUrlMappings() {
	Router = gin.Default()
	Router.Use(controllers.Cors())
	// v1 of the API
	v1 := Router.Group("/v1")
	{
		v1.GET("/websites/:id", controllers.GetAllWebsites)
		v1.GET("/SOS/name/:id", controllers.GetSOSUsername)
		v1.GET("/status/", controllers.StatusCode)
		v1.POST("/login", controllers.Login)
		v1.POST("/signup", controllers.SignUp)
	}
}