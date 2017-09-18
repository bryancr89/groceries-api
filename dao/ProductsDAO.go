package dao

import (
	"github.com/bryancr89/groceries-api/models"
)

type ProductsDAO struct { }

var products = []models.Product{
	{Id: "1", Name: "Arroz"},
	{Id: "2", Name: "Frijoles"},
}

func (p ProductsDAO) GetProducts() ([]models.Product, error) {
	return products, nil
}

func (p ProductsDAO) GetProduct(Id string) (models.Product, error) {
	for _, item := range products {
		if item.Id == Id {
			return item, nil
		}
	}
	return models.Product{}, error("The product doesn't exist")
}

func (p ProductsDAO) CreateProduct(product models.Product) (models.Product, error) {
	products = append(products, product)
	return models.Product{}, nil
}

func (p ProductsDAO) DeleteProduct(Id string) (models.Product, error) {
	for index, item := range products {
		if item.Id == Id {
			products = append(products[:index], products[index + 1:]...)
			return item, nil
		}
	}
	return models.Product{}, error("Oops")
}
