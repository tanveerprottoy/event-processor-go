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
	f, header, err := httpext.GetFile(r, constant.MaxFileSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	defer f.Close()

	log.Printf("Size: %d, Name: %s", header.Size, header.Filename)

	dto := fileapi.UploadDTO{
		File:   f,
		Header: header,
	}

	res, err := h.useCase.Upload(ctx, dto)

	if err != nil {
		err := errorext.ParseCustomError(err)
		http.Error(w, err.Error(), err.Code())
	}

	response.Respond(w, http.StatusOK, response.BuildData(res))
}
