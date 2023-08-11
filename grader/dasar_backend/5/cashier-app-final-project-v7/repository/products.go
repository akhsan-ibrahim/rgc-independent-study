package repository

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
)

type ProductRepository struct {
	db db.DB
}

func NewProductRepository(db db.DB) ProductRepository {
	return ProductRepository{db}
}

func (u *ProductRepository) ReadProducts() ([]model.Product, error) {
	jsonContent, err := u.db.Load("products")
	if err != nil {
		return []model.Product{}, err
	}

	lsProduct := []model.Product{}
	json.Unmarshal(jsonContent, &lsProduct)

	return lsProduct, nil // TODO: replace this
}

func (u *ProductRepository) ResetProducts() error {
	err := u.db.Reset("products", []byte("[]"))
	if err != nil {
		return err
	}

	return nil
}
