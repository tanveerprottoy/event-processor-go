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
	// http.MaxBytesReader limits the request body size
	r.Body = http.MaxBytesReader(w, r.Body, constant.MaxFileSize)
	if err := httpext.ParseMultiPartForm(r, constant.MaxFileSize); err != nil {
		http.Error(w, "the file is too large. the file must be less than 10MB in size", http.StatusBadRequest)
	}

	defer r.Body.Close()

	// get the file
	f, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "could not retrieve the file", http.StatusBadRequest)
	}

	defer f.Close()

	dto := fileapi.UploadDTO{
		File:   f,
		Header: header,
	}

	res, err := h.useCase.Upload(r.Context(), dto)

	if err != nil {
		err := errorext.ParseCustomError(err)
		http.Error(w, err.Error(), err.Code())
	}

	response.Respond(w, http.StatusOK, response.BuildData(res))
}

func (h *File) UploadMultiple(w http.ResponseWriter, r *http.Request) {
	log.Println("handler.File.MultipleUpload")
	err := httpext.ParseMultiPartForm(r, constant.MaxMemoryBytes)
	if err != nil {
		http.Error(w, "payload too large", http.StatusBadRequest)
	}

	var res response.Response[fileapi.ResponseMultiDTO]
	param := httpext.GetQueryParam(r, "outputProgress")
	ctx := r.Context()
	headers := r.MultipartForm.File["files"]

	if len(headers) == 0 {
		http.Error(w, "could not retrieve the files", http.StatusBadRequest)
	}

	dto := fileapi.UploadMultipleDTO{
		Headers: headers,
	}

	if param != "" && param == "true" {
		res, err = h.useCase.UploadMultipleOutputProgress(ctx, dto)
	} else {
		res, err = h.useCase.UploadMultiple(ctx, dto)
	}
	if err != nil {
		err := errorext.ParseCustomError(err)
		http.Error(w, err.Error(), err.Code())
	}

	response.Respond(w, http.StatusOK, response.BuildData(res))
}
