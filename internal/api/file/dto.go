package file

import (
	"net/http"
)

type UploadDTO struct {
	// pass the request
	Req *http.Request
}

type ResponseDTO struct {
	FilePath string `json:"filePath"`
}

type ResponseMultiDTO struct {
	FilePaths []string `json:"filePaths"`
}
