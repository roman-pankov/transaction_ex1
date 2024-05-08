package repo

import (
	"database/sql"
	"transaction_ex1/internal/entity"
)

type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) Save(order entity.Order) error {
	query := "INSERT INTO orders (id, user_id, product_id) VALUES ($1, $2, $3)"

	_, err := r.db.Exec(query, order.GetId(), order.GetUserId(), order.GetProductId())
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepo) GetNextId() (int, error) {
	var nextVal int
	err := r.db.QueryRow("SELECT nextval('orders_id_seq')").Scan(&nextVal)
	if err != nil {
		return 0, err
	}

	return nextVal, nil
}
