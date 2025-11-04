package service

import (
	"github.com/campuslink/campuslink/app/product-srv/internal/biz"
	pb "github.com/campuslink/campuslink/api/product/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewProductService)

type ProductService struct {
	pb.UnimplementedProductServer
	biz *biz.ProductBiz
}

func NewProductService(biz *biz.ProductBiz) *ProductService {
	return &ProductService{
		biz: biz,
	}
}
