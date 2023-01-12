package routes

import (
	"github.com/george4joseph/go-blog-backend/handlers"
	"github.com/gin-gonic/gin"
)

func BlogRoutes(router *gin.Engine) {

	router.GET("/blogs", handlers.BlogGetAll)
}
