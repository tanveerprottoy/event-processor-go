package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	fileapi "github.com/tanveerprottoy/event-processor-go/internal/api/file"
	"github.com/tanveerprottoy/event-processor-go/pkg/constant"
	"github.com/tanveerprottoy/event-processor-go/pkg/errorext"
	"github.com/tanveerprottoy/event-processor-go/pkg/file"
	"github.com/tanveerprottoy/event-processor-go/pkg/ioext"
	"github.com/tanveerprottoy/event-processor-go/pkg/response"
	"github.com/tanveerprottoy/event-processor-go/pkg/timeext"
)

// Service contains the business logic as well as calls to the
// repository to perform db operations
type Service struct {
}

// NewService initializes a new Service
func NewService() *Service {
	return &Service{}
}

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
	p, err := file.SaveFile(ctx, d.File, "uploads", fmt.Sprintf("%d", timeext.NowUnixMilli())+filepath.Ext(d.Header.Filename), nil)
	if err != nil {
		return response.Response[fileapi.ResponseDTO]{}, errorext.BuildCustomError(err)
	}
	return response.Response[fileapi.ResponseDTO]{Data: fileapi.ResponseDTO{FilePath: p}}, nil
}

func (s *Service) UploadMultiple(ctx context.Context, d fileapi.UploadMultipleDTO, args ...any) (response.Response[fileapi.ResponseMultiDTO], error) {
	paths := make([]string, 0)
	for _, header := range d.Headers {
		if header.Size > constant.MaxFileSize {
			return response.Response[fileapi.ResponseMultiDTO]{}, errorext.NewCustomError(http.StatusBadRequest, errors.New("the file is too large. the file must be less than 10MB in size"))
		}
		f, err := header.Open()
		if err != nil {
			return response.Response[fileapi.ResponseMultiDTO]{}, errorext.NewCustomError(http.StatusBadRequest, errors.New("one of the files cannot be processed"))
		}
		defer f.Close()
		// process file
		// get content/MIME type
		contentType, err := file.GetMultipartFileContentType(f, true)
		if err != nil {
			return response.Response[fileapi.ResponseMultiDTO]{}, errorext.BuildCustomError(err)
		}
		log.Println(contentType)
		// check mime type
		validMIME := file.IsAllowedMIMEType(contentType, constant.AllowedMimeTypes[:])
		if !validMIME {
			return response.Response[fileapi.ResponseMultiDTO]{}, errorext.BuildCustomError(err)
		}
		// proceed to save the file
		p, err := file.SaveFile(ctx, f, "uploads", fmt.Sprintf("%d", timeext.NowUnixMilli())+filepath.Ext(header.Filename), nil)
		if err != nil {
			return response.Response[fileapi.ResponseMultiDTO]{}, errorext.BuildCustomError(err)
		}
		paths = append(paths, p)
	}
	return response.Response[fileapi.ResponseMultiDTO]{Data: fileapi.ResponseMultiDTO{FilePaths: paths}}, nil
}

func (s *Service) UploadMultipleOutputProgress(ctx context.Context, d fileapi.UploadMultipleDTO, args ...any) (response.Response[fileapi.ResponseMultiDTO], error) {
	paths := make([]string, 0)
	for _, header := range d.Headers {
		if header.Size > constant.MaxFileSize {
			return response.Response[fileapi.ResponseMultiDTO]{}, errorext.NewCustomError(http.StatusBadRequest, errors.New("the file is too large. the file must be less than 10MB in size"))
		}
		f, err := header.Open()
		if err != nil {
			return response.Response[fileapi.ResponseMultiDTO]{}, errorext.NewCustomError(http.StatusBadRequest, errors.New("one of the files cannot be processed"))
		}
		defer f.Close()
		// process file
		// get content/MIME type
		contentType, err := file.GetMultipartFileContentType(f, true)
		if err != nil {
			return response.Response[fileapi.ResponseMultiDTO]{}, errorext.BuildCustomError(err)
		}
		log.Println(contentType)
		// check mime type
		validMIME := file.IsAllowedMIMEType(contentType, constant.AllowedMimeTypes[:])
		if !validMIME {
			return response.Response[fileapi.ResponseMultiDTO]{}, errorext.BuildCustomError(err)
		}
		prog := ioext.NewProgress(header.Size)
		// proceed to save the file
		p, err := file.SaveFile(ctx, f, "uploads", fmt.Sprintf("%d", timeext.NowUnixMilli())+filepath.Ext(header.Filename), prog)
		if err != nil {
			return response.Response[fileapi.ResponseMultiDTO]{}, errorext.BuildCustomError(err)
		}
		paths = append(paths, p)
	}
	return response.Response[fileapi.ResponseMultiDTO]{Data: fileapi.ResponseMultiDTO{FilePaths: paths}}, nil
}
