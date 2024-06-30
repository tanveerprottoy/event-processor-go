package file

import (
	"context"

	"github.com/tanveerprottoy/event-processor-go/pkg/response"
)

type UseCase interface {
	Upload(ctx context.Context, d UploadDTO, args ...any) (response.Response[ResponseDTO], error)
}
