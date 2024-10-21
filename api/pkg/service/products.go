package service

import (
	Telegram_Market "Telegram-Market"
	"Telegram-Market/api/pkg/repository"
)

type ProductsTelegram struct {
	repo repository.Products
}

func NewProductsTelegram(repo repository.Products) *ProductsTelegram {
	return &ProductsTelegram{repo: repo}
}

func (t *ProductsTelegram) CreateProduct(product Telegram_Market.Products) (int64, error) {
	return t.repo.CreateProduct(product)
}

func (t *ProductsTelegram) GetProductById(productId int64) (Telegram_Market.Products, error) {
	return t.repo.GetProductById(productId)
}

func (t *ProductsTelegram) GetAllProducts() ([]Telegram_Market.Products, error) {
	return t.repo.GetAllProducts()
}

func (t *ProductsTelegram) UpdateProduct(productId int64, product Telegram_Market.UpdateProducts) error {
	return t.repo.UpdateProduct(productId, product)
}

func (t *ProductsTelegram) DeleteProduct(productId int64) error {
	return t.repo.DeleteProduct(productId)
}
