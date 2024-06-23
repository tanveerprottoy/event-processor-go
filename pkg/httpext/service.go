package httpext

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// RequestWithContext makes a request with context
// can have optional args
// args[0] = retryInterval
// args[1] = retryCount
func RequestWithContext[T any](ctx context.Context, client *ClientProvider, method string, url string, header http.Header, body io.Reader, args ...any) (*T, []byte, error) {
	res, err := client.RequestWithContext(ctx, method, url, header, body)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	code := res.StatusCode
	if code >= http.StatusOK && code < http.StatusMultipleChoices {
		// res ok, parse response body to type
		var d T
		err := json.NewDecoder(res.Body).Decode(&d)
		if err != nil {
			return nil, nil, err
		}
		return &d, nil, nil
	} else {
		// res not ok
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, nil, err
		}
		return nil, b, errors.New("error response was returned")
	}
}
