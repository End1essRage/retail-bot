package api

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageurl"`
}
