package domain

type CreateVideoRequest struct {
	Title       string `json:"title" valid:"required~Title is required,stringlength(1|255)~Title must be between 1 and 255 characters"`
	Description string `json:"description" valid:"required~Description is required"`
	URL         string `json:"url" valid:"required~URL is required,url~URL must be valid"`
	UserID      int    `json:"userId" valid:"required~User ID is required"`
}
