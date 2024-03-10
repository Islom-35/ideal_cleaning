package domain

import "time"

type ProductRepository interface {
	ProductInsert(ProductInput) error
}

type ProductInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Count int    `json:"count"`
}

type Product struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Count      int       `json:"count"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at"`
}

type ProductUpdateInsert struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Count      int       `json:"count"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at"`
}

// CREATE TABLE "products" (
//     "id" SERIAL PRIMARY KEY,
//     "name" VARCHAR,
//     "price" int,
//     "count" int,
//     "created_at" TIMESTAMP DEFAULT current_timestamp,
//     "updated_at" TIMESTAMP DEFAULT current_timestamp,
//     "deleted_at" TIMESTAMP
// );
