package httpext

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// GetFile fetches the file
func GetFile(r *http.Request) string {
	// left shift 32 << 20 which results in 32*2^20 = 33554432
	// x << y, results in x*2^y
	// max allowed value is 32MB
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return ""
	}
	n := r.Form.Get("name")
	// Retrieve the file from form data
	f, h, err := r.FormFile("file")
	if err != nil {
		return ""
	}
	defer f.Close()
	
}
