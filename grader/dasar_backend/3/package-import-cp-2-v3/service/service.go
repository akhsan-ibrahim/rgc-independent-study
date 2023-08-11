package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"fmt"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Pay(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	items, _ := s.database.GetCartItems()
	prod, err := s.database.GetProductByName(productName)

	if len(items) == 0 {
		if err == nil && quantity > 0 {
			var detailProd = entity.CartItem{
				ProductName: prod.Name,
				Price:       prod.Price,
				Quantity:    quantity,
			}
			items = append(items, detailProd)
		} else if err == nil && quantity <= 0 {
			err = fmt.Errorf("invalid quantity")
			return err
		}
	} else {
		have := 0
		for i := 0; i < len(items); i++ {
			if items[i].ProductName == productName && quantity > 0 {
				items[i].Quantity = quantity
				have = 1
			} else if items[i].ProductName == productName && quantity <= 0 {
				s.RemoveCart(productName)
				err = fmt.Errorf("invalid quantity")
				return err
			}
		}
		if have == 0 {
			if err == nil && quantity > 0 {
				var detailProd = entity.CartItem{
					ProductName: prod.Name,
					Price:       prod.Price,
					Quantity:    quantity,
				}
				items = append(items, detailProd)
			} else if err == nil && quantity <= 0 {
				err = fmt.Errorf("invalid quantity")
				return err
			}
		}
	}

	s.database.SaveCartItems(items)
	return err // TODO: replace this
}

func (s *Service) RemoveCart(productName string) error {
	prod, err := s.database.GetCartItems()
	err = fmt.Errorf("product not found")
	for i := 0; i < len(prod); i++ {
		if prod[i].ProductName == productName {
			prod = append(prod[:i], prod[i+1:]...)
			err = nil
		}
	}
	s.database.SaveCartItems(prod)
	return err // TODO: replace this
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.GetCartItems()
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (s *Service) ResetCart() error {
	s.database.SaveCartItems([]entity.CartItem{})
	return nil // TODO: replace this
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	allProd := s.database.GetProductData()
	return allProd, nil // TODO: replace this
}

func (s *Service) Pay(money int) (entity.PaymentInformation, error) {
	items, _ := s.database.GetCartItems()
	var total int
	err := fmt.Errorf("money is not enough")
	for _, val := range items {
		// fmt.Println(val)
		total += val.Price * val.Quantity
	}

	if total > money {
		return entity.PaymentInformation{}, err
	}

	change := money - total
	var info = entity.PaymentInformation{
		ProductList: items,
		TotalPrice:  total,
		MoneyPaid:   money,
		Change:      change,
	}
	s.ResetCart()
	return info, nil // TODO: replace this
}
