package repoimpl

import (
	"fmt"

	"github.com/RohithER12/product-svc/pkg/db"
	"github.com/RohithER12/product-svc/pkg/models"
)

type ProductImpl struct {
	H db.Handler
}

func (p *ProductImpl) Create(product models.Product) error {

	if result := p.H.DB.Create(&product); result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *ProductImpl) FindOne(id int64) (models.Product, error) {
	var product models.Product
	if result := p.H.DB.First(&product, id); result.Error != nil {
		return models.Product{}, result.Error
	}

	return product, nil
}

func (p *ProductImpl) ListAll() ([]models.Product, error) {
	products := []models.Product{}
	if result := p.H.DB.Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil

}

func (p *ProductImpl) Search(target string) ([]models.Product, error) {
	fmt.Println("Searching for:", target)

	products := []models.Product{}
	if result := p.H.DB.Where("name ILIKE ?", target+"%").Find(&products); result.Error != nil {
		fmt.Println("Database error:", result.Error)
		return nil, result.Error
	}
	fmt.Println("Searching result:", products)
	return products, nil
}

func (p *ProductImpl) SortByPrice(asn bool) ([]models.Product, error) {
	products := []models.Product{}
	order := "ASC"
	if !asn {
		order = "DESC"
	}

	if result := p.H.DB.Order("price " + order).Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
