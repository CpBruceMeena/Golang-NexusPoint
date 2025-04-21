package main

import (
	"context"

	productv1 "github.com/CpBruceMeena/golang-nexuspoint/proto/gen/go/product/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// productServer implements product.ProductService
type productServer struct {
	productv1.UnimplementedProductServiceServer
}

// GetProducts implements product.ProductService
func (s *productServer) GetProducts(ctx context.Context, req *productv1.GetProductsRequest) (*productv1.GetProductsResponse, error) {
	products := getStaticProducts()
	return &productv1.GetProductsResponse{Products: products}, nil
}

// GetProduct implements product.ProductService
func (s *productServer) GetProduct(ctx context.Context, req *productv1.GetProductRequest) (*productv1.GetProductResponse, error) {
	products := getStaticProducts()
	for _, p := range products {
		if p.Id == req.ProductId {
			return &productv1.GetProductResponse{Product: p}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "product not found")
}

// getStaticProducts returns a static list of products
func getStaticProducts() []*productv1.Product {
	return []*productv1.Product{
		{
			Id:          1,
			Name:        "Laptop",
			Description: "High-performance laptop",
			Price:       999.99,
			Stock:       10,
		},
		{
			Id:          2,
			Name:        "Smartphone",
			Description: "Latest smartphone model",
			Price:       699.99,
			Stock:       20,
		},
		{
			Id:          3,
			Name:        "Headphones",
			Description: "Noise-cancelling headphones",
			Price:       199.99,
			Stock:       15,
		},
	}
}
