package httpext

import (
	"net/http"

	"github.com/tanveerprottoy/event-processor-go/pkg/constant"
)

func ParseMultiPartForm(r *http.Request, maxMemory int64) error {
	if maxMemory == 0 {
		// left shift 32 << 20 which results in 32*2^20 = 33554432
		// x << y, results in x*2^y
		maxMemory = constant.MaxMemoryBytes
	}
	return r.ParseMultipartForm(maxMemory)
}
