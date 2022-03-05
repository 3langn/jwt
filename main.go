package main

import (
	"github.com/3langn/learn-go/controller"
	"github.com/3langn/learn-go/models"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/3langn/learn-go/docs"
)


//https://stackoverflow.com/questions/24790175/when-is-the-init-function-run
func init() {

	models.Setup()
}
func main() {
	r := gin.Default()

	docs.SwaggerInfo_swagger.Title = "Swagger Example API"
	docs.SwaggerInfo_swagger.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo_swagger.Version = "1.0"
	docs.SwaggerInfo_swagger.Schemes = []string{"http", "https"}
	docs.SwaggerInfo_swagger.BasePath = "/api"


	authController := new(controller.AuthController)

	authRoute := r.Group("/api/auth")
	{

		authRoute.POST("/login", authController.Login)
		authRoute.POST("/register", authController.Register)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
