package main

import (
	"github.com/gin-gonic/gin"
	"github.com/george4joseph/go-blog-backend/routes"
)

func main() {

	router := gin.Default()

	routes.BlogRoutes(router)


}