package logic

import (
	"github.com/bryancr89/groceries-api/models"
	"github.com/bryancr89/groceries-api/dao"
	"encoding/json"
	"fmt"
)


var DAO dao.IProductDAO

func Initialize(projectsDAO dao.IProductDAO) {
	DAO = projectsDAO
}

func GetProducts() ([]models.Product, error) {
	return DAO.GetProducts()
}

func GetProduct(Id string) (models.Product, error) {
	return DAO.GetProduct(Id)
}

func CreateProduct(product models.Product) (models.Product, error) {
	return DAO.CreateProduct(product)
}

func DeleteProduct(Id string) (models.Product, error) {
	return DAO.DeleteProduct(Id)
}