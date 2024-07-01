package handler

import (
	"log"
	"net/http"

	fileapi "github.com/tanveerprottoy/event-processor-go/internal/api/file"
	"github.com/tanveerprottoy/event-processor-go/pkg/constant"
	"github.com/tanveerprottoy/event-processor-go/pkg/errorext"
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
	log.Println("handler.File.Upload")
	ctx := r.Context()
	// http.MaxBytesReader limits the request body size
	r.Body = http.MaxBytesReader(w, r.Body, constant.MaxFileSize)
	if err := httpext.ParseMultiPartForm(r, constant.MaxFileSize); err != nil {
		http.Error(w, "the file is too large. the file must be less than 10MB in size", http.StatusBadRequest)
	}

	res, err := h.useCase.Upload(ctx, fileapi.UploadDTO{Req: r})

	if err != nil {
		err := errorext.ParseCustomError(err)
		http.Error(w, err.Error(), err.Code())
	}

	response.Respond(w, http.StatusOK, response.BuildData(res))
}

func (h *File) MultipleUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("handler.File.MultipleUpload")
	// ctx := r.Context()
	if err := httpext.ParseMultiPartForm(r, constant.MaxMemoryBytes); err != nil {
		http.Error(w, "payload too large", http.StatusBadRequest)
	}

	response.Respond(w, http.StatusOK, response.BuildData("success"))
}
