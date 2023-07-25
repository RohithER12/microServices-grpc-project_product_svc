package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RohithER12/product-svc/pkg/db"
	"github.com/RohithER12/product-svc/pkg/models"
	pb "github.com/RohithER12/product-svc/pkg/pb"
	"github.com/RohithER12/product-svc/pkg/repo"
)

type Server struct {
	H db.Handler
	pb.UnimplementedProductServiceServer
	Product repo.Product
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product

	product.Name = req.Name
	product.Stock = req.Stock
	product.Price = req.Price

	if err := s.Product.Create(product); err != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, err
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.Id,
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {

	product, err := s.Product.FindOne(req.Id)
	if err != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:    product.Id,
		Name:  product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {

	products, err := s.Product.ListAll()
	if err != nil {
		return &pb.ListProductsResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	var response pb.ListProductsResponse
	for _, product := range products {
		data := &pb.FindOneData{
			Id:    product.Id,
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
		}
		response.Data = append(response.Data, data)
	}

	return &response, nil
}

func (s *Server) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchReponse, error) {
	target := req.Search
	fmt.Println("\ntarget\t", target)
	products, err := s.Product.Search(target)
	if err != nil {
		return &pb.SearchReponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}
	var response pb.SearchReponse
	for _, product := range products {
		fmt.Println("\nname:\t", product.Name)

		data := &pb.FindOneData{
			Id:    product.Id,
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
		}
		response.Data = append(response.Data, data)
	}

	return &response, nil
}

func (s *Server) SortByPrice(ctx context.Context, req *pb.SortByPriceRequest) (*pb.SortByPriceResponse, error) {
	order := true
	if req.Sort != "ASC" {
		order = false
	}
	fmt.Println(order)

	products, err := s.Product.SortByPrice(order)
	if err != nil {
		return &pb.SortByPriceResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}
	var response pb.SortByPriceResponse
	for _, product := range products {
		data := &pb.FindOneData{
			Id:    product.Id,
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
		}
		response.Data = append(response.Data, data)
	}

	return &response, nil

}

func (s *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var product models.Product

	if result := s.H.DB.First(&product, req.Id); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	if product.Stock <= 0 {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock too low",
		}, nil
	}

	var log models.StockDecreaseLog
	fmt.Println("\n\norderid\n\n", req.OrderId)
	if result := s.H.DB.Where(&models.StockDecreaseLog{OrderId: req.OrderId}).First(&log); result.Error == nil {
		fmt.Println(result)
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock already decreased",
		}, nil
	}

	product.Stock = product.Stock - req.Quantity

	s.H.DB.Save(&product)

	log.OrderId = req.OrderId
	log.ProductRefer = product.Id

	s.H.DB.Create(&log)

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}

// func (s *Server) mustEmbedUnimplementedProductServiceServer() {
// 	// Empty implementation
// }
