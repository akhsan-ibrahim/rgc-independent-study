package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (p *ProductRepository) AddProduct(product model.Product) error {
	err := p.db.Create(&product).Error

	return err // TODO: replace this
}

func (p *ProductRepository) ReadProducts() ([]model.Product, error) {
	result := []model.Product{}
	rows, err := p.db.Table("products").Select("*").Where("deleted_at IS NULL").Rows()
	if err != nil {
		return []model.Product{}, err
	}

	defer rows.Close()

	for rows.Next() {
		p.db.ScanRows(rows, &result)
	}
	return result, nil // TODO: replace this
}

func (p *ProductRepository) DeleteProduct(id uint) error {
	err := p.db.Where("id = ?", id).Delete([]model.Product{}).Error
	return err // TODO: replace this
}

func (p *ProductRepository) UpdateProduct(id uint, product model.Product) error {
	err := p.db.Table("products").Where("id = ?", id).Updates(product).Error

	return err // TODO: replace this
}
