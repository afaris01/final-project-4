package repositories

import (
	"final-project-4/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductByID(ID int) (models.Product, error)
	Update(newProduct models.Product) (models.Product, error)
	Save(product models.Product) (models.Product, error)
	Delete(ID int) (models.Product, error)
	GetAllProduct() ([]models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) Save(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (r *productRepository) GetAllProduct() ([]models.Product, error) {
	productResult := []models.Product{}

	err := r.db.Find(&productResult).Error

	if err != nil {
		return productResult, err
	}

	return productResult, nil
}

func (r *productRepository) GetProductByID(ID int) (models.Product, error) {
	productResult := models.Product{}

	err := r.db.Where("id = ?", ID).Find(&productResult).Error

	if err != nil {
		return models.Product{}, err
	}

	return productResult, nil
}

func (r *productRepository) Update(newProduct models.Product) (models.Product, error) {
	err := r.db.Where("id = ?", newProduct.ID).Updates(newProduct).Error

	if err != nil {
		return models.Product{}, err
	}

	return newProduct, nil
}

func (r *productRepository) Delete(ID int) (models.Product, error) {
	productDeleted := models.Product{
		ID: ID,
	}

	err := r.db.Where("id = ?", ID).Delete(&productDeleted).Error

	if err != nil {
		return models.Product{}, err
	}

	return productDeleted, err
}
