package repo

import (
	"github.com/RohithER12/product-svc/pkg/models"
	repoimpl "github.com/RohithER12/product-svc/pkg/repo/repoImpl"
)

type Product interface {
	Create(product models.Product) error
	FindOne(id int64) (models.Product, error)
	ListAll() ([]models.Product, error)
	Search(target string) ([]models.Product, error)
	SortByPrice(asn bool) ([]models.Product, error)
}

func NewProductImpl() Product {
	return &repoimpl.ProductImpl{}
}
