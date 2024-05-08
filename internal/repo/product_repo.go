package repo

import (
	"database/sql"
	"transaction_ex1/internal/entity"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) FindProduct(productId int) (*entity.Product, error) {
	row := r.db.QueryRow("SELECT id, name, price FROM products WHERE id = $1", productId)

	var (
		id    int
		name  string
		price int
	)

	if err := row.Scan(&id, &name, &price); err != nil {
		return nil, err
	}

	product := entity.NewProduct(id, name, price)

	return &product, nil
}
