package adapters

import (
	"ideal_cleaning/internal/domain"
	"time"

	"github.com/jackc/pgx"
)

type productRepo struct {
	db *pgx.Conn
}

func NewProductRepo(db *pgx.Conn) domain.ProductRepository {
	return &productRepo{db: db}
}

func (p *productRepo) ProductInsert(product domain.ProductInput) error {
	CreatedAt := time.Now()
	_, err := p.db.Exec("INSERT INTO products (name, price, count, created_at) values ($1,$2,$3,$4)",
		product.Name, product.Price, product.Count, CreatedAt)

	return err
}

// func (p *productRepo)UpdateProduct(product domain.ProductUpdateInsert) error{

// }
