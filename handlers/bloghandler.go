package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/george4joseph/go-blog-backend/config"
	"github.com/george4joseph/go-blog-backend/models"
	"github.com/gin-gonic/gin"
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

	writeblog, err := config.ClientConfig.Blog.Create().
		SetTitle(blog_item.Title).
		SetContent(blog_item.Content).
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
