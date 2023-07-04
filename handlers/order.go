package handlers

import (
	"sellerApp/models"
	"sellerApp/repository"
)

type OrderHandler struct {
	repo repository.OrderRepo
}

func New(repo repository.OrderRepo) OrderHandler {
	return OrderHandler{repo: repo}
}
func (oh *OrderHandler) Get(f models.Filter) ([]models.Order, error) {
	return oh.repo.Get(f)
}

func (oh *OrderHandler) Create(o models.Order) error {
	err := o.Validate()
	if err != nil {
		return err
	}
	err = oh.repo.Create(o)
	return err
}
