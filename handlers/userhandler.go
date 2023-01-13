package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/george4joseph/go-blog-backend/config"
	"github.com/george4joseph/go-blog-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser(ctx *gin.Context) {
	var user_item models.CreateUserRequest
	if err := ctx.BindJSON(&user_item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	writeuser, err := config.ClientConfig.User.Create().
		SetName(user_item.Name).
		SetEmail(user_item.Email).
		SetUserType("USER").
		Save(context.Background())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "success",
		"data":    writeuser,
	})

}

func DeleteUser(ctx *gin.Context) {
	idParam := ctx.Params.ByName("id")
	id_uuid, _ := uuid.Parse(idParam)
	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
	}

	err := config.ClientConfig.User.
		DeleteOneID(id_uuid).
		Exec(context.Background())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error})
		fmt.Println(err)
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "User-Deletion-success",
		})
	}

}
