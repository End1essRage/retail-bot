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
)

// переименовать
type Api interface {
	GetCategories() ([]Category, error)
	GetProducts(categoryId int) ([]Product, error)
	GetProductData(productId int) (Product, error)
}

const (
	categoriesRout = "menu"
	productsRout   = "menu/category/"
)

type MainApi struct {
	host     string
	basePath string
	client   http.Client
	scheme   string
}

func NewMainApi(host string, basePath string, scheme string) *MainApi {
	return &MainApi{host: host,
		basePath: basePath,
		scheme:   scheme,
		client:   http.Client{}}
}

func (a *MainApi) formatBaseUrl(rout string) url.URL {
	return url.URL{
		Scheme: a.scheme,
		Host:   a.host,
		Path:   path.Join(a.basePath, rout),
	}
}

func (a *MainApi) GetCategories() ([]Category, error) {
	u := a.formatBaseUrl(categoriesRout)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		logrus.Error("Error creating request")
	}

	resp, err := a.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

	var categories []Category
	err = json.Unmarshal(resp, &categories)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshall response: %w", err)
	}

	return categories, nil
}

func (a *MainApi) GetProducts(categoryId int) ([]Product, error) {
	u := a.formatBaseUrl(productsRout + strconv.Itoa(categoryId))

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := a.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

	var products []Product
	err = json.Unmarshal(resp, &products)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshall response: %w", err)
	}

	return products, nil
}

func (a *MainApi) GetProductData(productId int) (Product, error) {
	u := a.formatBaseUrl(productsRout + strconv.Itoa(productId))

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return Product{}, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := a.doRequest(req)
	if err != nil {
		return Product{}, fmt.Errorf("can't do request: %w", err)
	}

	var product Product
	err = json.Unmarshal(resp, &product)
	if err != nil {
		return Product{}, fmt.Errorf("can't unmarshall response: %w", err)
	}

	return product, nil
}

func (a *MainApi) doRequest(req *http.Request) ([]byte, error) {
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
