package file

type UploadDTO struct {
	Name string `json:"name" validate:"required"`
}

type ResponseDTO struct {
	Success bool `json:"success"`
}
