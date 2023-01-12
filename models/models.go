package models

type CreateBlogRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	User_id string `json:"user_id" binding:"required"`
}

type UpdateBlogRequest struct {
	Content string `json:"content" binding:"required"`
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
