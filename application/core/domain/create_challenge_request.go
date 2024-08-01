package domain

type CreateChallengeRequest struct {
	Title       string `json:"title" valid:"required~Title is required,stringlength(1|255)~Title must be between 1 and 255 characters"`
	Description string `json:"description" valid:"required~Description is required"`
	Difficulty  int    `json:"difficulty" valid:"range(1|5)~Difficulty must be between 1 and 5"`
	UserID      int    `json:"userId" valid:"required~User ID is required"`
}
