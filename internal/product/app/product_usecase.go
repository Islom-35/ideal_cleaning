package app

import (
	"context"
	"log"

	"example.com/m/internal/genproto/product/pb"
	"example.com/m/internal/product/domain"
)

type ProductUseCase interface {
	Create(context.Context, pb.ProductRequest) (*pb.EmptyResponse,error)
	GetByID(ctx context.Context, ID pb.ID) (*pb.ProductResponse, error)
	UpdateProduct(ctx context.Context, productInp pb.UpdateProductRequest) error
	DeleteByID(ctx context.Context, id pb.ID) error
	GetPageProducts(ctx context.Context, page *pb.PageRequest) (*pb.ProductResponseList, error)
}

func NewProductUseCase(repo domain.ProductRepository) ProductUseCase {
	return &productUseCase{repo: repo}
}

type productUseCase struct {
	repo domain.ProductRepository
}

func (p *productUseCase) Create(ctx context.Context, product pb.ProductRequest)(*pb.EmptyResponse,error) {

	err := p.repo.Insert(ctx, product)

	if err != nil {
		log.Println(err)
		return nil,err	
	}

	return &pb.EmptyResponse{},nil
}

func (p *productUseCase) GetByID(ctx context.Context, ID pb.ID) (*pb.ProductResponse, error) {

	product, err := p.repo.GetByID(ctx, int(ID.ID))

	if err != nil {
		return &pb.ProductResponse{}, domain.ErrorProductNotFound
	}

	return &product, nil
}

func (p *productUseCase) UpdateProduct(ctx context.Context, productInp pb.UpdateProductRequest) error {
	productID := int(productInp.Id)
	_, err := p.repo.GetByID(ctx, productID)

	if err != nil {
		if err == domain.ErrorProductNotFound {
			return domain.ErrorProductNotFound
		}
		return err
	}

	count := int(productInp.Count)
	price := int(productInp.Price)

	product := domain.ProductUpdate{
		Name:  &productInp.Name,
		Price: &price,
		Count: &count,
	}

	if err := p.repo.UpdateByID(ctx, productID, product); err != nil {
		return err
	}

	return nil
}

func (p *productUseCase) DeleteByID(ctx context.Context, id pb.ID) error {
	_, err := p.repo.GetByID(ctx, int(id.ID))

	if err != nil {
		if err == domain.ErrorProductNotFound {
			return domain.ErrorProductNotFound
		}
		return err
	}

	if err := p.repo.DeleteByID(ctx, int(id.ID)); err != nil {
		return err
	}

	return nil
}

func (p *productUseCase) GetPageProducts(ctx context.Context, page *pb.PageRequest) (*pb.ProductResponseList, error) {
	limit := int(page.PageSize)
	offset :=(int(page.PageNumber)-1) * limit
	products, err := p.repo.GetPage(offset,limit)

	if err!=nil{
		return nil,err 
	}

	return &products,nil
}
