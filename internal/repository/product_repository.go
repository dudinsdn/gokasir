package repository

import (
	"database/sql"

	"github.com/dudinsdn/gokasir/internal/entity"
)

type ProductRepository interface {
	List(filter string, sort string, page int, pageSize int) ([]entity.Product, int, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) List(filter string, sort string, page int, pageSize int) ([]entity.Product, int, error) {
	query := "SELECT id, name, price, stock, created_at, updated_at FROM products WHERE name ILIKE $1 ORDER BY " + sort + " LIMIT $2 OFFSET $3"
	rows, err := r.db.Query(query, "%"+filter+"%", pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var p entity.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}
	// Hitung total untuk pagination
	var total int
	r.db.QueryRow("SELECT COUNT(*) FROM products WHERE name ILIKE $1", "%"+filter+"%").Scan(&total)
	return products, total, nil
}
