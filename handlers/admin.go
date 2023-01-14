package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/george4joseph/go-blog-backend/config"
	"github.com/george4joseph/go-blog-backend/ent/blog"
	"github.com/george4joseph/go-blog-backend/ent/user"
	"github.com/george4joseph/go-blog-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckAdmin(idParam string, ctx *gin.Context) bool {

	id_admin, _ := uuid.Parse(idParam)
	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
	}
	readtype, err := config.ClientConfig.User.
		Query().Where(user.IDEQ(id_admin)).Select(user.FieldUserType).Strings(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
		})
	}
	if readtype[0] != "ADMIN" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Admin Not Found",
		})
		return false
	}
	return true
}

func GetAllUsers(ctx *gin.Context) {
	idParam := ctx.Params.ByName("id_admin")
	if CheckAdmin(idParam, ctx) {

		readusers, err := config.ClientConfig.User.
			Query().All(context.Background())

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": fmt.Sprint(err),
				"data":    nil,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "success",
				"users":   readusers,
			})
		}
	}

}

func CreateAdmin(ctx *gin.Context) {
	var user_item models.CreateUserRequest
	if err := ctx.BindJSON(&user_item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	writeuser, err := config.ClientConfig.User.Create().
		SetName(user_item.Name).
		SetEmail(user_item.Email).
		SetUserType("ADMIN").
		Save(context.Background())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Created Admin",
		"data":    writeuser,
	})
}

func AdminCreateUser(ctx *gin.Context) {
	idParam := ctx.Params.ByName("id_admin")
	if CheckAdmin(idParam, ctx) {
		CreateUser(ctx)
	}
}

func AdminDeleteUser(ctx *gin.Context) {
	idParam := ctx.Params.ByName("id_admin")
	if CheckAdmin(idParam, ctx) {
		DeleteUser(ctx)
	}
}

func AdminDeleteBlog(ctx *gin.Context) {
	idParam := ctx.Params.ByName("id_admin")
	if CheckAdmin(idParam, ctx) {
		idParam := ctx.Params.ByName("id")
		if idParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		}
		id_blog, _ := uuid.Parse(idParam)
		err := config.ClientConfig.Blog.
			DeleteOneID(id_blog).
			Exec(context.Background())

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error:": "Error Found"})
			fmt.Println(err)
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "Deletion-success",
			})
		}
	}
}

func AdminUpdateBlog(ctx *gin.Context) {
	idParam := ctx.Params.ByName("id_admin")
	if CheckAdmin(idParam, ctx) {
		var blog_update models.UpdateBlogRequest

		if err := ctx.BindJSON(&blog_update); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error})
			fmt.Println(err)
			return
		}

		idParam := ctx.Params.ByName("id")
		if idParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		}
		id_blog, _ := uuid.Parse(idParam)
		editblog, err := config.ClientConfig.Blog.Update().Where(blog.IDEQ(id_blog)).
			SetContent(blog_update.Content).
			Save(context.Background())

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error})
			fmt.Println(err)
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "success",
				"data":    editblog,
			})
		}
	}
}

func AssignAdmin(ctx *gin.Context) {
	idParam := ctx.Params.ByName("id_admin")
	if CheckAdmin(idParam, ctx) {
		idParam := ctx.Params.ByName("id")
		if idParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		}
		idu := ctx.Params.ByName("id")
		if idu == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		}
		id_user, _ := uuid.Parse(idu)

		addadmin, err := config.ClientConfig.User.UpdateOneID(id_user).
			SetUserType("ADMIN").
			Save(context.Background())

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error})
			fmt.Println(err)
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "Admin Added",
				"data":    addadmin,
			})
		}
	}
}
