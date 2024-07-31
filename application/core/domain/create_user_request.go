package domain

type CreateUserRequest struct {
	Name      string `json:"name" valid:"required~Name is required,stringlength(1|255)~Name must be between 1 and 255 characters"`
	Email     string `json:"email" valid:"required~Email is required,email~Email must be valid,stringlength(1|255)~Email must be between 1 and 255 characters"`
	ImagePath string `json:"image_path" valid:"url~ImagePath must be a valid URL,stringlength(0|255)~ImagePath must be up to 255 characters"`
}
