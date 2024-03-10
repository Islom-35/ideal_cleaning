package app

import "ideal_cleaning/internal/domain"

type ProductService interface{
	CreateProduct(product domain.ProductInput)error
}

type productService struct{
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository)ProductService{
	return &productService{
		repo: repo,
	}
}

func (p *productService)CreateProduct(product domain.ProductInput)error{
	return p.repo.ProductInsert(product)
}