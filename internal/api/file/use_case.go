package file

import (
	"context"

)

type UseCase interface {
	Upload(ctx context.Context, d UploadDTO, args ...any) (ResponseDTO, error)
	
	MultiUpload(ctx context.Context, d UploadDTO, args ...any) (ResponseDTO, error)
}
