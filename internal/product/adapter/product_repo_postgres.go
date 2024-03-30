package adapter

import (
	"context"
	"fmt"
	"strings"
	"time"

	"example.com/m/internal/genproto/product/pb"
	"example.com/m/internal/product/domain"

	"github.com/jackc/pgx"
)

type productRepo struct {
	db *pgx.Conn
}

func NewProductRepo(db *pgx.Conn) domain.ProductRepository {
	return &productRepo{db: db}
}

func (p *productRepo) Insert(ctx context.Context, product pb.ProductRequest) error {
	CreatedAt := time.Now()
	_, err := p.db.Exec("INSERT INTO products (name, price, count, created_at) values ($1,$2,$3,$4)",
		product.Name, int(product.Price), int(product.Count), CreatedAt)

	return err
}

func (p *productRepo) GetByID(ctx context.Context, ID int) (pb.ProductResponse, error) {
	var product pb.ProductResponse
	err := p.db.QueryRow(`
        SELECT id, name, price, count, created_at, updated_at
        FROM products
        WHERE id=$1 AND deleted_at IS NULL
    `, ID).
		Scan(&product.ID, &product.Name, &product.Price, &product.Count, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return pb.ProductResponse{}, err
	}

	return product, nil
}

func (p *productRepo) UpdateByID(ctx context.Context, Id int, productInp domain.ProductUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if productInp.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=%d", argId))
		args = append(args, productInp.Name)
		argId++
	}

	if productInp.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=%d", argId))
		args = append(args, productInp.Price)
		argId++
	}

	if productInp.Count != nil {
		setValues = append(setValues, fmt.Sprintf("count=%d", argId))
		args = append(args, productInp.Count)
		argId++
	}

	updatedAt := time.Now()

	setValues = append(setValues, fmt.Sprintf("updated_at=%d", argId))
	args = append(args, updatedAt)
	argId++

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE products SET %s WHERE id=$%d`, setQuery, argId+1)
	args = append(args, Id)
	_, err := p.db.Exec(query, args...)

	if err != nil {
		return err
	}

	return nil
}

func (p *productRepo) DeleteByID(ctx context.Context, id int) error {

	deleted_at := time.Now()

	_, err := p.db.Exec(`
		UPDATE products SET deleted_at=$2 WHERE id=$1`,
		id, deleted_at)

	if err != nil {
		return err
	}

	return nil
}

func (p *productRepo) GetPage(offset, limit int) (pb.ProductResponseList, error) {
	query := `
			SELECT id, name, price, count, created_at, updated_at
			FROM products WHERE deleted_at IS NULL
			LIMIT $1 OFFSET $2
	`
	rows, err := p.db.Query(query, limit, offset)

	if err != nil {
		return pb.ProductResponseList{}, err
	}

	productsList := pb.ProductResponseList{}

	for rows.Next() {

		var product pb.ProductResponse

		err := rows.Scan(&product.ID, &product.Name, &product.Price,
			&product.Count, product.CreatedAt, &product.UpdatedAt,
		)

		if err != nil {
			return pb.ProductResponseList{}, err
		}
		productsList.Products = append(productsList.Products, &product)

	}

	
	rows.Close()

	return productsList, rows.Err()
}
