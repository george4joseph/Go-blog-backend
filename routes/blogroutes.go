package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/george4joseph/go-blog-backend/handlers"
)


func BlogRoutes(router *gin.Engine) {

	router.GET("/blogs", handlers.BlogGetAll)
}