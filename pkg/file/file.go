package file

import (
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// FilePathWalkDir walks the file path and returns the files
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// ReadDir reads the directory and returns the files
func ReadDir(root string) ([]string, error) {
	var files []string
	dirEntries, err := os.ReadDir(root)
	if err != nil {
		return files, err
	}
	for _, file := range dirEntries {
		files = append(files, file.Name())
	}
	return files, nil
}

// ReadDir1 reads the directory and returns the files
func ReadDir1(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfos, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfos {
		files = append(files, file.Name())
	}
	return files, nil
}

// CreateDirIfNotExists creates a directory if it does not exist
func CreateDirIfNotExists(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return os.Mkdir(path, os.ModePerm)
	} else {
		return err
	}
}

// ReadFile reads a file from the root directory
func ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

// SaveFile saves a file to the root directory
func SaveFile(multipartFile multipart.File, dirName string, fileName string) (string, error) {
	path := filepath.Join(".", dirName)
	_ = os.MkdirAll(path, os.ModePerm)
	fullPath := path + "/" + fileName
	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer file.Close()
	// Copy the file to the destination path
	_, err = io.Copy(file, multipartFile)
	if err != nil {
		return "", err
	}
	return fullPath, nil
}

func t() {
	path := filepath.Join(".", "files")
	_ = os.MkdirAll(path, os.ModePerm)
	fullPath := path + "/" + n + filepath.Ext(h.Filename)
	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return ""
	}
	defer file.Close()
	// Copy the file to the destination path
	_, err = io.Copy(file, f)
	if err != nil {
		return ""
	}
	return fullPath
}

// GetFileContentType returns the content type of the file
func GetFileContentType(file *os.File) (string, error) {
	// to sniff the content type only the first
	// 512 bytes are used.
	buf := make([]byte, 512)

	_, err := file.Read(buf)

	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buf)

	return contentType, nil
}

// GetExtension returns the extension of the file
func GetExtension(fileName string) string {
	return strings.ToLower(filepath.Ext(fileName))
}

// GetMIMEType returns the mime type of the file
func GetMIMEType(fileName string) string {
	ext := filepath.Ext(fileName)
	return mime.TypeByExtension(ext)
}

// IsAllowedMIMEType checks if the file is of an allowed mime type
func IsAllowedMIMEType(fileName string, allowedMimeTypes []string) bool {
	mimeType := GetMIMEType(fileName)
	for _, allowedType := range allowedMimeTypes {
		if strings.HasPrefix(mimeType, allowedType) {
			return true
		}
	}
	return false
}

// IsAudioMIMEType checks if the file is an audio file
func IsAudioMIMEType(fileName string) bool {
	ext := filepath.Ext(fileName)
	mimeType := mime.TypeByExtension(ext)
	return strings.HasPrefix(mimeType, "audio/")
}

// IsImageMIMEType checks if the file is an image file
func IsImageMIMEType(fileName string) bool {
	ext := filepath.Ext(fileName)
	mimeType := mime.TypeByExtension(ext)
	return strings.HasPrefix(mimeType, "image/")
}

// IsVideoMIMEType checks if the file is a video file
func IsVideoMIMEType(fileName string) bool {
	ext := filepath.Ext(fileName)
	mimeType := mime.TypeByExtension(ext)
	return strings.HasPrefix(mimeType, "video/")
}

// IsTargetMIMEType checks if the file is of the target mime type
func IsTargetMIMEType(fileName string, targetMimeType string) bool {
	ext := filepath.Ext(fileName)
	mimeType := mime.TypeByExtension(ext)
	return strings.HasPrefix(mimeType, targetMimeType)
}
