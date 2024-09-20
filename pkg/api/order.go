package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
)

func (a *MainApi) CreateOrder(order CreateOrderRequest) error {
	u := a.formatBaseUrl(orderRout)

	body, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		logrus.Error("Error creating request")
	}

	resp, err := a.doRequest(req)
	if err != nil {
		return fmt.Errorf("can't do request: %w", err)
	}

	logrus.Info(resp)

	return nil
}

func (a *MainApi) GetOrder(orderId int) (Order, error) {
	u := a.formatBaseUrl(orderRout + "/" + strconv.Itoa(orderId))

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		logrus.Error("Error creating request")
	}

	resp, err := a.doRequest(req)
	if err != nil {
		return Order{}, fmt.Errorf("can't do request: %w", err)
	}

	var order Order
	err = json.Unmarshal(resp, &order)
	if err != nil {
		return Order{}, fmt.Errorf("can't unmarshall response: %w", err)
	}

	return order, nil
}

func (a *MainApi) GetOrders(userName string) ([]OrderShort, error) {
	u := a.formatBaseUrl(orderRout)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		logrus.Error("Error creating request")
	}

	var params = url.Values{}
	params.Add("userName", userName)
	req.URL.RawQuery = params.Encode()

	resp, err := a.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

	var orders []OrderShort
	err = json.Unmarshal(resp, &orders)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshall response: %w", err)
	}

	return orders, nil
}
