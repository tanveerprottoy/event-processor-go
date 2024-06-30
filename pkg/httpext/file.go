package httpext

import (
	"errors"
	"mime/multipart"
	"net/http"
)

func ParseMultiPartForm(r *http.Request, maxMemory int64) error {
	if maxMemory == 0 {
		// left shift 32 << 20 which results in 32*2^20 = 33554432
		// x << y, results in x*2^y
		// max allowed value is 32MB
		maxMemory = 32 << 20
	}
	return r.ParseMultipartForm(maxMemory)
}

func GetFile(r *http.Request, maxFileSize int64) (multipart.File, *multipart.FileHeader, error) {
	if err := ParseMultiPartForm(r, maxFileSize); err != nil {
		return nil, nil, errors.New("the file is too large. the file must be less than 10MB in size")
	}
	// get file
	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, nil, err
	}
	return file, header, nil
}
