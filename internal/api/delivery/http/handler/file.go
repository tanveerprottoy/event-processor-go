package handler

import (
	"log"
	"net/http"

	fileapi "github.com/tanveerprottoy/event-processor-go/internal/api/file"
	"github.com/tanveerprottoy/event-processor-go/pkg/constant"
	"github.com/tanveerprottoy/event-processor-go/pkg/file"
	"github.com/tanveerprottoy/event-processor-go/pkg/httpext"
	"github.com/tanveerprottoy/event-processor-go/pkg/response"
)

// File handles incoming requests
type File struct {
	useCase fileapi.UseCase
}

// NewFile initializes a new Handler
func NewFile(u fileapi.UseCase) *File {
	return &File{useCase: u}
}

func (h *File) Upload(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	// http.MaxBytesReader limits the request body size
	r.Body = http.MaxBytesReader(w, r.Body, constant.MaxFileSize)
	f, head, err := httpext.GetFile(r, constant.MaxFileSize)
	log.Panicln(head)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer f.Close()

	// get content/MIME type
	contentType, err := file.GetMultipartFileContentType(f, true)
	log.Println(contentType)
	// check mime type
	/* validMIME := filepkg.IsAllowedMIMEType(head.Filename, constant.AllowedMimeTypes[:])
	if !validMIME {
		return dto, errorext.HTTPError{Code: http.StatusBadRequest, Err: errors.New(constant.UnsupportedFileType)}
	}

	_, err := file.Seek(0, io.SeekStart) */
	response.Respond(http.StatusCreated, response.BuildData(http.StatusCreated, ""), w)
}
