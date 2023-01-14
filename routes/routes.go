package routes

import (
	"github.com/george4joseph/go-blog-backend/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/blogs", handlers.BlogGetAll)
	router.POST("/users", handlers.CreateUser)
	router.DELETE("/user/:id", handlers.DeleteUser)
	router.GET("/user/:id/blogs", handlers.UsersBlog)                 // Getting all blogs of a user with id
	router.POST("/user/:id/create", handlers.CreateBlog)              // User creating a blog
	router.PATCH("/user/:id_user/edit_blog/:id", handlers.UpdateBlog) // User Updating Blog
	router.POST("/user/:id/delete_blog", handlers.DeleteBlog)         // User Deletes Blog

	// Admin roles
	router.GET("/admin/:id_admin/users", handlers.GetAllUsers) // Admin create users
	router.POST("/create_admin", handlers.CreateAdmin)         // Create Admin
	router.POST("/admin/:id_admin/assign_admin/:id", handlers.AssignAdmin)
	router.POST("/admin/:id_admin/create_user/:id", handlers.AdminCreateUser)   // Admin Create User
	router.DELETE("/admin/:id_admin/delete_user/:id", handlers.AdminDeleteUser) // Admin Delete User
	router.DELETE("/admin/:id_admin/delete_blog/:id", handlers.AdminDeleteBlog) // Admin delete blog
	router.PATCH("/admin/:id_admin/edit_blog/:id", handlers.AdminUpdateBlog)    // Admin update blog

}
