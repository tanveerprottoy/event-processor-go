package handler

import (
	"net/http"

	"github.com/tanveerprottoy/event-processor-go/internal/api/file"
)

// File handles incoming requests
type File struct {
	useCase file.UseCase
}

// NewFile initializes a new Handler
func NewFile(u file.UseCase) *File {
	return &File{useCase: u}
}

func (h *File) Upload(w http.ResponseWriter, r *http.Request) {
	/* ctx := r.Context()
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return
	}
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" { {
		http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
		return
	}

	_, err := file.Seek(0, io.SeekStart)
	var v activitytype.CreateDTO
	// validate the request body
	err = httpext.ParseRequestBody(r.Body, &v)
	if err != nil {
		response.RespondError(http.StatusBadRequest, response.BuildError(http.StatusBadRequest, err), w)
		return
	}
	if companyID != "" {
		v.CompanyID = companyID
	}
	// validate the request body
	validationErrs := h.validater.Validate(&v)
	if validationErrs != nil {
		response.RespondError(http.StatusBadRequest, response.BuildError(http.StatusBadRequest, validationErrs), w)
		return
	}
	d, httpErr := h.useCase.Upload(ctx, v)
	if httpErr.Err != nil {
		response.RespondError(httpErr.Code, response.BuildError(httpErr.Code, httpErr.Err), w)
		return
	}
	response.Respond(http.StatusCreated, response.BuildData(http.StatusCreated, d), w) */
}
