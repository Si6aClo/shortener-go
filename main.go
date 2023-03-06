package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "shortener/docs"
	r "shortener/router"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
	}
}

// @title           Shortener service
// @version         1.0
// @description     Сервис, позволяющий укорачивать ссылки.
func main() {
	router := gin.Default()
	r.SetupRouter(router)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = router.Run(":8080")
}
