package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/george4joseph/go-blog-backend/config"
	"github.com/george4joseph/go-blog-backend/ent/blog"
	"github.com/george4joseph/go-blog-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func BlogGetAll(ctx *gin.Context) {
	readBlog, err := config.ClientConfig.Blog.Query().All(context.Background())

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
			"data":    readBlog,
		})
	}

}

func CreateBlog(ctx *gin.Context) {
	var blog_item models.CreateBlogRequest
	if err := ctx.BindJSON(&blog_item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	idParam := ctx.Params.ByName("id")
	id_user, _ := uuid.Parse(idParam)
	writeblog, err := config.ClientConfig.Blog.Create().
		SetTitle(blog_item.Title).
		SetContent(blog_item.Content).
		SetUserID(id_user).
		Save(context.Background())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "success",
			"data":    writeblog,
		})
	}

}

func UpdateBlog(ctx *gin.Context) {
	var blog_update models.UpdateBlogRequest

	if err := ctx.BindJSON(&blog_update); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error})
		fmt.Println(err)
		return
	}

	idParam := ctx.Params.ByName("id_user")
	id_user, _ := uuid.Parse(idParam)
	idb := ctx.Params.ByName("id")
	id_blog, _ := uuid.Parse(idb)
	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
	}
	if idb == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
	}
	editblog, err := config.ClientConfig.Blog.Update().Where(blog.IDEQ(id_blog), blog.UserID(id_user)).
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

func DeleteBlog(ctx *gin.Context) {

	var blog_del models.DeleteBlogRequest

	if err := ctx.BindJSON(&blog_del); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error})
		fmt.Println(err)
		return
	}
	id_blog, _ := uuid.Parse(blog_del.Blog_id)
	idParam := ctx.Params.ByName("id")
	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
	}
	id_user, _ := uuid.Parse(idParam)

	res, err := config.ClientConfig.Blog.
		Delete().Where(blog.IDEQ(id_blog), blog.UserIDEQ(id_user)).
		Exec(context.Background())

	if err != nil || res == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error:": "Error Found"})
		fmt.Println(err)
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Deletion-success",
			"data":    res,
		})
	}

}

func UsersBlog(ctx *gin.Context) {
	idParam := ctx.Params.ByName("id")
	id_uuid, _ := uuid.Parse(idParam)
	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
	}

	readBlog, err := config.ClientConfig.Blog.Query().Where(blog.UserID(id_uuid)).All(ctx)
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
			"data":    readBlog,
		})
	}

}
