package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

const (
	categoriesRout = "menu"
	productRout    = "menu/product/"
	orderRout      = "order"
)

type Api struct {
	host     string
	basePath string
	client   http.Client
	scheme   string
}

func NewApi(host string, basePath string, scheme string) *Api {
	return &Api{host: host,
		basePath: basePath,
		scheme:   scheme,
		client:   http.Client{}}
}

func (a *Api) formatBaseUrl(rout string) url.URL {
	return url.URL{
		Scheme: a.scheme,
		Host:   a.host,
		Path:   path.Join(a.basePath, rout),
	}
}

func (a *Api) doRequest(req *http.Request) ([]byte, error) {
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read response: %w", err)
	}

	return body, nil
}
