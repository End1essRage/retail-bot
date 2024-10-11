package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
)

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
	u := a.formatBaseUrl(productRout)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	var params = url.Values{}
	params.Add("id", strconv.Itoa(categoryId))
	req.URL.RawQuery = params.Encode()

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
	u := a.formatBaseUrl(productRout + strconv.Itoa(productId))

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
