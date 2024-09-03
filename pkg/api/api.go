package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type IApi interface {
	GetCategories() ([]Category, error)
	GetProducts(categoryId int) ([]Product, error)
}

type Api struct {
	host     string
	basePath string
	client   http.Client
}

func NewApi(host string) *Api {
	return &Api{host: host,
		basePath: viper.GetString("api_basepath"),
		client:   http.Client{}}
}

func (a *Api) GetCategories() ([]Category, error) {
	u := a.formatBaseUrl("menu")
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		logrus.Error("Error creating request")
	}

	logrus.Info("path is " + req.URL.Path)

	resp, err := a.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

	logrus.Warn(req.Host)

	var categories []Category
	err = json.Unmarshal(resp, &categories)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshall response: %w", err)
	}

	return categories, nil
}

func (a *Api) GetProducts(categoryId int) ([]Product, error) {
	u := a.formatBaseUrl("menu/category/" + strconv.Itoa(categoryId))
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		logrus.Error("Error creating request")
	}

	resp, err := a.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}
	logrus.Info(req.URL.Path)
	var products []Product
	err = json.Unmarshal(resp, &products)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshall response: %w", err)
	}

	return products, nil
}

func (a *Api) formatBaseUrl(rout string) url.URL {
	return url.URL{
		Scheme: viper.GetString("api_sheme"),
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
