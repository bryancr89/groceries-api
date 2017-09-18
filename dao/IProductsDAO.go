package dao

import "github.com/bryancr89/groceries-api/models"

type IProductDAO interface {
	GetProducts() ([]models.Product, error)
	GetProduct(Id string) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	DeleteProduct(Id string) (models.Product, error)
}
