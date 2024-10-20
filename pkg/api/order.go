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

func (a *Api) CreateOrder(order CreateOrderRequest) error {
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

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.doRequest(req)
	if err != nil {
		return fmt.Errorf("can't do request: %w", err)
	}

	logrus.Info(string(resp))

	return nil
}

func (a *Api) GetOrder(orderId int) (Order, error) {
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

func (a *Api) GetOrders(userName string) ([]OrderShort, error) {
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

func (a *Api) ChangeOrderStatus(orderId, targetStatus int) error {
	u := a.formatBaseUrl(orderRout + "/" + strconv.Itoa(orderId) + "/status")

	req, err := http.NewRequest(http.MethodPatch, u.String(), nil)
	if err != nil {
		logrus.Error("Error creating request")
	}

	var params = url.Values{}
	params.Add("targetStatus", strconv.Itoa(targetStatus))
	req.URL.RawQuery = params.Encode()

	resp, err := a.doRequest(req)
	if err != nil {
		return fmt.Errorf("can't do request: %w", err)
	}

	var order Order
	err = json.Unmarshal(resp, &order)
	if err != nil {
		return fmt.Errorf("can't unmarshall response: %w", err)
	}

	return nil
}
