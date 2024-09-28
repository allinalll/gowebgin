package services

import (
	"926new/models"
	"errors"
	"strconv"
)

var products []models.Product

func GetProductByID(id string) *models.Product {
	for _, p := range products {
		if strconv.Itoa(int(p.ID)) == id {
			return &p
		}
	}
	return nil
}

func UpdateProduct(id string, updatedProduct *models.Product) error {
	for i, p := range products {
		if strconv.Itoa(int(p.ID)) == id {
			updatedProduct.ID = p.ID // 保留原 ID
			products[i] = *updatedProduct
			return nil
		}
	}
	return errors.New("product not found")
}

func DeleteProduct(id string) error {
	for i, p := range products {
		if strconv.Itoa(int(p.ID)) == id {
			products = append(products[:i], products[i+1:]...) // 删除商品
			return nil
		}
	}
	return errors.New("product not found")
}
