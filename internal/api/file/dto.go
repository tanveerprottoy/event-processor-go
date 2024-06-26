package file

type UploadDTO struct {
	Name string
	File 
}

type ResponseDTO struct {
	Success bool `json:"success"`
}
