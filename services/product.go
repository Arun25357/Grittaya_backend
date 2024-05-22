package services

import (
	"errors"

	"github.com/Pure227/Grittaya_backend/models"
	uuid "github.com/satori/go.uuid"
)

var products = []models.Product{}

func GetAllProducts() ([]models.Product, error) {
	return products, nil
}

func CreateProduct(product *models.Product) error {
	products = append(products, *product)
	return nil
}

func GetProductByID(id uuid.UUID) (models.Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return models.Product{}, errors.New("product not found")
}

func UpdateProduct(updatedProduct *models.Product) error {
	for i, product := range products {
		if product.ID == updatedProduct.ID {
			products[i] = *updatedProduct
			return nil
		}
	}
	return errors.New("product not found")
}

func DeleteProduct(id uuid.UUID) error {
	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}
