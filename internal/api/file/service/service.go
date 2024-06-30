package service

import (
	"context"

	"github.com/tanveerprottoy/event-processor-go/internal/api/file"
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
func (s *Service) Upload(ctx context.Context, d file.UploadDTO, args ...any) (file.ResponseDTO, error) {
	return file.ResponseDTO{}, nil
}

func (s *Service) MultiUpload(ctx context.Context, d file.UploadDTO, args ...any) (file.ResponseDTO, error) {
	return file.ResponseDTO{}, nil
}

