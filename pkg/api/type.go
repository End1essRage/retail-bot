package api

import "time"

type Category struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Parent int    `json:"parent"`
}

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryId  int     `json:"categoryId"`
}

type Position struct {
	Product Product `json:"product"`
	Count   int     `json:"count"`
}

type Order struct {
	Id        int        `json:"id"`
	Positions []Position `json:"positions"`
	Status    string     `json:"status"`
}

type OrderShort struct {
	Id           int       `json:"id"`
	DateCreation time.Time `json:"dateCreation"`
	Status       int       `json:"status"`
	StatusName   string    `json:"statusName"`
}

type CreateOrderRequest struct {
	UserName  string      `json:"userName"`
	Positions map[int]int `json:"positions"`
}
