package httpext

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)

type ClientProvider struct {
	HTTPClient *http.Client
}

func NewClientProvider(timeout time.Duration, transport *http.Transport, checkRedirectFunc func(req *http.Request, via []*http.Request) error) *ClientProvider {
	c := new(ClientProvider)
	c.HTTPClient = &http.Client{
		Timeout: timeout,
	}
	if transport != nil {
		c.HTTPClient.Transport = transport
	}
	if checkRedirectFunc != nil {
		c.HTTPClient.CheckRedirect = checkRedirectFunc
	}
	return c
}

func (c *ClientProvider) RequestWithContext(ctx context.Context, method string, url string, header http.Header, body io.Reader) (*http.Response, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.HTTPClient.Timeout)
		defer cancel()
	}
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ClientProvider) Request(method string, url string, header http.Header, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ex:
// resp, err := http.PostForm("http://example.com/form",
// url.Values{"key": {"Value"}, "id": {"123"}})
func (c *ClientProvider) PostForm(url string, header http.Header, values url.Values) (*http.Response, error) {
	res, err := http.PostForm(url, values)
	if err != nil {
		return nil, err
	}
	return res, nil
}
