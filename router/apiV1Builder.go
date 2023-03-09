package router

import (
	"github.com/gin-gonic/gin"
	"shortener/api/v1"
)

func buildV1Api(router *gin.Engine) {
	server := v1.NewTaskServerV1()
	// add the v1 api group
	v1Group := router.Group("/api/v1")
	{
		// add the ping route
		v1Group.POST("/make_shorter", server.MakeShorterHandler)
		v1Group.DELETE("/admin/:secretKey", server.DeleteLinkHandler)
		v1Group.GET("/admin/:secretKey", server.GetLinkInfoHandler)
		v1Group.GET("/get_urls_list/:token", server.GetUrlsListHandler)
		v1Group.POST("/registration", server.RegistrationHandler)
		v1Group.POST("/auth", server.AuthHandler)
	}

	router.GET("/:shortUrl", server.RedirectToLongHandler)
}
