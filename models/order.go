package models

import "sellerApp/errors"

type Order struct {
	ID        string  `json:"id"`
	UserID    string  `json:"userId"`
	ProductID string  `json:"productId"`
	Price     float64 `json:"price"`
	Qty       int     `json:"qty"`
	CreatedAt string  `json:"createdAt"`
}

type Filter struct {
	UserID    string
	ProductID string
	Price     string
	Qty       string
}

func (o *Order) Validate() error {
	var param []string

	if o.ProductID == "" {
		param = append(param, "productId")
	}
	if o.UserID == "" {
		param = append(param, "userId")
	}
	if o.Price <= 0 {
		param = append(param, "price")
	}
	if o.Qty <= 0 {
		param = append(param, "qty")
	}

	if len(param) == 0 {
		return nil
	}

	return errors.InvalidParam{Param: param}
}
