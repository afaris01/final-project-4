package repositories

import (
	"final-project-4/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(input models.Category) (models.Category, error)
	CheckType(tipe string) (models.Category, error)
	GetByID(ID int) (models.Category, error)
	Update(ID int, category models.Category) (models.Category, error)
	Delete(ID int) (models.Category, error)
	GetCategory() ([]models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Save(category models.Category) (models.Category, error) {
	err := r.db.Create(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) CheckType(tipe string) (models.Category, error) {
	typeSame := models.Category{}

	err := r.db.Where("type = ?", tipe).Find(&typeSame).Error

	if err != nil {
		return models.Category{}, err
	}

	return typeSame, nil
}

func (r *categoryRepository) GetByID(ID int) (models.Category, error) {
	categoryResult := models.Category{}

	err := r.db.Where("id = ?", ID).Find(&categoryResult).Error

	if err != nil {
		return models.Category{}, err
	}

	return categoryResult, nil
}

func (r *categoryRepository) GetCategory() ([]models.Category, error) {
	categoryResult := []models.Category{}

	err := r.db.Find(&categoryResult).Error

	if err != nil {
		return categoryResult, err
	}

	return categoryResult, nil
}

func (r *categoryRepository) Update(ID int, category models.Category) (models.Category, error) {
	err := r.db.Where("id = ?", ID).Updates(&category).Error

	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (r *categoryRepository) Delete(ID int) (models.Category, error) {

	Deleted := models.Category{
		ID: ID,
	}

	err := r.db.Where("id = ?", ID).Delete(&Deleted).Error

	if err != nil {
		return models.Category{}, err
	}

	return Deleted, err
}
