package services

import (
	"errors"
	"final-project-4/models"
	"final-project-4/models/inputs"
	"final-project-4/repositories"
)

type ProductService interface {
	CreateProduct(input inputs.CreateProduct) (models.Product, error)
	GetProductByID(ID int) (models.Product, error)
	GetProduct() ([]models.Product, error)
	UpdateProduct(ID int, input inputs.UpdateProduct) (models.Product, error)
	DeleteProduct(ID int) (models.Product, error)
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *productService {
	return &productService{productRepository}
}

func (s *productService) CreateProduct(input inputs.CreateProduct) (models.Product, error) {
	newProduct := models.Product{}
	newProduct.Title = input.Title
	newProduct.Price = input.Price
	newProduct.Stock = input.Stock
	newProduct.CategoryID = input.CategoryID

	created, err := s.productRepository.Save(newProduct)

	if err != nil {
		return models.Product{}, err
	}

	return created, nil
}

func (s *productService) GetProductByID(ID int) (models.Product, error) {
	product, err := s.productRepository.GetProductByID(ID)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *productService) GetProduct() ([]models.Product, error) {
	product, err := s.productRepository.GetAllProduct()

	if err != nil {
		return product, err
	}

	return product, nil
}
func (s *productService) UpdateProduct(ID int, input inputs.UpdateProduct) (models.Product, error) {
	Result, err := s.productRepository.GetProductByID(ID)

	if err != nil {
		return models.Product{}, err
	}

	if Result.ID == 0 {
		return models.Product{}, errors.New("not found!")
	}

	updated := models.Product{
		ID:    ID,
		Title: input.Title,
		Price: input.Price,
		Stock: input.Stock,
	}

	productUpdate, err := s.productRepository.Update(updated)

	if err != nil {
		return productUpdate, err
	}

	return productUpdate, nil
}

func (s *productService) DeleteProduct(ID int) (models.Product, error) {
	product, err := s.productRepository.GetProductByID(ID)

	if err != nil {
		return models.Product{}, err
	}

	if product.ID == 0 {
		return models.Product{}, nil
	}

	Deleted, err := s.productRepository.Delete(ID)

	if err != nil {
		return models.Product{}, err
	}

	return Deleted, nil
}
