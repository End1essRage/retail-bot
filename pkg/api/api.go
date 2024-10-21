package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	c "github.com/end1essrage/retail-bot/pkg"
)

const (
	categoriesRout = "menu/"
	productRout    = "menu/product/"
	orderRout      = "order/"
	userRout       = "user/"
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

func (a *Api) GetUserRole(userName string) (int, error) {
	u := a.formatBaseUrl(userRout + userName)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return int(c.Client), fmt.Errorf("error creating request: %w", err)
	}

	resp, err := a.doRequest(req)
	if err != nil {
		return int(c.Client), fmt.Errorf("can't do request: %w", err)
	}

	var roleId int
	err = json.Unmarshal(resp, &roleId)
	if err != nil {
		return int(c.Client), fmt.Errorf("can't unmarshall response: %w", err)
	}

	return roleId, nil
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
