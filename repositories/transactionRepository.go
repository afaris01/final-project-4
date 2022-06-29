package repositories

import (
	"final-project-4/models"

	"gorm.io/gorm"
)

type TransactionHistoryRepository interface {
	GetTransactionByIDUser(ID int) ([]models.TransactionHistory, error)
	GetAllTransaction() ([]models.TransactionHistory, error)
	Save(input models.TransactionHistory) (models.TransactionHistory, error)
}

type transactionHistoryRepository struct {
	db *gorm.DB
}

func NewTransactionHistoryRepository(db *gorm.DB) *transactionHistoryRepository {
	return &transactionHistoryRepository{db}
}

func (r *transactionHistoryRepository) GetTransactionByIDUser(ID int) ([]models.TransactionHistory, error) {
	allTransactions := []models.TransactionHistory{}
	err := r.db.Where("user_id", ID).Find(&allTransactions).Error

	if err != nil {
		return []models.TransactionHistory{}, err
	}

	return allTransactions, nil
}

func (r *transactionHistoryRepository) GetAllTransaction() ([]models.TransactionHistory, error) {
	allTransactions := []models.TransactionHistory{}
	err := r.db.Find(&allTransactions).Error

	if err != nil {
		return []models.TransactionHistory{}, err
	}

	return allTransactions, nil
}

func (r *transactionHistoryRepository) Save(input models.TransactionHistory) (models.TransactionHistory, error) {
	err := r.db.Create(&input).Error

	if err != nil {
		return models.TransactionHistory{}, err
	}

	return input, nil
}
