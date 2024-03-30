package domain

import (
	"example.com/m/internal/genproto/product/pb"
	"context"
)

type ProductRepository interface {
	Insert(context.Context, pb.ProductRequest) error
	GetByID(ctx context.Context,ID int)(pb.ProductResponse,error)
	UpdateByID(ctx context.Context, Id int, productInp ProductUpdate) error
	DeleteByID(ctx context.Context, id int)error
	GetPage(offset, limit int) (pb.ProductResponseList, error)
}
