package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

// переименовать
type Api interface {
	GetCategories() ([]Category, error)
	GetProducts(categoryId int) ([]Product, error)
	GetProductData(productId int) (Product, error)

	GetOrders(userName string) ([]OrderShort, error)
	GetOrder(orderId int) (Order, error)
	CreateOrder(order CreateOrderRequest) error

	ChangeOrderStatus(orderId, targetStatus int) error
}

const (
	categoriesRout = "menu"
	productRout    = "menu/product/"
	orderRout      = "order"
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
