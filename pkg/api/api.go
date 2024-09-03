package api

type Category struct {
	Id   int
	Name string
}

type Product struct {
	Id         int
	Name       string
	CategoryId int
}

func GetCategories() []Category {
	result := make([]Category, 0)

	result = append(result, Category{Id: 0, Name: "Пицца"})
	result = append(result, Category{Id: 1, Name: "Круасаны"})
	result = append(result, Category{Id: 2, Name: "Напитки"})
	result = append(result, Category{Id: 3, Name: "ПИВО"})
	result = append(result, Category{Id: 4, Name: "ПИВО2"})
	result = append(result, Category{Id: 5, Name: "ПИВО3"})
	result = append(result, Category{Id: 6, Name: "ПИВО4"})
	result = append(result, Category{Id: 7, Name: "ПИВО5"})
	result = append(result, Category{Id: 8, Name: "ПИВО6"})

	return result
}

func GetProductsByCategoryId(categoryId int) []Product {
	result := make([]Product, 0)

	result = append(result, Product{Id: 0, Name: "Пицца", CategoryId: 0})
	result = append(result, Product{Id: 1, Name: "Пицца1", CategoryId: 0})
	result = append(result, Product{Id: 2, Name: "Пицца2", CategoryId: 0})
	result = append(result, Product{Id: 3, Name: "Пицца3", CategoryId: 0})

	return result
}
