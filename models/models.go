package models



type CreateBlogRequest struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateBlogRequest struct {
	Content string `json:"content" binding:"required"`
}