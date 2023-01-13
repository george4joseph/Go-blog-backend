package routes

import (
	"github.com/george4joseph/go-blog-backend/handlers"
	"github.com/gin-gonic/gin"
)



func Routes(router *gin.Engine) {
	router.GET("/blogs", handlers.BlogGetAll)
	router.GET("/users", handlers.GetAllUsers)
	router.POST("/users", handlers.CreateUser)
	router.DELETE("/user/:id", handlers.DeleteUser)
	router.GET("/user/:id/blogs", handlers.UsersBlog)        // Getting all blogs of a user with id
	router.POST("/user/:id/create", handlers.CreateBlog)     // User creating a blog
	router.PATCH("/user/:id/edit_blog", handlers.UpdateBlog) // User Updating Blog
	router.POST("/user/:id/del_blog", handlers.DeleteBlog)
}
