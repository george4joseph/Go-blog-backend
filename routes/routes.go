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
	router.POST("/user/:id/del_blog", handlers.DeleteBlog)            // User Deletes Blog

	// Admin roles
	router.GET("/admin/:id_admin/users", handlers.GetAllUsers)
	router.POST("/create_admin", handlers.CreateAdmin)
	// router.POST("/admin/:id/assign_admin", handlers.AssignAdmin)
	router.POST("/admin/:id_admin/create_user", handlers.AdminCreateUser)
	router.DELETE("/admin/:id_admin/delete_user/:id", handlers.AdminDeleteUser)
	router.DELETE("/admin/:id_admin/delete_blog/:id", handlers.AdminDeleteBlog)
	router.PATCH("/admin/:id_admin/update_blog/:id", handlers.AdminUpdateBlog)

}
