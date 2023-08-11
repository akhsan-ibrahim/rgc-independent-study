package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return CartRepository{db}
}

func (c *CartRepository) ReadCart() ([]model.JoinCart, error) {
	result := []model.JoinCart{}
	err := c.db.Table("carts").Select("carts.id, carts.product_id, products.name, carts.quantity, carts.total_price").Joins("INNER JOIN products ON products.id = carts.product_id").Find(&result).Error
	if err != nil {
		return []model.JoinCart{}, err
	}
	return result, nil // TODO: replace this
}

func (c *CartRepository) AddCart(product model.Product) error {
	c.db.Transaction(func(tx *gorm.DB) error {
		lsChart := model.Cart{}
		c.db.Table("carts").Select("*").Where("product_id = ?", product.ID).Find(&lsChart)

		if lsChart.ID == 0 {
			err := tx.Table("carts").Create(&model.Cart{
				ProductID:  product.ID,
				Quantity:   1,
				TotalPrice: product.Price - (product.Price * product.Discount / 100),
			}).Error
			if err != nil {
				return err
			}

			err = tx.Table("products").Where("id = ?", product.ID).Update("stock", product.Stock-1).Error
			if err != nil {
				return err
			}
		} else {
			err := tx.Table("carts").Where("product_id = ?", product.ID).Updates(&model.Cart{
				Quantity:   lsChart.Quantity + 1,
				TotalPrice: product.Price*(lsChart.Quantity+1) - (product.Price * product.Discount / 100 * (lsChart.Quantity + 1)),
			}).Error
			if err != nil {
				return err
			}

			err = tx.Table("products").Where("id = ?", product.ID).Update("stock", product.Stock-1).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
	return nil // TODO: replace this
}

func (c *CartRepository) DeleteCart(id uint, productID uint) error {
	c.db.Transaction(func(tx *gorm.DB) error {
		lsChr := model.Cart{}
		c.db.Table("carts").Select("*").Where("id = ?", id).Find(&lsChr)

		pr := model.Product{}
		c.db.Table("products").Select("*").Where("id = ?", productID).Find(&pr)

		err := tx.Delete([]model.Cart{}, id).Error
		if err != nil {
			return err
		}

		err = tx.Table("products").Where("id = ?", productID).Update("stock", pr.Stock+int(lsChr.Quantity)).Error
		if err != nil {
			return err
		}

		return nil
	})
	return nil // TODO: replace this
}

func (c *CartRepository) UpdateCart(id uint, cart model.Cart) error {
	err := c.db.Table("carts").Where("id = ?", id).Updates(cart).Error
	return err // TODO: replace this
}
