package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/george4joseph/go-blog-backend/config"
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


