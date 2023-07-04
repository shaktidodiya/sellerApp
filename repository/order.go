package repository

import (
	"database/sql"
	"sellerApp/errors"
	"sellerApp/models"
	"time"
)

type OrderRepo struct {
	client *sql.DB
}

func New(db *sql.DB) *OrderRepo {
	return &OrderRepo{client: db}
}

func (db *OrderRepo) Create(order models.Order) error {
	query := "INSERT INTO orders (product_id, user_id, price, qty, created_at) values(?, ?, ?, ?, ?)"
	_, err := db.client.Exec(query, order.ProductID, order.UserID, order.Price, order.Qty, time.Now().UTC())
	if err != nil {
		return errors.DBError{Err: err}
	}

	return nil
}

func (db *OrderRepo) Get(filter models.Filter) ([]models.Order, error) {
	query := "SELECT id, product_id, user_id, price, qty, created_at from orders where user_id = ? "
	args := []interface{}{filter.UserID}

	if filter.ProductID != "" {
		query += "AND product_id = ? "
		args = append(args, filter.ProductID)
	}
	if filter.Price != "" {
		query += "AND price = ? "
		args = append(args, filter.Price)
	}
	if filter.Qty != "" {
		query += "AND qty = ? "
		args = append(args, filter.Qty)
	}

	rows, err := db.client.Query(query, args...)
	if err != nil {
		return nil, errors.DBError{Err: err}
	}
	defer rows.Close()
	var orders []models.Order

	for rows.Next() {
		order := models.Order{}
		err = rows.Scan(&order.ID, &order.ProductID, &order.UserID, &order.Price, &order.Qty, &order.CreatedAt)

		if err != nil {
			return nil, errors.DBError{Err: err}
		}
		orders = append(orders, order)
	}

	return orders, nil
}
