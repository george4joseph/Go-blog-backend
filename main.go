package main

import (
	"github.com/george4joseph/go-blog-backend/config"
	"github.com/george4joseph/go-blog-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	config.NewEntClient()
	// defer config.ClientConfig.Close()

	router := gin.Default()

	routes.BlogRoutes(router)
	router.Run(":8000")

}
