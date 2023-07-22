package main

import (
	"fmt"
	"log"
	"net"

	"github.com/RohithER12/product-svc/pkg/config"
	"github.com/RohithER12/product-svc/pkg/db"
	pb "github.com/RohithER12/product-svc/pkg/pb"
	"github.com/RohithER12/product-svc/pkg/repo"
	repoimpl "github.com/RohithER12/product-svc/pkg/repo/repoImpl"
	services "github.com/RohithER12/product-svc/pkg/services"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", c.Port)
	product := InitializeProductImpl(&h)
	s := services.Server{
		H:       h,
		Product: product,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

func InitializeProductImpl(h *db.Handler) repo.Product {
	wire.Build(repo.NewProductImpl)
	return &repoimpl.ProductImpl{
		H: *h,
	}
}
