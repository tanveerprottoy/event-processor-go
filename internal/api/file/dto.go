package file

import "mime/multipart"

type UploadDTO struct {
	File   multipart.File
	Header *multipart.FileHeader
}

type ResponseDTO struct {
	FilePath string `json:"filePath"`
}
