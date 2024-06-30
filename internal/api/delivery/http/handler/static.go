package handler

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Static handles incoming requests
type Static struct {
	formFilePath string
}

// NewStatic initializes a new Handler
func NewStatic() *Static {
	pwd, err := os.Getwd()

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Println("pwd: ", pwd)

	return &Static{formFilePath: filepath.Join(pwd, "/html/form.html")}
}

func (h *Static) ServeHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")

	http.ServeFile(w, r, h.formFilePath)
}
