package service

import (
	"context"
	"log"

	fileapi "github.com/tanveerprottoy/event-processor-go/internal/api/file"
	"github.com/tanveerprottoy/event-processor-go/pkg/constant"
	"github.com/tanveerprottoy/event-processor-go/pkg/errorext"
	"github.com/tanveerprottoy/event-processor-go/pkg/file"
	"github.com/tanveerprottoy/event-processor-go/pkg/response"
)

// Service contains the business logic as well as calls to the
// repository to perform db operations
type Service struct {
}

// NewService initializes a new Service
func NewService() *Service {
	return &Service{}
}

// readOneInternal fetches one entity from db
func (s *Service) Upload(ctx context.Context, d fileapi.UploadDTO, args ...any) (response.Response[fileapi.ResponseDTO], error) {
	// get content/MIME type
	contentType, err := file.GetMultipartFileContentType(d.File, true)
	if err != nil {
		return response.Response[fileapi.ResponseDTO]{}, errorext.BuildCustomError(err)
	}
	log.Println(contentType)
	// check mime type
	validMIME := file.IsAllowedMIMEType(contentType, constant.AllowedMimeTypes[:])
	if !validMIME {
		return response.Response[fileapi.ResponseDTO]{}, errorext.BuildCustomError(err)
	}
	// proceed to save the file
	p, err := file.SaveFile(ctx, d.File, "uploads", d.Header.Filename)
	if err != nil {
		return response.Response[fileapi.ResponseDTO]{}, errorext.BuildCustomError(err)
	}
	return response.Response[fileapi.ResponseDTO]{Data: fileapi.ResponseDTO{FilePath: p}}, nil
}
