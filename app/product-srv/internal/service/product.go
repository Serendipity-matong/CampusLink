package service

import (
	"context"

	pb "github.com/campuslink/campuslink/api/product/v1"
)

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductReply, error) {
	return &pb.CreateProductReply{
		Message: "Product created successfully",
	}, nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductReply, error) {
	return &pb.GetProductReply{
		Product: &pb.ProductInfo{
			ProductId:      req.ProductId,
			Name:           "Sample Product",
			Description:    "Sample Description",
			Price:          99.99,
			Stock:          100,
			Category:       "book",
			ProductType:    "sale",
			Status:         "online",
		},
	}, nil
}

func (s *ProductService) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsReply, error) {
	return &pb.ListProductsReply{
		Products: []*pb.ProductInfo{},
		Total:    0,
	}, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductReply, error) {
	return &pb.UpdateProductReply{
		Message: "Product updated successfully",
	}, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductReply, error) {
	return &pb.DeleteProductReply{
		Message: "Product deleted successfully",
	}, nil
}

func (s *ProductService) PublishProduct(ctx context.Context, req *pb.PublishProductRequest) (*pb.PublishProductReply, error) {
	return &pb.PublishProductReply{
		Message: "Product published successfully",
	}, nil
}

func (s *ProductService) UnpublishProduct(ctx context.Context, req *pb.UnpublishProductRequest) (*pb.UnpublishProductReply, error) {
	return &pb.UnpublishProductReply{
		Message: "Product unpublished successfully",
	}, nil
}

func (s *ProductService) DeductStock(ctx context.Context, req *pb.DeductStockRequest) (*pb.DeductStockReply, error) {
	return &pb.DeductStockReply{
		RemainingStock: 90,
	}, nil
}

func (s *ProductService) AddStock(ctx context.Context, req *pb.AddStockRequest) (*pb.AddStockReply, error) {
	return &pb.AddStockReply{
		NewStock: 110,
	}, nil
}
