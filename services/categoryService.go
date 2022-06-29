package services

import (
	"errors"
	"final-project-4/models"
	"final-project-4/models/inputs"
	"final-project-4/repositories"
)

type CategoryService interface {
	CreateCategory(input inputs.CreateCategory) (models.Category, error)
	GetCategoryByID(ID int) (models.Category, error)
	UpdateCategory(ID int, input inputs.UpdateCategory) (models.Category, error)
	DeleteCategory(ID int) (models.Category, error)
	GetCategory() ([]models.Category, error)
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) *categoryService {
	return &categoryService{categoryRepository}
}

func (s *categoryService) CreateCategory(input inputs.CreateCategory) (models.Category, error) {
	newCategory := models.Category{}
	newCategory.Type = input.Type

	created, err := s.categoryRepository.Save(newCategory)

	if err != nil {
		return models.Category{}, err
	}

	return created, nil
}

func (s *categoryService) GetCategoryByID(ID int) (models.Category, error) {
	category, err := s.categoryRepository.GetByID(ID)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *categoryService) GetCategory() ([]models.Category, error) {
	category, err := s.categoryRepository.GetCategory()

	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *categoryService) UpdateCategory(ID int, input inputs.UpdateCategory) (models.Category, error) {
	Result, err := s.categoryRepository.GetByID(ID)

	if err != nil {
		return models.Category{}, err
	}

	if Result.ID == 0 {
		return models.Category{}, errors.New("not found!")
	}

	updated := models.Category{
		Type: input.Type,
	}

	categoryUpdate, err := s.categoryRepository.Update(ID, updated)

	if err != nil {
		return categoryUpdate, err
	}

	return categoryUpdate, nil
}

func (s *categoryService) DeleteCategory(ID int) (models.Category, error) {
	category, err := s.categoryRepository.GetByID(ID)

	if err != nil {
		return models.Category{}, err
	}

	if category.ID == 0 {
		return models.Category{}, nil
	}

	Deleted, err := s.categoryRepository.Delete(ID)

	if err != nil {
		return models.Category{}, err
	}

	return Deleted, nil
}
